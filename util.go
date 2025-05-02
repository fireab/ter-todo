package main

import "fmt"

func formattedString(tasks []string) string {
	var res string
	for i, v := range tasks {
		res += fmt.Sprintf("%d %s\n", i+1, v)
	}
	return res
}
