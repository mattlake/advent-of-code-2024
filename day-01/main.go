package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	l1, l2 := FormatData("./testData1.txt")
	distance, _ := CalculateDistance(l1, l2)
	similarity, _ := CalculateSimilarity(l1, l2)
	fmt.Printf("Distance: %d,\n", distance)
	fmt.Printf("Similarity: %d, \n", similarity)
}

func CalculateDistance(list1, list2 []int) (int, error) {
	if len(list1) != len(list2) {
		return 0, errors.New("Mismatched array lengths provided")
	}

	result := 0

	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			result += (list1[i] - list2[i])
		} else {
			result += (list2[i] - list1[i])
		}

	}

	return result, nil
}

func FormatData(filePath string) (list1 []int, list2 []int) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), "   ")
		int1, _ := strconv.Atoi(temp[0])
		int2, _ := strconv.Atoi(temp[1])

		list1 = append(list1, int1)
		list2 = append(list2, int2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	return
}

func CalculateSimilarity(list1, list2 []int) (int, error) {

	if len(list1) != len(list2) {
		return 0, errors.New("Mismatched array lengths provided")
	}

	result := 0

	for _, key := range list1 {
		occurances := 0
		for _, value := range list2 {
			if value == key {
				occurances += 1
			}
		}

		result += (occurances * key)
	}

	return result, nil
}
