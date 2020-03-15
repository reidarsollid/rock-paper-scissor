package main

import (
	"math/rand"
	"time"
)

func Play(value string) (string, string) {
	me := get()
	return me.String(), rockPaperScissor(ParseItem(value), me)
}

func get() Item {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Item(r1.Intn(3))
}

func rockPaperScissor(you Item, me Item) string {
	switch you {
	case Rock:
		if me == Paper {
			return "I win"
		}
		if me == Scissor {
			return "You win"
		}
	case Paper:
		if me == Scissor {
			return "I win"
		}
		if me == Rock {
			return "You win"
		}
	case Scissor:
		if me == Rock {
			return "I win"
		}
		if me == Paper {
			return "You win"
		}
	}
	return "Tie"
}
