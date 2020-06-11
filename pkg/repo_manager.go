package pkg

import (
	"github.com/blang/semver/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// RepoContext for storing repo handle
type RepoContext struct {
	repo *git.Repository
}

//MakeRepoContext Initialize a repo
func MakeRepoContext(path string) (*RepoContext, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	return &RepoContext{repo: repo}, nil
}

func getVersions(tags []*plumbing.Reference) ([]semver.Version, error) {
	var versions []semver.Version
	for _, tag := range tags {
		ver, err := semver.ParseTolerant(tag.Name().Short())
		if err != nil {
			return nil, err
		}
		versions = append(versions, ver)
	}

	return versions, nil
}

func getTagsRefs(repo *git.Repository) ([]*plumbing.Reference, error) {
	tagRefs, err := repo.Tags()
	var allTags []*plumbing.Reference

	if err != nil {
		return allTags, err
	}

	tagRefs.ForEach(func(t *plumbing.Reference) error {
		allTags = append(allTags, t)
		return nil
	})

	return allTags, nil
}

//GetLastTag Get the next version to tag
func (context RepoContext) GetLastTag(path string) (*semver.Version, error) {

	tagRefs, err := getTagsRefs(context.repo)
	if err != nil {
		return nil, err
	}

	versions, err := getVersions(tagRefs)
	if err != nil {
		return nil, err
	}

	semver.Sort(versions)
	lastVersion := versions[len(versions)-1]

	return &lastVersion, nil
}

//GetCurrentBranch lint
func (context RepoContext) GetCurrentBranch() (*string, error) {
	head, err := context.repo.Head()
	if err != nil {
		return nil, err
	}

	branchName := head.Name().Short()

	return &branchName, nil
}
