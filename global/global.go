package global

import (
	"github.com/keepgoing/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
	"golang.org/x/sync/singleflight"
)

var (
	DB   *gorm.DB
	Conf config.Server
	VP   *viper.Viper
	Concurrency_Control= &singleflight.Group{}
	
)

type CommonModel struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}