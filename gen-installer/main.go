package main

import (
	"fmt"
	"net/url"
	"os"
	"path"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/labels"

	api "go.bytebuilders.dev/installer/apis/installer/v1alpha1"
	passgen "gomodules.xyz/password-generator"
	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"
)

// Steps to do manually
// buckets to create
// DNS record configure

type AceOptionsSpec struct {
	Release       types.NamespacedName `json:"release"`
	Hosted        bool                 `json:"hosted"`
	Billing       ComponentSpec        `json:"billing"`
	PlatformUi    ComponentSpec        `json:"platform-ui"`
	AccountsUi    ComponentSpec        `json:"accounts-ui"`
	ClusterUi     ComponentSpec        `json:"cluster-ui"`
	DeployUi      ComponentSpec        `json:"deploy-ui"`
	Grafana       ComponentSpec        `json:"grafana"`
	KubedbUi      ComponentSpec        `json:"kubedb-ui"`
	MarketplaceUi ComponentSpec        `json:"marketplace-ui"`
	PlatformApi   ComponentSpec        `json:"platform-api"`
	PromProxy     ComponentSpec        `json:"prom-proxy"`
	Ingress       IngressNginx         `json:"ingress"`
	Nats          NatsSettings         `json:"nats"`
	Global        AceGlobalValues      `json:"global"`
	Settings      Settings             `json:"settings"`
}

type ComponentSpec struct {
	Enabled bool `json:"enabled"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
}

// +kubebuilder:validation:Enum=LoadBalancer;HostPort
type ServiceType string

const (
	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
	ServiceTypeHostPort     ServiceType = "HostPort"
)

const (
	DefaultPasswordLength = 16
)

type IngressNginx struct {
	ExposeVia ServiceType `json:"exposeVia"`
	//+optional
	Resources    core.ResourceRequirements `json:"resources"`
	NodeSelector map[string]string         `json:"nodeSelector"`
}

type NatsSettings struct {
	ExposeVia ServiceType `json:"exposeVia"`
	Replics   int         `json:"replicas"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`

	// ShardCount int `json:"shardCount"`
	// MountPath       string `json:"mountPath"`
	// OperatorCreds   string `json:"operatorCreds"`
	// OperatorJwt     string `json:"operatorJwt"`
	// SystemCreds     string `json:"systemCreds"`
	// SystemJwt       string `json:"systemJwt"`
	// SystemPubKey    string `json:"systemPubKey"`
	// SystemUserCreds string `json:"systemUserCreds"`
	// AdminCreds      string `json:"adminCreds"`
	// AdminUserCreds  string `json:"adminUserCreds"`
}

type AceGlobalValues struct {
	License          string               `json:"license"`
	Registry         string               `json:"registry"`
	RegistryFQDN     string               `json:"registryFQDN"`
	ImagePullSecrets []string             `json:"imagePullSecrets"`
	Monitoring       api.GlobalMonitoring `json:"monitoring"`
	Infra            PlatformInfra        `json:"infra"`
}

type PlatformInfra struct {
	StorageClass api.LocalObjectReference `json:"storageClass"`
	TLS          InfraTLS                 `json:"tls"`
	DNS          api.InfraDns             `json:"dns"`
	Objstore     InfraObjstore            `json:"objstore"`
	Kms          InfraKms                 `json:"kms"`
	// Kubepack     InfraKubepack            `json:"kubepack"`
	// Badger       InfraBadger              `json:"badger"`
	// Invoice      InfraInvoice             `json:"invoice"`
}

type InfraTLS struct {
	// AcmeServer string `json:"acmeServer"`
	Email string `json:"email"`
}

type DNSProviderAuth struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type InfraObjstore struct {
	Bucket   string `json:"bucket"`
	Provider string `json:"provider"`
	// MountPath string       `json:"mountPath"`
	Auth api.ObjstoreAuth `json:"auth"`
}

type InfraKms struct {
	Provider string `json:"provider"`
	// MountPath    string `json:"mountPath"`
	MasterKeyURL string `json:"masterKeyURL"`
}

type Settings struct {
	DB       DBSettings       `json:"db"`
	Cache    CacheSettings    `json:"cache"`
	Smtp     SmtpSettings     `json:"smtp"`
	Platform PlatformSettings `json:"platform"`
	// Security api.SecuritySettings `json:"security"`
}

type DBSettings struct {
	Persistence api.PersistenceSpec       `json:"persistence"`
	Resources   core.ResourceRequirements `json:"resources"`
}

type CacheSettings struct {
	Persistence api.PersistenceSpec       `json:"persistence"`
	Resources   core.ResourceRequirements `json:"resources"`
}

type SmtpSettings struct {
	Host       string `json:"host"`
	TlsEnabled bool   `json:"tlsEnabled"`
	// From            string `json:"from"`
	Username string `json:"username"`
	Password string `json:"password"`
	// SubjectPrefix   string `json:"subjectPrefix"`
	SendAsPlainText bool `json:"sendAsPlainText"`
}

type PlatformSettings struct {
	Domain string `json:"domain"`
	// AppName                         string  `json:"appName"`
	// RunMode                         string  `json:"runMode"`
	// ExperimentalFeatures            bool    `json:"experimentalFeatures"`
	// ForcePrivate                    bool    `json:"forcePrivate"`
	// DisableHttpGit                  bool    `json:"disableHttpGit"`
	// InstallLock                     bool    `json:"installLock"`
	// RepositoryUploadEnabled         bool    `json:"repositoryUploadEnabled"`
	// RepositoryUploadAllowedTypes    *string `json:"repositoryUploadAllowedTypes"`
	// RepositoryUploadMaxFileSize     int     `json:"repositoryUploadMaxFileSize"`
	// RepositoryUploadMaxFiles        int     `json:"repositoryUploadMaxFiles"`
	// ServiceEnableCaptcha            bool    `json:"serviceEnableCaptcha"`
	// ServiceRegisterEmailConfirm     bool    `json:"serviceRegisterEmailConfirm"`
	// ServiceDisableRegistration      bool    `json:"serviceDisableRegistration"`
	// ServiceRequireSignInView        bool    `json:"serviceRequireSignInView"`
	// ServiceEnableNotifyMail         bool    `json:"serviceEnableNotifyMail"`
	// SessionProvider                 string  `json:"sessionProvider"`
	// SessionProviderConfig           *string `json:"sessionProviderConfig"`
	// CookieName                      string  `json:"cookieName"`
	// ServerLandingPage               string  `json:"serverLandingPage"`
	// LogMode                         string  `json:"logMode"`
	// LogLevel                        string  `json:"logLevel"`
	// OtherShowFooterBranding         bool    `json:"otherShowFooterBranding"`
	// OtherShowFooterVersion          bool    `json:"otherShowFooterVersion"`
	// OtherShowFooterTemplateLoadTime bool    `json:"otherShowFooterTemplateLoadTime"`
	// EnableCSRFCookieHttpOnly        bool    `json:"enableCSRFCookieHttpOnly"`
}

func main() {
	var v api.AceSpec
	data, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func NewOptions() *AceOptionsSpec {
	hosted := false
	return &AceOptionsSpec{
		Release: types.NamespacedName{
			Name:      "ace",
			Namespace: "ace",
		},
		Hosted: hosted,
		Billing: ComponentSpec{
			Enabled: hosted,
		},
		PlatformUi: ComponentSpec{
			Enabled: true,
		},
		AccountsUi: ComponentSpec{
			Enabled: true,
		},
		ClusterUi: ComponentSpec{
			Enabled: true,
		},
		DeployUi: ComponentSpec{
			Enabled: hosted,
		},
		Grafana: ComponentSpec{
			Enabled: true,
		},
		KubedbUi: ComponentSpec{
			Enabled: true,
		},
		MarketplaceUi: ComponentSpec{
			Enabled: hosted,
		},
		PlatformApi: ComponentSpec{
			Enabled: true,
		},
		PromProxy: ComponentSpec{
			Enabled: true,
		},
		Ingress: IngressNginx{
			ExposeVia: ServiceTypeLoadBalancer,
			// Resources:    core.ResourceRequirements{},
			// NodeSelector: nil,
		},
		Nats: NatsSettings{
			ExposeVia: ServiceTypeLoadBalancer,
			Replics:   1,
			//Resources:    core.ResourceRequirements{
			//	Limits:   nil,
			//	Requests: nil,
			//},
			//NodeSelector: nil,
		},
		Global: AceGlobalValues{
			License:          "",
			Registry:         "",
			RegistryFQDN:     "",
			ImagePullSecrets: nil,
			Monitoring:       api.GlobalMonitoring{},
			Infra: PlatformInfra{
				StorageClass: api.LocalObjectReference{
					Name: "standard",
				},
				//TLS: InfraTLS{
				//	Email: "",
				//},
				//DNS: InfraDns{
				//	Provider: "",
				//	Auth:     DNSProviderAuth{},
				//},
				//Objstore: InfraObjstore{
				//	Provider: "",
				//	Auth:     ObjstoreAuth{},
				//},
				//Kms:     InfraKms{
				//	Provider:     "",
				//	MasterKeyURL: "",
				//},
				//Avatars: InfraAvatars{
				//	Bucket:
				//},
			},
		},
		Settings: Settings{
			DB: DBSettings{
				Persistence: api.PersistenceSpec{
					Size: resource.MustParse("20Gi"),
				},
				Resources: core.ResourceRequirements{
					Limits: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
					Requests: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
				},
			},
			Cache: CacheSettings{
				Persistence: api.PersistenceSpec{
					Size: resource.MustParse("10Gi"),
				},
				Resources: core.ResourceRequirements{
					Limits: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
					Requests: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
				},
			},
			Smtp:     SmtpSettings{},
			Platform: PlatformSettings{},
			//Security: api.SecuritySettings{
			//	Oauth2JWTSecret: "",
			//	CsrfSecretKey:   "",
			//},
		},
	}
}

func InitComponents(in *AceOptionsSpec, out *api.AceSpec) error {
	out.Reloader = api.AceReloader{
		Enabled: true,
	}

	if in.Billing.Enabled {
		out.Billing = api.AceBilling{
			Enabled: true,
			BillingSpec: &api.BillingSpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.Billing.Resources,
				NodeSelector: in.Billing.NodeSelector,
			},
		}
	}
	if in.PlatformUi.Enabled {
		out.PlatformUi = api.AcePlatformUi{
			Enabled: true,
			PlatformUiSpec: &api.PlatformUiSpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.PlatformUi.Resources,
				NodeSelector: in.PlatformUi.NodeSelector,
			},
		}
	}
	if in.AccountsUi.Enabled {
		out.AccountsUi = api.AceAccountsUi{
			Enabled: true,
			AccountsUiSpec: &api.AccountsUiSpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.AccountsUi.Resources,
				NodeSelector: in.AccountsUi.NodeSelector,
			},
		}
	}
	if in.ClusterUi.Enabled {
		out.ClusterUi = api.AceClusterUi{
			Enabled: true,
			ClusterUiSpec: &api.ClusterUiSpec{
				//PodAnnotations: map[string]string{
				//	"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				//},
				Resources:    in.ClusterUi.Resources,
				NodeSelector: in.ClusterUi.NodeSelector,
			},
		}
	}
	if in.DeployUi.Enabled {
		out.DeployUi = api.AceDeployUi{
			Enabled: true,
			DeployUiSpec: &api.DeployUiSpec{
				//PodAnnotations: map[string]string{
				//	"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				//},
				Resources:    in.DeployUi.Resources,
				NodeSelector: in.DeployUi.NodeSelector,
			},
		}
	}
	if in.Grafana.Enabled {
		out.Grafana = api.AceGrafana{
			Enabled: true,
			GrafanaSpec: &api.GrafanaSpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.Grafana.Resources,
				NodeSelector: in.Grafana.NodeSelector,
			},
		}
	}
	if in.KubedbUi.Enabled {
		out.KubedbUi = api.AceKubedbUi{
			Enabled: true,
			KubedbUiSpec: &api.KubedbUiSpec{
				//PodAnnotations: map[string]string{
				//	"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				//},
				Resources:    in.KubedbUi.Resources,
				NodeSelector: in.KubedbUi.NodeSelector,
			},
		}
	}
	if in.MarketplaceUi.Enabled {
		out.MarketplaceUi = api.AceMarketplaceUi{
			Enabled: true,
			MarketplaceUiSpec: &api.MarketplaceUiSpec{
				//PodAnnotations: map[string]string{
				//	"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				//},
				Resources:    in.MarketplaceUi.Resources,
				NodeSelector: in.MarketplaceUi.NodeSelector,
			},
		}
	}
	if in.PlatformApi.Enabled {
		out.PlatformApi = api.AcePlatformApi{
			Enabled: true,
			PlatformApiSpec: &api.PlatformApiSpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.PlatformApi.Resources,
				NodeSelector: in.PlatformApi.NodeSelector,
			},
		}
	}
	if in.PromProxy.Enabled {
		out.PromProxy = api.AcePromProxy{
			Enabled: true,
			PromProxySpec: &api.PromProxySpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.PromProxy.Resources,
				NodeSelector: in.PromProxy.NodeSelector,
			},
		}
	}
	return nil
}

func GenerateIngress(in *AceOptionsSpec, out *api.AceSpec) error {
	if in.Ingress.ExposeVia == ServiceTypeLoadBalancer {
		out.IngressNginx = api.AceIngressNginx{
			Enabled: true,
			IngressNginxSpec: &api.IngressNginxSpec{
				Controller: api.IngressNginxController{
					IngressClassByName: true,
					IngressClassResource: api.IngressNginxControllerIngressClassResource{
						Enabled:         true,
						ControllerValue: fmt.Sprintf("k8s.io/ingress-nginx-%s", in.Release.Name),
						Name:            fmt.Sprintf("nginx-%s", in.Release.Name),
					},
					//HostPort: &api.IngressNginxControllerHostPort{
					//	Enabled: false,
					//},
					//Kind:         "",
					NodeSelector: in.Ingress.NodeSelector,
					Service: &api.IngressNginxControllerService{
						External: api.IngressNginxControllerServiceExternal{
							Enabled: true,
						},
					},
				},
			},
		}
	} else {
		out.IngressNginx = api.AceIngressNginx{
			Enabled: true,
			IngressNginxSpec: &api.IngressNginxSpec{
				Controller: api.IngressNginxController{
					IngressClassByName: true,
					IngressClassResource: api.IngressNginxControllerIngressClassResource{
						Enabled:         true,
						ControllerValue: fmt.Sprintf("k8s.io/ingress-nginx-%s", in.Release.Name),
						Name:            fmt.Sprintf("nginx-%s", in.Release.Name),
					},
					HostPort: &api.IngressNginxControllerHostPort{
						Enabled: true,
					},
					Kind:         "DaemonSet",
					NodeSelector: in.Ingress.NodeSelector,
					Service: &api.IngressNginxControllerService{
						External: api.IngressNginxControllerServiceExternal{
							Enabled: false,
						},
					},
				},
			},
		}
	}

	out.IngressDns = api.AceIngressDns{
		Enabled: true,
		ExternalDnsSpec: &api.ExternalDnsSpec{
			DomainFilters: []string{
				in.Settings.Platform.Domain,
			},
			// ref: https://github.com/kubernetes-sigs/external-dns/pull/2718
			Image: api.ExternalDnsImageReference{
				Repository: "appscode/external-dns",
				Tag:        "external-dns-helm-chart-1.9.0-1-gbd1bb40c",
				PullPolicy: "IfNotPresent",
			},
			LogLevel:   "debug",
			Sources:    []string{"ingress"},
			ExtraArgs:  []string{"--ignore-ingress-tls-spec"},
			Policy:     "sync",
			Registry:   "txt",
			TxtOwnerID: "ingress-dns",
		},
	}

	// TODO: Add additional DNS providers
	if in.Global.Infra.DNS.Provider == "cloudflare" {
		out.IngressDns.Provider = "cloudflare"
		out.IngressDns.Env = []core.EnvVar{
			{
				Name:  "CF_API_TOKEN",
				Value: in.Global.Infra.DNS.Auth.Token,
			},
		}
	}

	return nil
}

func GenerateNats(in *AceOptionsSpec, out *api.AceSpec) error {
	if in.Nats.Replics != 1 && in.Nats.Replics != 3 {
		return errors.Errorf("nats replicas can be 1 or 3, found %d", in.Nats.Replics)
	}

	if err := os.RemoveAll(confDir()); err != nil {
		return err
	}

	nc, err := genNatsCredentials()
	if err != nil {
		return err
	}

	out.Settings.Nats = api.NatsSettings{
		ShardCount:      128, // reduce to 32
		Replics:         in.Nats.Replics,
		MountPath:       "/nats/creds",
		OperatorCreds:   nc["Operator.creds"],
		OperatorJwt:     nc["Operator.jwt"],
		SystemCreds:     nc["SYS.creds"],
		SystemJwt:       nc["SYS.jwt"],
		SystemPubKey:    nc["SYS.pub"],
		SystemUserCreds: nc["sys.creds"],
		AdminCreds:      nc["ADMIN.creds"],
		AdminUserCreds:  nc["admin.creds"],
	}

	out.Nats = api.AceNats{
		Enabled: true,
		NatsSpec: &api.NatsSpec{
			NodeSelector: in.Nats.NodeSelector,
			StatefulSetPodLabels: map[string]string{
				"secret.reloader.stakater.com/reload": tplNATSTLSSecret(in),
			},
			Nats: api.NatsServerSpec{
				Advertise: false,
				// ExternalAccess: true, // true means HostPost
				Limits: api.NatsServerLimitsSpec{
					MaxPayload: pointer.StringP("4Mb"),
				},
				Logging: api.NatsLoggingSpec{
					Debug: pointer.FalseP(),
					Trace: pointer.FalseP(),
				},
				Jetstream: api.JetstreamSpec{
					Enabled: true,
					FileStorage: api.JetstreamFileStorage{
						Enabled:          true,
						StorageDirectory: "/nats/jetstream",
						Size:             resource.MustParse("10Gi"), // TODO: high?
						StorageClassName: in.Global.Infra.StorageClass.Name,
					},
				},
				Resources: core.ResourceRequirements{
					Limits: core.ResourceList{
						core.ResourceMemory: resource.MustParse("2Gi"),
					},
					Requests: core.ResourceList{
						core.ResourceMemory: resource.MustParse("2Gi"),
					},
				},
				TLS: &api.NatsServerTLSSpec{
					AllowNonTLS: false,
					Secret: api.LocalObjectReference{
						Name: tplNATSTLSSecret(in),
					},
					// Ca:          "",
					Cert: core.TLSCertKey,
					Key:  core.TLSPrivateKeyKey,
				},
			},
			Natsbox: api.NatsboxSpec{
				Enabled: false,
			},
			Exporter: api.NatsExporterSpec{
				Enabled: true,
				ServiceMonitor: api.NatsExporterServiceMonitorSpec{
					Enabled:   true,
					Namespace: "", // use nats namespace
					Labels:    in.Global.Monitoring.ServiceMonitor.Labels,
					Path:      "/metrics",
				},
			},
			// Affinity:  nil,
			// Cluster:   api.NatsClusterSpec{},
			Auth: api.NatsAuthSpec{
				Enabled: true,
				Operatorjwt: &api.NatsOperatorJWTSpec{
					ConfigMap: api.ConfigMapKeySelector{
						Name: tplNATSCredSecret(in),
						Key:  "Operator.jwt",
					},
				},
				SystemAccount: pointer.StringP(nc["SYS.pub"]), // account or user?
				Resolver: api.NatsResolverSpec{
					Type:          "full",
					Operator:      pointer.StringP(nc["Operator.jwt"]),
					SystemAccount: pointer.StringP(nc["SYS.pub"]), // account or user
					Store: api.NatsResolverStoreSpec{
						Dir:              "/etc/nats-config/accounts/jwt",
						Size:             resource.MustParse("10Gi"),
						StorageClassName: in.Global.Infra.StorageClass.Name,
					},
					ResolverPreload: map[string]string{
						nc["SYS.pub"]:   nc["SYS.jwt"],
						nc["ADMIN.pub"]: nc["ADMIN.jwt"], // TODO: skip?
					},
				},
			},
			Websocket: api.NatsWebsocketSpec{
				Enabled: true,
				Port:    443,
				AllowedOrigins: []string{
					fmt.Sprintf("https://%s", in.Settings.Platform.Domain),
					fmt.Sprintf("https://console.%s", in.Settings.Platform.Domain),
					fmt.Sprintf("https://kubedb.%s", in.Settings.Platform.Domain),
					fmt.Sprintf("https://grafana.%s", in.Settings.Platform.Domain),
				},
				TLS: &api.TLSSpec{
					Secret: api.LocalObjectReference{
						Name: tplNATSTLSSecret(in),
					},
					// Ca:          "",
					Cert: core.TLSCertKey,
					Key:  core.TLSPrivateKeyKey,
				},
			},
			UseFQDN: false,
		},
	}
	if in.Nats.Replics > 1 {
		natsPodSelector := &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app.kubernetes.io/instance": "nats-server",
				"app.kubernetes.io/name":     "nats",
			},
		}
		out.Nats.Affinity = &core.Affinity{
			PodAntiAffinity: &core.PodAntiAffinity{
				PreferredDuringSchedulingIgnoredDuringExecution: []core.WeightedPodAffinityTerm{
					// Prefer to not schedule multiple pods on the same node
					{
						Weight: 100,
						PodAffinityTerm: core.PodAffinityTerm{
							Namespaces:    []string{in.Release.Namespace},
							LabelSelector: natsPodSelector,
							TopologyKey:   core.LabelHostname,
						},
					},
					// Prefer to not schedule multiple pods on the node with same zone
					{
						Weight: 50,
						PodAffinityTerm: core.PodAffinityTerm{
							Namespaces:    []string{in.Release.Namespace},
							LabelSelector: natsPodSelector,
							TopologyKey:   core.LabelTopologyZone,
						},
					},
				},
			},
		}

		out.Nats.Cluster = api.NatsClusterSpec{
			Enabled:  true,
			Replicas: in.Nats.Replics,
			TLS: &api.TLSSpec{
				Secret: api.LocalObjectReference{
					Name: tplNATSTLSSecret(in),
				},
				// Ca:          "",
				Cert: core.TLSCertKey,
				Key:  core.TLSPrivateKeyKey,
			},
		}
	}

	if in.Nats.ExposeVia == ServiceTypeLoadBalancer {
		out.Nats.Nats.ExternalAccess = false
		// out.Nats.Websocket.Port = 9222

		// ingress TCP
		// expose NATS client port via TCP
		out.IngressNginx.IngressNginxSpec.TCP = map[string]string{
			"4222": fmt.Sprintf("%s/nats-server:4222", in.Release.Namespace),
		}
	} else {
		// out.Nats.Websocket.Port = 443
		out.Nats.Nats.ExternalAccess = true

		out.NatsDns = api.AceNatsDns{
			Enabled: true,
			ExternalDnsSpec: &api.ExternalDnsSpec{
				Sources: []string{"node"},
				Image: api.ExternalDnsImageReference{
					Repository: "appscode/external-dns",
					Tag:        "external-dns-helm-chart-1.9.0-1-gbd1bb40c",
					PullPolicy: "IfNotPresent",
				},
				DomainFilters: []string{in.Settings.Platform.Domain},
				LogLevel:      "debug",
				ExtraArgs: []string{
					fmt.Sprintf("--fqdn-template=nats.%s", in.Settings.Platform.Domain),
					fmt.Sprintf("--label-filter=%s", labels.Set(in.Nats.NodeSelector).String()),
				},
				Policy:     "sync",
				Registry:   "txt",
				TxtOwnerID: "nats-dns",
			},
		}

		// TODO: Add additional DNS providers
		if in.Global.Infra.DNS.Provider == "cloudflare" {
			out.NatsDns.Provider = "cloudflare"
			out.NatsDns.Env = []core.EnvVar{
				{
					Name:  "CF_API_TOKEN",
					Value: in.Global.Infra.DNS.Auth.Token,
				},
			}
		}
	}

	return nil
}

func GeneratePlatformValues(in *AceOptionsSpec, out *api.AceSpec) error {
	out.Global = api.AceGlobalValues{
		NameOverride: in.Release.Name,
		// FullnameOverride: "",
		// License:          "",
		// Registry:         "",
		// RegistryFQDN:     "",
		// ImagePullSecrets: nil,
		// ServiceAccount:   api.NatsServiceAccountSpec{},
		Monitoring: in.Global.Monitoring,
		Infra: api.PlatformInfra{
			StorageClass: in.Global.Infra.StorageClass,
			TLS: api.InfraTLS{
				// TODO: prod URL: https://acme-v02.api.letsencrypt.org/directory
				AcmeServer: "https://acme-staging-v02.api.letsencrypt.org/directory",
				Email:      in.Global.Infra.TLS.Email,
			},
			DNS: in.Global.Infra.DNS,
			Objstore: api.InfraObjstore{
				Provider:  in.Global.Infra.Objstore.Provider,
				MountPath: "/data/credentials",
				Auth:      in.Global.Infra.Objstore.Auth,
			},
			Kms: api.InfraKms{
				Provider:     in.Global.Infra.Objstore.Provider,
				MasterKeyURL: fmt.Sprintf("base64key://%s", passgen.GenerateForCharset(64, passgen.AlphaNum)),
			},
			Avatars: api.InfraAvatars{
				Bucket: mustBucketName(in.Global.Infra.Objstore.Bucket, "avatars"),
			},
			// TODO: bucket proxy
			//Kubepack: api.InfraKubepack{
			//	Host:   "",
			//	Bucket: "",
			//},
			// TODO: skip Customer install vs appscode install
			//Badger: api.InfraBadger{
			//	MountPath: "/badger",
			//	Levels:    7,
			//},
			//Invoice: api.InfraInvoice{
			//	MountPath:    "/billing",
			//	Bucket:       mustBucketName(in.Global.Infra.Objstore.Bucket, "invoices"),
			//	TrackerEmail: "",
			//},
		},
	}
	if in.Hosted {
		// TODO: bucket proxy
		//out.Global.Infra.Kubepack = api.InfraKubepack{
		//	Host:   "",
		//	Bucket: "gs://",
		//}
		out.Global.Infra.Badger = api.InfraBadger{
			MountPath: "/badger",
			Levels:    7,
		}
		out.Global.Infra.Invoice = api.InfraInvoice{
			MountPath:    "/billing",
			Bucket:       mustBucketName(in.Global.Infra.Objstore.Bucket, "invoices"),
			TrackerEmail: "",
		}
	}

	out.Settings = api.Settings{
		DB: api.DBSettings{
			Version:           "13.2",
			DatabaseName:      in.Release.Name,
			TerminationPolicy: "Delete", // TODO: change for prod mode
			Persistence:       in.Settings.DB.Persistence,
			Resources:         in.Settings.DB.Resources,
			Auth: api.BasicAuth{
				Username: "postgres",
				Password: passgen.Generate(DefaultPasswordLength),
			},
		},
		Cache: api.CacheSettings{
			Version:           "6.0.6",
			TerminationPolicy: "Delete",
			Persistence:       in.Settings.Cache.Persistence,
			Resources:         in.Settings.Cache.Resources,
			Auth: api.BasicAuth{
				Username: "root",
				Password: passgen.Generate(DefaultPasswordLength),
			},
			CacheInterval: 60,
		},
		Smtp: api.SmtpSettings{
			Host:       in.Settings.Smtp.Host,
			TlsEnabled: in.Settings.Smtp.TlsEnabled,
			From:       fmt.Sprintf("no-reply@%s", in.Settings.Platform.Domain), // TODO: configure?
			Username:   in.Settings.Smtp.Username,
			Password:   in.Settings.Smtp.Password,
			SubjectPrefix: func() string {
				if in.Hosted {
					return "ByteBuilders |"
				}
				return "ACE |"
			}(),
			SendAsPlainText: in.Settings.Smtp.SendAsPlainText,
		},
		// Nats:        api.NatsSettings{},
		Platform: api.PlatformSettings{
			Domain: in.Settings.Platform.Domain,
			AppName: func() string {
				if in.Hosted {
					return "ByteBuilders: Kubernetes Native Data Platform"
				}
				return "ACE: Kubernetes Native Data Platform"
			}(),
			RunMode:                         "prod",
			ExperimentalFeatures:            false,
			ForcePrivate:                    false,
			DisableHttpGit:                  false,
			InstallLock:                     true, // TODO: why?
			RepositoryUploadEnabled:         true,
			RepositoryUploadAllowedTypes:    nil,
			RepositoryUploadMaxFileSize:     3,
			RepositoryUploadMaxFiles:        5,
			ServiceEnableCaptcha:            true,
			ServiceRegisterEmailConfirm:     false,
			ServiceDisableRegistration:      false,
			ServiceRequireSignInView:        false,
			ServiceEnableNotifyMail:         true,
			CookieName:                      "i_like_bytebuilders",
			ServerLandingPage:               "home",
			LogMode:                         "console",
			LogLevel:                        "Info", // Trace
			OtherShowFooterBranding:         false,
			OtherShowFooterVersion:          true,
			OtherShowFooterTemplateLoadTime: true,
			EnableCSRFCookieHttpOnly:        false,
		},
		// Stripe:   api.StripeSettings{},
		Security: api.SecuritySettings{
			Oauth2JWTSecret: passgen.GenerateForCharset(43, passgen.AlphaNum),
			CsrfSecretKey:   passgen.GenerateForCharset(64, passgen.AlphaNum),
		},
		// Searchlight: api.SearchlightSettings{},
		Grafana: api.GrafanaSettings{
			AppMode:        "production",
			CacheAdapter:   "",
			CacheInterval:  0,
			CacheHost:      nil,
			SkipMigrations: false,
		},
	}
	if in.Hosted {
		// out.Settings.Stripe = api.StripeSettings{}
	}

	return nil
}

func tplPlatformTLSSecret(in *AceOptionsSpec) string {
	return fmt.Sprintf("%s-cert", in.Release.Name)
}

func tplNATSCredSecret(in *AceOptionsSpec) string {
	return fmt.Sprintf("%s-nats-cred", in.Release.Name)
}

func tplNATSTLSSecret(in *AceOptionsSpec) string {
	if in.Nats.ExposeVia == ServiceTypeLoadBalancer {
		return fmt.Sprintf("%s-nats-cert", in.Release.Name)
	}
	return fmt.Sprintf("%s-cert", in.Release.Name)
}

func tplPlatformConfig(in *AceOptionsSpec) string {
	return fmt.Sprintf("%s-config", in.Release.Name)
}

func getBucket(bucket string, elem ...string) (string, error) {
	u, err := url.Parse(bucket)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(append([]string{u.Path}, elem...)...)
	return u.String(), nil
}

func mustBucketName(bucket string, elem ...string) string {
	if name, err := getBucket(bucket, elem...); err != nil {
		panic(errors.Wrap(err, "failed to generate bucket name"))
	} else {
		return name
	}
}
