package engines

import (
	"fmt"

	"github.com/jacobweinstock/cake-interfaces/pkg/config"
)

type Engine interface {
	Create(*ConfigSpec) error
	Pivot(*ConfigSpec) error
}

type ConfigSpec struct {
	config.ConfigSpec `yaml:",inline" json:",inline" mapstructure:",squash"`
	EngineContext     Engine `yaml:"ProviderContext,omitempty" json:"providercontext,omitempty"`
}

func (e *ConfigSpec) RunEngine() error {
	var err error

	fmt.Printf("RunEngine: %+v\n", e)
	err = e.EngineContext.Create(e)
	if err != nil {
		return err
	}
	err = e.EngineContext.Pivot(e)
	if err != nil {
		return err
	}

	return err
}
