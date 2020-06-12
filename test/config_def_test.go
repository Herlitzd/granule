package main_test

import (
	"strings"
	"testing"

	"github.com/Herlitzd/granule/pkg"
)

// TestBranchConfigDefaults test
func TestBranchConfigDefaults(t *testing.T) {
	branchConfig := pkg.BranchConfig{Branch: "master", Bump: pkg.Major}

	if branchConfig.BranchIsRegex == true {
		t.Errorf("BranchIsRegex must default to `false`")
	}

}

func makeConfigDef(branchConfigs ...pkg.BranchConfig) pkg.ConfigDef {
	repoConfig := pkg.ConfigDef{Branches: branchConfigs}
	return repoConfig
}

// TestBranchConfigDefaults test
func TestGetBranchConfigWithoutRegex(t *testing.T) {
	masterConfig := pkg.BranchConfig{Branch: "master", Bump: pkg.Major}
	developConfig := pkg.BranchConfig{Branch: "develop", Bump: pkg.Minor}

	repoConfig := makeConfigDef(masterConfig, developConfig)

	for _, branchName := range [...]string{"master", "develop"} {

		branchConfig, err := repoConfig.GetBranchConfig(&branchName)
		if err != nil {
			t.Error(err)
		}

		if branchConfig.Branch != branchName {
			t.Errorf("Expected matching branch to be `%s` got %s",
				branchName, branchConfig.Branch)
		}

	}

	notConfiguredBranchName := "not_configured"
	branchConfig, err := repoConfig.GetBranchConfig(&notConfiguredBranchName)

	if branchConfig != nil {
		t.Error("an un-configured branch should not return a value")
	}

	if err == nil {
		t.Error("an un-configured branch should return a error")
	}

}

// TestBranchConfigDefaults test
func TestGetBranchConfigWithRegexTrivial(t *testing.T) {
	masterConfig := pkg.BranchConfig{Branch: "^master$", BranchIsRegex: true, Bump: pkg.Major}
	developConfig := pkg.BranchConfig{Branch: "^develop$", BranchIsRegex: true, Bump: pkg.Minor}

	repoConfig := makeConfigDef(masterConfig, developConfig)

	for _, branchName := range [...]string{"master", "develop"} {

		branchConfig, err := repoConfig.GetBranchConfig(&branchName)
		if err != nil {
			t.Error(err)
		}

		if strings.Contains(branchName, branchConfig.Branch) {
			t.Errorf("Expected matching branch to be `%s` got %s",
				branchName, branchConfig.Branch)
		}

	}

}

// TestGetBranchConfigWithRegexNonTrivial test
func TestGetBranchConfigWithRegexNonTrivial(t *testing.T) {
	featurePattern := "^feature/.+$"
	masterConfig := pkg.BranchConfig{Branch: featurePattern, BranchIsRegex: true, Bump: pkg.Minor}
	developConfig := pkg.BranchConfig{Branch: "^develop$", BranchIsRegex: true, Bump: pkg.Patch}

	repoConfig := makeConfigDef(masterConfig, developConfig)

	featureBranchName := "feature/addSomeFeature"
	branchConfig, err := repoConfig.GetBranchConfig(&featureBranchName)
	if err != nil {
		t.Error(err)
	}
	if branchConfig.Branch != featurePattern {
		t.Errorf("Expected matching branch to be `%s` got %s",
			featurePattern, branchConfig.Branch)
	}
}

// TestDuplicateMatchingBranches test
func TestDuplicateMatchingBranches(t *testing.T) {
	masterConfig := pkg.BranchConfig{Branch: "master", Bump: pkg.Major}
	secondMasterConfig := pkg.BranchConfig{Branch: "master", Bump: pkg.Minor}

	repoConfig := makeConfigDef(masterConfig, secondMasterConfig)

	branchName := "master"
	_, err := repoConfig.GetBranchConfig(&branchName)

	if err == nil {
		t.Error("expecting error for overlapping branches")
	}
}
