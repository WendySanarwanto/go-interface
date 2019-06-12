package main

type englishBot struct {}
type koreanBot struct {}

func main() {
	eb := englishBot{}
	kb := koreanBot{}

	printGreeting(eb)
	printGreeting(kb)
}

func (eb englishBot) getGreeting() string {
	return "Hi There"
}

func (kb koreanBot) getGreeting() string {
	return "안녕하세요"
}
