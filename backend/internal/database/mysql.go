package database

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"stockai/backend/internal/config"
)

const (
	mysqlMaxOpenConns    = 20
	mysqlMaxIdleConns    = 5
	mysqlConnMaxLifetime = 30 * time.Minute
	mysqlPingTimeout     = 5 * time.Second
)

func NewMySQL(ctx context.Context, cfg config.MySQLConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("open mysql: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get mysql connection pool: %w", err)
	}

	sqlDB.SetMaxOpenConns(mysqlMaxOpenConns)
	sqlDB.SetMaxIdleConns(mysqlMaxIdleConns)
	sqlDB.SetConnMaxLifetime(mysqlConnMaxLifetime)

	pingCtx, cancel := context.WithTimeout(ctx, mysqlPingTimeout)
	defer cancel()
	if err := sqlDB.PingContext(pingCtx); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping mysql: %w", err)
	}

	return db, nil
}
