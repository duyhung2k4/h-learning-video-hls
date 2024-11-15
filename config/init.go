package config

import (
	"flag"
)

func init() {
	db := flag.Bool("db", false, "")

	flag.Parse()

	// connect
	initJwt()
	loadEnv()
	makeVariable()
	connectPostgresql(*db)
	connectRedis()
	initSmptAuth()
	connectRabbitmq()
}
