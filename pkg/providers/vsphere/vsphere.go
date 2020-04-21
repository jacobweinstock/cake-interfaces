package vsphere

import (
	"fmt"

	"github.com/jacobweinstock/cake-interfaces/pkg/config"
	"github.com/jacobweinstock/cake-interfaces/pkg/providers"
)

type ConfigFile struct {
	providers.ConfigSpec `yaml:",inline" json:",inline" mapstructure:",squash"`
	ProviderVsphere      `yaml:",inline" json:",inline" mapstructure:",squash"`
}

type ProviderVsphere struct {
	config.ProviderVsphere `yaml:",inline" json:",inline" mapstructure:",squash"`
}

func (v *ProviderVsphere) Prep(b *providers.ConfigSpec) error {
	fmt.Println("doing vsphere prep")
	n := new(ConfigFile)
	n.ConfigSpec = *b
	n.ProviderVsphere = *v
	fmt.Printf("datacenter: %v, URL: %v\n", n.Datacenter, n.URL)
	fmt.Printf("version: %v, count: %v\n", n.KubernetesVersion, n.ControlPlaneCount)
	return nil
}

func (v *ProviderVsphere) Deploy(b *providers.ConfigSpec) error {
	fmt.Println("doing vsphere deploy")
	n := new(ConfigFile)
	n.ConfigSpec = *b
	n.ProviderVsphere = *v
	fmt.Printf("datacenter: %v, URL: %v\n", n.Datacenter, n.URL)
	fmt.Printf("version: %v, count: %v\n", n.KubernetesVersion, n.ControlPlaneCount)
	fmt.Printf("calling to engine: %v\n", n.Engine)

	return nil
}
