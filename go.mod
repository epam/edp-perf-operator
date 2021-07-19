module github.com/epam/edp-perf-operator/v2

go 1.14

replace (
	git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
	github.com/kubernetes-incubator/reference-docs => github.com/kubernetes-sigs/reference-docs v0.0.0-20170929004150-fcf65347b256
	github.com/markbates/inflect => github.com/markbates/inflect v1.0.4
	github.com/openshift/api => github.com/openshift/api v0.0.0-20210416130433-86964261530c
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20210112165513-ebc401615f47
	k8s.io/api => k8s.io/api v0.20.7-rc.0
)

require (
	github.com/epam/edp-codebase-operator/v2 v2.3.0-95.0.20210719112602-046af47ad1c0
	github.com/epam/edp-component-operator v0.1.1-0.20210712140516-09b8bb3a4cff
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
