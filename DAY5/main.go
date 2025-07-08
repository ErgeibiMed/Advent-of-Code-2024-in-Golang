package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes_raw, err := (os.ReadFile("./example.txt"))
	//bytes_raw, err := (os.ReadFile("./input.txt"))
	if err != nil {
		fmt.Println("ERROR: could not read file because of: ", err)
		os.Exit(1)

	}
	pg_ord_rules, pg_num_upd, _ := strings.Cut(string(bytes_raw), "\n\n")
	pg_num_updt, _ := strings.CutSuffix(pg_num_upd, "\n")
	sum, unorderder := valid_update(pg_num_updt, orderng_rules(pg_ord_rules))
	part_1(sum)
	fmt.Println(orderng_rules(pg_ord_rules))
	fmt.Println(unorderder)
}

func part_1(s int) {
	fmt.Println("we get :", s)
}

func orderng_rules(in string) map[int][]int {
	result := make(map[int][]int)
	s := strings.Split(in, "\n")
	for _, line := range s {
		index := strings.Index(line, "|")
		left, _ := strconv.ParseInt(line[:index], 10, 64)
		tmp := new([]int)
		for _, value := range s {
			index2 := strings.Index(value, "|")
			right, _ := strconv.ParseInt(value[index2+1:], 10, 64)
			if left != right {
				*tmp = append(*tmp, int(right))
			}
		}

		result[int(left)] = *tmp
	}
	r := make(map[int][]int)
	for k := range result {
		tmp := new([]int)
		for _, line := range s {
			index := strings.Index(line, "|")
			left, _ := strconv.ParseInt(line[:index], 10, 64)
			if int(left) == k {
				index2 := strings.Index(line, "|")
				right, _ := strconv.ParseInt(line[index2+1:], 10, 64)
				*tmp = append(*tmp, int(right))

			}

		}
		r[k] = *tmp
	}
	return r
}
func valid_update(in string, valid_ord_map map[int][]int) (int, string) {
	sum := 0
	unordered := ""
	inp := strings.Split(in, "\n")
	for _, line := range inp {
		num_pages := strings.Split(line, ",")
		contains := true
		for i := range num_pages {
			num_i, _ := strconv.ParseInt(num_pages[i], 10, 64)
			table, _ := valid_ord_map[int(num_i)]
			for j := i + 1; j < len(num_pages); j++ {
				num_j, _ := strconv.ParseInt(num_pages[j], 10, 64)
				contains = slices.Contains(table, int(num_j))
				if contains == false {
					break
				}
			}
			if contains == false {
				break
			}
		}
		if contains {
			//fmt.Println(line)
			w := (len(line) + 1) / 2
			x, _ := strconv.ParseInt(line[w-1:w+1], 10, 64)
			//fmt.Println("x:", x)
			sum += int(x)
		} else {
			unordered = unordered + line + "\n"
			//fmt.Println("unordered ones: ", unordered)
		}
	}
	unordered, _ = strings.CutSuffix(unordered, "\n")
	return sum, unordered

}
