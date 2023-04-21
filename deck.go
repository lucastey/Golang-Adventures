package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck', which is a slice of strings
type deck []string

// Function to create a new deck
func newDeck() deck {
	//initialize empty deck (string slice)
	cards := deck{}

	//initialize range of suits and values for individual card combi
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//In Golang, if we don't use the index or value, we can use the underscore as placeholder
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	// return up till size of hand and return from size of hand to end
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	//iterate through every card in the range and it's index
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	//utilizing write-file library func to save deck to file as string
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	//generate new source for random generator
	source := rand.NewSource(time.Now().UnixNano())
	//declare new random generator based on new rand source
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
