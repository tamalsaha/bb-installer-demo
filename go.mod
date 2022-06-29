module github.com/tamalsaha/bb-installer-demo

go 1.18

require (
	github.com/alessio/shellescape v1.4.1
	github.com/nats-io/jwt/v2 v2.2.1-0.20220330180145-442af02fd36a
	github.com/nats-io/nats-server/v2 v2.8.4
	github.com/nats-io/nkeys v0.3.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.57.0
	go.bytebuilders.dev/installer v0.0.0-20220613195146-011e295a85ea
	gomodules.xyz/encoding v0.0.7
	gomodules.xyz/go-sh v0.1.0
	gomodules.xyz/homedir v0.1.0
	gomodules.xyz/password-generator v0.2.8
	gomodules.xyz/pointer v0.1.0
	k8s.io/api v0.24.1
	k8s.io/apimachinery v0.24.1
	kubepack.dev/lib-helm v0.3.2-0.20220620214914-ae2ddd1c076d
	sigs.k8s.io/yaml v1.3.0
)

replace go.bytebuilders.dev/installer => ../../../go.bytebuilders.dev/installer

require (
	github.com/BurntSushi/toml v1.0.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/Masterminds/sprig/v3 v3.2.2 // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/cyphar/filepath-securejoin v0.2.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1-0.20220316001817-d5090ed65664 // indirect
	github.com/emicklei/go-restful v2.9.5+incompatible // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.14.4 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	go.openviz.dev/installer v0.0.0-20220604082341-49cce25117ff // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	golang.org/x/net v0.0.0-20220531201128-c960675eff93 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220411224347-583f2d630306 // indirect
	gomodules.xyz/jsonpath v0.0.2 // indirect
	gomodules.xyz/mergo v0.3.13-0.20220214162359-48efe39fd402 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
	helm.sh/helm/v3 v3.0.0-00010101000000-000000000000 // indirect
	k8s.io/apiextensions-apiserver v0.24.1 // indirect
	k8s.io/client-go v0.24.1 // indirect
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/kube-openapi v0.0.0-20220413171646-5e7f5fdc6da6 // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	kubeops.dev/installer v0.0.0-20220627095147-82e789070444 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
)

replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.24.2-rc.0.0.20220603191800-1c7484099dee

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20220603223637-59dad1716c43

replace k8s.io/kubernetes => github.com/kmodules/kubernetes v1.25.0-alpha.0.0.20220603172133-1c9d09d1b3b8

replace sigs.k8s.io/controller-runtime => github.com/kmodules/controller-runtime v0.12.2-0.20220603144237-6cd001896bf3

replace github.com/Masterminds/sprig/v3 => github.com/gomodules/sprig/v3 v3.2.3-0.20220405051441-0a8a99bac1b8

replace helm.sh/helm/v3 => github.com/kubepack/helm/v3 v3.9.1-0.20220603235400-7882cd0f0948

replace sigs.k8s.io/application => github.com/kmodules/application v0.8.4-0.20220603221517-a62565381f64
