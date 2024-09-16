package config

type Config struct {
	DB  DB  `json:"db"`
	API API `json:"api"`
}

type API struct {
	Port int64 `json:"port"`
}

type DB struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}
