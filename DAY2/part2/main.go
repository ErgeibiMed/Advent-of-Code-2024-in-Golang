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
    reports:=bytes_to_slices(bytes)
    safe:=0
    for _,v:=range reports{
        if safe_report(v) {
            safe+=1
        } else {
            for i:=0;i< len(v) ;i++{
                t:=remov_elm(v,i)
                if safe_report(t){
                    safe+=1
                    break
                }
            }
        }
    }

    fmt.Printf("%d reports are safe\n",safe)

}

func bytes_to_slices(input []byte) [][]int {
    reports:=strings.Split(strings.TrimSpace(string(input)),"\n")
    output:=make([][]int,len(reports))
    for i,v:=range reports{
        data:=strings.Split(v, " ")
        for _,elm:=range data {
            x,err:=strconv.Atoi(elm)
            if err!=nil {
                fmt.Println(err)
            }
            output[i] = append(output[i], x)
        }
    }

    return output

}

func safe_report(v []int) bool{
    if is_all_incr_decr(v)==true && differ_by(v)==true {
        return true
    }
    return false

}

func is_all_incr_decr(in []int)bool{
    var flag bool= in[0]>in[1]// false if ascending order
    for i:=0;i< len(in)-1 ;i++{
        if in[i+1]==in[i] {
            return false
        }else if in[i+1]>in[i] && flag==true {
            return false
        }else if in[i+1]<in[i] && flag==false {
            return false

        }

    }
    return true

}
func differ_by(in []int)bool{
    temp:=in
    sort.Ints(temp)
    for i:=0;i< len(in)-1 ;i++{
        d:=in[i+1]- in[i]
        if d>3 || d<1 {
            return false
        }

    }
    return true
}


func remov_elm(in []int,x int) []int {
    n:=len(in)
    temp:= []int{}
    for i:=0;i<n ;i++{
        if i==x { continue
        }else {
            temp = append(temp, in[i])
        }

    }
    return temp
}
