package bootstrap

type Env struct {
	SERVERenv  string
	SERVERport string
	DBport     string
	DBhost     string
	DBname     string
	DBsslmode  string
	DBusername string
	DBpassword string
}

func NewEnv() (*Env, error) {
	env := &Env{}
	return env, nil
}
