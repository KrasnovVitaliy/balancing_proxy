package config

import (
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Routes Routes `json:"routes"`
}

func (c *Config) Load(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully opened %s", path)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &c)

	log.Printf("Loaded config: %s", c.Routes)
}
