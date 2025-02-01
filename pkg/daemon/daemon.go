package daemon

import (
	"github.com/stefx99/deucalion/pkg/config"
)

type Daemon struct {
	config *config.Config

	Services []Service
}

func (d *Daemon) parseServiceConfig() {

}

func New() *Daemon {
	daemon := &Daemon{
		config: config.Get(),
	}
	daemon.parseServiceConfig()

	return daemon
}
