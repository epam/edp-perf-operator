module github.com/epam/edp-perf-operator/v2

go 1.14

replace (
	github.com/epam/edp-cd-pipeline-operator/v2 => ../cd-pipeline-operator
	github.com/epam/edp-codebase-operator/v2 => ../codebase-operator
	github.com/epam/edp-gerrit-operator/v2 => ../gerrit-operator
	github.com/epam/edp-jenkins-operator/v2 => ../jenkins-operator
	github.com/epam/edp-perf-operator/v2 => ../perf-operator

	git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
	github.com/openshift/api => github.com/openshift/api v0.0.0-20210416130433-86964261530c
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20210112165513-ebc401615f47
	k8s.io/api => k8s.io/api v0.20.7-rc.0
)

require (
	//EPMDEDP dependencies on operators
	github.com/epam/edp-codebase-operator/v2 v2.0.1
	github.com/epam/edp-component-operator v0.1.1-0.20210413101042-1d8f823f27cc

	//third party dependencies
	github.com/go-logr/logr v0.4.0
	github.com/go-openapi/spec v0.19.5
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	gopkg.in/resty.v1 v1.12.0
	k8s.io/api v0.21.0-rc.0
	k8s.io/apimachinery v0.21.0-rc.0
	k8s.io/client-go v0.20.2
	k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7
	sigs.k8s.io/controller-runtime v0.8.3
)
