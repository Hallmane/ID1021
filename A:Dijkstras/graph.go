package main

const HASHMOD = 541

type Map struct {
	cities [HASHMOD]*City
}

type City struct {
	name        string
	connections []Connection
}

type Connection struct {
	city     *City
	distance int
}
