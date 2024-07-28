package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Sonny")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
	fmt.Println(data["name"])
}
