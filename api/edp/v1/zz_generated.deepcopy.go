//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSourceGitLabConfig) DeepCopyInto(out *DataSourceGitLabConfig) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Branches != nil {
		in, out := &in.Branches, &out.Branches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSourceGitLabConfig.
func (in *DataSourceGitLabConfig) DeepCopy() *DataSourceGitLabConfig {
	if in == nil {
		return nil
	}
	out := new(DataSourceGitLabConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSourceJenkinsConfig) DeepCopyInto(out *DataSourceJenkinsConfig) {
	*out = *in
	if in.JobNames != nil {
		in, out := &in.JobNames, &out.JobNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSourceJenkinsConfig.
func (in *DataSourceJenkinsConfig) DeepCopy() *DataSourceJenkinsConfig {
	if in == nil {
		return nil
	}
	out := new(DataSourceJenkinsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSourceSonarConfig) DeepCopyInto(out *DataSourceSonarConfig) {
	*out = *in
	if in.ProjectKeys != nil {
		in, out := &in.ProjectKeys, &out.ProjectKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSourceSonarConfig.
func (in *DataSourceSonarConfig) DeepCopy() *DataSourceSonarConfig {
	if in == nil {
		return nil
	}
	out := new(DataSourceSonarConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceGitLab) DeepCopyInto(out *PerfDataSourceGitLab) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceGitLab.
func (in *PerfDataSourceGitLab) DeepCopy() *PerfDataSourceGitLab {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceGitLab)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerfDataSourceGitLab) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceGitLabList) DeepCopyInto(out *PerfDataSourceGitLabList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerfDataSourceGitLab, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceGitLabList.
func (in *PerfDataSourceGitLabList) DeepCopy() *PerfDataSourceGitLabList {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceGitLabList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerfDataSourceGitLabList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceGitLabSpec) DeepCopyInto(out *PerfDataSourceGitLabSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceGitLabSpec.
func (in *PerfDataSourceGitLabSpec) DeepCopy() *PerfDataSourceGitLabSpec {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceGitLabSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceGitLabStatus) DeepCopyInto(out *PerfDataSourceGitLabStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceGitLabStatus.
func (in *PerfDataSourceGitLabStatus) DeepCopy() *PerfDataSourceGitLabStatus {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceGitLabStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceJenkins) DeepCopyInto(out *PerfDataSourceJenkins) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceJenkins.
func (in *PerfDataSourceJenkins) DeepCopy() *PerfDataSourceJenkins {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceJenkins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerfDataSourceJenkins) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceJenkinsList) DeepCopyInto(out *PerfDataSourceJenkinsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerfDataSourceJenkins, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceJenkinsList.
func (in *PerfDataSourceJenkinsList) DeepCopy() *PerfDataSourceJenkinsList {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceJenkinsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerfDataSourceJenkinsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceJenkinsSpec) DeepCopyInto(out *PerfDataSourceJenkinsSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceJenkinsSpec.
func (in *PerfDataSourceJenkinsSpec) DeepCopy() *PerfDataSourceJenkinsSpec {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceJenkinsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceJenkinsStatus) DeepCopyInto(out *PerfDataSourceJenkinsStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceJenkinsStatus.
func (in *PerfDataSourceJenkinsStatus) DeepCopy() *PerfDataSourceJenkinsStatus {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceJenkinsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceSonar) DeepCopyInto(out *PerfDataSourceSonar) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceSonar.
func (in *PerfDataSourceSonar) DeepCopy() *PerfDataSourceSonar {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceSonar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerfDataSourceSonar) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceSonarList) DeepCopyInto(out *PerfDataSourceSonarList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerfDataSourceSonar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceSonarList.
func (in *PerfDataSourceSonarList) DeepCopy() *PerfDataSourceSonarList {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceSonarList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerfDataSourceSonarList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceSonarSpec) DeepCopyInto(out *PerfDataSourceSonarSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceSonarSpec.
func (in *PerfDataSourceSonarSpec) DeepCopy() *PerfDataSourceSonarSpec {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceSonarSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfDataSourceSonarStatus) DeepCopyInto(out *PerfDataSourceSonarStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfDataSourceSonarStatus.
func (in *PerfDataSourceSonarStatus) DeepCopy() *PerfDataSourceSonarStatus {
	if in == nil {
		return nil
	}
	out := new(PerfDataSourceSonarStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfServer) DeepCopyInto(out *PerfServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfServer.
func (in *PerfServer) DeepCopy() *PerfServer {
	if in == nil {
		return nil
	}
	out := new(PerfServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerfServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfServerList) DeepCopyInto(out *PerfServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerfServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfServerList.
func (in *PerfServerList) DeepCopy() *PerfServerList {
	if in == nil {
		return nil
	}
	out := new(PerfServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerfServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfServerSpec) DeepCopyInto(out *PerfServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfServerSpec.
func (in *PerfServerSpec) DeepCopy() *PerfServerSpec {
	if in == nil {
		return nil
	}
	out := new(PerfServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerfServerStatus) DeepCopyInto(out *PerfServerStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerfServerStatus.
func (in *PerfServerStatus) DeepCopy() *PerfServerStatus {
	if in == nil {
		return nil
	}
	out := new(PerfServerStatus)
	in.DeepCopyInto(out)
	return out
}
