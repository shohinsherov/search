package search

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

// Result описывает один результать поиска
type Result struct {
	Phrase  string
	Line    string
	LineNum int64
	ColNum  int64
}

// All ...
func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	ch := make(chan []Result,100000)
	defer close(ch)

	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(_ctx context.Context, index int) {
			defer wg.Done()
			src, err := os.Open(files[index])
			if err != nil {
				log.Print(err)
			}
			defer func() {
				if cerr := src.Close(); cerr != nil {
					log.Print(cerr)
				}
			}()

			reader := bufio.NewReader(src)
			lineNum := int64(1)
			for {
				line, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Print(err)
				}

				cont := strings.Contains(line, phrase)
				if cont {
					ind := strings.Index(line, phrase)

					res := []Result{{
						Phrase:  phrase,
						Line:    string(line[:len(line)-1]),
						LineNum: lineNum,
						ColNum:  int64(ind)},
					}
					log.Print(res)
					mu.Lock()
					ch <- res
					mu.Unlock()
				}
			}

		}(ctx, i)
	}

	wg.Wait()
	return ch
}

// AllTest ....
func AllTest(phrase string, files []string) []Result {
	result := []Result{}

	for i := 0; i < len(files); i++ {
		src, err := os.Open(files[i])
		if err != nil {
			log.Print(err)
		}

		defer func() {
			if cerr := src.Close(); cerr != nil {
				log.Print(cerr)
			}
		}()

		reader := bufio.NewReader(src)
		lineNum := int64(1)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				log.Print(line)
				break
			}

			if err != nil {
				log.Print(err)
				break
			}
			cantain := strings.Contains(line, "Shohin")
			//log.Print(cantain)
			if cantain {
				abs := strings.Index(line, "Shohin")
				//val := strings.Split(line, "\n")
				val := string(line[:len(line)-1])
				//val = val[:len(val)-1]
				log.Print("-----")
				log.Print(val)
				res := Result{
					Phrase:  phrase,
					Line:    val,
					LineNum: lineNum,
					ColNum:  int64(abs + 1),
				}
				result = append(result, res)
			}
			lineNum++
			//start := strings.HasPrefix(line, "Shohin")
			//log.Print(start)

		}
		//log.Print("----------------------------")
	}
	return result
}
