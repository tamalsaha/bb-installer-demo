package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	m := map[string]interface{}{
		"persons": []Person{
			{
				Name: "John",
				Age:  20,
			},
			{
				Name: "Jane",
				Age:  24,
			},
		},
	}
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
