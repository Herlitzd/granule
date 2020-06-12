package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
)

// ConfigDef Top level config object
type ConfigDef struct {
	Branches []BranchConfig `json:"branches"`
}

// BranchConfig Rule for branches
type BranchConfig struct {
	Branch        string    `json:"branch"`
	BranchIsRegex bool      `json:"branchIsRegex"`
	Bump          BumpLevel `json:"bump"`
}

// BumpLevel Enum for Version Levels
type BumpLevel string

// Version Levels
const (
	Major BumpLevel = "major"
	Minor BumpLevel = "minor"
	Patch BumpLevel = "patch"
)

func (branchConfig *BranchConfig) matchBranch(branchName *string) (*bool, error) {
	if branchConfig.BranchIsRegex {
		re, err := regexp.Compile(branchConfig.Branch)

		if err != nil {
			return nil, err
		}

		matched := re.MatchString(*branchName)
		return &matched, nil
	}

	matched := branchConfig.Branch == *branchName
	return &matched, nil
}

// GetBranchConfig Find the branch config for a given branch
func (c ConfigDef) GetBranchConfig(branchName *string) (*BranchConfig, error) {
	var matchedBranch *BranchConfig
	log.Printf("active branch: %s", *branchName)
	for i, branch := range c.Branches {

		matched, err := branch.matchBranch(branchName)

		if err != nil {
			return nil, err
		}

		log.Printf("trying pattern: %v, matched: %t", branch.Branch, *matched)

		if *matched && matchedBranch != nil {
			return nil, fmt.Errorf(
				"multiple matching branch configurations found: %s and %s",
				matchedBranch.Branch, branch.Branch)
		}

		if *matched {
			log.Print("setting matched branch")
			matchedBranch = &c.Branches[i]
		}
	}

	if matchedBranch == nil {
		return nil, errors.New("branch configuration not found")
	}
	return matchedBranch, nil
}

func toJSON(i interface{}) string {
	s, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	return string(s)
}

func (c ConfigDef) String() string {
	return toJSON(c)
}

func (c BranchConfig) String() string {
	return toJSON(c)
}
