package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGet(t *testing.T) {
	value := get()
	t.Logf("Value %s", value)
	sett := itemSet{}
	sett.addAll([]Item{Rock, Paper, Scissor})
	assert.Equal(t, sett.has(value), true)
}

func TestRockPaperScissor(t *testing.T) {
	t.Log("Me = scissor, you = rock You win")
	result := rockPaperScissor(Rock, Scissor)
	assert.Equal(t, result, "You win")
	t.Log("Me = rock, you = rock Tie")
	result2 := rockPaperScissor(Rock, Rock)
	assert.Equal(t, result2, "Tie")
	t.Log("Me = paper, you = rock I win")
	result3 := rockPaperScissor(Rock, Paper)
	assert.Equal(t, result3, "I win")
}

type itemSet map[Item]struct{}

func (sett itemSet) add(item Item) {
	sett[item] = struct{}{}
}

func (sett itemSet) addAll(items []Item) {
	for _, item := range items {
		sett.add(item)
	}
}

func (sett itemSet) remove(item Item) {
	delete(sett, item)
}

func (sett itemSet) has(item Item) bool {
	_, ok := sett[item]
	return ok
}
