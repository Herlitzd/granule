package pkg

import (
	"log"

	"github.com/blang/semver/v4"
)

// GetNextVersion Get the next version number
func GetNextVersion(branchConfig *BranchConfig, lastVersion *semver.Version) string {
	log.Printf("Using branchConfig %v", branchConfig)
	switch branchConfig.Bump {
	case Major:
		lastVersion.IncrementMajor()
	case Minor:
		lastVersion.IncrementMinor()
	case Patch:
		lastVersion.IncrementPatch()
	}

	return lastVersion.String()
}
