package main

import (
	"log"
	"context"
	"github.com/shohinsherov/search/pkg/search"
)

//"regexp"

func main() {

	files := []string{"cmd/data/text.txt", "cmd/data/text1.txt"}

	root := context.Background()
	ctx, _ := context.WithCancel(root)

	aa := search.All(ctx, "Shohin", files)

	log.Print(aa)

}

/*log.Print(files)

	gert, _ := os.Getwd()

	log.Print(gert)
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

			//ss, _ := regexp.Compile("Shohin")

			//inds := ss.FindAllStringIndex(line, -1)
		//	log.Print(inds)
			//log.Print(ss)
			cantain := strings.Contains(line, "Shohin")
			//log.Print(cantain)
			if cantain {
			ind := strings.Index(line, "Shohin")
			log.Print(ind)
			}
			//start := strings.HasPrefix(line, "Shohin")
			//log.Print(start)


		}
		log.Print("----------------------------")
	}

}
*/
