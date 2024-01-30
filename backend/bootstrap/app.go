package bootstrap

import "log"

type Application struct {
	Env *Env
}

func App() Application {
	app := &Application{}
	env, err := NewEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	app.Env = env

	return *app
}
