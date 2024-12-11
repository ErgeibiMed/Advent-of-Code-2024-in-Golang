package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//bytes, err := (os.ReadFile("./example.txt"))
	bytes,err := (os.ReadFile("../input.txt"))
	if err != nil {
		fmt.Println("ERROR: could not read file because of: ", err)
		os.Exit(1)
	}
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)


    input:=handle_ins(bytes)
	matches := re.FindAllString(input, -1)
    sum:=0
    for _,v:=range matches {
        sum+=parse(v)
    }
    fmt.Println("the sum of enabled instructions is :",sum)
}

func parse(input string) int {
	first, _ := strings.CutPrefix(input, "mul(")
	second, _ := strings.CutSuffix(first, ")")
	third := strings.Split(second, ",")
	t1, _ := strconv.Atoi(third[0])
	t2, _ := strconv.Atoi(third[1])
	return t1 * t2
}

func handle_ins(in []byte) string {
    input:=string(in)
    result:=""
    index_of_dont := strings.Index(input, "don't()")
    result+=input[0:index_of_dont]
    temp:=input[index_of_dont:]
    index_of_do := strings.Index(temp, "do()")
    flag:=true
    for flag{
        temp=temp[index_of_do:]
        index_of_dont = strings.Index(temp, "don't()")
        if index_of_dont==-1{
            result+=temp
            break
        }
        result+=temp[:index_of_dont]
        temp=temp[index_of_dont:]
        index_of_do = strings.Index(temp, "do()")
        if index_of_do==-1 {
            break
        }

    }
    return result
}
