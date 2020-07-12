package global

import (
	"ctfm_backend/config"

	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	CTFM_DB     *gorm.DB
	CTFM_CONFIG config.Server
	CTFM_VP     *viper.Viper
	CTFM_LOG    *oplogging.Logger
)
