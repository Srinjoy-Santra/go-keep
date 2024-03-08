package config

type Configuration struct {
	Server   Server
	Database Database
}

type Server struct {
	HTTP HTTP `yaml:"http"`
}

type HTTP struct {
	Address string `yaml:"address"`
}

type Database struct {
	Source     string `yaml:"source"`
	InMemory   InMemory
	Relational Relational
}

type InMemory struct {
}

type Relational struct {
	Connection string `yaml:"connection"`
}
