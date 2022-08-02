package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo")

	ii := []int {1,2,3,4}
	fooVar := foo(ii...)
	barVarInt := bar(ii)


	fmt.Println("FOO: ", fooVar)
	fmt.Println("BAR: ", barVarInt)

}

func foo(a ...int) int {
	total := 0
	for _,v := range a {
		total+= v
	}
	return total
}

func bar(b []int) (int) {
	total := 0
	for _,v := range b {
		total+= v
	}
	return total
}