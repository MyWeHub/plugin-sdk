package util

import (
	"context"
	"fmt"
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

func getContextData(ctx context.Context) (*string, bool) {
	clientId := ctx.Value("clientId").(string)
	superAdmin := ctx.Value("superAdmin").(bool)

	return &clientId, superAdmin
}
