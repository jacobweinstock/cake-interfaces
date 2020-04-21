package kvm

import (
	"fmt"

	"github.com/jacobweinstock/cake-interfaces/pkg/providers"
)

type KVMBootstrap struct {
	Host     string
	Password string
}

func (k *KVMBootstrap) Prep(b *providers.ConfigSpec) error {
	fmt.Println("doing kvm prep")
	fmt.Printf("Host: %v, Password: %v\n", k.Host, k.Password)
	fmt.Printf("version: %v, count: %v\n", b.KubernetesVersion, b.ControlPlaneCount)
	return nil
}

func (k *KVMBootstrap) Deploy(b *providers.ConfigSpec) error {
	fmt.Println("doing kvm deploy")
	fmt.Printf("Host: %v, Password: %v\n", k.Host, k.Password)
	fmt.Printf("version: %v, count: %v\n", b.KubernetesVersion, b.ControlPlaneCount)
	fmt.Printf("calling to engine: %v\n", b.Engine)
	return nil
}
