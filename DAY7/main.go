package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	//bytes, err := (os.ReadFile("exemple.txt"))
	bytes, err := (os.ReadFile("input.txt"))
	if err != nil {
		fmt.Println("ERROR: could not read file because of: ", err)
		os.Exit(1)
	}

	data := strings.Split(string(bytes), "\n")
	totCalibrationresult := int64(0)
	for _, v := range data {
		testData := trueEq(v)
		if testData != -1 {
			totCalibrationresult += testData
		}
	}
	fmt.Println("the total calibration result is", totCalibrationresult)

}

func power(n int, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func possibleCombination(line string) map[string]string {
	left, rightWithTrailingSpace, _ := strings.Cut(line, ":")
	right := strings.TrimSpace(rightWithTrailingSpace)
	n := strings.Count(right, " ")
	numberOfcomb := power(2, n)
	AddCom := strings.ReplaceAll(right, " ", "+")
	MultCom := strings.ReplaceAll(right, " ", "*")
	Acomb := new([]string)
	Mcomb := new([]string)
	for _, v := range splitStringtoSliceOFString(AddCom) {
		*Acomb = append(*Acomb, v)
	}
	for _, v := range splitStringtoSliceOFString(MultCom) {
		*Mcomb = append(*Mcomb, v)
	}
	possComb := map[string]string{}
	possComb[strings.Join(*Acomb, "")] = left
	possComb[strings.Join(*Mcomb, "")] = left
	i := 0
	j := 0
	for range numberOfcomb {
		if i < len(*Acomb) {
			comb1 := strings.Replace(AddCom, "+", "*", i)
			possComb[comb1] = left
			n := slices.Index((*Acomb)[i:len(*Acomb)], "+")
			if n != -1 {
				possComb[strings.Join(*Acomb, "")] = left
				(*Acomb)[n+i] = "*"
				possComb[strings.Join(*Acomb, "")] = left
				(*Acomb)[n+i] = "+"
			}
			i++
		}
		if j < len(*Mcomb) {
			comb2 := strings.Replace(MultCom, "+", "*", i)
			possComb[comb2] = left

			m := slices.Index((*Mcomb)[j:len(*Mcomb)], "*")
			if m != -1 {
				possComb[strings.Join(*Mcomb, "")] = left
				(*Mcomb)[m+j] = "+"
				possComb[strings.Join(*Mcomb, "")] = left
				(*Mcomb)[m+j] = "*"
			}
			j++
		}

	}
	return possComb
}

func splitStringtoSliceOFString(s string) []string {
	temp := ""
	for _, v := range s {
		if v == ' ' {
			continue
		}
		if v == '+' || v == '*' {
			temp += " "
			temp += string(v)
			temp += " "

		} else {
			temp += string(v)
		}
	}
	return strings.Split(temp, " ")
}

func trueEq(line string) int64 {
	possComb := possibleCombination(line)
	for k, v := range possComb {
		kToString := splitStringtoSliceOFString(k)
		test, _ := strconv.ParseInt(kToString[0], 10, 0)
		for i := 1; i < len(kToString)-1; i++ {
			if kToString[i] == string('+') {
				a, _ := strconv.ParseInt(kToString[i+1], 10, 0)
				test += a
			}
			if kToString[i] == string('*') {
				m, _ := strconv.ParseInt(kToString[i+1], 10, 0)
				test *= m
			}
		}
		val, _ := strconv.ParseInt(v, 10, 0)
		if test == val {
			return test
		}
	}
	return -1
}
