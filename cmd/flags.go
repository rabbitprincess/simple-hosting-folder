package main

import (
	"os"

	"github.com/urfave/cli"
)

var (
	// example flag
	mode string
	// input additional flag variable here
	dataPath string
	certPath string
	keyPath  string
)

func flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "mode",                      // flag name
			Value:       getEnv("MODE", "debug"),     // from env or default value
			Usage:       "run mode (debug, release)", // flag description
			Destination: &mode,                       // variable to store the flag value
		},
		cli.StringFlag{
			Name:        "data",
			Value:       getEnv("DATA_PATH", "./data"),
			Usage:       "data path",
			Destination: &dataPath,
		},
		cli.StringFlag{
			Name:        "cert",
			Value:       getEnv("CERT_PATH", "./cert.pem"),
			Usage:       "cert path",
			Destination: &certPath,
		},
		cli.StringFlag{
			Name:        "key",
			Value:       getEnv("KEY_PATH", "./key.pem"),
			Usage:       "key path",
			Destination: &keyPath,
		},
		// input additional flag here
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
