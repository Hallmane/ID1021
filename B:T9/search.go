package main

import (
	"strconv"
)

// function wrapper
func (t *Trie) Search(sequence string) []string {
	possible := t.root._Search(sequence, "")
	return possible
}

// finds traversable "letters" from any path
func (n *Node) Path_Finder(key int) ([]Letter, bool) {
	alternatives := KeyPad[key]
	var possible_paths []Letter
	for _, k := range alternatives {
		if n.next[rune_To_Letter(k)] != nil {
			possible_paths = append(possible_paths, rune_To_Letter(k))
		}
	}
	return possible_paths, n.word
}

// called with decreasing length of possible word called path
func (n *Node) _Search(sequence string, path string) []string {
	traversed := make([]string, 0)

	if len(sequence) > 0 {
		key, _ := strconv.Atoi(string(sequence[0]))
		possible_paths, _ := n.Path_Finder(key)

		for _, pp := range possible_paths {
			if n.next[pp] != nil {
				path += string(letter_To_Rune(pp))
				traversed = append(traversed, n.next[pp]._Search(sequence[1:], path)...)
				path = ""
			}
		}
	} else {
		if n.word {
			traversed = append(traversed, path)
			return traversed
		}
	}
	return traversed
}
