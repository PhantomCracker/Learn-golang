package main

import "fmt"

func main() {
	// part 1
	s := "Hello World";
	fmt.Println(s);
	fmt.Printf("%T\n", s);

	// part 2
	bs := []byte(s);
	fmt.Println(bs);
	fmt.Printf("%T\n", bs);

	// part 3
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i]);
	}

	// part 4
	fmt.Println("");
	for i, v := range s {
		fmt.Println(i, v);
	}
}
