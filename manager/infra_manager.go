package manager

import (
	"fmt"

	"github.com/enigma-mart/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type InfraManager interface {
	SqlDb() *sqlx.DB
}

type infraManager struct {
	db  *sqlx.DB
	cfg config.Config
}

func (i *infraManager) SqlDb() *sqlx.DB {
	return i.db
}

func (i *infraManager) initDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", i.cfg.Host, i.cfg.User, i.cfg.Password, i.cfg.DbName, i.cfg.Port)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	i.db = db
}

func NewInfraManager(cfg config.Config) InfraManager {
	infra := infraManager{cfg: cfg}
	infra.initDB()
	return &infra
}
