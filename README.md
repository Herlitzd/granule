# Granule

Granule is a small utility that can be used to easily version software releases. Granule works by looking for existing version tags in a repository, and finding the next semantic version that would be appropriate given the configuration file in the repo.

At the moment, granule simply computes the next semantic version and outputs it to `stdout`. It does not currently tag the repo for you. Additionally, support for metadata following the semver has not been added yet.

## Configuration

By default, granule will look for a configuration file called `granule.config.json` in the current working directory. The configuration file name can be configured, see runtime flags.

### Schema

```jsonc
{
  "branches": [
    {

      // the exact branch name to match
      "branch": "branchName",

      // semver part to bump, can be
      // "major", "minor", "patch"
      "bump": "major"
    },
    {
      // [Optional, default `false`] when true,
      // the branch field is processed as a regex
      "BranchIsRegex: true
      // a regex to match against the branch name
      "branch": "^branchName$",

      // semver part to bump, can be
      // "major", "minor", "patch"
      "bump": "major"
    },
}
```

## Runtime Arguments

These runtime arguments and flags are enumerated with `./granule -h`.

```sh
$ ./granule [flags] [config-file]
```

### Flags

```txt
  -repo, -r string
    	path to git repo root (default ".")
  -verbose, -v
    	enable verbose logging
```


