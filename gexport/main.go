// gexport --name='User' <data.sql
// gexport < data.sql
// gexport < data.json > data.go
// gexport --outfile < data.json
// gexport --outfile=data.file < data.json

package main

import (
	"github.com/kingzcheung/gexport/cmd"
	_ "github.com/kingzcheung/gexport/driver"
)

func main() {

	//f := bufio.NewReader(os.Stdin)
	//input, _ := f.ReadString('\n')
	//fmt.Println(input)

	//_examples.Example()
	cmd.Execute()
}
