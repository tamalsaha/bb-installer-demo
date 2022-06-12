package main

import (
	"fmt"

	core "k8s.io/api/core/v1"

	api "go.bytebuilders.dev/installer/apis/installer/v1alpha1"
	"sigs.k8s.io/yaml"
)

type Options struct {
	Billing       ComponentSpec   `json:"billing"`
	PlatformUi    ComponentSpec   `json:"platform-ui"`
	AccountsUi    ComponentSpec   `json:"accounts-ui"`
	ClusterUi     ComponentSpec   `json:"cluster-ui"`
	DeployUi      ComponentSpec   `json:"deploy-ui"`
	Grafana       ComponentSpec   `json:"grafana"`
	KubedbUi      ComponentSpec   `json:"kubedb-ui"`
	MarketplaceUi ComponentSpec   `json:"marketplace-ui"`
	PlatformApi   ComponentSpec   `json:"platform-api"`
	PromProxy     ComponentSpec   `json:"prom-proxy"`
	Ingress       IngressNginx    `json:"ingress"`
	Nats          NatsSettings    `json:"nats"`
	Global        AceGlobalValues `json:"global"`
	Settings      Settings        `json:"settings"`
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

type IngressNginx struct {
	ExposeVia    ServiceType       `json:"exposeVia"`
	NodeSelector map[string]string `json:"nodeSelector"`
}

type NatsSettings struct {
	ExposeVia ServiceType `json:"exposeVia"`
	Replics   int         `json:"replicas"`

	// ShardCount int `json:"shardCount"`
	//MountPath       string `json:"mountPath"`
	//OperatorCreds   string `json:"operatorCreds"`
	//OperatorJwt     string `json:"operatorJwt"`
	//SystemCreds     string `json:"systemCreds"`
	//SystemJwt       string `json:"systemJwt"`
	//SystemPubKey    string `json:"systemPubKey"`
	//SystemUserCreds string `json:"systemUserCreds"`
	//AdminCreds      string `json:"adminCreds"`
	//AdminUserCreds  string `json:"adminUserCreds"`
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
	DNS          InfraDns                 `json:"dns"`
	Objstore     InfraObjstore            `json:"objstore"`
	Kms          InfraKms                 `json:"kms"`
	Avatars      InfraAvatars             `json:"avatars"`
	// Kubepack     InfraKubepack            `json:"kubepack"`
	// Badger       InfraBadger              `json:"badger"`
	// Invoice      InfraInvoice             `json:"invoice"`
}

type InfraTLS struct {
	// AcmeServer string `json:"acmeServer"`
	Email string `json:"email"`
}

type InfraDns struct {
	Provider string          `json:"provider"`
	Auth     DNSProviderAuth `json:"auth"`
}

type DNSProviderAuth struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type InfraObjstore struct {
	Provider string `json:"provider"`
	// MountPath string       `json:"mountPath"`
	Auth ObjstoreAuth `json:"auth"`
}

type ObjstoreAuth struct {
	ServiceAccountJson string `json:"serviceAccountJson"`
}

type InfraKms struct {
	Provider string `json:"provider"`
	// MountPath    string `json:"mountPath"`
	MasterKeyURL string `json:"masterKeyURL"`
}

type InfraAvatars struct {
	Bucket string `json:"bucket"`
}

type InfraKubepack struct {
	Host   string `json:"host"`
	Bucket string `json:"bucket"`
}

type InfraBadger struct {
	MountPath string `json:"mountPath"`
	Levels    int    `json:"levels"`
}

type InfraInvoice struct {
	MountPath    string `json:"mountPath"`
	Bucket       string `json:"bucket"`
	TrackerEmail string `json:"trackerEmail"`
}

type Settings struct {
	DB       DBSettings           `json:"db"`
	Cache    CacheSettings        `json:"cache"`
	Smtp     SmtpSettings         `json:"smtp"`
	Platform PlatformSettings     `json:"platform"`
	Security api.SecuritySettings `json:"security"`
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
