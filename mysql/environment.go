package mysql

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql support
)

const (
	DefaultMaximumOpenConnections = 32
	DefaultIdleConnections        = 4
)

func URL(prefix string) string {
	databaseURL := os.Getenv(fmt.Sprintf("%s_DATABASE_URL", prefix))
	if databaseURL != "" {
		return databaseURL
	}

	address := url.URL{
		User: url.UserPassword(
			os.Getenv(fmt.Sprintf("%s_DATABASE_USERNAME", prefix)),
			os.Getenv(fmt.Sprintf("%s_DATABASE_PASSWORD", prefix)),
		),
		Host: net.JoinHostPort(
			os.Getenv(fmt.Sprintf("%s_DATABASE_SERVICE_HOST", prefix)),
			os.Getenv(fmt.Sprintf("%s_DATABASE_SERVICE_PORT", prefix)),
		),
		Path: os.Getenv(fmt.Sprintf("%s_DATABASE_NAME", prefix)),
	}

	return address.String()
}

// Connect returns new MySQL connection
func Connect(prefix string) *gorm.DB {
	address := URL(prefix)
	connection, err := gorm.Open("mysql", address)
	if err != nil {
		fmt.Printf("Couldn't connect to %s databse: %s - %s\n", prefix, address, err)
		os.Exit(1)
		return nil
	}

	connection.DB().Ping()
	connection.DB().SetMaxIdleConns(idle(prefix))
	connection.DB().SetMaxOpenConns(max(prefix))
	connection.LogMode(log(prefix))

	return connection
}

func max(prefix string) int {
	value := os.Getenv(fmt.Sprintf("%s_DATABASE_MAXIMUM_OPEN_CONNECTIONS", prefix))
	if value == "" {
		value = os.Getenv("DATABASE_MAXIMUM_OPEN_CONNECTIONS")
	}

	if value == "" {
		return DefaultMaximumOpenConnections
	}

	if max, err := strconv.Atoi(value); err == nil {
		return max
	}

	return DefaultMaximumOpenConnections
}

func idle(prefix string) int {
	value := os.Getenv(fmt.Sprintf("%s_DATABASE_IDLE_CONNECTIONS", prefix))
	if value == "" {
		value = os.Getenv("DATABASE_IDLE_CONNECTIONS")
	}

	if value == "" {
		return DefaultIdleConnections
	}

	if idle, err := strconv.Atoi(value); err == nil {
		return idle
	}

	return DefaultIdleConnections
}

// log включение логгирования запросов к БД
func log(prefix string) bool {
	value := os.Getenv(fmt.Sprintf("%s_DATABASE_LOG", prefix))
	if value == "" {
		value = os.Getenv("DATABASE_LOG")
	}

	switch value {
	case "TRUE", "true", "t", "1":
		return true
	default:
		return false
	}
}
