package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){

    bytes,err := (os.ReadFile("../input.txt"))
    if err!=nil {
        fmt.Println("ERROR: could not read file because of: ",err)
        os.Exit(1)
    }
    left_list,right_list:= processing_data(bytes)
      sum:=0
     for i:=0;i< len(left_list); i++{
        c:=left_list[i]-right_list[i]
        if c>0{
            sum+=c
        }else{
            sum-=c
        }

    }
    fmt.Println("The total distance between the two lists is : ",sum)



}

func processing_data( bytes []byte) ([]int ,[]int){
    input:=strings.Split(strings.TrimSpace(string(bytes)),"\n")
    data:=strings.Fields(strings.Join(input," "))
    left_list  := []int{}
    right_list := []int{}
    for i:=0; i<len(data);i++{
        t:=data[i]
        x,e:=strconv.Atoi(t)
        if e!=nil {
            fmt.Println("ERROR: could not convert string to int because of: ",e)
            os.Exit(1)
        }
        if i%2==0{
            left_list= append(left_list,x)
        } else{
            right_list=append(right_list, x)
        }

    }
    sort.Ints(left_list)
    sort.Ints(right_list)
    return left_list,right_list

}



