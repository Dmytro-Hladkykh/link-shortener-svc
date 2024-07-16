package config

import (
	"net"
	"sync"

	"gitlab.com/distributed_lab/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/pgdb"
)

type Config interface {
    comfig.Logger
    pgdb.Databaser
    comfig.Listenerer
    types.Copus
}

type config struct {
    comfig.Logger
    pgdb.Databaser
    comfig.Listenerer
    copus      types.Copus
    log        *logan.Entry
    getter     kv.Getter
    copusOnce  sync.Once
}

func New(getter kv.Getter) Config {
    return &config{
        Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
        Databaser:  pgdb.NewDatabaser(getter),
        Listenerer: comfig.NewListenerer(getter),
        getter:     getter,
    }
}

func (c *config) Copus() types.Copus {
    c.copusOnce.Do(func() {
        c.copus = copus.Must(c.getter)
    })
    return c.copus
}

func (c *config) Log() *logan.Entry {
    return c.Logger.Log()
}

func (c *config) Listener() net.Listener {
    return c.Listenerer.Listener()
}
