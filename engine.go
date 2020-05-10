package packx

import (
	"sync"

	"github.com/eddycjy/packx/driver"
)

type Engine struct {
	drivers []driver.DriverIface
}

func New() *Engine {
	return &Engine{}
}

func (engine *Engine) Use(driver ...driver.DriverIface) {
	engine.drivers = append(engine.drivers, driver...)
}

func (engine *Engine) Run() {
	wg := sync.WaitGroup{}
	for _, handler := range engine.drivers {
		wg.Add(1)
		go func(driver driver.DriverIface) {
			defer wg.Done()
			_ = driver.Serve()
		}(handler)
	}

	wg.Wait()
}
