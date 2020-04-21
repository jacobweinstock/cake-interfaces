package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jacobweinstock/cake-interfaces/pkg/config"
	"github.com/jacobweinstock/cake-interfaces/pkg/engines"
	"github.com/jacobweinstock/cake-interfaces/pkg/providers"
	"github.com/jacobweinstock/cake-interfaces/pkg/providers/kvm"
	"github.com/jacobweinstock/cake-interfaces/pkg/providers/vsphere"
)

func main() {
	fmt.Println("==================== PROVIDERS")
	providerVsphere := vsphere.ProviderVsphere{
		ProviderVsphere: config.ProviderVsphere{
			Datacenter: "HCI-DATACENTER-01",
			URL:        "https://172.60.0.150",
		},
	}
	providerKVM := kvm.KVMBootstrap{
		Host:     "172.60.0.151",
		Password: "secret",
	}

	bootstrapGo := providers.ConfigSpec{
		ConfigSpec: config.ConfigSpec{
			Cluster: config.Cluster{
				KubernetesVersion: "v1.17.3",
				ControlPlaneCount: 5,
			},
			Local: true,
		},
	}

	bootstrapGo.ProviderContext = &providerVsphere
	bootstrapGo.Engine = "RKE"
	err := bootstrapGo.RunBootstrap()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("====================")

	bootstrapGo.ProviderContext = &providerKVM
	bootstrapGo.Engine = "RKE"
	err = bootstrapGo.RunBootstrap()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("==================== ENGINES")

	engineGo := engines.ConfigSpec{
		ConfigSpec: config.ConfigSpec{
			Cluster: config.Cluster{
				KubernetesVersion: "v1.17.3",
				ControlPlaneCount: 5,
			},
			Local: true,
		},
	}
	fmt.Printf("%+v\n", engineGo)
	fmt.Println("====================")
}
