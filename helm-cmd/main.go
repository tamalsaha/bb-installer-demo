package main

import (
	"fmt"
	pkglib "kubepack.dev/kubepack/pkg/lib"
	"kubepack.dev/lib-helm/pkg/values"
)

func main() {
	if err := gen(); err != nil {
		panic(err)
	}
}

func gen() error {
	chrt, err := pkglib.DefaultRegistry.GetChart("", "opscenter-config", "")
	if err != nil {
		return err
	}
	cmd, err := values.GetChangedValues(chrt.Values, nil)
	if err != nil {
		return err
	}
	fmt.Println(cmd)
	return nil
}
