package postgresql

import (
	"time"
)

type Config struct {
	DSN                   string        `envconfig:"POSTGRE_DSN" required:"true"`
	MaxIdleConnections    uint          `envconfig:"POSTGRE_MAX_IDLE_CONS" default:"10"`
	MaxOpenConnections    uint          `envconfig:"POSTGRE_MAX_OPEN_CONS" default:"10"`
	MaxLifetimeConnection time.Duration `envconfig:"POSTGRE_MAX_LIFETIME_CON" default:"0"`
}
