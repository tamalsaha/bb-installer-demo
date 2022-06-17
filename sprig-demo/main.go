package main

import "fmt"

type Values map[string]interface{}

func main() {
	var v2 interface{} = new(Values)
	v, ok := v2.(map[string]interface{})
	fmt.Println(v, ok)

	// dig("global", "missing", Values{})
}

func dig(ps ...interface{}) (interface{}, error) {
	if len(ps) < 3 {
		panic("dig needs at least three arguments")
	}
	dict := ps[len(ps)-1].(map[string]interface{})
	def := ps[len(ps)-2]
	ks := make([]string, len(ps)-2)
	for i := 0; i < len(ks); i++ {
		ks[i] = ps[i].(string)
	}

	return digFromDict(dict, def, ks)
}

func digFromDict(dict map[string]interface{}, d interface{}, ks []string) (interface{}, error) {
	k, ns := ks[0], ks[1:]
	step, has := dict[k]
	if !has || step == nil {
		return d, nil
	}
	if len(ns) == 0 {
		return step, nil
	}
	return digFromDict(step.(map[string]interface{}), d, ns)
}
