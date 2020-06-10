package pkg

import (
	"log"

	"github.com/blang/semver/v4"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func getVersions(tags []*plumbing.Reference) []*semver.Version {
	var versions []*semver.Version
	for _, tag := range tags {
		ver, err := semver.Make(tag.String())
		if err != nil {
			log.Fatal(err)
		}
		versions = append(versions, &ver)
		log.Print(ver)
	}
	return versions
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

func getRepo() error {
	repo, err := git.PlainOpen("/tmp/foo")
	if err != nil {
		return err
	}

	tagRefs, err := getTagsRefs(repo)
	if err != nil {
		return err
	}

	getVersions(tagRefs)

	return nil
}

// func string GetNextVersion(){

// 	return ""
// }
