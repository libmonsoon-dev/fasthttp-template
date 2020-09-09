package config

import (
	"fmt"
	"github.com/libmonsoon-dev/fasthttp-template/app"
	"github.com/pkg/errors"
	"os"
)

const (
	ServerAddressKey = "SERVER_ADDRESS"
	JwtSecretKey = "JWT_SECRET"
)

func EnvironmentProvider() (*app.Config, error) {
	serverAddress := os.Getenv(ServerAddressKey)
	if serverAddress == "" {
		return nil, errors.WithStack(fmt.Errorf("enviroment variable %v not defined", ServerAddressKey))
	}

	jwtSecret := os.Getenv(JwtSecretKey)
	if jwtSecret == "" {
		return nil, errors.WithStack(fmt.Errorf("enviroment variable %v not defined", JwtSecretKey))
	}

	config := &app.Config{
		serverAddress,
		[]byte(jwtSecret),
	}

	return config, nil
}
