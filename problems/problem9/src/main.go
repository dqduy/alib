package main

import (
	"fmt"
	"sort"
)

func swap(a int, b int) (int, int) {
	return b, a
}

func testMultipleResultFunc() {
	first, second := swap(7, 3)
	fmt.Println(first, second)
}

func testArray() {
	countries := [5]string{
		"Australia",
		"Germany",
		"Brazil",
		"Egypt",
		"Japan"}
	fmt.Println(countries)
}

func testSlice() {
	countries := [5]string{
		"Australia",
		"Germany",
		"Brazil",
		"Egypt",
		"Japan"}
	fmt.Println("Countries list: ", countries)

	slice1 := countries[0:2]
	slice2 := countries[1:4]
	fmt.Printf("Slice1: len=%d, cap=%d, %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("Slice2: len=%d, cap=%d, %v\n", len(slice2), cap(slice2), slice2)

	slice2[0] = "Korea"
	fmt.Printf("Slice1: len=%d, cap=%d, %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("Slice2: len=%d, cap=%d, %v\n", len(slice2), cap(slice2), slice2)
}

func testMap() {
	capital_by_countries := map[string]string{
		"Australia": "Canberra",
		"Germany":   "Berlin",
		"Brazil":    "Brasília",
		"Egypt":     "Cairo",
		"Japan":     "Tokyo",
		"China":     "赔款"}

	fmt.Println(capital_by_countries)

	fmt.Println(capital_by_countries["China"])
}

func swapByPtr(x *int) {
	*x = 10
}

func testPtr() {
	a := 5
	fmt.Println(a)
	swapByPtr(&a)
	//fmt.Println(&a)
}

func testSort() {
	points := []int{4, 6, 1}
	sort.Ints(points)
	fmt.Println(points)

	for i, v := range points {
		fmt.Println(i, " - ", v)
	}
}

func main() {
	testSort()
}
