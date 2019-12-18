package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	infoMap := make(map[string]string)

	fmt.Print("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	name = strings.TrimSuffix(name, "\n")
	infoMap["name"] = name

	fmt.Print("Enter your address: ")
	address, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Erro: %s", err)
	}

	address = strings.TrimSuffix(address, "\n")
	infoMap["address"] = address

	json, err := json.Marshal(infoMap)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Println(string(json))
}
