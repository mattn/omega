package main

import (
	"fmt"
	"github.com/mattn/M"
	"strings"
	"time"
)

func main() {
	values := []interface{}{"333", 2, "foo", "boo", "zoo"}

	r := M.A(values).Map(M.S).
		Map(func(v M.Value) M.Value {
			return v.(string) + " " + time.Now().String()
		}).Filter(func(v M.Value) bool {
			return strings.Index(v.(string), "o") != -1
		}).N([]string(nil))

	for _, v := range r.([]string) {
		fmt.Println(v)
	}
}
