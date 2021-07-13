package config

type Mysql struct {
	Path 			string 	`mapstructure:"path" json:"path" yaml:"path"`
	Config 			string	`mapstructure:"config" json:"config" yaml:"config"`
	Dbname 			string	`mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username 		string	`mapstructure:"username" json:"username" yaml:"username"`
	Password 		string	`mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns 	int		`mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns 	int		`mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
}
