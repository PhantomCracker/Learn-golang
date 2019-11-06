package main

import "fmt"

func main() {
	p1 := struct {
		first string;
		last string;
	}{
		first: "Firstname",
		last: "Lastname",
	}
	fmt.Println(p1);
}
