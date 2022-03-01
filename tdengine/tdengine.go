package tdengine

import (
	"database/sql"
	"fmt"
)

type TdEngineConf struct {
	InsName      string `json:"ins_name"`
	Driver       string `json:"driver"`
	Addr         string `json:"addr"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Db           string `json:"db"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxIdleTime  int    `json:"max_idle_time"`
	MaxLifeTime  int    `json:"max_life_time"`
	MaxOpenConns int    `json:"max_open_conns"`
}

func InitTdEngine(tdConf []TdEngineConf) {

	for _, conf := range tdConf {

		url := fmt.Sprintf("%s:%s@/tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Addr, conf.Port, conf.Db)
		db, err := sql.Open(conf.Driver, url)
		if err != nil {
			panic("init td engine err : " + err.Error())
		}

		db.SetMaxIdleConns(conf.MaxIdleConns)
		//db.SetConnMaxIdleTime(time.Duration(pgsqlConfig.MaxIdleTime) * time.Second)
		//db.SetConnMaxLifetime(time.Duration(conf.MaxLifeTime) * time.Second)
		db.SetMaxOpenConns(conf.MaxOpenConns)

		tdEnginePool[conf.InsName] = db
	}
}
