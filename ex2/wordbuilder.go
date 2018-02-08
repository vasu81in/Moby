/**
 * wordbuilder.go -- Builds word frequency
 * from moby txt files
 * 
 *
 * @author Vasu Mahalingam <vasu.uky@gmail.com>
 *
 *
*/

package main

import (
	"./moby"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

const (
	dir string = "files"
	outfile string = "result.txt"
)

func main() {
	abspath, _ := filepath.Abs(dir)
	log.Println("Reading moby files from " + abspath)
	files, err := ioutil.ReadDir(abspath)
	if err != nil {
		log.Fatal(err)
		return
	}
	resc, errc := make(chan map[string]int), make(chan error)
	for _, file := range files {
		if file.IsDir() == true {
			log.Fatal("Nested directory not allowed")
			return
		}
		filename := abspath + string(filepath.Separator) + file.Name()
		go func(filename string) {
			wb := new(moby.WordBuilder)
			wb.Init()
			err := wb.Parse(filename)
			if err != nil {
				errc <- err
				return
			}
			resc <- wb.GetWordCount()
		}(filename)
	}

	result := make(map[string]int)
	for i := 0; i < len(files); i++ {
		select {
		case res := <-resc:
			for i, j := range res {
				if _, ok := result[i]; !ok {
					result[i] = j
				} else {
					result[i] += j
				}
			}
		case err := <-errc:
			fmt.Println(err)
		}
	}

	// Save the result  
	wb := new(moby.WordBuilder)
	wb.WordCount = result
	abspath, _ = filepath.Abs(outfile)
	log.Println("Saving results to " + abspath)
	wb.SaveToFile(outfile)
}
