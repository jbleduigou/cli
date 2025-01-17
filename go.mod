module github.com/kyma-project/cli

go 1.12

replace github.com/census-instrumentation/opencensus-proto v0.1.0-0.20181214143942-ba49f56771b8 => github.com/census-instrumentation/opencensus-proto v0.0.3-0.20181214143942-ba49f56771b8

// Docker client has an issue on windows with the latest sys package, we have to fix the version
replace golang.org/x/sys => golang.org/x/sys v0.0.0-20190710143415-6ec70d6a5542

require (
	github.com/Masterminds/semver v1.5.0
	github.com/briandowns/spinner v1.7.0
	github.com/containerd/continuity v0.0.0-20190827140505-75bee3e2ccb6 // indirect
	github.com/daviddengcn/go-colortext v0.0.0-20180409174941-186a3d44e920
	github.com/fatih/color v1.7.0
	github.com/fsouza/go-dockerclient v1.4.5-0.20191009031337-a958d2e31b6c
	github.com/golangplus/bytes v0.0.0-20160111154220-45c989fe5450 // indirect
	github.com/golangplus/fmt v0.0.0-20150411045040-2a5d6d7d2995 // indirect
	github.com/golangplus/testing v0.0.0-20180327235837-af21d9c3145e // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/kyma-incubator/hydroform v0.0.0-20191015145515-4977abd24fbc
	github.com/kyma-incubator/octopus v0.0.0-20191009105757-2e9d86cd9967
	github.com/kyma-project/kyma v0.5.1-0.20190909070658-69599d4a33a2
	github.com/mitchellh/mapstructure v1.1.2
	github.com/olekukonko/tablewriter v0.0.1
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20191002192127-34f69633bfdc // indirect
	golang.org/x/net v0.0.0-20191003171128-d98b1b443823 // indirect
	google.golang.org/genproto v0.0.0-20191002211648-c459b9ce5143 // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/yaml.v2 v2.2.4
	k8s.io/api v0.0.0-20190620084959-7cf5895f2711
	k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
	k8s.io/client-go v12.0.0+incompatible
)
