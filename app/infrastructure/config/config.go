package config

import (
	"fmt"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/pkg/errors"
	"os"
)

const (
	ServerAddressKey = "SERVER_ADDRESS"
)

func EnvironmentProvider() (*app.Config, error) {
	serverAddress := os.Getenv(ServerAddressKey)
	if serverAddress == "" {
		return nil, errors.WithStack(fmt.Errorf("enviroment variable %v not defined", ServerAddressKey))
	}

	config := &app.Config{
		serverAddress,
	}

	return config, nil
}
