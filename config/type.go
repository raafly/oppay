package config

type App struct {
	Postgres Postgres
}

type Postgres struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}