package bootstrap

type Env struct{}

func NewEnv() (*Env, error) {
	env := &Env{}
	return env, nil
}
