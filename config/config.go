package config

import (
	"os"
	osuser "os/user"
	"path"

	"github.com/pelletier/go-toml"
	"github.com/scott-wilson/godo/user"
)

type Config struct {
	User       user.User
	RootPath   string
	ConfigPath string
}

func Read() (Config, error) {
	c := Config{}

	rootPath, err := _rootPath()

	if err != nil {
		return c, err
	}

	configPath, err := _configPath()

	if err != nil {
		return c, err
	}

	_, err = os.Stat(configPath)

	if os.IsNotExist(err) {
		f, err := os.Create(configPath)
		defer f.Close()

		if err != nil {
			return c, err
		}
	}

	config, err := toml.LoadFile(configPath)

	if err != nil {
		return c, err
	}

	username := config.Get("user.name").(string)
	email := config.Get("user.email").(string)

	c.User.Name = username
	c.User.Email = email
	c.RootPath = rootPath
	c.ConfigPath = configPath

	return c, nil
}

func (c *Config) Write() {}

func _rootPath() (string, error) {
	sysUser, err := osuser.Current()

	if err != nil {
		return "", err
	}

	rootPath := path.Join(sysUser.HomeDir, ".godo")
	err = os.Mkdir(rootPath, 0755)

	if !os.IsExist(err) {
		return "", err
	}

	return rootPath, nil
}

func _configPath() (string, error) {
	// Get config
	rootPath, err := _rootPath()

	if err != nil {
		return "", err
	}

	configPath := path.Join(rootPath, ".config")

	return configPath, nil
}
