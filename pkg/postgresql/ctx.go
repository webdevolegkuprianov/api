package postgresql

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type ctxKey string

const (
	gormCtxKey      ctxKey = "psql_gorm"
	sqlDriverCtxKey ctxKey = "psql_driver"
)

func ContextWithGorm(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, gormCtxKey, db)
}

func GormFromContext(ctx context.Context) (db *gorm.DB, ok bool) {
	if v := ctx.Value(gormCtxKey); v != nil {
		db, ok = v.(*gorm.DB)
	}

	return
}

func ContextWithSqlDriver(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, sqlDriverCtxKey, db)
}

func SqlDriverFromContext(ctx context.Context) (db *sql.DB, ok bool) {
	if v := ctx.Value(gormCtxKey); v != nil {
		db, ok = v.(*sql.DB)
	}
	return
}
