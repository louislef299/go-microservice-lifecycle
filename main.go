package main

import (
	"flag"
	"log"

	v1 "example.com/pipelines/v1"
)

var (
	version string
)

func main() {
	flag.Parse()

	switch version {
	case "v1":
		v1.RunV1()
	default:
		log.Fatalf("%s is not a valid version option", version)
	}
}

func init() {
	flag.StringVar(&version, "v", "v1", "specify the version that the program should execute")
}
