package main

// import "fmt"
func main() {
	cards := newDeckFromFile("my_caards")
	// cards := newDeckFromFile("my") :To check error handling
	cards.print()

}
