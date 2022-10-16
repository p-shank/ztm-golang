//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Vichlelr interface {
	reDirect()
}

type Motorcycles struct {
	name string
}

type Cars struct {
	name string
}
type Trucks struct {
	name string
}

func (m Motorcycles) reDirect() {
	fmt.Println("small lift, model:", m.name)
}

func (c Cars) reDirect() {
	fmt.Println("standard lift, model:", c.name)
}

func (t Trucks) reDirect() {
	fmt.Println("large lift, model:", t.name)
}

func selectLift(i Vichlelr) {
	i.reDirect()
}

func main() {
	motorcycle := Motorcycles{"honda"}
	car := Cars{"toyota"}
	truck := Trucks{"tata"}

	selectLift(motorcycle)
	selectLift(car)
	selectLift(truck)
}
