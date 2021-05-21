// gexport --name='User' <data.sql
// gexport < data.sql
// gexport < data.json > data.go
// gexport --outfile=data.file < data.json

// connect sql
// gexport --mysql='' users > data.go
package main

import (
	"context"
	"github.com/alexedwards/scs/v2"
	"github.com/kingzcheung/gexport/internal"
	"log"
	"time"
)

func main() {

	session := scs.New()
	session.Lifetime = 24 * time.Hour

	serve := internal.NewServer(session)
	ctx := context.Background()
	log.Fatalln(serve.ListenAndServe(ctx, ":5210"))
}
