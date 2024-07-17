package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
    comfig.Logger
    pgdb.Databaser
    types.Copuser
    comfig.Listenerer

    DB() *pgdb.DB
}

type config struct {
    comfig.Logger
    pgdb.Databaser
    types.Copuser
    comfig.Listenerer
    getter kv.Getter
    db     *pgdb.DB
}

func New(getter kv.Getter) Config {
    db := pgdb.NewDatabaser(getter).DB() 
    return &config{
        getter:     getter,
        db:         db,
        Databaser:  pgdb.NewDatabaser(getter),
        Copuser:    copus.NewCopuser(getter),
        Listenerer: comfig.NewListenerer(getter),
        Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
    }
}

func (c *config) DB() *pgdb.DB {
    return c.db
}
