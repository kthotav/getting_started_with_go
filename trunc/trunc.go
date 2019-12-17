package main

import (
	"fmt"
	"log"
)

func main(){
	var floatingNum float64

	fmt.Printf("Enter a floating point number: ")

	_, err := fmt.Scanln(&floatingNum)
	if err != nil {
		log.Fatalf("Can't read floating point number, reason %s", err)
	}
}
