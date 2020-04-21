package kvm

import (
	"fmt"

	"github.com/jacobweinstock/cake-interfaces/pkg/engines"
)

type ConfigSpec struct {
	engines.ConfigSpec
}

func (c *ConfigSpec) Create(e *engines.ConfigSpec) error {
	var err error
	fmt.Printf("in RKE:kvm engine CREATE: %+v", c)
	fmt.Printf("in RKE:kvm engine CREATE: %+v", e)
	return err
}

func (c *ConfigSpec) Pivot(e *engines.ConfigSpec) error {
	var err error
	fmt.Printf("in RKE:kvm engine PIVOT: %+v", c)
	fmt.Printf("in RKE:kvm engine PIVOT: %+v", e)
	return err
}
