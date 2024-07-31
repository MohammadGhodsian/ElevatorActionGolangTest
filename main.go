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

func Order(startingFloor int, queue []Person) []int {
	if len(queue) == 0 {
		return []int{startingFloor}
	}

	resultStops := []int{}
	passengers := []Person{}
	currentFloor := startingFloor

	for hasWork(queue, passengers) {
		dropOffPassengers(&passengers, currentFloor, &resultStops)
		if !hasWork(queue, passengers) {
			break
		}

		directionElev := determineDirection(queue, passengers, currentFloor)
		pickUpPassengers(&queue, &passengers, currentFloor, directionElev, &resultStops)
		moveElevator(&currentFloor, directionElev)
	}

	return resultStops
}

func hasWork(queue []Person, passengers []Person) bool {
	return len(queue) > 0 || len(passengers) > 0
}

func dropOffPassengers(passengers *[]Person, currentFloor int, resultStops *[]int) {
	for i := 0; i < len(*passengers); i++ {
		if (*passengers)[i].To == currentFloor {
			fmt.Printf("Dropped off %v at floor %d\n", (*passengers)[i], currentFloor)
			*passengers = append((*passengers)[:i], (*passengers)[i+1:]...)
			i--
			addStop(currentFloor, resultStops)
		}
	}
}

func determineDirection(queue []Person, passengers []Person, currentFloor int) Direction {
	if len(passengers) == 0 {
		if currentFloor != queue[0].From {
			return direction(currentFloor, queue[0].From)
		}
		return direction(currentFloor, queue[0].To)
	}
	return direction(currentFloor, passengers[0].To)
}

func pickUpPassengers(queue *[]Person, passengers *[]Person, currentFloor int, directionElev Direction, resultStops *[]int) {
	for i := 0; i < len(*queue); i++ {
		if (*queue)[i].direction() == directionElev && (*queue)[i].From == currentFloor {
			fmt.Printf("Picked up %v at floor %d\n", (*queue)[i], currentFloor)
			*passengers = append(*passengers, (*queue)[i])
			*queue = append((*queue)[:i], (*queue)[i+1:]...)
			i--
			addStop(currentFloor, resultStops)
		}
	}
}

func moveElevator(currentFloor *int, directionElev Direction) {
	if directionElev == Up {
		*currentFloor++
	} else {
		*currentFloor--
	}
}

func addStop(floor int, resultStops *[]int) {
	if len(*resultStops) == 0 || (*resultStops)[len(*resultStops)-1] != floor {
		*resultStops = append(*resultStops, floor)
	}
}
