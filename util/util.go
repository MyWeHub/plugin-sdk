package util

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	EnvPROD = "PROD"
	EnvDEV  = "DEV"
)

func LoadEnvironment() string {
	return GetEnv("ENVIRONMENT", false, EnvPROD, false)
}

func GetStringPtr(s string) *string {
	return &s
}

func GetEnv(name string, validatePort bool, def string, mandatory bool) string {
	env := os.Getenv(name)
	switch {
	case env == "":
		if mandatory {
			panic(fmt.Errorf("env %s not found", name))
		} else {
			return def
		}
	case validatePort:
		_, err := strconv.ParseInt(env, 10, 64)
		if err != nil {
			panic("Wrong port value")
		}
		return env
	default:
		return env
	}
	return ""
}

func GetClientID(ctx context.Context) string {
	if clientId := ctx.Value("clientId"); clientId != nil {
		if stClientID, ok := clientId.(string); ok {
			return stClientID
		} else if !ok {
			log.Println("Error: clientID is not a string")
		}
	}
	return ""
}
