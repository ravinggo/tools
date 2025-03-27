module github.com/ravinggo/tools

go 1.23.3

require (
	github.com/emicklei/proto v1.14.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/gofuzz v1.2.0
	github.com/huandu/xstrings v1.5.0
	github.com/mailru/easyjson v0.9.0
	github.com/ravinggo/objectpool v1.0.1
	github.com/ravinggo/zerolog v0.0.2
	github.com/spf13/pflag v1.0.6
	github.com/stretchr/testify v1.10.0
	github.com/xuri/excelize/v2 v2.9.0
	go.mongodb.org/mongo-driver/v2 v2.1.0
	google.golang.org/protobuf v1.35.1
	k8s.io/apimachinery v0.32.2
	k8s.io/gengo/v2 v2.0.0-20250207200755-1244d31929d7
	k8s.io/klog/v2 v2.130.1
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.4 // indirect
	github.com/xuri/efp v0.0.0-20250227110027-3491fafc2b79 // indirect
	github.com/xuri/nfp v0.0.0-20250226145837-86d5fc24b2ba // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/tools v0.26.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace k8s.io/gengo/v2 v2.0.0-20250207200755-1244d31929d7 => github.com/ravinggo/gengo/v2 v2.0.1
