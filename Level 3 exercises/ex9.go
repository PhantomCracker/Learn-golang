package main

import "fmt"

func main() {
	var sport string = "favSport";
	switch sport {
	case "favSport":
		fmt.Println("Football");
	case "basket":
		fmt.Println("Basketball");
	default:
		fmt.Println("I do not like sports at all");
	}
}
