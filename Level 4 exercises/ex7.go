package main

import "fmt"

func main() {
	mySlice := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Hello, James"},
	}
	for index, value := range mySlice {
		fmt.Println(index, value);
		for index2, value2 := range value {
			fmt.Println(index2, value2);
		}
	}
}
