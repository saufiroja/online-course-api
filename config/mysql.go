package config

import "os"

func initMysql(conf *AppConfig) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	dbname := os.Getenv("MYSQL_NAME")
	ssl := os.Getenv("MYSQL_SSL")

	conf.Mysql.Host = host
	conf.Mysql.Port = port
	conf.Mysql.User = user
	conf.Mysql.Pass = pass
	conf.Mysql.Name = dbname
	conf.Mysql.Ssl = ssl
}
