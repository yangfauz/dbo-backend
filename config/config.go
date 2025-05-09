package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/pelletier/go-toml/v2"
)

var AppConfig Config

func Load() Config {

	// Find file location
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		panic("unable to get the current filename")
	}
	filePath := filepath.Join(filename, "../")
	configPath := fmt.Sprintf("%s/%s", path.Join(path.Dir(filePath)), "config.toml")

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteRead, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Scan config to Struct
	var zConfig Config

	err = toml.Unmarshal(byteRead, &zConfig)
	if err != nil {
		log.Fatal(err)
	}

	AppConfig = zConfig

	return zConfig
}
