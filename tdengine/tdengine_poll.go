package tdengine

import "database/sql"

var tdEnginePool map[string]*sql.DB

type TdEngine struct {
	Db        *sql.DB
	InsName   string
	DbName    string
	TableName string
}

func (pg *TdEngine) Conn() {
	if pg.InsName == "" {
		panic("conn tdEngine err : ins name don't is empty")
	}

	if pg.DbName == "" {
		panic("conn tdEngine err : db name don't is empty")
	}

	if pg.TableName == "" {
		panic("conn tdEngine err : table name don't is empty")
	}

	if conn, ok := tdEnginePool[pg.InsName]; ok {
		pg.Db = conn
	} else {
		panic("conn tdEngine err : ins name not found")
	}
}
