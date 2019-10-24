package main

import "fmt"

func main() {
	if true {
		fmt.Println("Hello, true");
	}
	if false {
		fmt.Println("Hello, false");
	}
	if !true {
		fmt.Println("Hello, !true");
	}
	if !false {
		fmt.Println("Hello, !false");
	}
	if (2 == 2) {
		fmt.Println("2 == 2");
	}
	if (2 != 2) {
		fmt.Println("2 != 2");
	}
	if x := 42; (x == 42) {
		fmt.Println(x);
	}
	 x := 40;
	 if (x == 40) {
	 	fmt.Println("X is 40");
	 } else if (x == 41) {
	 	fmt.Println("X is 41");
	 } else {
	 	fmt.Println("What is X?");
	 }
}
