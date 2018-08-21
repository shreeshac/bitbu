package main

import (
	"fmt"
)

type st struct {
	Name string
}

func main() {
	arr := make([]interface{}, 0)
	fmt.Println("Shreesha")
	s := st{}
	s.Name = "shreesha"
	fmt.Println(s.Name)
	arr = append(arr, &s.Name)
	fmt.Println(1, arr[0])
	v1 := arr[0].(*string)
	fmt.Println(2, arr[0], v1, *v1)
	sneha := "sneha"
	s.Name = sneha
	fmt.Println(2, arr[0], v1, *v1)
}
