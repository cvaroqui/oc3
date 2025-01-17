package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"
	_ "time/tzdata"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// initDatabase setup database handler.
func newDatabase() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 viper.GetString("db.username"),
		Passwd:               viper.GetString("db.password"),
		Net:                  "tcp",
		Addr:                 viper.GetString("db.host") + ":" + viper.GetString("db.port"),
		DBName:               "opensvc",
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  time.Local,
	}
	slog.Info(fmt.Sprintf("db addr=%s", cfg.Addr))
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
