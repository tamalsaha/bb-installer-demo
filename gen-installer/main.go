package main

import (
	"fmt"

	api "go.bytebuilders.dev/installer/apis/installer/v1alpha1"
	"sigs.k8s.io/yaml"
)

func main() {
	var v api.AceSpec
	data, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
