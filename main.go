package main

import (
	"flag"
	"log"
	"os"

	"github.com/willfaught/drudge/drudge"
)

var w drudge.Worker

func main() {
	log.SetFlags(0)
	log.SetPrefix(os.Args[0] + ": ")

	var dry = flag.Bool("d", false, "Dry run mode. Does not make changes.")
	var quiet = flag.Bool("q", false, "Quiet mode. Prints only the log errors.")
	var verbose = flag.Bool("v", false, "Verbose mode. Prints the log.")

	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalln("error: no verb")
	}

	w.Dry = *dry
	w.Quiet = *quiet
	w.Verbose = *verbose

	var args = flag.Args()

	if err := w.Do(args[0], args[1:]...); err != nil {
		log.Fatalln("error:", err)
	}
}
