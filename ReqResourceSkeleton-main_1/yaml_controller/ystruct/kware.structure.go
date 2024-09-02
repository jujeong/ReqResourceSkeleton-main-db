package ystruct

type ReqResource struct {
	Version string  `yaml:"version,omitempty"`
	Request Request `yaml:"request,omitempty"`
}

type Request struct {
	Name       string      `yaml:"name,omitempty"`
	ID         string      `yaml:"id,omitempty"`
	Date       string      `yaml:"date,omitempty"`
	Containers []Container `yaml:"containers,omitempty"`
	Attribute  Attribute   `yaml:"attribute,omitempty"`
}

type Attribute struct {
	WorkloadType           string  `yaml:"workloadType,omitempty"`
	IsCronJob              bool    `yaml:"isCronJob,omitempty"`
	DevOpsType             string  `yaml:"devOpsType,omitempty"`
	CudaVersion            float64 `yaml:"cudaVersion,omitempty"`
	GPUDriverVersion       float64 `yaml:"gpuDriverVersion,omitempty"`
	MaxReplicas            int     `yaml:"maxReplicas,omitempty"`
	IsNetworking           bool    `yaml:"isNetworking,omitempty"`
	TotalSize              int     `yaml:"totalSize,omitempty"`
	PredictedExecutionTime int     `yaml:"predictedExecutionTime,omitempty"`
	UserID                 string  `yaml:"userId,omitempty"`
	Yaml                   string  `yaml:"yaml,omitempty"`
}

type RespResource struct {
	Response Response `yaml:"response,omitempty"`
}

type Response struct {
	ID               string      `yaml:"id,omitempty"`
	Date             string      `yaml:"date,omitempty"`
	Container        []Container `yaml:"container,omitempty"`
	PriorityClass    string      `yaml:"priorityClass,omitempty"`
	Priority         string      `yaml:"priority,omitempty"`
	PreemptionPolicy string      `yaml:"preemptionPolicy,omitempty"`
}

type Result struct {
	Cluster          string `yaml:"cluster,omitempty"`
	Node             string `yaml:"node,omitempty"`
	PriorityClass    string `yaml:"priorityClass,omitempty"`
	Priority         string `yaml:"priority,omitempty"`
	PreemptionPolicy string `yaml:"preemptionPolicy,omitempty"`
}
