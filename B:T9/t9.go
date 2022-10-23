package main

// á la deni
func (t *Trie) Add(word string) {
	cur := t.root
	runes := []rune(word)
	var code Letter
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		code = rune_To_Letter(char)
		if cur.next[code] == nil {
			cur.next[code] = new_node()
		}
		cur = cur.next[code]
	}
	cur.word = true
}

func letter_To_Rune(letter Letter) rune {
	return rune(letter + 'a')
}

func rune_To_Letter(char rune) Letter {
	if char == 'å' || char == 'ä' || char == 'ö' {
		return åäöCases(char)
	}
	return Letter(char - 'a')
}

func new_node() *Node {
	return &Node{word: false}
}
