package main

import "fmt"

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

type Elevator struct {
	CurrentFloor int
	Direction    Direction
}

func (e *Elevator) move() {
	if e.Direction == Up {
		e.CurrentFloor++
	} else if e.Direction == Down {
		e.CurrentFloor--
	}
}

func Order(startingFloor int, queue []Person) []int {
	return []int{}
}

func main() {
	queue := []Person{
		{From: 5, To: 4},  // 1st passenger
		{From: 5, To: 3},  // 2nd passenger
		{From: 3, To: 4},  // 3rd passenger
		{From: 0, To: 2},  // 5th passenger
		{From: 3, To: -4}, // 4th passenger
		{From: 1, To: 2},  // 6th passenger
	}
	startingFloor := 5
	result := Order(startingFloor, queue)
	fmt.Println(result) // Expected: []int{5, 4, 3, 4, 3, -4, 0, 1, 2}
}
