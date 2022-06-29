package main

import (
	"fmt"
	pkglib "kubepack.dev/kubepack/pkg/lib"
	chartlib "kubepack.dev/lib-helm/pkg/chart"
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
	cmd, err := chartlib.GetChangedValues(chrt.Values, nil)
	if err != nil {
		return err
	}
	fmt.Println(cmd)
	return nil
}
