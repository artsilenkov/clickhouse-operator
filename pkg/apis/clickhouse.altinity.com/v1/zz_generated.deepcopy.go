// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiCluster) DeepCopyInto(out *ChiCluster) {
	*out = *in
	in.Layout.DeepCopyInto(&out.Layout)
	out.Templates = in.Templates
	out.Address = in.Address
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiCluster.
func (in *ChiCluster) DeepCopy() *ChiCluster {
	if in == nil {
		return nil
	}
	out := new(ChiCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiClusterAddress) DeepCopyInto(out *ChiClusterAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiClusterAddress.
func (in *ChiClusterAddress) DeepCopy() *ChiClusterAddress {
	if in == nil {
		return nil
	}
	out := new(ChiClusterAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiConfiguration) DeepCopyInto(out *ChiConfiguration) {
	*out = *in
	in.Zookeeper.DeepCopyInto(&out.Zookeeper)
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make(map[string]interface{}, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val
			}
		}
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make(map[string]interface{}, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val
			}
		}
	}
	if in.Quotas != nil {
		in, out := &in.Quotas, &out.Quotas
		*out = make(map[string]interface{}, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val
			}
		}
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make(map[string]interface{}, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val
			}
		}
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]ChiCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiConfiguration.
func (in *ChiConfiguration) DeepCopy() *ChiConfiguration {
	if in == nil {
		return nil
	}
	out := new(ChiConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiDefaults) DeepCopyInto(out *ChiDefaults) {
	*out = *in
	out.DistributedDDL = in.DistributedDDL
	out.Templates = in.Templates
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiDefaults.
func (in *ChiDefaults) DeepCopy() *ChiDefaults {
	if in == nil {
		return nil
	}
	out := new(ChiDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiDistributedDDL) DeepCopyInto(out *ChiDistributedDDL) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiDistributedDDL.
func (in *ChiDistributedDDL) DeepCopy() *ChiDistributedDDL {
	if in == nil {
		return nil
	}
	out := new(ChiDistributedDDL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiLayout) DeepCopyInto(out *ChiLayout) {
	*out = *in
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = make([]ChiShard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiLayout.
func (in *ChiLayout) DeepCopy() *ChiLayout {
	if in == nil {
		return nil
	}
	out := new(ChiLayout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiPodTemplate) DeepCopyInto(out *ChiPodTemplate) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiPodTemplate.
func (in *ChiPodTemplate) DeepCopy() *ChiPodTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiPodTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiReplica) DeepCopyInto(out *ChiReplica) {
	*out = *in
	out.Templates = in.Templates
	out.Address = in.Address
	out.Config = in.Config
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiReplica.
func (in *ChiReplica) DeepCopy() *ChiReplica {
	if in == nil {
		return nil
	}
	out := new(ChiReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiReplicaAddress) DeepCopyInto(out *ChiReplicaAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiReplicaAddress.
func (in *ChiReplicaAddress) DeepCopy() *ChiReplicaAddress {
	if in == nil {
		return nil
	}
	out := new(ChiReplicaAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiReplicaConfig) DeepCopyInto(out *ChiReplicaConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiReplicaConfig.
func (in *ChiReplicaConfig) DeepCopy() *ChiReplicaConfig {
	if in == nil {
		return nil
	}
	out := new(ChiReplicaConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiShard) DeepCopyInto(out *ChiShard) {
	*out = *in
	out.Templates = in.Templates
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]ChiReplica, len(*in))
		copy(*out, *in)
	}
	out.Address = in.Address
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiShard.
func (in *ChiShard) DeepCopy() *ChiShard {
	if in == nil {
		return nil
	}
	out := new(ChiShard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiShardAddress) DeepCopyInto(out *ChiShardAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiShardAddress.
func (in *ChiShardAddress) DeepCopy() *ChiShardAddress {
	if in == nil {
		return nil
	}
	out := new(ChiShardAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiSpec) DeepCopyInto(out *ChiSpec) {
	*out = *in
	out.Defaults = in.Defaults
	in.Configuration.DeepCopyInto(&out.Configuration)
	in.Templates.DeepCopyInto(&out.Templates)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiSpec.
func (in *ChiSpec) DeepCopy() *ChiSpec {
	if in == nil {
		return nil
	}
	out := new(ChiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiStatus) DeepCopyInto(out *ChiStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiStatus.
func (in *ChiStatus) DeepCopy() *ChiStatus {
	if in == nil {
		return nil
	}
	out := new(ChiStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiTemplateNames) DeepCopyInto(out *ChiTemplateNames) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiTemplateNames.
func (in *ChiTemplateNames) DeepCopy() *ChiTemplateNames {
	if in == nil {
		return nil
	}
	out := new(ChiTemplateNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiTemplates) DeepCopyInto(out *ChiTemplates) {
	*out = *in
	if in.PodTemplates != nil {
		in, out := &in.PodTemplates, &out.PodTemplates
		*out = make([]ChiPodTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]ChiVolumeClaimTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiTemplates.
func (in *ChiTemplates) DeepCopy() *ChiTemplates {
	if in == nil {
		return nil
	}
	out := new(ChiTemplates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiVolumeClaimTemplate) DeepCopyInto(out *ChiVolumeClaimTemplate) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiVolumeClaimTemplate.
func (in *ChiVolumeClaimTemplate) DeepCopy() *ChiVolumeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiVolumeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiZookeeperConfig) DeepCopyInto(out *ChiZookeeperConfig) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]ChiZookeeperNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiZookeeperConfig.
func (in *ChiZookeeperConfig) DeepCopy() *ChiZookeeperConfig {
	if in == nil {
		return nil
	}
	out := new(ChiZookeeperConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiZookeeperNode) DeepCopyInto(out *ChiZookeeperNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiZookeeperNode.
func (in *ChiZookeeperNode) DeepCopy() *ChiZookeeperNode {
	if in == nil {
		return nil
	}
	out := new(ChiZookeeperNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallation) DeepCopyInto(out *ClickHouseInstallation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallation.
func (in *ClickHouseInstallation) DeepCopy() *ClickHouseInstallation {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallationList) DeepCopyInto(out *ClickHouseInstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClickHouseInstallation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallationList.
func (in *ClickHouseInstallationList) DeepCopy() *ClickHouseInstallationList {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}
