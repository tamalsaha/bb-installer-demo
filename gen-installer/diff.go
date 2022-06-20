package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/alessio/shellescape"
)

func GetValuesDiff(original, modified map[string]interface{}) (map[string]interface{}, error) {
	return getValuesDiff(original, modified, "", nil)
}

func getValuesDiff(original, modified map[string]interface{}, prefix string, diff map[string]interface{}) (map[string]interface{}, error) {
	if diff == nil {
		diff = map[string]interface{}{}
	}

	for k, v := range modified {
		curKey := ""
		if prefix == "" {
			curKey = escapeKey(k)
		} else {
			curKey = prefix + "." + escapeKey(k)
		}

		switch val := v.(type) {
		case map[string]interface{}:
			oVal, ok := original[k].(map[string]interface{})
			if !ok {
				oVal = map[string]interface{}{}
			}

			d2, err := getValuesDiff(oVal, val, curKey, nil)
			if err != nil {
				return nil, err
			}
			if len(d2) > 0 {
				diff[k] = d2
			}
		case []interface{}, string, int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64, bool, json.Number, nil:
			if !reflect.DeepEqual(original[k], val) {
				diff[k] = val
			}
		default:
			return nil, fmt.Errorf("unknown type %v with value %v", reflect.TypeOf(v), v)
		}
	}

	// https://github.com/kubepack/lib-helm/blob/32de2acacbfb84f57d4a66c6d896360eb664399c/pkg/values/options.go#L133
	for k, v := range original {
		if _, found := modified[k]; !found {
			curKey := ""
			if prefix == "" {
				curKey = escapeKey(k)
			} else {
				curKey = prefix + "." + escapeKey(k)
			}

			// TODO: how does Helm merge --values remove keys?
			// diff[k] = nil
			return nil, fmt.Errorf("key %s is missing in the modified values, original values %v", curKey, v)
		}
	}
	return diff, nil
}

// kubernetes.io/role becomes "kubernetes\.io/role"
func escapeKey(s string) string {
	return shellescape.Quote(strings.ReplaceAll(strings.ReplaceAll(s, `\`, `\\`), `.`, `\.`))
}
