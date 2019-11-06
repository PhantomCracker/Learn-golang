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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range m {
		fmt.Println("Key: ", k);
		fmt.Println("Value: ", v.firstName);
		for i, val := range v.favoriteIcecream {
			fmt.Println(i, val);
		}
	}

	fmt.Println("Map: ", m);

	fmt.Println(p1);
	fmt.Println(p2);
}
