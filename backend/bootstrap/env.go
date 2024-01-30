package bootstrap

type Env struct {
	SERVERenv  string
	SERVERport int
	DBport     int
	DBhost     string
	DBname     string
	DBsslmode  string
	DBusername string
}

func NewEnv() (*Env, error) {
	env := &Env{}
	return env, nil
}
