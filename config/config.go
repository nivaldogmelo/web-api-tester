package config

type Config struct {
	Server   Server
	Database DatabaseConfiguration
}

type Server struct {
	Port string
}

type DatabaseConfiguration struct {
	Filename string
}
