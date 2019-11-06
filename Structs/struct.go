package main

import "fmt"

type person struct {
	first string;
	last string;
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
		sa1 := secretAgent{
			person: person{
				first: "James",
				last: "Bond",
			},
			ltk: false,
		}
		p2 := person {
			first: "Firstname",
			last: "Lastname",
		}
		fmt.Println(sa1);
		fmt.Println(sa1.person.first);
		fmt.Println(sa1.last);
		fmt.Println(sa1.ltk);
		fmt.Println(p2);
		fmt.Println(p2.first);
		fmt.Println(p2.last);
}