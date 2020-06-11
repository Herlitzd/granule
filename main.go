package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Herlitzd/granule/pkg"
)

var configPath = "granule.config.json"
var repoPath = flag.String("repo", ".", "path to git repo root")

func init() {
	flag.StringVar(repoPath, "r", ".", "path to git repo root")
	flag.Parse()

	maybeConfigPath := flag.CommandLine.Arg(0)

	// If we got a loose arg passed to us, try and use it
	// as the config
	if len(maybeConfigPath) > 0 {
		configPath = maybeConfigPath
	}

}

func main() {
	log.SetOutput(ioutil.Discard)

	config, err := pkg.ParseConfig(&configPath)
	if err != nil {
		log.Fatal(err)
	}

	context, err := pkg.MakeRepoContext(*repoPath)
	if err != nil {
		log.Fatal(err)
	}

	lastVersion, err := context.GetLastTag(*repoPath)
	if err != nil {
		log.Fatal(err)
	}

	branchName, err := context.GetCurrentBranch()
	if err != nil {
		log.Fatal(err)
	}

	branchConfig, err := config.GetBranchConfig(branchName)
	if err != nil {
		log.Fatal(err)
	}

	nextVersion := pkg.GetNextVersion(branchConfig, lastVersion)

	log.Printf("\n%s", lastVersion)
	fmt.Println(nextVersion)
}
