package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	deck := deck {}
	cardSuits := []string { "Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string { "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, value + " of " + suite)
		}
	}

	return deck
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toFile(filename string) {
	ioutil.WriteFile(filename + ".csv", []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	deckString := strings.Join(d, ",")
	return deckString
}

func deckFromFile(filepath string) deck {
	var cards deck
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	commaString := string(bs)
	cards = strings.Split(commaString, ",")
	return deck(cards)
}

func (d deck) shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for  i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

