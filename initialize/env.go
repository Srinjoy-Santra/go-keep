package initialize

import "os"

func Getenv() (string, error) {
	env := os.Getenv("GO_ENV")
	if env == "" {
		err := os.Setenv("GO_ENV", "development")
		if err != nil {
			return "", err
		}
		env = os.Getenv("GO_ENV")
	}

	return env, nil
}
