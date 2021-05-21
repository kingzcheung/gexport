// gexport --name='User' <data.sql
// gexport < data.sql
// gexport < data.json > data.go
// gexport --outfile=data.file < data.json

// connect sql
// gexport --mysql='' users > data.go
package main

import (
	"context"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/kingzcheung/gexport/internal"
	"log"
	"os/exec"
	"runtime"
	"time"
)

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func main() {

	session := scs.New()
	session.Lifetime = 24 * time.Hour

	serve := internal.NewServer(session)
	ctx := context.Background()
	go func() {
		err := Open("http://127.0.0.1:5210")
		if err != nil {
			return
		}
	}()
	log.Fatalln(serve.ListenAndServe(ctx, ":5210"))
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}
