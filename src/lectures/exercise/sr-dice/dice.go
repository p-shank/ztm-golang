//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice(side int) int {
	return rand.Intn(side)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	noOfDice := getInput("No Of Dice: ")
	noOfSides := getInput("No Of Sides: ")
	noOfRolls := getInput("No Of Rolls: ")

	fmt.Println("No Of Dice:", noOfDice)
	fmt.Println("No Of Sides:", noOfSides)
	fmt.Println("No Of Rolls:", noOfRolls)

	for r := 1; r <= noOfRolls; r++ {
		sum := 0
		for d := 1; d <= noOfDice; d++ {
			rolled := rollDice(noOfSides)
			sum += rolled
		}

		fmt.Println("sum #", sum)

		switch sum := sum; {
		case sum == 2 && noOfDice == 2:
			fmt.Println("Sanke eyes")
		case sum == 7:
			fmt.Println("Lucky 7")
		case sum%2 == 0:
			fmt.Println("Even")
		case sum%2 == 1:
			fmt.Println("Odd")
		}
	}

}

func getInput(msg string) int {
	var input int
	fmt.Print(msg)
	fmt.Scanf("%d", &input)
	return input
}
