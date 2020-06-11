package pkg

import (
	"encoding/json"
	"errors"
	"log"
	"regexp"
)

// ConfigDef Top level config object
type ConfigDef struct {
	Branches []BranchConfig `json:"branches"`
}

// BranchConfig Rule for branches
type BranchConfig struct {
	Branch string    `json:"branch"`
	Bump   BumpLevel `json:"bump"`
}

// BumpLevel Enum for Version Levels
type BumpLevel string

// Version Levels
const (
	Major BumpLevel = "major"
	Minor BumpLevel = "minor"
	Patch BumpLevel = "patch"
)

// GetBranchConfig lint
func (c ConfigDef) GetBranchConfig(branchName *string) (*BranchConfig, error) {
	var matchedBranch *BranchConfig
	for _, branch := range c.Branches {
		matched, err := regexp.MatchString(branch.Branch, *branchName)
		log.Printf(*branchName)
		if err != nil {
			return nil, err
		}
		if matched {
			matchedBranch = &branch
		}
	}
	if matchedBranch == nil {
		return nil, errors.New("Branch not configured")
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
