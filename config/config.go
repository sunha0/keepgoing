package config

const (
	ConfigEnv         = "CONFIG"
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigDebugFile   = "config.debug.yaml"
	ConfigReleaseFile = "config.release.yaml"
)

type Server struct {
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mongo   Mongo   `mapstructure:"mongo" json:"mongo" yaml:"mongo"`

	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`


	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
