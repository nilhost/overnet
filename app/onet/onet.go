/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2021-01-25
 */
package onet


//go:generate errorgen

import (
	"context"
	"sync"

	"fmt"

	"github.com/nilhost/overnet/common"
	"github.com/nilhost/overnet/common/log"
)

// Instance is a log.Handler that handles logs.
type Instance struct {
	sync.RWMutex
	config *Config
	active bool
}

// New creates a new log.Instance based on the given config.
func New(ctx context.Context, config *Config) (*Instance, error) {
	g := &Instance{
		config: config,
		active: false,
	}
	return g, nil
}

// Type implements common.HasType.
func (*Instance) Type() interface{} {
	return (*Instance)(nil)
}

// Start implements common.Runnable.Start().
func (g *Instance) Start() error {
	fmt.Println("this is ONET Instance")
	return nil
}

// Handle implements log.Handler.
func (g *Instance) Handle(msg log.Message) {
	g.RLock()
	defer g.RUnlock()

	if !g.active {
		return
	}
}

// Close implements common.Closable.Close().
func (g *Instance) Close() error {
	newError("Logger closing").AtDebug().WriteToLog()
	g.Lock()
	defer g.Unlock()
	return nil
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return New(ctx, config.(*Config))
	}))
}
