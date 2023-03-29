package config

type MySql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`
	Usernamne    string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Singular     bool   `mapstructure:"sigular" json:"sigular" yaml:"sigular"`
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode      string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	LogZap       bool   `mapstructure:"log_zap" json:"log_zap" yaml:"log_zap"`
}

func (m *MySql) Dsn() string {
	return m.Usernamne + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *MySql) GetLogMode() string {
	return m.LogMode
}
