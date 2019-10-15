package main

import "fmt"

func main() {
	nr1 := 3;
	nr2 := 5;
	if nr1 == nr2 {
		fmt.Println("Equal numbers");
	} else if nr1 <= nr2 {
		fmt.Println("Smaller or equal");
	} else if nr1 >= nr2 {
		fmt.Println("Greater or equal");
	} else if nr1 != nr2 {
		fmt.Println("Different numbers");
	} else if nr1 < nr2 {
		fmt.Println("Smaller");
	} else if nr1 > nr2 {
		fmt.Println("Bigger");
	}
}
