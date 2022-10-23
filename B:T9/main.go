package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	quick_add()
}

// Sheiße funkade finally
func quick_add() {
	trie := Trie{new_node()}
	f, err := os.Open("kelly.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		trie.Add(s)
	}

	word := trie.Search("12323676353")
	fmt.Println(word)
}
