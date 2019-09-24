package main;

import (
	"fmt"
	"runtime"
)

var z int8 = -128;

func main() {
	// part 1
	x := 3;
	y := 5.123;
	fmt.Println(x);
	fmt.Println(y);
	fmt.Printf("%T\n", x);
	fmt.Printf("%T\n", y);

	// part 2
	fmt.Println(z);
	fmt.Printf("%T\n", z);

	// part 3
	fmt.Println(runtime.GOOS);
	fmt.Println(runtime.GOARCH);
}
