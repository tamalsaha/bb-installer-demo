/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	monitoring "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindIngressNginx = "IngressNginx"
	ResourceIngressNginx     = "ingressnginx"
	ResourceIngressNginxs    = "ingressnginxs"
)

// IngressNginx defines the schama for IngressNginx Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=ingressnginxs,singular=ingressnginx,categories={kubeops,appscode}
type IngressNginx struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IngressNginxSpec `json:"spec,omitempty"`
}

// IngressNginxSpec is the schema for IngressNginx Operator values file
type IngressNginxSpec struct {
	CommonLabels map[string]string `json:"commonLabels"`
	Controller   struct {
		Name  string `json:"name"`
		Image struct {
			Chroot                   bool   `json:"chroot"`
			Registry                 string `json:"registry"`
			Image                    string `json:"image"`
			Tag                      string `json:"tag"`
			Digest                   string `json:"digest"`
			DigestChroot             string `json:"digestChroot"`
			PullPolicy               string `json:"pullPolicy"`
			RunAsUser                int    `json:"runAsUser"`
			AllowPrivilegeEscalation bool   `json:"allowPrivilegeEscalation"`
		} `json:"image"`
		ExistingPsp   string `json:"existingPsp"`
		ContainerName string `json:"containerName"`
		ContainerPort struct {
			Http  int `json:"http"`
			Https int `json:"https"`
		} `json:"containerPort"`
		Config                   struct{}          `json:"config"`
		ConfigAnnotations        map[string]string `json:"configAnnotations"`
		ProxySetHeaders          struct{}          `json:"proxySetHeaders"`
		AddHeaders               struct{}          `json:"addHeaders"`
		DnsConfig                struct{}          `json:"dnsConfig"`
		Hostname                 struct{}          `json:"hostname"`
		DnsPolicy                string            `json:"dnsPolicy"`
		ReportNodeInternalIp     bool              `json:"reportNodeInternalIp"`
		WatchIngressWithoutClass bool              `json:"watchIngressWithoutClass"`
		IngressClassByName       bool              `json:"ingressClassByName"`
		AllowSnippetAnnotations  bool              `json:"allowSnippetAnnotations"`
		HostNetwork              bool              `json:"hostNetwork"`
		HostPort                 struct {
			Enabled bool `json:"enabled"`
			Ports   struct {
				Http  int `json:"http"`
				Https int `json:"https"`
			} `json:"ports"`
		} `json:"hostPort"`
		ElectionID           string `json:"electionID"`
		IngressClassResource struct {
			Name            string   `json:"name"`
			Enabled         bool     `json:"enabled"`
			Default         bool     `json:"default"`
			ControllerValue string   `json:"controllerValue"`
			Parameters      struct{} `json:"parameters"`
		} `json:"ingressClassResource"`
		IngressClass       string                   `json:"ingressClass"`
		PodLabels          map[string]string        `json:"podLabels"`
		PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
		Sysctls            struct{}                 `json:"sysctls"`
		PublishService     struct {
			Enabled      bool   `json:"enabled"`
			PathOverride string `json:"pathOverride"`
		} `json:"publishService"`
		Scope struct {
			Enabled           bool   `json:"enabled"`
			Namespace         string `json:"namespace"`
			NamespaceSelector string `json:"namespaceSelector"`
		} `json:"scope"`
		ConfigMapNamespace string `json:"configMapNamespace"`
		Tcp                struct {
			ConfigMapNamespace string            `json:"configMapNamespace"`
			Annotations        map[string]string `json:"annotations"`
		} `json:"tcp"`
		Udp struct {
			ConfigMapNamespace string            `json:"configMapNamespace"`
			Annotations        map[string]string `json:"annotations"`
		} `json:"udp"`
		MaxmindLicenseKey             string                          `json:"maxmindLicenseKey"`
		ExtraArgs                     struct{}                        `json:"extraArgs"`
		ExtraEnvs                     []core.EnvVar                   `json:"extraEnvs"`
		Kind                          string                          `json:"kind"`
		Annotations                   map[string]string               `json:"annotations"`
		Labels                        map[string]string               `json:"labels"`
		UpdateStrategy                struct{}                        `json:"updateStrategy"`
		MinReadySeconds               int                             `json:"minReadySeconds"`
		Tolerations                   []core.Toleration               `json:"tolerations"`
		Affinity                      *core.Affinity                  `json:"affinity"`
		TopologySpreadConstraints     []core.TopologySpreadConstraint `json:"topologySpreadConstraints"`
		TerminationGracePeriodSeconds *int64                          `json:"terminationGracePeriodSeconds"`
		NodeSelector                  map[string]string               `json:"nodeSelector"`
		LivenessProbe                 *core.Probe                     `json:"livenessProbe"`
		ReadinessProbe                *core.Probe                     `json:"readinessProbe"`
		HealthCheckPath               string                          `json:"healthCheckPath"`
		HealthCheckHost               string                          `json:"healthCheckHost"`
		PodAnnotations                map[string]string               `json:"podAnnotations"`
		ReplicaCount                  int                             `json:"replicaCount"`
		MinAvailable                  int                             `json:"minAvailable"`
		Resources                     core.ResourceRequirements       `json:"resources"`
		Autoscaling                   struct {
			Enabled                           bool     `json:"enabled"`
			MinReplicas                       int      `json:"minReplicas"`
			MaxReplicas                       int      `json:"maxReplicas"`
			TargetCPUUtilizationPercentage    int      `json:"targetCPUUtilizationPercentage"`
			TargetMemoryUtilizationPercentage int      `json:"targetMemoryUtilizationPercentage"`
			Behavior                          struct{} `json:"behavior"`
		} `json:"autoscaling"`
		AutoscalingTemplate []interface{} `json:"autoscalingTemplate"`
		Keda                struct {
			ApiVersion                    string `json:"apiVersion"`
			Enabled                       bool   `json:"enabled"`
			MinReplicas                   int    `json:"minReplicas"`
			MaxReplicas                   int    `json:"maxReplicas"`
			PollingInterval               int    `json:"pollingInterval"`
			CooldownPeriod                int    `json:"cooldownPeriod"`
			RestoreToOriginalReplicaCount bool   `json:"restoreToOriginalReplicaCount"`
			ScaledObject                  struct {
				Annotations map[string]string `json:"annotations"`
			} `json:"scaledObject"`
			Triggers []interface{} `json:"triggers"`
			Behavior struct{}      `json:"behavior"`
		} `json:"keda"`
		EnableMimalloc bool `json:"enableMimalloc"`
		CustomTemplate struct {
			ConfigMapName string `json:"configMapName"`
			ConfigMapKey  string `json:"configMapKey"`
		} `json:"customTemplate"`
		Service struct {
			Enabled                  bool              `json:"enabled"`
			AppProtocol              bool              `json:"appProtocol"`
			Annotations              map[string]string `json:"annotations"`
			Labels                   map[string]string `json:"labels"`
			ExternalIPs              []string          `json:"externalIPs"`
			LoadBalancerIP           string            `json:"loadBalancerIP"`
			LoadBalancerSourceRanges []string          `json:"loadBalancerSourceRanges"`
			EnableHttp               bool              `json:"enableHttp"`
			EnableHttps              bool              `json:"enableHttps"`
			IpFamilyPolicy           string            `json:"ipFamilyPolicy"`
			IpFamilies               []string          `json:"ipFamilies"`
			Ports                    struct {
				Http  int `json:"http"`
				Https int `json:"https"`
			} `json:"ports"`
			TargetPorts struct {
				Http  string `json:"http"`
				Https string `json:"https"`
			} `json:"targetPorts"`
			Type      string `json:"type"`
			NodePorts struct {
				Http  string   `json:"http"`
				Https string   `json:"https"`
				Tcp   struct{} `json:"tcp"`
				Udp   struct{} `json:"udp"`
			} `json:"nodePorts"`
			External struct {
				Enabled bool `json:"enabled"`
			} `json:"external"`
			Internal struct {
				Enabled                  bool              `json:"enabled"`
				Annotations              map[string]string `json:"annotations"`
				LoadBalancerSourceRanges []string          `json:"loadBalancerSourceRanges"`
			} `json:"internal"`
		} `json:"service"`
		ShareProcessNamespace bool                                    `json:"shareProcessNamespace"`
		ExtraContainers       []core.Container                        `json:"extraContainers"`
		ExtraVolumeMounts     []core.VolumeMount                      `json:"extraVolumeMounts"`
		ExtraVolumes          []core.Volume                           `json:"extraVolumes"`
		ExtraInitContainers   []core.Container                        `json:"extraInitContainers"`
		ExtraModules          []NginxModule                           `json:"extraModules"`
		AdmissionWebhooks     IngressNginxControllerAdmissionWebhooks `json:"admissionWebhooks"`
		Metrics               IngressNginxControllerMetrics           `json:"metrics"`
		Lifecycle             *core.Probe                             `json:"lifecycle"`
		PriorityClassName     string                                  `json:"priorityClassName"`
	} `json:"controller"`
	RevisionHistoryLimit int                               `json:"revisionHistoryLimit"`
	DefaultBackend       IngressNginxDefaultBackend        `json:"defaultBackend"`
	Rbac                 IngressNginxRbacSpec              `json:"rbac"`
	PodSecurityPolicy    IngressNginxPodSecurityPolicySpec `json:"podSecurityPolicy"`
	ServiceAccount       IngressNginxServiceAccountSpec    `json:"serviceAccount"`
	ImagePullSecrets     []string                          `json:"imagePullSecrets"`
	Tcp                  map[string]string                 `json:"tcp"`
	Udp                  map[string]string                 `json:"udp"`
	PortNamePrefix       string                            `json:"portNamePrefix"`
	DhParam              *string                           `json:"dhParam"`
}

type NginxModule struct {
	Name  string
	Image string
}

type IngressNginxControllerAdmissionWebhooks struct {
	Annotations       map[string]string                                      `json:"annotations"`
	Enabled           bool                                                   `json:"enabled"`
	FailurePolicy     string                                                 `json:"failurePolicy"`
	Port              int                                                    `json:"port"`
	Certificate       string                                                 `json:"certificate"`
	Key               string                                                 `json:"key"`
	NamespaceSelector map[string]string                                      `json:"namespaceSelector"`
	ObjectSelector    map[string]string                                      `json:"objectSelector"`
	Labels            map[string]string                                      `json:"labels"`
	ExistingPsp       string                                                 `json:"existingPsp"`
	Service           IngressNginxControllerAdmissionWebhooksService         `json:"service"`
	CreateSecretJob   IngressNginxControllerAdmissionWebhooksCreateSecretJob `json:"createSecretJob"`
	PatchWebhookJob   IngressNginxControllerAdmissionWebhooksPatchWebhookJob `json:"patchWebhookJob"`
	Patch             IngressNginxControllerAdmissionWebhooksPatch           `json:"patch"`
}

type IngressNginxControllerAdmissionWebhooksService struct {
	Annotations              map[string]string `json:"annotations"`
	ExternalIPs              []string          `json:"externalIPs"`
	LoadBalancerSourceRanges []string          `json:"loadBalancerSourceRanges"`
	ServicePort              int               `json:"servicePort"`
	Type                     string            `json:"type"`
}

type IngressNginxControllerAdmissionWebhooksCreateSecretJob struct {
	Resources core.ResourceRequirements `json:"resources"`
}

type IngressNginxControllerAdmissionWebhooksPatchWebhookJob struct {
	Resources core.ResourceRequirements `json:"resources"`
}

type IngressNginxControllerAdmissionWebhooksPatch struct {
	Enabled           bool                                         `json:"enabled"`
	Image             IngressNginxControllerAdmissionWebhooksImage `json:"image"`
	PriorityClassName string                                       `json:"priorityClassName"`
	PodAnnotations    map[string]string                            `json:"podAnnotations"`
	NodeSelector      map[string]string                            `json:"nodeSelector"`
	Tolerations       []core.Toleration                            `json:"tolerations"`
	Labels            map[string]string                            `json:"labels"`
	RunAsUser         int                                          `json:"runAsUser"`
	FsGroup           int                                          `json:"fsGroup"`
}
type IngressNginxControllerAdmissionWebhooksImage struct {
	Registry   string `json:"registry"`
	Image      string `json:"image"`
	Tag        string `json:"tag"`
	Digest     string `json:"digest"`
	PullPolicy string `json:"pullPolicy"`
}

type IngressNginxControllerMetrics struct {
	Port           int                                         `json:"port"`
	Enabled        bool                                        `json:"enabled"`
	Service        IngressNginxControllerMetricsService        `json:"service"`
	ServiceMonitor IngressNginxControllerMetricsServiceMonitor `json:"serviceMonitor"`
	PrometheusRule IngressNginxControllerMetricsPrometheusRule `json:"prometheusRule"`
}

type IngressNginxControllerMetricsService struct {
	Annotations              map[string]string `json:"annotations"`
	ExternalIPs              []string          `json:"externalIPs"`
	LoadBalancerSourceRanges []string          `json:"loadBalancerSourceRanges"`
	ServicePort              int               `json:"servicePort"`
	Type                     string            `json:"type"`
}

type IngressNginxControllerMetricsServiceMonitor struct {
	Enabled           bool                        `json:"enabled"`
	AdditionalLabels  map[string]string           `json:"additionalLabels"`
	Namespace         string                      `json:"namespace"`
	NamespaceSelector map[string]string           `json:"namespaceSelector"`
	ScrapeInterval    string                      `json:"scrapeInterval"`
	TargetLabels      []string                    `json:"targetLabels"`
	Relabelings       []*monitoring.RelabelConfig `json:"relabelings"`
	MetricRelabelings []*monitoring.RelabelConfig `json:"metricRelabelings"`
}

type IngressNginxControllerMetricsPrometheusRule struct {
	Enabled          bool              `json:"enabled"`
	AdditionalLabels map[string]string `json:"additionalLabels"`
	Rules            []monitoring.Rule `json:"rules"`
}

type IngressNginxDefaultBackend struct {
	Enabled                  bool                                         `json:"enabled"`
	Name                     string                                       `json:"name"`
	Image                    IngressNginxDefaultBackendImageReference     `json:"image"`
	ExistingPsp              string                                       `json:"existingPsp"`
	ExtraArgs                map[string]string                            `json:"extraArgs"`
	ServiceAccount           IngressNginxDefaultBackendServiceAccountSpec `json:"serviceAccount"`
	ExtraEnvs                []core.EnvVar                                `json:"extraEnvs"`
	Port                     int                                          `json:"port"`
	LivenessProbe            *core.Probe                                  `json:"livenessProbe"`
	ReadinessProbe           *core.Probe                                  `json:"readinessProbe"`
	Tolerations              []core.Toleration                            `json:"tolerations"`
	Affinity                 *core.Affinity                               `json:"affinity"`
	PodSecurityContext       *core.PodSecurityContext                     `json:"podSecurityContext"`
	ContainerSecurityContext *core.SecurityContext                        `json:"containerSecurityContext"`
	PodLabels                map[string]string                            `json:"podLabels"`
	NodeSelector             map[string]string                            `json:"nodeSelector"`
	PodAnnotations           map[string]string                            `json:"podAnnotations"`
	ReplicaCount             int                                          `json:"replicaCount"`
	MinAvailable             int                                          `json:"minAvailable"`
	Resources                core.ResourceRequirements                    `json:"resources"`
	ExtraVolumeMounts        []core.VolumeMount                           `json:"extraVolumeMounts"`
	ExtraVolumes             []core.Volume                                `json:"extraVolumes"`
	Autoscaling              IngressNginxAutoscalingSpec                  `json:"autoscaling"`
	Service                  IngressNginxServiceSpec                      `json:"service"`
	PriorityClassName        string                                       `json:"priorityClassName"`
	Labels                   map[string]string                            `json:"labels"`
}

type IngressNginxDefaultBackendImageReference struct {
	Registry                 string `json:"registry"`
	Image                    string `json:"image"`
	Tag                      string `json:"tag"`
	PullPolicy               string `json:"pullPolicy"`
	RunAsUser                int    `json:"runAsUser"`
	RunAsNonRoot             bool   `json:"runAsNonRoot"`
	ReadOnlyRootFilesystem   bool   `json:"readOnlyRootFilesystem"`
	AllowPrivilegeEscalation bool   `json:"allowPrivilegeEscalation"`
}

type IngressNginxDefaultBackendServiceAccountSpec struct {
	Create                       bool   `json:"create"`
	Name                         string `json:"name"`
	AutomountServiceAccountToken bool   `json:"automountServiceAccountToken"`
}

type IngressNginxAutoscalingSpec struct {
	Annotations                       map[string]string `json:"annotations"`
	Enabled                           bool              `json:"enabled"`
	MinReplicas                       int               `json:"minReplicas"`
	MaxReplicas                       int               `json:"maxReplicas"`
	TargetCPUUtilizationPercentage    int               `json:"targetCPUUtilizationPercentage"`
	TargetMemoryUtilizationPercentage int               `json:"targetMemoryUtilizationPercentage"`
}

type IngressNginxServiceSpec struct {
	Annotations              map[string]string `json:"annotations"`
	ExternalIPs              []string          `json:"externalIPs"`
	LoadBalancerSourceRanges []string          `json:"loadBalancerSourceRanges"`
	ServicePort              int               `json:"servicePort"`
	Type                     string            `json:"type"`
}

type IngressNginxRbacSpec struct {
	Create bool `json:"create"`
	Scope  bool `json:"scope"`
}

type IngressNginxPodSecurityPolicySpec struct {
	Enabled bool `json:"enabled"`
}

type IngressNginxServiceAccountSpec struct {
	Create                       bool              `json:"create"`
	Name                         string            `json:"name"`
	AutomountServiceAccountToken bool              `json:"automountServiceAccountToken"`
	Annotations                  map[string]string `json:"annotations"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IngressNginxList is a list of IngressNginxs
type IngressNginxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IngressNginx CRD objects
	Items []IngressNginx `json:"items,omitempty"`
}
