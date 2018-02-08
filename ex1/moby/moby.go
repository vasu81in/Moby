/**
 * moby.go -- Builds word frequency
 * 
 *
 * @author Vasu Mahalingam <vasu.uky@gmail.com>
 *
 *
*/


package moby

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

type WordBuilder struct {
	WordCount map[string]int
}

func (this *WordBuilder) Init() {
	this.WordCount = make(map[string]int)
}

func (this *WordBuilder) GetWordCount() map[string]int {
	return this.WordCount
}

func (this *WordBuilder) SaveToFile(filename string) error {
	if len(this.WordCount) == 0 {
		err := errors.New("Nothing to write, \"" + filename + "\"  not created")
		log.Fatal(err)
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer func() error {
		if err := f.Close(); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}()

	w := bufio.NewWriter(f)
	bw := bufio.NewWriterSize(w, 4096)

	for k, v := range this.WordCount {
		bw.WriteString(k + ":" + strconv.Itoa(v) + "\n")
	}
	w.Flush()

	return nil
}

func (this *WordBuilder) Parse(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return errors.New("File not found")
	}

	defer func() error {
		if err = f.Close(); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}()

	if this.WordCount == nil {
		return errors.New("Word map not initialized")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > 0 {
			if _, ok := this.WordCount[word]; !ok {
				this.WordCount[word] = 1
			} else {
				this.WordCount[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return errors.New("Scanner read error")
	}

	return nil
}
