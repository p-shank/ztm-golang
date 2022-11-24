//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}
	totalSum := 0

	channel := make(chan int, 5)

	for i := 0; i < len(files); i++ {
		go readFile(files[i], channel)
	}

	for {
		select {
		case sum := <-channel:
			totalSum += sum
		case <-time.After(100 * time.Millisecond):
			fmt.Println(totalSum)
			return
		}
	}
}

func readFile(fileName string, channel chan<- int) {

	f, err := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	totalSum := 0

	if err != nil {
		fmt.Println(fileName + " not found")
	}

	defer f.Close()

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err == nil {
			totalSum += number
		}
	}
	channel <- totalSum
}
