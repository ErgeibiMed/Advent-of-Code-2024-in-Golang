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
      simularity_score:=0
     for i:=0;i< len(left_list); i++{
        same:=0
        for j:=0; j< len(left_list);j++{
            if left_list[i]==right_list[j]{
                same+=1
            }

        }
        simularity_score=simularity_score+(same*left_list[i])

    }
    fmt.Println("The simularity score between the two list is : ",simularity_score)



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



