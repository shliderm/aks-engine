package api

// the orchestrators supported by vlabs
const (
	// Mesos is the string constant for MESOS orchestrator type
	Mesos OrchestratorType = "Mesos"
	// DCOS is the string constant for DCOS orchestrator type and defaults to DCOS188
	DCOS OrchestratorType = "DCOS"
	// DCOS190 is the string constant for DCOS 1.9.0 orchestrator type
	DCOS190 OrchestratorType = "DCOS190"
	// DCOS188 is the string constant for DCOS 1.8.8 orchestrator type
	DCOS188 OrchestratorType = "DCOS188"
	// DCOS187 is the string constant for DCOS 1.8.7 orchestrator type
	DCOS187 OrchestratorType = "DCOS187"
	// DCOS184 is the string constant for DCOS 1.8.4 orchestrator type
	DCOS184 OrchestratorType = "DCOS184"
	// DCOS173 is the string constant for DCOS 1.7.3 orchestrator type
	DCOS173 OrchestratorType = "DCOS173"
	// Swarm is the string constant for the Swarm orchestrator type
	Swarm OrchestratorType = "Swarm"
	// Kubernetes is the string constant for the Kubernetes orchestrator type
	Kubernetes OrchestratorType = "Kubernetes"
	// SwarmMode is the string constant for the Swarm Mode orchestrator type
	SwarmMode OrchestratorType = "SwarmMode"
)

// the OSTypes supported by vlabs
const (
	Windows OSType = "Windows"
	Linux   OSType = "Linux"
)

// validation values
const (
	// MinAgentCount are the minimum number of agents per agent pool
	MinAgentCount = 1
	// MaxAgentCount are the maximum number of agents per agent pool
	MaxAgentCount = 100
	// MinPort specifies the minimum tcp port to open
	MinPort = 1
	// MaxPort specifies the maximum tcp port to open
	MaxPort = 65535
	// MaxDisks specifies the maximum attached disks to add to the cluster
	MaxDisks = 4
)

// Availability profiles
const (
	// AvailabilitySet means that the vms are in an availability set
	AvailabilitySet = "AvailabilitySet"
	// VirtualMachineScaleSets means that the vms are in a virtual machine scaleset
	VirtualMachineScaleSets = "VirtualMachineScaleSets"
)

// storage profiles
const (
	// StorageAccount means that the nodes use raw storage accounts for their os and attached volumes
	StorageAccount = "StorageAccount"
	// ManagedDisks means that the nodes use managed disks for their os and attached volumes
	ManagedDisks = "ManagedDisks"
)

// KubernetesVersion defines supported Kubernetes versions
type KubernetesVersion string

// DCOSVersion defines supported Kubernetes versions
type DCOSVersion string

const (
	// Kubernetes153 is the string constant for Kubernetes 1.5.3
	Kubernetes153 OrchestratorVersion = "1.5.3"
	// Kubernetes160 is the string constant for Kubernetes 1.6.0
	Kubernetes160 OrchestratorVersion = "1.6.0"
	// Kubernetes162 is the string constant for Kubernetes 1.6.2
	Kubernetes162 OrchestratorVersion = "1.6.2"
	// KubernetesLatest is the string constant for latest Kubernetes version
	KubernetesLatest OrchestratorVersion = Kubernetes162
)

const (
	// DCOS190Version is the string constant for DCOS 1.9.0
	DCOS190Version OrchestratorVersion = "1.9.0"
	// DCOS188Version is the string constant for DCOS 1.8.8
	DCOS188Version OrchestratorVersion = "1.8.8"
	// DCOS187Version is the string constant for DCOS 1.8.7
	DCOS187Version OrchestratorVersion = "1.8.7"
	// DCOS184Version is the string constant for DCOS 1.8.4
	DCOS184Version OrchestratorVersion = "1.8.4"
	// DCOS173Version is the string constant for DCOS 1.7.3
	DCOS173Version OrchestratorVersion = "1.7.3"
)
