//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func calculateArea(rect Rectangle) {
	area := rect.length * rect.breadth
	perimeter := 2 * (rect.length + rect.breadth)
	fmt.Println("Area :", area)
	fmt.Println("Perimeter :", perimeter)
}

func main() {
	rect := Rectangle{4, 5}
	calculateArea(rect)
	rect.length = rect.length * 2
	rect.breadth = rect.breadth * 2
	calculateArea(rect)
}
