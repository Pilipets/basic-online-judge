package config

import (
	"encoding/json"
	"io/ioutil"
)

type ServerConfiguration struct {
	Port string
}

type InternalConfiguration struct {
	UploadFilesDir      string
	MaxAllowedFilesSize int
	LoggingDir          string
	DbPath              string
	PythonScriptPath    string
}

type MinioConfiguration struct {
	ConnectionString string
	AccessKeyID      string
	SecretAccessKey  string
	UseSSL           bool
	BucketName       string
}

type Configuration struct {
	Server   ServerConfiguration
	Internal InternalConfiguration
	Minio    MinioConfiguration
}

var (
	Server   ServerConfiguration
	Internal InternalConfiguration
	Minio    MinioConfiguration
)

func ReadConfig(config_path string) error {
	content, err := ioutil.ReadFile(config_path)
	if err != nil {
		return err
	}

	var conf Configuration
	if err = json.Unmarshal(content, &conf); err != nil {
		return err
	}

	Server = conf.Server
	Internal = conf.Internal
	Minio = conf.Minio
	return nil
}
