package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var cfg *Config

// Config struct contains the configuration of the server
type Config struct {
	Address      string `json:"address"`
	Port         int    `json:"port"`
	Workdir      string `json:"workdir"`
	ServerKey    string `json:"server_key"`
	ServerCert   string `json:"server_cert"`
	Tls          bool   `json:"tls"`
	User         string `json:"user"`
	PasswordHash string `json:"password_hash"` // sha256 hash of the password
}

func LoadConfig(c string) error {
	file, err := os.Open(c)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	cfg = &config

	return nil
}

func SaveConfig(c string) error {
	file, err := os.Create(c)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func DefaultConfig() *Config {

	// Http handler for the root path "/"
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Warning: cannot get current directory, using \".\"")
		wd = "." // Current directory by default
	}

	return &Config{
		Address:      "0.0.0.0",
		Port:         8080,
		Workdir:      wd,
		ServerKey:    "",
		ServerCert:   "",
		Tls:          false,
		User:         "guest",
		PasswordHash: "b7a875fc1ea228b9061041b7cec4bd3c52ab3ce3c1b2a6f1b5c6f5f8d9d4c4c2", // sha256 hash of the password "guest"
	}
}
