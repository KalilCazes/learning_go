package main

import "fmt"

func main() {

	names := map[string][]string{
		"Hawk_Tony":       {"Skate", "Surf"},
		"Dylan_Bob":       {"Guitar", "Poetry"},
		"Mercury_Freddie": {"Sing", "Piano"},
	}

	names["Parker_Peter"] = []string{"Fight", "Save the world"}

	delete(names, "Dylan_Bob")
	for i, v := range names {
		fmt.Println(i)
		for item, h := range v {
			fmt.Printf("\tHobbie:%d -> %s\t\n", item, h)
		}
	}
}
