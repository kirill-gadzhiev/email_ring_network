package main

import (
	"strings"
	"strconv"
	"fmt"
	"sort"
)

const (
	inf = "1000"
	n  = 7
 	k  = 4
 	generate  = "1011"
 	maxNumber = 128
 	)

var syndrom = map[string]int{
		"001": 0,
		"010": 1,
		"100": 2,
		"011": 3,
		"110": 4,
		"111": 5,
		"101": 6,
	}

var noError = [...]int{0,0,0}

type table struct {
	allErr int
	correctErr int
	ck float32
}

func getArray(data string) []int {
	length := len(data)
	intArr := make([]int, length)

	arr := strings.Split(data, "")
	for i, s := range arr {
		number, _ := strconv.Atoi(s)
		intArr[length - i - 1] = number
	}

	return intArr

}

func reverseArray(data []int) []int {
	length := len(data)
	newData := make([]int, length)
	for i, s := range data {
		newData[length - i - 1] = s
	}

	return newData
}

func getCodeStr(data []int) string {
	data = reverseArray(data)
	var str string
	for _, s := range data {
		str += strconv.Itoa(s)
	}

	return str
}


func getRest(data []int) []int {
	devider := getArray(generate)

	devider = reverseArray(devider)
	data = reverseArray(data)

	for len(data) >= len(devider) {
		for i, s := range data {
			if i < len(devider) {
				if s == devider[i] {
					data[i] = 0
				} else {
					data[i] = 1
				}
			}
		}
		if data[0] == 0 {
			data = append(data[:0], data[1:]...)
		}
	}


	return reverseArray(data)
}

func getErrDigit(data []int) int {
	code := getCodeStr(data)
	return syndrom[code]
}

func isErr(data []int) bool {
	for i, s := range data {
		if s != noError[i] {
			return true
		}
	}
	return false
}

func countDifferentDigits(data []int, err []int) int {
	if len(data) != len(err) {
		return -1
	}

	var count int
	for i, s := range data {
		if s != err[i] {
			count++
		}
	}

	return count
}

func encode(data []int) []int {
	newData := make([]int, n-k, n)
	newData = append(newData, data...)

	rest := getRest(newData)

	var res []int
	res = append(res, rest...)
	res = append(res, data...)

	return res
}

func decrypt(data []int) []int {
	err := getRest(data)

	if isErr(data) {
		errDigit := getErrDigit(err)
		if data[errDigit] == 0 {
			data[errDigit] = 1
		} else {
			data[errDigit] = 0
		}
	}

	return data
}

func printTable(data map[int]*table) {
	keys := make([]int, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	borderUppp := "┌───┬─────┬─────┬─────┐"
	borderMidd := "├───┼─────┼─────┼─────┤"
	borderDown := "└───┴─────┴─────┴─────┘"

	nameColumns := "│ i │  C  │  N  │  Ck │"


	fmt.Println(borderUppp)
	fmt.Println(nameColumns)

	for _, k := range keys {
		fmt.Println(borderMidd)
		fmt.Printf("│%3d│%5d│%5d│%5.1f│\n", k, data[k].allErr, data[k].correctErr, data[k].ck)
	}

	fmt.Println(borderDown)
}

func printInfo(data []int) {
	infError := make(map[int]*table)

	for i := 0; i < maxNumber; i++ {
		shortCode := getArray(fmt.Sprintf("%b", i))
		code := make([]int, n)
		copy(code[0:len(shortCode)], shortCode)

		difference := countDifferentDigits(data, code)
		if difference <= 0 {
			continue
		}

		if _, ok := infError[difference]; !ok {
			infError[difference] = &table{
				allErr: 0,
				correctErr: 0,
				ck: 0,
			}
		}

		infError[difference].allErr += 1

		if difference == 1 {
			newData := decrypt(code)

			if getCodeStr(data) == getCodeStr(newData) {
				infError[difference].correctErr += 1
			}

			infError[difference].ck = float32(infError[difference].correctErr / infError[difference].allErr)
		}




	}
	printTable(infError)

}


func test() {
	fmt.Println("Исходный вектор -", inf)

	data := getArray(inf)
	encrypted := encode(data)

	fmt.Println("Закодированный вектор -", getCodeStr(encrypted))

	printInfo(encrypted)
}

