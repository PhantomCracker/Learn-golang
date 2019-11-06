package main

import "fmt"

type person struct {
	firstName string
	lastName string
	favoriteIcecream []string
}

func main() {
	p1 := person{
		firstName: "Jobs",
		lastName: "Steve",
		favoriteIcecream: []string{"Hazelnut"},
	}
	p2 := person{
		firstName:        "Gates",
		lastName:         "Bill",
		favoriteIcecream: []string{"Caramel"},
	}
	fmt.Println(p1);
	fmt.Println(p2);
}
