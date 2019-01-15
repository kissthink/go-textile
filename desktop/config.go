package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

// Config is used to load textile desktop config files.
type Config struct {
	BackupFolder string `json:"backup_folder,omitempty"` // desktop photo backup folder path
}

// ReadConfig reads a config from disk
func ReadConfig(repoPath string) (*Config, error) {
	var conf *Config
	data, err := ioutil.ReadFile(path.Join(repoPath, "desktop"))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// Write writes an on-disk version of config
func WriteConfig(repoPath string, conf *Config) error {
	f, err := os.Create(path.Join(repoPath, "desktop"))
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}