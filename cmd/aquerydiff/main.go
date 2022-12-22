package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	"github.com/stackb/bazel-aquery-differ/pkg/protobuf"
)

func main() {
	log.Printf("Hello world")
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	var config config

	flags := flag.NewFlagSet("aquerydiff", flag.ExitOnError)
	flags.StringVar(&config.beforeFile, "before", "", "filepath to aquery file (before)")
	flags.StringVar(&config.afterFile, "after", "", "filepath to aquery file (after)")
	if err := flags.Parse(os.Args[1:]); err != nil {
		return err
	}

	if config.beforeFile == "" {
		return fmt.Errorf("--before <filename> is required")
	}

	if config.afterFile == "" {
		return fmt.Errorf("--after <filename> is required")
	}

	var before anpb.ActionGraphContainer
	var after anpb.ActionGraphContainer

	if err := protobuf.ReadFile(config.beforeFile, &before); err != nil {
		return err
	}

	if err := protobuf.ReadFile(config.afterFile, &after); err != nil {
		return err
	}

	log.Println("diffing %s <> %s", config.beforeFile, config.afterFile)

	return nil
}
