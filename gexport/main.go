// gexport --name='User' <data.sql
// gexport < data.sql
// gexport < data.json > data.go
// gexport --outfile=data.file < data.json

// connect sql
// gexport --mysql='' users > data.go
package main

import (
	"github.com/kingzcheung/gexport/cmd"
)

func main() {

	cmd.Execute()
}
