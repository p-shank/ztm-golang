//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	commands := make([]string, 0, 5)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("enter text: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) > 0 {
			if strings.ToLower(text) == "q" {
				fmt.Println("exiting the program")
				break
			}
			commands = append(commands, text)
			printMessage(text)
		} else {
			fmt.Println("enter a valid text")
		}
	}

	fmt.Println("no of commands: ", len(commands))
}

func printMessage(text string) {
	textMesage := map[string]string{
		"hi":  "hello how are you",
		"bye": "bye take care",
	}

	if val, ok := textMesage[text]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("please enter something else")
	}
}
