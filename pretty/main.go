package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/pretty"
)

const (
	jsonstr = `{"name":  {"first":"Tom","last":"Anderson"},  "age":37,
	"children": ["Sara","Alex","Jack"],
	"fav.movie": "Deer Hunter", "friends": [
		{"first": "Janet", "last": "Murphy", "age": 44}
	  ]}`
)

var (
	log = logrus.New()
)

func init() {
	log.Out = os.Stdout
}

func main() {

	colorres := pretty.Color([]byte(jsonstr), nil)
	log.Info(string(colorres))

	prettyres := pretty.Pretty([]byte(jsonstr))
	log.Info(string(prettyres))

	uglyres := pretty.Ugly([]byte(jsonstr))
	log.Info(string(uglyres))
}
