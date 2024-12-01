package main

import (
	"reflect"
	"testing"
)

func TestDistanceCalculator(t *testing.T) {
	t.Run("Using test data", func(t *testing.T) {

		list1 := []int{1, 2, 3, 3, 3, 4}
		list2 := []int{3, 3, 3, 4, 5, 9}

		expect := 11
		actual, err := CalculateDistance(list1, list2)

		if err != nil {
			t.Fatal(err.Error())
		}

		if expect != actual {
			t.Errorf("Test failure: expected %d but got %d", expect, actual)
		}
	})
}

func TestCalculateSimilarity(t *testing.T) {
	t.Run("Using exampleData", func(t *testing.T) {
		list1 := []int{1, 2, 3, 3, 3, 4}
		list2 := []int{3, 3, 3, 4, 5, 9}

		expect := 31
		actual, err := CalculateSimilarity(list1, list2)

		if err != nil {
			t.Fatal(err.Error())
		}

		if expect != actual {
			t.Errorf("Test failure: expected %d but got %d", expect, actual)
		}
	})
}

func TestFormatData(t *testing.T) {
	t.Run("Using example data", func(t *testing.T) {

		wanted1 := []int{1, 2, 3, 3, 3, 4}
		wanted2 := []int{3, 3, 3, 4, 5, 9}

		got1, got2 := FormatData("./exampleData.txt")
		if !reflect.DeepEqual(got1, wanted1) {
			t.Errorf("Test failure: expected %v but got %v", wanted1, got1)
		}

		if !reflect.DeepEqual(got2, wanted2) {
			t.Errorf("Test failure: expected %v but got %v", wanted2, got2)
		}
	})
}
