package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type thaiBot struct{}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}
	// tb := thaiBot{}

	printGreeting(eb)
	// printGreeting(sb)
	// printGreeting(tb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// func printGreeting(tb thaiBot) {
// 	fmt.Println(tb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (thaiBot) getGreeting() string {
	return "สบายดีนะไอ้สัส!"
}