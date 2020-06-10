package pkg

import (
	"encoding/json"
	"log"
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

func (c ConfigDef) String() string {
	s, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	return string(s)
}
