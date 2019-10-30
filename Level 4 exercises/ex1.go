package main

import "fmt"

func main() {
	var arr[5] int;
	arr[0] = 0;
	arr[1] = 10;
	arr[2] = 20;
	arr[3] = 30;
	arr[4] = 40;
	for i, v := range arr {
		fmt.Println(i, v);
	}
	fmt.Printf("%T", arr);
}
