package main

import "testing"

func TestNewDeck(t *testing.T)  {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck lenght of 52 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card King of clubs but got %v", d[len(d)-1])
	}
}