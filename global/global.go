package global

import (
	"github.com/keepgoing/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	Conf config.Server
	VP   *viper.Viper
	
)
