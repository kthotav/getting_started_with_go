package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string: ")

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	text = strings.Replace(text, "\n", "", -1)
	lowerCasedStr := strings.ToLower(text)

	startsWithI := strings.Index(lowerCasedStr, "i") == 0
	containsI := strings.Contains(lowerCasedStr, "a")
	endsWithN := strings.LastIndex(lowerCasedStr, "n") == len(lowerCasedStr)-1

	if startsWithI && containsI && endsWithN {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}

}
