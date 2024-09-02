package ystruct

type Container struct {
	Name      string    `yaml:"name,omitempty"`
	Resources Resources `yaml:"resources,omitempty"`
	Cluster   string    `yaml:"cluster,omitempty"`
	Node      string    `yaml:"node,omitempty"`
}

type Resources struct {
	Requests ResourceDetails `yaml:"requests,omitempty"`
	Limits   ResourceDetails `yaml:"limits,omitempty"`
}

type ResourceDetails struct {
	CPU              string `yaml:"cpu,omitempty"`
	Memory           string `yaml:"memory,omitempty"`
	GPU              string `yaml:"gpu,omitempty"`
	EphemeralStorage string `yaml:"ephemeral-storage,omitempty"`
}
