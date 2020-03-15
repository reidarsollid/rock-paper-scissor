package main

type Item int

const (
	Rock Item = iota
	Paper
	Scissor
)

func (item Item) String() string {
	items := [...]string{
		"Rock",
		"Paper",
		"Scissor"}
	if item > Rock || item < Scissor {
		return "Unknown"
	}
	return items[item]
}

func ParseItem(item string) Item {
	items := map[string]Item{
		"Rock":    Rock,
		"Paper":   Paper,
		"Scissor": Scissor}

	return items[item]
}

