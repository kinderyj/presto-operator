// +build !ignore_autogenerated

/*
Copyright 2020 yujunwang.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogConfig) DeepCopyInto(out *CatalogConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogConfig.
func (in *CatalogConfig) DeepCopy() *CatalogConfig {
	if in == nil {
		return nil
	}
	out := new(CatalogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoordinatorConfig) DeepCopyInto(out *CoordinatorConfig) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.DynamicArgs != nil {
		in, out := &in.DynamicArgs, &out.DynamicArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoordinatorConfig.
func (in *CoordinatorConfig) DeepCopy() *CoordinatorConfig {
	if in == nil {
		return nil
	}
	out := new(CoordinatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoCluster) DeepCopyInto(out *PrestoCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoCluster.
func (in *PrestoCluster) DeepCopy() *PrestoCluster {
	if in == nil {
		return nil
	}
	out := new(PrestoCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrestoCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoClusterList) DeepCopyInto(out *PrestoClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrestoCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoClusterList.
func (in *PrestoClusterList) DeepCopy() *PrestoClusterList {
	if in == nil {
		return nil
	}
	out := new(PrestoClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrestoClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoClusterSpec) DeepCopyInto(out *PrestoClusterSpec) {
	*out = *in
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = new(int32)
		**out = **in
	}
	in.CoordinatorConfig.DeepCopyInto(&out.CoordinatorConfig)
	in.WorkerConfig.DeepCopyInto(&out.WorkerConfig)
	out.CatalogConfig = in.CatalogConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoClusterSpec.
func (in *PrestoClusterSpec) DeepCopy() *PrestoClusterSpec {
	if in == nil {
		return nil
	}
	out := new(PrestoClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoClusterStatus) DeepCopyInto(out *PrestoClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoClusterStatus.
func (in *PrestoClusterStatus) DeepCopy() *PrestoClusterStatus {
	if in == nil {
		return nil
	}
	out := new(PrestoClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerConfig) DeepCopyInto(out *WorkerConfig) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.DynamicArgs != nil {
		in, out := &in.DynamicArgs, &out.DynamicArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerConfig.
func (in *WorkerConfig) DeepCopy() *WorkerConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerConfig)
	in.DeepCopyInto(out)
	return out
}
