package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/blang/semver/v4"
)

var (
	level string = "patch"
)

func increment(version, level string) (semver.Version, error) {
	v, err := semver.Make(version)
	if err != nil {
		return v, fmt.Errorf("unable to parse input: %s", err)
	}

	switch level {
	case "patch":
		v.IncrementPatch()
	case "minor":
		v.IncrementMinor()
	case "major":
		v.IncrementMajor()
	default:
		return v, fmt.Errorf("invalid level argument")
	}

	return v, nil
}

func main() {
	flag.StringVar(&level, "i", level, "level to increment the provided version")
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("unexpected number of args")
	}
	version := flag.Arg(0)

	out, err := increment(version, level)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out)
}
