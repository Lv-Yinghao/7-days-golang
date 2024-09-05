package GeeORM

import (
	"database/sql"
	"geeorm/log"
	"geeorm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (*Engine, error) {
	db, err := sql.Open(driver, source)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}

	e := Engine{db: db}
	log.Info("Successfully connected to database")

	return &e, nil
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error(err)
	}
	log.Info("Successfully closed database")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
