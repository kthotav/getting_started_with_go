package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const sliceSize = 3

func main() {

	nums := make([]int, sliceSize)
	initialIndex := 1

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Slice shell")
	fmt.Println("--------")

	for {
		fmt.Print("Enter an integer or 'x' to exit -> ")

		text, err := reader.ReadString('\n')
		if err != nil {
			continue
		}

		text = strings.TrimSuffix(text, "\n")
		if text == "x" {
			break
		}

		num, err := strconv.Atoi(text)
		if err != nil {
			continue
		}

		if initialIndex <= sliceSize {
			nums[len(nums)-initialIndex] = num
			initialIndex++
		} else {
			nums = append(nums, num)
		}

		sort.Ints(nums)
		fmt.Println(nums)
	}
}

func logUnknownInput() {
	log.Print("Unknown input. Type integer of 'x' to exit")
}
