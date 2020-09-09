package main

import (
	"fmt"
	"github.com/kingzcheung/gexport"
)

func json_example() {
	json := `{
    "bar": "foo",
    "group": {
        "username": "zfq",
        "age": 12
    }
}`
	gx := gexport.New(json, gexport.JSON)
	fmt.Println(gx.Parse().Output())
}
