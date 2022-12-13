package postgresql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

func InitGORM(cnf Config, l logger.Interface) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(cnf.DSN),
		&gorm.Config{
			Logger: l,
		},
	)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if cnf.MaxIdleConnections != 0 {
		sqlDB.SetMaxIdleConns(int(cnf.MaxIdleConnections))
	}
	if cnf.MaxOpenConnections != 0 {
		sqlDB.SetMaxOpenConns(int(cnf.MaxOpenConnections))
	}
	if cnf.MaxLifetimeConnection != 0 {
		sqlDB.SetConnMaxLifetime(cnf.MaxLifetimeConnection)
	}

	return db, nil
}

type ZerologGORM struct {
	Log                       zerolog.Logger
	LogLevel                  logger.LogLevel
	IgnoreRecordNotFoundError bool
	SlowThreshold             time.Duration
}

func (l *ZerologGORM) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *ZerologGORM) Info(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Log.Debug().Msgf(s, i...)
	}
}

func (l *ZerologGORM) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Log.Warn().Msgf(s, i...)
	}
}

func (l *ZerologGORM) Error(ctx context.Context, s string, i ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Log.Error().Msgf(s, i...)
	}
}

func (l *ZerologGORM) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()

	log := l.Log.With().
		Str("File", utils.FileWithLineNum()).
		Int64("Rows", rows).
		Float64("Duration", float64(elapsed.Nanoseconds())/1e6).
		Logger()

	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		log.Trace().Err(err).Msg(sql)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		ll := log.With().Str("SLOW SQL", fmt.Sprintf(" threshold: %v", l.SlowThreshold)).Logger()
		ll.Warn().Msg(sql)
	case l.LogLevel == logger.Info:
		log.Debug().Msg(sql)
	}
}
