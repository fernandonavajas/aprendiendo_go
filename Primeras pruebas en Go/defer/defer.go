package main

import "fmt"

func main() {

	myList := []string{"dog", "cat", "hedgehog"}

	// for {key}, {value} := range {list}
	for _, animal := range myList {
		fmt.Println("My animal is:", animal)
	}
}
