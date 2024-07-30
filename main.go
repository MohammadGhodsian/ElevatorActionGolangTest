package main

type Direction bool

const (
	Up   Direction = true
	Down Direction = false
)

type Person struct {
	From int
	To   int
}

func (p Person) direction() Direction {
	if p.From < p.To {
		return Up
	}
	return Down
}
