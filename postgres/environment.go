/*
Package postgres

<NAME>_DATABASE_USERNAME
<NAME>_DATABASE_PASSWORD
<NAME>_DATABASE_NAME
<NAME>_DATABASE_SERVICE_HOST
<NAME>_DATABASE_SERVICE_PORT
<NAME>_DATABASE_CONNECTION_POOL_SIZE
DATABASE_CONNECTION_POOL_SIZE
<NAME>_DATABASE_CONNECT_TIMEOUT
DATABASE_CONNECT_TIMEOUT
<NAME>_DATABASE_READ_TIMEOUT
DATABASE_READ_TIMEOUT
<NAME>_DATABASE_WRITE_TIMEOUT
DATABASE_WRITE_TIMEOUT
<NAME>_DATABASE_CONNECTION_POOL_TIMEOUT
DATABASE_CONNECTION_POOL_TIMEOUT
<NAME>_DATABASE_CONNECTION_IDLE_TIMEOUT
DATABASE_CONNECTION_IDLE_TIMEOUT
<NAME>_DATABASE_IDLE_CHECK
DATABASE_IDLE_CHECK
<NAME>_DATABASE_LOG
DATABASE_LOG
*/
package postgres

import (
	"fmt"
	// "log"
	"os"
	"strconv"
	"time"

	"gopkg.in/pg.v5"
)

// func init() {
// 	pg.SetQueryLogger(log.New(os.Stdout, "", log.LstdFlags))
// }

// Connect return new PostrgeSQL connection
func Connect(prefix string) *pg.DB {
	return pg.Connect(Options(prefix))
}

func Options(prefix string) *pg.Options {
	options := &pg.Options{
		User:     os.Getenv(fmt.Sprintf("%s_DATABASE_USERNAME", prefix)),
		Password: os.Getenv(fmt.Sprintf("%s_DATABASE_PASSWORD", prefix)),
		Database: os.Getenv(fmt.Sprintf("%s_DATABASE_NAME", prefix)),
		Addr: fmt.Sprintf("%s:%s",
			os.Getenv(fmt.Sprintf("%s_DATABASE_SERVICE_HOST", prefix)),
			os.Getenv(fmt.Sprintf("%s_DATABASE_SERVICE_PORT", prefix))),
	}

	setPoolSize(prefix, options)
	setPoolTimeout(prefix, options)
	setIdleTimeout(prefix, options)
	setIdleCheckFrequency(prefix, options)
	setConnectionTimeout(prefix, options)
	setReadTimeout(prefix, options)
	setWriteTimeout(prefix, options)

	return options
}

// setPoolSize sets max active connections in pool
func setPoolSize(prefix string, options *pg.Options) {
	value := os.Getenv(fmt.Sprintf("%s_DATABASE_CONNECTION_POOL_SIZE", prefix))

	if value == "" {
		if value = os.Getenv("DATABASE_CONNECTION_POOL_SIZE"); value == "" {
			return
		}
	}

	if size, err := strconv.Atoi(value); err == nil {
		options.PoolSize = size
	}
}

// connectionTimeout sets connect timeout
func setConnectionTimeout(prefix string, options *pg.Options) {
	value := os.Getenv(fmt.Sprintf("%s_DATABASE_CONNECT_TIMEOUT", prefix))

	if value == "" {
		value = os.Getenv("DATABASE_CONNECT_TIMEOUT")
	}

	if timeout, err := time.ParseDuration(value); err == nil {
		options.DialTimeout = timeout
	}
}

// readTimeout set read timeout
func setReadTimeout(prefix string, options *pg.Options) {
	value := os.Getenv(fmt.Sprintf("%s_DATABASE_READ_TIMEOUT", prefix))

	if value == "" {
		value = os.Getenv("DATABASE_READ_TIMEOUT")
	}

	if timeout, err := time.ParseDuration(value); err == nil {
		options.ReadTimeout = timeout
	}
}

// writeTimeout sets write timeout
func setWriteTimeout(prefix string, options *pg.Options) {
	value := os.Getenv(fmt.Sprintf("%s_DATABASE_WRITE_TIMEOUT", prefix))

	if value == "" {
		value = os.Getenv("DATABASE_WRITE_TIMEOUT")
	}

	if timeout, err := time.ParseDuration(value); err == nil {
		options.WriteTimeout = timeout
	}
}

// poolTimeout sets connection checkout timeout
func setPoolTimeout(prefix string, options *pg.Options) {
	value := os.Getenv(fmt.Sprintf("%s_DATABASE_CONNECTION_POOL_TIMEOUT", prefix))

	if value == "" {
		value = os.Getenv("DATABASE_CONNECTION_POOL_TIMEOUT")
	}

	if timeout, err := time.ParseDuration(value); err == nil {
		options.PoolTimeout = timeout
	}
}

// idleTimeout sets connection TTL
func setIdleTimeout(prefix string, options *pg.Options) {
	value := os.Getenv(fmt.Sprintf("%s_DATABASE_CONNECTION_IDLE_TIMEOUT", prefix))

	if value == "" {
		value = os.Getenv("DATABASE_CONNECTION_IDLE_TIMEOUT")
	}

	if timeout, err := time.ParseDuration(value); err == nil {
		options.IdleTimeout = timeout
	}
}

// idleTimeout sets connection check frequency
func setIdleCheckFrequency(prefix string, options *pg.Options) {
	value := os.Getenv(fmt.Sprintf("%s_DATABASE_IDLE_CHECK", prefix))

	if value == "" {
		value = os.Getenv("DATABASE_IDLE_CHECK")
	}

	if frequency, err := time.ParseDuration(value); err == nil {
		options.IdleCheckFrequency = frequency
	}
}

func logs(prefix string) bool {
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
