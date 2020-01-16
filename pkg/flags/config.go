package flags

type Config struct {
	UseSyslog int    `long:"use-syslog" env:"USE_SYSLOG" description:"UseSyslog" default:"1"`
	LogDir    string `long:"log-dir" env:"LOG_DIR" required:"true"`

	AppPort   int    `long:"port" env:"APP_PORT" default:"8000"`
	AppSecret string `long:"secret" env:"APP_SECRET" required:"true"`
}
