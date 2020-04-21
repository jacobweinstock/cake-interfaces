package vsphere

import (
	"fmt"

	"github.com/jacobweinstock/cake-interfaces/pkg/engines"
)

type ConfigSpec struct {
	EngineBase engines.ConfigSpec `yaml:",inline" json:",inline" mapstructure:",squash"`
}

func (c *ConfigSpec) Create(e *engines.ConfigSpec) error {
	var err error
	fmt.Printf("in RKE:vsphere engine CREATE: %+v\n", c)
	fmt.Printf("in RKE:vsphere engine CREATE: %+v\n", e)
	return err
}

func (c *ConfigSpec) Pivot(e *engines.ConfigSpec) error {
	var err error
	fmt.Printf("in RKE:vsphere engine Pivot: %+v\n", c)
	fmt.Printf("in RKE:vsphere engine Pivot: %+v\n", e)
	return err
}
