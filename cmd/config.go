package main

import "github.com/ilyakaznacheev/cleanenv"

type HTTPConfig struct {
	Addr         string `json:"addr"`
	Port         string `json:"port"`
	ReadTimeout  int64  `json:"read_timeout"`
	WriteTimeout int64  `json:"write_timeout"`
}

func (c *HTTPConfig) ReadConfig(filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
