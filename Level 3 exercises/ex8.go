package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("Case true");
	case false:
		fmt.Println("Case false");
	default:
		fmt.Println("Default");
	}
}
