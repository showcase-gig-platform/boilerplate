package main

import (
	"os"

	"github.com/showcase-gig-platform/boilerplate/cmd"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("boilerplate", "Generate code from boilerplate")
	cmd.RegisterGen(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
