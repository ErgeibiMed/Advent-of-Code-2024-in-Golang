package main

import (
	"fmt"
	"os"
)

func main() {
	bytes_raw, err := (os.ReadFile("./example.txt"))
	if err != nil {
		fmt.Println("ERROR: could not read file because of: ", err)
		os.Exit(1)
	}

	s := string(bytes_raw)
	for i := range s {

		fmt.Printf("% s", string(s[i]))
	}
}
