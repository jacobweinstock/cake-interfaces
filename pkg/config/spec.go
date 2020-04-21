package config

// ConfigSpec holds information needed to provision a K8s management cluster
type ConfigSpec struct {
	Provider string `yaml:"Provider" json:"provider"`
	Engine   string `yaml:"Engine" json:"engine"`
	SSH      SSH    `yaml:"SSH" json:"ssh"`
	Local    bool   `yaml:"Local" json:"local"`
	LogFile  string `yaml:"LogFile,omitempty" json:"logfile,omitempty"`
	Addons   Addons `yaml:"Addons,omitempty" json:"addons,omitempty"`
	Cluster  `yaml:",inline" json:",inline" mapstructure:",squash"`
}

// Cluster specifies the details about the management cluster
type Cluster struct {
	ClusterName           string `yaml:"ClusterName" json:"clustername"`
	ControlPlaneCount     int    `yaml:"ControlPlaneCount" json:"controlplanecount"`
	ControlPlaneSize      string `yaml:"ControlPlaneSize" json:"controlplanesize"`
	WorkerCount           int    `yaml:"WorkerCount" json:"workercount"`
	WorkerSize            string `yaml:"WorkerSize" json:"workersize"`
	KubernetesVersion     string `yaml:"KubernetesVersion" json:"kubernetesversion"`
	KubernetesPodCidr     string `yaml:"KubernetesPodCidr" json:"kubernetespodcidr"`
	KubernetesServiceCidr string `yaml:"KubernetesServiceCidr" json:"kubernetesservicecidr"`
}

// SSH holds ssh info
type SSH struct {
	Username      string `yaml:"Username" json:"username"`
	AuthorizedKey string `yaml:"AuthorizedKey" json:"authorizedkey"`
}

// Addons holds optional configuration values
type Addons struct {
	Observability ObservabilitySpec `yaml:"Observability,omitempty" json:"observability,omitempty"`
}

// ObservabilitySpec holds values for the observability archive file
type ObservabilitySpec struct {
	Enable          bool   `yaml:"Enable" json:"enable"`
	ArchiveLocation string `yaml:"ArchiveLocation" json:"archivelocation"`
}

// ProviderVsphere is vsphere specifc data
type ProviderVsphere struct {
	URL               string  `yaml:"URL" json:"url"`
	Username          string  `yaml:"Username" json:"username"`
	Password          string  `yaml:"Password" json:"password"`
	Datacenter        string  `yaml:"Datacenter" json:"datacenter"`
	ResourcePool      string  `yaml:"ResourcePool" json:"resourcepool"`
	Datastore         string  `yaml:"Datastore" json:"datastore"`
	ManagementNetwork string  `yaml:"ManagementNetwork" json:"managementnetwork"`
	StorageNetwork    string  `yaml:"StorageNetwork" json:"storagenetwork"`
	OVA               OVASpec `yaml:"OVA" json:"ova"`
}

// OVASpec sets OVA information used for virtual machine templates
type OVASpec struct {
	NodeTemplate         string `yaml:"NodeTemplate" json:"nodetemplate"`
	LoadbalancerTemplate string `yaml:"LoadbalancerTemplate" json:"loadbalancertemplate"`
}
