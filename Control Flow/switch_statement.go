package main

import "fmt"

func main () {
	switch  {
	case false:
		fmt.Println("This should not print");
	case (3 == 3):
		fmt.Println("This will be printed");
		fallthrough;
	default:
		fmt.Println("Default");
	}
}
