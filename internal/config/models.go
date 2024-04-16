package config

type Configuration struct {
	Server   Server
	Database Database
	Auth     Auth
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
	Relational Relational `yaml:"relational"`
}

type InMemory struct {
}

type Relational struct {
	Connection string `yaml:"connection"`
}

type Auth struct {
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
	Domain       string `yaml:"domain"`
	CallbackURL  string `yaml:"callbackURL"`
}
