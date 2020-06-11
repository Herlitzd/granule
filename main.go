package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Herlitzd/granule/pkg"
)

var configPath = "granule.config.json"
var repoPath = flag.String("repo", ".", "path to git repo root")
var verbose = flag.Bool("verbose", false, "enable verbose logging")

func init() {
	flag.StringVar(repoPath, "r", ".", "path to git repo root")
	flag.BoolVar(verbose, "v", false, "enable verbose logging")
	flag.Parse()

	maybeConfigPath := flag.CommandLine.Arg(0)

	// If we got a loose arg passed to us, try and use it
	// as the config
	if len(maybeConfigPath) > 0 {
		configPath = maybeConfigPath
	}

}

func exit(err error) {
	log.SetOutput(os.Stderr)
	log.Fatal(err)
}

func main() {
	log.SetFlags(0)
	if !(*verbose) {
		log.SetOutput(ioutil.Discard)
	}

	config, err := pkg.ParseConfig(&configPath)
	if err != nil {
		exit(err)
	}

	context, err := pkg.MakeRepoContext(*repoPath)
	if err != nil {
		exit(err)
	}

	lastVersion, err := context.GetLastTag(*repoPath)
	if err != nil {
		exit(err)
	}

	branchName, err := context.GetCurrentBranch()
	if err != nil {
		exit(err)
	}

	branchConfig, err := config.GetBranchConfig(branchName)
	if err != nil {
		exit(err)
	}

	log.Printf("Last version found: %s\n", lastVersion)
	nextVersion := pkg.GetNextVersion(branchConfig, lastVersion)

	fmt.Println(nextVersion)
}
