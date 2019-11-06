package main

import "fmt"

func main() {
	as := struct{
		firstName string;
		lastName string;
	}{
		firstName: "Bill",
		lastName: "Gates",
	}
	fmt.Println(as);
}
