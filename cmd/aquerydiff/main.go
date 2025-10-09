package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	"github.com/stackb/bazel-aquery-differ/pkg/action"
	"github.com/stackb/bazel-aquery-differ/pkg/protobuf"
	"github.com/stackb/bazel-aquery-differ/pkg/report"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	var config config

	flags := flag.NewFlagSet("aquerydiff", flag.ExitOnError)
	flags.StringVar(&config.target, "target", "", "the target under analysis")
	flags.StringVar(&config.beforeFile, "before", "", "filepath to aquery file (before)")
	flags.StringVar(&config.afterFile, "after", "", "filepath to aquery file (after)")
	flags.StringVar(&config.matchingStrategy, "match", "output_files", "method used to build mapping of before & after actions (output_files|mnemonic)")
	flags.StringVar(&config.reportDir, "report_dir", "", "path to directory where report files should be written")
	flags.StringVar(&config.port, "port", "8000", "port number to use when serving content")
	flags.BoolVar(&config.unidiff, "unidiff", false, "compute unidiffs (can be slow)")
	flags.BoolVar(&config.cmpdiff, "cmpdiff", true, "compute go-cmp diffs (usually fast)")
	flags.BoolVar(&config.serve, "serve", false, "start webserver")
	flags.BoolVar(&config.open, "open", false, "open browser to webserver URL")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if config.beforeFile == "" {
		return fmt.Errorf("--before <filename> is required")
	}

	if config.afterFile == "" {
		return fmt.Errorf("--after <filename> is required")
	}

	if config.reportDir == "" {
		return fmt.Errorf("--report_dir <filename> is required")
	}

	var before anpb.ActionGraphContainer
	if err := protobuf.ReadFile(config.beforeFile, &before); err != nil {
		return err
	}
	log.Printf("Loaded %s (%d actions)", config.beforeFile, len(before.Actions))

	var after anpb.ActionGraphContainer
	if err := protobuf.ReadFile(config.afterFile, &after); err != nil {
		return err
	}
	log.Printf("Loaded %s (%d actions)", config.afterFile, len(after.Actions))

	beforeGraph, err := action.NewGraph("before", &before)
	if err != nil {
		return err
	}
	afterGraph, err := action.NewGraph("after", &after)
	if err != nil {
		return err
	}

	var mapper action.ActionMapper
	switch config.matchingStrategy {
	case "output_files":
		mapper = action.NewOutputFilesMap
	case "mnemonic":
		mapper = action.NewMnemonicFileMap
	default:
		return fmt.Errorf("unknown matching strategy '%s'", config.matchingStrategy)
	}

	beforeOnly, afterOnly, both := action.Partition(
		mapper(beforeGraph.Actions),
		mapper(afterGraph.Actions),
	)

	var equal action.OutputPairs
	var nonEqual action.OutputPairs

	log.Printf("Partitioning complete: (only before: %d, only after: %d, both: %d)", len(beforeOnly), len(afterOnly), len(both))

	for i, v := range both {
		log.Printf("Diffing %s (%d/%d)", v.Action.Mnemonic, i+1, len(both))

		if v.Diff() == "" {
			equal = append(equal, v)
		} else {
			nonEqual = append(nonEqual, v)
		}
	}

	r := report.Html{
		Target:     config.target,
		BeforeFile: config.beforeFile,
		AfterFile:  config.afterFile,
		Before:     beforeGraph,
		After:      afterGraph,
		BeforeOnly: beforeOnly,
		AfterOnly:  afterOnly,
		Equal:      equal,
		NonEqual:   nonEqual,
		Unidiff:    config.unidiff,
		Cmpdiff:    config.cmpdiff,
	}

	log.Printf("Generating report in: %s", config.reportDir)

	if err := r.Emit(config.reportDir); err != nil {
		return fmt.Errorf("generating report: %w", err)
	}

	log.Printf("aquerydiff report available at <%s>", config.reportDir)

	if config.serve {
		log.Printf("Starting webserver on port %s, serving %s", config.port, config.reportDir)
		http.Handle("/", http.FileServer(http.Dir(config.reportDir)))

		if config.open {
			url := "http://localhost:" + config.port
			log.Printf("Opening browser to %s", url)
			if err := openBrowser(url); err != nil {
				log.Printf("Failed to open browser: %v", err)
			}
		}

		return http.ListenAndServe(":"+config.port, nil)
	}

	return nil
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
