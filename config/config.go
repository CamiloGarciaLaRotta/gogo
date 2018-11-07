package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// Config contains the configuration of gogoGadget
type Config struct {
	Postgres Postgres `json:"postgres"`
}

// Postgres contains the configuration for a postgres DB
type Postgres struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

// FromFile returns a configuration parsed from the given file
func FromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %q", path)
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal %q", path)
	}

	if cfg.Postgres.Host == "" {
		return nil, errors.New("Host variable required but not set")
	}
	if cfg.Postgres.Port == "" {
		return nil, errors.New("Port variable required but not set")
	}
	if cfg.Postgres.User == "" {
		return nil, errors.New("User variable required but not set")
	}
	if cfg.Postgres.Pass == "" {
		return nil, errors.New("Pass variable required but not set")
	}
	if cfg.Postgres.Name == "" {
		return nil, errors.New("Name variable required but not set")
	}

	return &cfg, nil
}

// FromEnv returns a configuration parsed from the OS ENV vars
func FromEnv() (*Config, error) {
	host, ok := os.LookupEnv("DBHOST")
	if !ok {
		return nil, errors.New("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv("DBPORT")
	if !ok {
		return nil, errors.New("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv("DBUSER")
	if !ok {
		return nil, errors.New("DBUSER environment variable required but not set")
	}
	pass, ok := os.LookupEnv("DBPASS")
	if !ok {
		return nil, errors.New("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv("DBNAME")
	if !ok {
		return nil, errors.New("DBNAME environment variable required but not set")
	}

	cfg := Config{
		Postgres{
			Host: host,
			Port: port,
			User: user,
			Pass: pass,
			Name: name,
		},
	}
	return &cfg, nil
}
