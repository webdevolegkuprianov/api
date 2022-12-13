package app

import "fmt"

type Environment string

const (
	LocalEnv Environment = "local"
	StageEnv Environment = "stage"
	QAEnv    Environment = "qa"
	ProdEnv  Environment = "prod"
)

func (env Environment) IsValid() error {
	switch env {
	case LocalEnv, StageEnv, ProdEnv, QAEnv:
		return nil
	}

	return fmt.Errorf("неизвестная среда (environment) запуска приложения: %s", env)
}

func (env Environment) String() string {
	return string(env)
}
