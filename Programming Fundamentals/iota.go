package main

import "fmt"

const (
	a = iota;
	b // this is also iota
	c // this is also iota
)

func main() {
	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
	fmt.Printf("%T\n", a);
	fmt.Printf("%T\n", b);
	fmt.Printf("%T\n", c);
}
