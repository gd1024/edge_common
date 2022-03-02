package edgeCommon

import (
	"github.com/gd1024/edge_common/config"
	logger2 "github.com/gd1024/edge_common/logger"
	"github.com/gd1024/edge_common/mqtt"
	"github.com/gd1024/edge_common/pgsql"
)

type engine struct {
}

func New() *engine {
	return &engine{}
}

func (e *engine) RegisterConfig(filePath string, conf interface{}) {
	config.InitConf(filePath, conf)
}

func (e *engine) RegisterLogger(logPath string) {
	logger2.InitLog(logPath)
}

func (e *engine) RegisterMqtt(confs []mqtt.MqttConf) {
	mqtt.InitMqtt(confs)
}

func (e *engine) RegisterPgsql(pgConf []pgsql.PgConf) {
	pgsql.InitPgsql(pgConf)
}
