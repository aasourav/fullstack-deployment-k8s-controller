//go:build !ignore_autogenerated

/*
Copyright 2025 Ahsan Amin.

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
func (in *FullStackDeploy) DeepCopyInto(out *FullStackDeploy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FullStackDeploy.
func (in *FullStackDeploy) DeepCopy() *FullStackDeploy {
	if in == nil {
		return nil
	}
	out := new(FullStackDeploy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FullStackDeploy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FullStackDeployList) DeepCopyInto(out *FullStackDeployList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FullStackDeploy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FullStackDeployList.
func (in *FullStackDeployList) DeepCopy() *FullStackDeployList {
	if in == nil {
		return nil
	}
	out := new(FullStackDeployList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FullStackDeployList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FullStackDeploySpec) DeepCopyInto(out *FullStackDeploySpec) {
	*out = *in
	if in.FrontendEnv != nil {
		in, out := &in.FrontendEnv, &out.FrontendEnv
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.BackendEnv != nil {
		in, out := &in.BackendEnv, &out.BackendEnv
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FullStackDeploySpec.
func (in *FullStackDeploySpec) DeepCopy() *FullStackDeploySpec {
	if in == nil {
		return nil
	}
	out := new(FullStackDeploySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FullStackDeployStatus) DeepCopyInto(out *FullStackDeployStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FullStackDeployStatus.
func (in *FullStackDeployStatus) DeepCopy() *FullStackDeployStatus {
	if in == nil {
		return nil
	}
	out := new(FullStackDeployStatus)
	in.DeepCopyInto(out)
	return out
}
