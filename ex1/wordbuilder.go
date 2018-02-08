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
	"fmt"
	moby "./moby"
	"log"
	"path/filepath"
)

const (
	outfile string = "result.txt"
	infile  string = "moby-000.txt"
)

func main() {
	wb := new(moby.WordBuilder)
	wb.Init()
	abspath, _ := filepath.Abs(infile)
	log.Println("Parsing " + abspath)
	err := wb.Parse(infile)
	if err != nil {
		fmt.Println(err)
		return
	}
	abspath, _ = filepath.Abs(outfile)
	log.Println("Saving results to " + abspath)
	wb.SaveToFile(outfile)
}
