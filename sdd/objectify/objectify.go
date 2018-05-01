package objectify

import (
	"reflect"
	"strconv"
	"strings"
)

type Object struct {
	Name  string
	Type  string
	Size  float64
	Color string
}

// Objectify takes an array of strings and returns them as objects of struct type.
func Objectify(input []string) ([]*Object, bool) {
	v := new(Object)
	s := reflect.ValueOf(v).Elem()
	if len(input) > 0 && len(input) < s.NumField() {
		return nil, false
	}
	res := []*Object{}
	for _, items := range input {
		r := new(Object)
		items := strings.Split(items, ",")
		for i, item := range items {
			item = strings.TrimSpace(item)
			if i == 2 {
				b, err := strconv.ParseFloat(string(item), 64)
				if err == nil {
					reflect.ValueOf(r).Elem().Field(i).SetFloat(b)
				}
			} else {
				reflect.ValueOf(r).Elem().Field(i).SetString(item)
			}
			//fmt.Println(r)
		}
		res = append(res, r)
	}

	return res, true
}
