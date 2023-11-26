package main

import (
	"fmt"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (journal *Journal) AddEntry(text string) {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	journal.entries = append(journal.entries, entry)
}

func (journal *Journal) RemoveEntry(index int) {
	//remove entry
}

//breaks srp

func (journal *Journal) Save(filename string) {
	_ = os.WriteFile(filename,
		[]byte(strings.Join(journal.entries, "\n")), 0644)
}

func (journal *Journal) Load(filename string) {
	content, _ := os.ReadFile(filename)
	fmt.Printf("content:%v\n", string(content))
}

func main() {
	j := Journal{
		entries: []string{
			"Good Morning",
			"This is srp",
		},
	}
	j.AddEntry("Welcome")

	j.Save("File.txt")
	j.Load("File.txt")
}
