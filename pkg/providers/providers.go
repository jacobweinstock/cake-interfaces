package providers

import (
	"fmt"

	"github.com/jacobweinstock/cake-interfaces/pkg/config"
)

type Bootstrap interface {
	Prep(*ConfigSpec) error
	Deploy(*ConfigSpec) error
}

type ConfigSpec struct {
	config.ConfigSpec `yaml:",inline" json:",inline" mapstructure:",squash"`
	ProviderContext   Bootstrap `yaml:"ProviderContext,omitempty" json:"providercontext,omitempty"`
}

func (c *ConfigSpec) RunBootstrap() error {
	var err error

	fmt.Printf("RunBootstrap: %+v\n", c)
	err = c.ProviderContext.Prep(c)
	if err != nil {
		return err
	}
	err = c.ProviderContext.Deploy(c)
	if err != nil {
		return err
	}

	return err
}
