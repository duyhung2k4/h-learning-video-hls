package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectPostgresql(migrate bool) error {
	var err error
	dns := fmt.Sprintf(
		`
			host=%s
			user=%s
			password=%s
			dbname=%s
			port=%s
			sslmode=disable`,
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	dbPsql, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if migrate {
		errMigrate := dbPsql.AutoMigrate()

		if errMigrate != nil {
			return errMigrate
		}
	}

	return err
}

func connectRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisUrl,
	})
}
