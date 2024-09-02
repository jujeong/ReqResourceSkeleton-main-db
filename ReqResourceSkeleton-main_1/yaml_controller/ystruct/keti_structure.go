package ystruct

type Workflow struct {
	APIVersion string   `yaml:"apiVersion,omitempty"`
	Kind       string   `yaml:"kind,omitempty"`
	Metadata   Metadata `yaml:"metadata,omitempty"`
	Spec       Spec     `yaml:"spec,omitempty"`
}

type Metadata struct {
	GenerateName string            `yaml:"generateName,omitempty"`
	Annotations  map[string]string `yaml:"annotations,omitempty"`
	Labels       map[string]string `yaml:"labels,omitempty"`
}

type Spec struct {
	Entrypoint         string     `yaml:"entrypoint,omitempty"`
	Templates          []Template `yaml:"templates,omitempty"`
	Arguments          Arguments  `yaml:"arguments,omitempty"`
	ServiceAccountName string     `yaml:"serviceAccountName,omitempty"`
}

type Template struct {
	Name         string     `yaml:"name,omitempty"`
	Container    *Container `yaml:"container,omitempty"`
	Metadata     *Metadata  `yaml:"metadata,omitempty"`
	NodeSelector NodeSelect `yaml:"nodeSelector,omitempty"`
	DAG          *DAG       `yaml:"dag,omitempty"`
}

type ContainerResources struct {
	Limits   map[string]string `yaml:"limits,omitempty"`
	Requests map[string]string `yaml:"requests,omitempty"`
}

type DAG struct {
	Tasks []Task `yaml:"tasks,omitempty"`
}

type Task struct {
	Name         string   `yaml:"name,omitempty"`
	Template     string   `yaml:"template,omitempty"`
	Dependencies []string `yaml:"dependencies,omitempty"`
}

type Arguments struct {
	Parameters []interface{} `yaml:"parameters,omitempty"`
}

type NodeSelect struct {
	Node string `yaml:"kubernetes.io/hostname,omitempty"`
}
