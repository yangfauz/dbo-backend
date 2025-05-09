package config

type Config struct {
	App        App        `toml:"app"`
	Token      Token      `toml:"token"`
	Connection Connection `toml:"connection"`
}

type App struct {
	Mode  string `toml:"mode"`
	Debug bool   `toml:"debug"`
	Name  string `toml:"name"`
	URL   string `toml:"url"`
	Port  int    `toml:"port"`
}

// Token
type Token struct {
	JWT JWT `toml:"jwt"`
}

type JWT struct {
	SecretKey      string `toml:"key"`
	BaseExpiration int32  `toml:"expired"`
}

// Connection
type Connection struct {
	Postgresql Postgresql `toml:"postgresql"`
}

type Postgresql struct {
	DSN                    string `toml:"dsn"`
	MaxOpenConnections     int    `toml:"max_open_connections"`
	MaxIdleConnections     int    `toml:"max_idle_connections"`
	MaxLifetimeConnections int    `toml:"max_lifetime_connections"`
}
