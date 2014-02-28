package main

import (
	"fmt"
	"github.com/mattn/omega"
	"strings"
	"time"
)

func main() {
	values := []interface{}{"333", 2, "foo", "boo", "zoo"}

	r := ω.A(values).Map(ω.S).
		Map(func(v ω.Value) ω.Value {
			return v.(string) + " " + time.Now().String()
		}).Filter(func(v ω.Value) bool {
			return strings.Index(v.(string), "o") != -1
		}).N()

	for _, v := range r.([]string) {
		fmt.Println(v)
	}
}
