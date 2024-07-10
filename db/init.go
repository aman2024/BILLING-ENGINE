package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL() (*sql.DB, error) {

	Host := os.Getenv("MYSQL_HOST")
	Port := os.Getenv("MYSQL_PORT")
	User := os.Getenv("MYSQL_USER")
	Password := os.Getenv("MYSQL_PASSWORD")
	DBName := os.Getenv("MYSQL_DB_NAME")
	MaxIdleConns, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	if err != nil {
		MaxIdleConns = 100
	}
	MaxOpenConns, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	if err != nil {
		MaxOpenConns = 100
	}

	MaxConnIdleTimeMin, err := strconv.Atoi(os.Getenv("MAX_CONN_IDLE_MIN"))
	if err != nil {
		MaxConnIdleTimeMin = 30
	}

	MaxConnLifeTimeMin, err := strconv.Atoi(os.Getenv("MAX_CONN_LIFE_TIME_MIN"))
	if err != nil {
		MaxConnLifeTimeMin = 30
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", User, Password, Host, Port, DBName)
	connection, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	connection.SetMaxIdleConns(MaxIdleConns)
	connection.SetConnMaxIdleTime(time.Duration(MaxConnIdleTimeMin))
	connection.SetConnMaxLifetime(time.Duration(MaxConnLifeTimeMin))
	if MaxOpenConns > 0 {
		connection.SetMaxOpenConns(MaxOpenConns)
	}

	if err := connection.Ping(); err != nil {
		return nil, err
	}

	return connection, nil
}
