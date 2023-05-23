package config

type Config struct {
	HttpConfig HttpConfig
}

type HttpConfig struct {
	Host string
	Port string
}

type ParserConfig struct {
	ParserType string
	FileParse  string
}

func LoadConfig(file string) Config {
	return Config{}
}
