package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func main() {
    //bytes,err := (os.ReadFile("./example.txt"))
    bytes,err := (os.ReadFile("../input.txt"))
    if err!=nil {
        fmt.Println("ERROR: could not read file because of: ",err)
        os.Exit(1)
    }
    re:= regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

    matches:=re.FindAll(bytes,-1)
    sum:=0
    for _,v:=range matches {
        sum+=parse(v)
    }

    fmt.Printf("the result of adding up all the instructions is: %d\n",sum)

}



func parse (input []byte) int {
    first,_:=strings.CutPrefix(string(input),"mul(")
    second,_:=strings.CutSuffix(first,")")
    third:=strings.Split(second,",")
    t1,_:=strconv.Atoi(third[0])
    t2,_:=strconv.Atoi(third[1])

    return t1*t2



}
