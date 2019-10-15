package main;

import "fmt"

func main() {
	myVariable := 7;
	fmt.Printf("%d %b %#x\n", myVariable, myVariable, myVariable);
	shiftedBits := myVariable << 1;
	fmt.Printf("%d %b %#x", shiftedBits, shiftedBits, shiftedBits);
}
