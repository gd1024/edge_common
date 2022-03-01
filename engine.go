package edgeCommon

import "github.com/gd1024/edge_common/component"

type engine struct {
}

func New() *engine {
	return &engine{}
}

func (e *engine) RegisterConfig(filePath string, conf interface{}) {
	component.InitConf(filePath, conf)
}
