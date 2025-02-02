package daemon

import (
	"github.com/stefx99/deucalion/internal/config"
)

type Monitor struct {
	config *config.Config

	Services []Service
}

func (d *Monitor) parseServiceConfig() {

}

func New() *Monitor {
	daemon := &Monitor{
		config: config.Get(),
	}
	daemon.parseServiceConfig()

	return daemon
}
