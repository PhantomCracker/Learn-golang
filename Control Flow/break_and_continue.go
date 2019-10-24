package main

import "fmt"

func main() {
	x := 43 / 40;
	y := 43 % 40;
	fmt.Println(x, y);

	z := 0;
	for {
		z++;
		if z > 100 {
			break;
		}
		if z % 2 != 0 {
			continue
		}
		fmt.Println(z);
	}
}