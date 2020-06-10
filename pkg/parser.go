package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ParseConfig Load a configuration
func ParseConfig(path *string) *ConfigDef {
	file, err := ioutil.ReadFile(*path)

	if err != nil {
		log.Fatal(err)
	}

	var configDef ConfigDef
	err = json.Unmarshal(file, &configDef)
	if err != nil {
		log.Fatal(err)
	}

	return &configDef
}
