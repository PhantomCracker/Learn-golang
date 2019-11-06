package main

import "fmt"

type vehicle struct {
	doors int;
	color string;
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		luxury:  true,
	}
	fmt.Println("Truck: ", t);
	fmt.Println("Truck color: ", t.vehicle.color);
	fmt.Println("Sedan: ", s);
	fmt.Println("Sedan color: ", s.vehicle.color);
}
