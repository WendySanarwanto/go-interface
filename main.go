package main

import "fmt"

type englishBot struct {}
type koreanBot struct {}

func main() {
	eb := englishBot{}
	kb := koreanBot{}

	eb.printGreeting()
	kb.printGreeting()
}

func (eb englishBot) printGreeting() {
	fmt.Println(eb.getGreeting())
}

func (kb koreanBot) printGreeting() {
	fmt.Println(kb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi There"
}

func (kb koreanBot) getGreeting() string {
	return "안녕하세요"
}
