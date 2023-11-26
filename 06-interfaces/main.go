package main

import "fmt"

// concrete type: map, struct, int, string, ...
// interface type
// implict
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type koreanBot struct{}
type japaneseBot struct{}

func main() {
	eb := englishBot{}
	kb := koreanBot{}
	jb := japaneseBot{}

	printGreeting(eb)
	printGreeting(kb)
	printGreeting(jb)
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (koreanBot) getGreeting() string {
	return "안녕"
}

func (japaneseBot) getGreeting() string {
	return "こんにちは"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
