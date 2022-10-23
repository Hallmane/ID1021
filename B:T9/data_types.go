package main

type Letter rune
type Key int

const NODESIZE = 27

type Trie struct {
	root *Node
}

type Node struct {
	next [NODESIZE]*Node
	word bool
}

var KeyPad = map[int]string{
	1: "abc",
	2: "def",
	3: "ghi",
	4: "jkl",
	5: "mno",
	6: "prs",
	7: "tuv",
	8: "xyz",
	9: "åäö",
}

func åäöCases(val rune) Letter {
	switch {
	case val == 'å':
		return 24
	case val == 'ä':
		return 25
	case val == 'ö':
		return 26
	}
	return 0
}

const (
	a Letter = iota
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	p
	r
	s
	t
	u
	v
	x
	y
	z
	å
	ä
	ö
)
