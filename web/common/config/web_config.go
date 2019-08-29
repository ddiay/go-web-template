package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type WebConfig struct {
	Debug      bool
	Port       string
	SSLPort    string
	SSLCert    string
	SSLCertKey string
	UseSSL     bool
	Mysql      []MysqlConfig
}

var WebCfg WebConfig

func init() {
	WebCfg.Debug = true
	WebCfg.Port = "80"
	WebCfg.SSLPort = "443"
	WebCfg.SSLCert = ""
	WebCfg.SSLCertKey = ""
	WebCfg.UseSSL = false
	WebCfg.Mysql = []MysqlConfig{}
}

func SaveConfig(path string) error {
	bytes, err := json.MarshalIndent(&WebCfg, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig(path string) error {
	_, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return SaveConfig(path)
	}

	readBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(readBytes, &WebCfg)
	if err != nil {
		return err
	}

	return nil
}
