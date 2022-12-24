package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	"github.com/stackb/bazel-aquery-differ/pkg/action"
	"github.com/stackb/bazel-aquery-differ/pkg/protobuf"
)

func main() {
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
	if err := protobuf.ReadFile(config.beforeFile, &before); err != nil {
		return err
	}

	var after anpb.ActionGraphContainer
	if err := protobuf.ReadFile(config.afterFile, &after); err != nil {
		return err
	}

	log.Println("diffing %s <> %s", config.beforeFile, config.afterFile)

	beforeGraph, err := action.NewGraph(&before)
	if err != nil {
		return err
	}
	afterGraph, err := action.NewGraph(&after)
	if err != nil {
		return err
	}

	beforeOnly, afterOnly, both := action.Partition(beforeGraph.OutputMap, afterGraph.OutputMap)

	for _, v := range beforeOnly {
		log.Println("only in --before:", v.Output)
	}
	for _, v := range afterOnly {
		log.Println("only in --after:", v.Output)
	}
	for _, v := range both {
		if v.Diff() == "" {
			log.Printf("in both, no change: %s\n%s", v.Output)
		} else {
			log.Printf("in both, changed: %s\n%s", v.Output, v.UnifiedDiff(config.beforeFile, config.afterFile))
		}
	}

	return nil
}
