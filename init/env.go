package init

import "os"

func Getenv() (String, error) {
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
