package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestParseItem(t *testing.T) {
	assert.Equal(t, ParseItem("Rock"), Rock)
	assert.Equal(t, ParseItem("Paper"), Paper)
	assert.Equal(t, ParseItem("Scissor"), Scissor)
}

func TestPlay(t *testing.T) {
	play, s := Play("Rock")
	switch play {
	case Paper.String():
		assert.Equal(t, s, "I win")
	case Rock.String():
		assert.Equal(t, s, "Tie")
	case Scissor.String():
		assert.Equal(t, s, "You win")
	}
}

func TestItem_String(t *testing.T) {
	assert.Equal(t, Rock.String(), "Rock")
	assert.Equal(t, Paper.String(), "Paper")
	assert.Equal(t, Scissor.String(), "Scissor")
}
