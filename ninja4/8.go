package main

import "fmt"

func main() {

	names := map[string][]string{
		"Hawk_Tony":       {"Skate", "Surf"},
		"Dylan_Bob":       {"Guitar", "Poetry"},
		"Mercury_Freddie": {"Sing", "Piano"},
	}

	for i, v := range names {
		fmt.Println(i)
		for item, h := range v {
			fmt.Printf("\tHobbie:%d -> %s\t\n", item, h)
		}
	}
}
