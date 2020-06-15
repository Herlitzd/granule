package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ParseConfig Load a configuration
func ParseConfig(path *string) (*ConfigDef, error) {
	file, err := ioutil.ReadFile(*path)

	if err != nil {
		return nil, fmt.Errorf("configuration file not found at '%s'\n%w", *path, err)
	}

	var configDef ConfigDef
	err = json.Unmarshal(file, &configDef)
	if err != nil {
		return nil, err
	}

	return &configDef, nil
}
