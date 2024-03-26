package global

import (
	"github.com/go_demo/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	Conf config.Server
	VP   *viper.Viper
	
)
