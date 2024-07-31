package main

import (
	"fmt"
)

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
	return direction(p.From, p.To)
}

func direction(from, to int) Direction {
	if from < to {
		return Up
	}
	return Down
}

// type Elevator struct {
// 	CurrentFloor int
// 	Direction    Direction
// }

// func (e *Elevator) move() {
// 	if e.Direction == Up {
// 		e.CurrentFloor++
// 	} else if e.Direction == Down {
// 		e.CurrentFloor--
// 	}
// }

func Order(startingFloor int, queue []Person) []int {
	if len(queue) == 0 {
		return []int{startingFloor}
	}
	// rresult list of floors
	resultStops := []int{}
	// current passengers in elevator
	passengers := []Person{}
	// current floor of elevator in each step
	currentFloor := startingFloor
	// intial the direction of elevator by first person in list
	var directionElev Direction = direction(currentFloor, queue[0].From)
	// // create queue of up and down (passengers directions)
	// downQueue := []Person{}
	// upQueue := []Person{}
	// for _, person := range queue {
	// 	if person.direction() == Up {
	// 		upQueue = append(upQueue, person)
	// 	} else {
	// 		downQueue = append(downQueue, person)
	// 	}
	// }

	addStop := func(floor int) {
		if len(resultStops) == 0 || resultStops[len(resultStops)-1] != floor {
			resultStops = append(resultStops, floor)
		}
	}
	elevatorMustMove := func() bool {
		return len(queue) > 0 || len(passengers) > 0
	}

	for elevatorMustMove() {

		// Drop off passengers
		droppedOff := false
		for i := 0; i < len(passengers); i++ {
			if passengers[i].To == currentFloor {
				fmt.Printf("droppedOff  %v   %v  currentFloor %d     queue %+v   passengers %+v\n", droppedOff, passengers[i], currentFloor, queue, passengers)
				passengers = append(passengers[:i], passengers[i+1:]...)
				i--
				droppedOff = true
			}
		}
		if droppedOff {
			addStop(currentFloor)
		}

		if !elevatorMustMove() {
			break
		}

		// // Manage the direction based on the passenger queue and destination
		if len(passengers) == 0 {
			if currentFloor != queue[0].From {
				directionElev = direction(currentFloor, queue[0].From)
			} else {
				directionElev = direction(currentFloor, queue[0].To)
			}
		} else {
			directionElev = direction(currentFloor, passengers[0].To)
		}

		// Pick up passengers
		pickedUp := false
		for i := 0; i < len(queue); i++ {
			if queue[i].direction() == directionElev && queue[i].From == currentFloor {
				fmt.Printf("pickedUp  %v   %v    currentFloor %d    queue %+v   passengers %+v\n", pickedUp, queue[i], currentFloor, queue, passengers)
				passengers = append(passengers, queue[i])
				queue = append(queue[:i], queue[i+1:]...)
				i--
				pickedUp = true
			}
		}
		if pickedUp {
			addStop(currentFloor)
		}

		// Determine the next floor to move to
		if elevatorMustMove() {
			if directionElev == Up {
				currentFloor++
			} else {
				currentFloor--
			}
		}
	}

	return resultStops
}

func main() {
	queue := []Person{{From: 5, To: 4}, // 1st passenger
		{From: 5, To: 3},  // 2nd passenger
		{From: 3, To: 4},  // 3rd passenger
		{From: 0, To: 2},  // 5th passenger
		{From: 3, To: -4}, // 4th passenger
		{From: 1, To: 2}}
	startingFloor := 5
	// queue := []Person{
	// 	{From: 3, To: 2}, // Al
	// 	{From: 5, To: 2}, // Betty
	// 	{From: 2, To: 1}, // Charles
	// 	{From: 2, To: 5}, // Dan
	// 	{From: 4, To: 3}, // Ed
	// }
	// startingFloor := 1

	result := Order(startingFloor, queue)
	fmt.Println(result) // Expected: []int{5, 4, 3, 4, 3, -4, 0, 1, 2}
}
