package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func (d deck) toString() string {
	return (strings.Join([]string(d), ","))
}
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
func readFile(filename string) (deck, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return deck(strings.Split(string(data), ",")), nil
}
func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
func (d deck) printeven() {
	for i, card := range d {
		if i%2 == 0 {
			fmt.Println(card)
		}
	}
}
