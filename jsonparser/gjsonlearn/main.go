package main

import (
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
)

const jsonstr = `{
	"name": {"first": "Tom", "last": "Anderson"},
	"age":37,
	"children": ["Sara","Alex","Jack"],
	"fav.movie": "Deer Hunter",
	"friends": [
	  {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
	  {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
	  {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	]
  }`

func main() {
	fmt.Printf("First name : %s\n", gjson.Get(jsonstr, "name.first").String())
	fmt.Printf("age : %d\n", gjson.Get(jsonstr, "age").Int())
	fmt.Printf("fav.movie : %s\n", gjson.Get(jsonstr, "fav\\.movie").String())

	input := []byte(jsonstr)
	reduceresult := pretty.Pretty(input)

	fmt.Println(string(reduceresult))

	listresult := gjson.Get(jsonstr, "friends.#.first")

	for key, part := range listresult.Array() {
		fmt.Printf("key : %d, part : %s\n", key, part)
	}
}
