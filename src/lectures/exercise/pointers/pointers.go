//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Item struct {
	itemName string
	status   bool
}

func createItems() []Item {
	items := []Item{
		{"TV", true},
		{"Washing Machine", true},
		{"PS", true},
		{"Fan", true},
	}

	return items
}

func checkout(items []Item) {

	for i := 0; i < len(items); i++ {
		items[i].status = false
		fmt.Println(items)
	}

}

func main() {
	items := createItems()
	items[0].status = false
	fmt.Println(items)
	checkout(items)
}
