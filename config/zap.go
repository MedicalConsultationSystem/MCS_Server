package config

type Zap struct {
	Dir      string `mapstructure:"dir" json:"dir" yaml:"dir"`
	Header   string `mapstructure:"header" json:"header" yaml:"header"`
	Format   string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`
}
