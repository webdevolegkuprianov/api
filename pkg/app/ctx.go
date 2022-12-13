package app

import (
	"context"
)

const ctxKey = "application_env"

func ContextWithEnvironment(ctx context.Context, appEnv Environment) context.Context {
	return context.WithValue(ctx, ctxKey, appEnv)
}

func EnvironmentFromContext(ctx context.Context) (Environment, bool) {
	var d Environment
	var ok bool
	if v := ctx.Value(ctxKey); v != nil {
		d, ok = v.(Environment)
	}

	if !ok {
		d = LocalEnv
		ok = true
	}
	if err := d.IsValid(); err != nil {
		d = LocalEnv
		ok = true
	}

	return d, ok
}
