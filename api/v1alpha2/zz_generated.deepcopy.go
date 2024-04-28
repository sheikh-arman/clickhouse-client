//go:build !ignore_autogenerated

/*
Copyright 2024.

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

package v1alpha2

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
	kubedbv1alpha2 "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouse) DeepCopyInto(out *ClickHouse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouse.
func (in *ClickHouse) DeepCopy() *ClickHouse {
	if in == nil {
		return nil
	}
	out := new(ClickHouse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseKeeperNode) DeepCopyInto(out *ClickHouseKeeperNode) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseKeeperNode.
func (in *ClickHouseKeeperNode) DeepCopy() *ClickHouseKeeperNode {
	if in == nil {
		return nil
	}
	out := new(ClickHouseKeeperNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseList) DeepCopyInto(out *ClickHouseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClickHouse, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseList.
func (in *ClickHouseList) DeepCopy() *ClickHouseList {
	if in == nil {
		return nil
	}
	out := new(ClickHouseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseSpec) DeepCopyInto(out *ClickHouseSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AuthSecret != nil {
		in, out := &in.AuthSecret, &out.AuthSecret
		*out = new(kubedbv1alpha2.SecretReference)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(apiv1.TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
	if in.ServiceTemplates != nil {
		in, out := &in.ServiceTemplates, &out.ServiceTemplates
		*out = make([]kubedbv1alpha2.NamedServiceTemplateSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClickHousekeeper != nil {
		in, out := &in.ClickHousekeeper, &out.ClickHousekeeper
		*out = new(ClickhousekeeperConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseSpec.
func (in *ClickHouseSpec) DeepCopy() *ClickHouseSpec {
	if in == nil {
		return nil
	}
	out := new(ClickHouseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseStatus) DeepCopyInto(out *ClickHouseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(kubedbv1alpha2.Gateway)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseStatus.
func (in *ClickHouseStatus) DeepCopy() *ClickHouseStatus {
	if in == nil {
		return nil
	}
	out := new(ClickHouseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickhouseApp) DeepCopyInto(out *ClickhouseApp) {
	*out = *in
	if in.ClickHouse != nil {
		in, out := &in.ClickHouse, &out.ClickHouse
		*out = new(ClickHouse)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickhouseApp.
func (in *ClickhouseApp) DeepCopy() *ClickhouseApp {
	if in == nil {
		return nil
	}
	out := new(ClickhouseApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickhousekeeperConfig) DeepCopyInto(out *ClickhousekeeperConfig) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]ClickHouseKeeperNode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickhousekeeperConfig.
func (in *ClickhousekeeperConfig) DeepCopy() *ClickhousekeeperConfig {
	if in == nil {
		return nil
	}
	out := new(ClickhousekeeperConfig)
	in.DeepCopyInto(out)
	return out
}