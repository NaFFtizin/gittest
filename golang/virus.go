package main

import "fmt"

var g int = 0

func fc(i int) {
	for i >= -1 {
		i++
		fmt.Println(i)
	}
}	

func main() {
	for g >= -1 {
		go fc(1)
	}
}

