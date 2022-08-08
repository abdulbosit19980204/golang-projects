package config

import "fmt"

type Config struct {
	pgPassword     string
	pgUserName     string
	pgDatabaseName string
	pgPort         int
	pgHost         string
}

func New(pass, uName, dName, h string, p int) Config {
	return Config{
		pgPassword:     pass,
		pgUserName:     uName,
		pgDatabaseName: dName,
		pgPort:         p,
		pgHost:         h,
	}
}

func (c Config) GetPgURL() string {
	pgURL := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

	return fmt.Sprintf(pgURL, c.pgHost, c.pgPort, c.pgUserName, c.pgPassword, c.pgDatabaseName)
}
