package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	//if length of our deck is not 16, then throw an error message
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	//if the 1st card in the deck isn't ace of spades, throw error msg
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	//if last card isn't 4 of clubs, throw error msg
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}
