package wegospider

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type MongoConfiguration struct {
	Host string
	Port string
}

type Configuration struct {
	Mongo MongoConfiguration
}


func (c *Configuration) LoadFromFile() error {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fileName := filepath.Join(dir, "conf/main.json")
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&c)
	log.Println(c)
	if err != nil {
		return err
	}
	return nil
}
