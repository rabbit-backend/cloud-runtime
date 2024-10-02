package env

import "os"

func (env *EnvModule) Getenv(key string) string {
	return os.Getenv(key)
}
