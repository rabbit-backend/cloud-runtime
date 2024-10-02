package env

func (env *EnvModule) Exports() map[string]interface{} {
	return map[string]interface{}{
		"getEnv": env.Getenv,
	}
}
