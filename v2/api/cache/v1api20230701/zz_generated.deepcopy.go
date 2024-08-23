//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20230701

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProperties_ARM) DeepCopyInto(out *ClusterProperties_ARM) {
	*out = *in
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(ClusterProperties_MinimumTlsVersion_ARM)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProperties_ARM.
func (in *ClusterProperties_ARM) DeepCopy() *ClusterProperties_ARM {
	if in == nil {
		return nil
	}
	out := new(ClusterProperties_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProperties_STATUS_ARM) DeepCopyInto(out *ClusterProperties_STATUS_ARM) {
	*out = *in
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(ClusterProperties_MinimumTlsVersion_STATUS_ARM)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS_ARM)
		**out = **in
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ResourceState != nil {
		in, out := &in.ResourceState, &out.ResourceState
		*out = new(ResourceState_STATUS_ARM)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProperties_STATUS_ARM.
func (in *ClusterProperties_STATUS_ARM) DeepCopy() *ClusterProperties_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(ClusterProperties_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseProperties_ARM) DeepCopyInto(out *DatabaseProperties_ARM) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(DatabaseProperties_ClientProtocol_ARM)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(DatabaseProperties_ClusteringPolicy_ARM)
		**out = **in
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(DatabaseProperties_EvictionPolicy_ARM)
		**out = **in
	}
	if in.GeoReplication != nil {
		in, out := &in.GeoReplication, &out.GeoReplication
		*out = new(DatabaseProperties_GeoReplication_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseProperties_ARM.
func (in *DatabaseProperties_ARM) DeepCopy() *DatabaseProperties_ARM {
	if in == nil {
		return nil
	}
	out := new(DatabaseProperties_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseProperties_GeoReplication) DeepCopyInto(out *DatabaseProperties_GeoReplication) {
	*out = *in
	if in.GroupNickname != nil {
		in, out := &in.GroupNickname, &out.GroupNickname
		*out = new(string)
		**out = **in
	}
	if in.LinkedDatabases != nil {
		in, out := &in.LinkedDatabases, &out.LinkedDatabases
		*out = make([]LinkedDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseProperties_GeoReplication.
func (in *DatabaseProperties_GeoReplication) DeepCopy() *DatabaseProperties_GeoReplication {
	if in == nil {
		return nil
	}
	out := new(DatabaseProperties_GeoReplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseProperties_GeoReplication_ARM) DeepCopyInto(out *DatabaseProperties_GeoReplication_ARM) {
	*out = *in
	if in.GroupNickname != nil {
		in, out := &in.GroupNickname, &out.GroupNickname
		*out = new(string)
		**out = **in
	}
	if in.LinkedDatabases != nil {
		in, out := &in.LinkedDatabases, &out.LinkedDatabases
		*out = make([]LinkedDatabase_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseProperties_GeoReplication_ARM.
func (in *DatabaseProperties_GeoReplication_ARM) DeepCopy() *DatabaseProperties_GeoReplication_ARM {
	if in == nil {
		return nil
	}
	out := new(DatabaseProperties_GeoReplication_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseProperties_GeoReplication_STATUS) DeepCopyInto(out *DatabaseProperties_GeoReplication_STATUS) {
	*out = *in
	if in.GroupNickname != nil {
		in, out := &in.GroupNickname, &out.GroupNickname
		*out = new(string)
		**out = **in
	}
	if in.LinkedDatabases != nil {
		in, out := &in.LinkedDatabases, &out.LinkedDatabases
		*out = make([]LinkedDatabase_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseProperties_GeoReplication_STATUS.
func (in *DatabaseProperties_GeoReplication_STATUS) DeepCopy() *DatabaseProperties_GeoReplication_STATUS {
	if in == nil {
		return nil
	}
	out := new(DatabaseProperties_GeoReplication_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseProperties_GeoReplication_STATUS_ARM) DeepCopyInto(out *DatabaseProperties_GeoReplication_STATUS_ARM) {
	*out = *in
	if in.GroupNickname != nil {
		in, out := &in.GroupNickname, &out.GroupNickname
		*out = new(string)
		**out = **in
	}
	if in.LinkedDatabases != nil {
		in, out := &in.LinkedDatabases, &out.LinkedDatabases
		*out = make([]LinkedDatabase_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseProperties_GeoReplication_STATUS_ARM.
func (in *DatabaseProperties_GeoReplication_STATUS_ARM) DeepCopy() *DatabaseProperties_GeoReplication_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(DatabaseProperties_GeoReplication_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseProperties_STATUS_ARM) DeepCopyInto(out *DatabaseProperties_STATUS_ARM) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(DatabaseProperties_ClientProtocol_STATUS_ARM)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(DatabaseProperties_ClusteringPolicy_STATUS_ARM)
		**out = **in
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(DatabaseProperties_EvictionPolicy_STATUS_ARM)
		**out = **in
	}
	if in.GeoReplication != nil {
		in, out := &in.GeoReplication, &out.GeoReplication
		*out = new(DatabaseProperties_GeoReplication_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS_ARM)
		**out = **in
	}
	if in.ResourceState != nil {
		in, out := &in.ResourceState, &out.ResourceState
		*out = new(ResourceState_STATUS_ARM)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseProperties_STATUS_ARM.
func (in *DatabaseProperties_STATUS_ARM) DeepCopy() *DatabaseProperties_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(DatabaseProperties_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedDatabase) DeepCopyInto(out *LinkedDatabase) {
	*out = *in
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedDatabase.
func (in *LinkedDatabase) DeepCopy() *LinkedDatabase {
	if in == nil {
		return nil
	}
	out := new(LinkedDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedDatabase_ARM) DeepCopyInto(out *LinkedDatabase_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedDatabase_ARM.
func (in *LinkedDatabase_ARM) DeepCopy() *LinkedDatabase_ARM {
	if in == nil {
		return nil
	}
	out := new(LinkedDatabase_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedDatabase_STATUS) DeepCopyInto(out *LinkedDatabase_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LinkedDatabase_State_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedDatabase_STATUS.
func (in *LinkedDatabase_STATUS) DeepCopy() *LinkedDatabase_STATUS {
	if in == nil {
		return nil
	}
	out := new(LinkedDatabase_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedDatabase_STATUS_ARM) DeepCopyInto(out *LinkedDatabase_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LinkedDatabase_State_STATUS_ARM)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedDatabase_STATUS_ARM.
func (in *LinkedDatabase_STATUS_ARM) DeepCopy() *LinkedDatabase_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(LinkedDatabase_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module_ARM) DeepCopyInto(out *Module_ARM) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module_ARM.
func (in *Module_ARM) DeepCopy() *Module_ARM {
	if in == nil {
		return nil
	}
	out := new(Module_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module_STATUS) DeepCopyInto(out *Module_STATUS) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module_STATUS.
func (in *Module_STATUS) DeepCopy() *Module_STATUS {
	if in == nil {
		return nil
	}
	out := new(Module_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module_STATUS_ARM) DeepCopyInto(out *Module_STATUS_ARM) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module_STATUS_ARM.
func (in *Module_STATUS_ARM) DeepCopy() *Module_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(Module_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence) DeepCopyInto(out *Persistence) {
	*out = *in
	if in.AofEnabled != nil {
		in, out := &in.AofEnabled, &out.AofEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofFrequency != nil {
		in, out := &in.AofFrequency, &out.AofFrequency
		*out = new(Persistence_AofFrequency)
		**out = **in
	}
	if in.RdbEnabled != nil {
		in, out := &in.RdbEnabled, &out.RdbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbFrequency != nil {
		in, out := &in.RdbFrequency, &out.RdbFrequency
		*out = new(Persistence_RdbFrequency)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence.
func (in *Persistence) DeepCopy() *Persistence {
	if in == nil {
		return nil
	}
	out := new(Persistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence_ARM) DeepCopyInto(out *Persistence_ARM) {
	*out = *in
	if in.AofEnabled != nil {
		in, out := &in.AofEnabled, &out.AofEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofFrequency != nil {
		in, out := &in.AofFrequency, &out.AofFrequency
		*out = new(Persistence_AofFrequency_ARM)
		**out = **in
	}
	if in.RdbEnabled != nil {
		in, out := &in.RdbEnabled, &out.RdbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbFrequency != nil {
		in, out := &in.RdbFrequency, &out.RdbFrequency
		*out = new(Persistence_RdbFrequency_ARM)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence_ARM.
func (in *Persistence_ARM) DeepCopy() *Persistence_ARM {
	if in == nil {
		return nil
	}
	out := new(Persistence_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence_STATUS) DeepCopyInto(out *Persistence_STATUS) {
	*out = *in
	if in.AofEnabled != nil {
		in, out := &in.AofEnabled, &out.AofEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofFrequency != nil {
		in, out := &in.AofFrequency, &out.AofFrequency
		*out = new(Persistence_AofFrequency_STATUS)
		**out = **in
	}
	if in.RdbEnabled != nil {
		in, out := &in.RdbEnabled, &out.RdbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbFrequency != nil {
		in, out := &in.RdbFrequency, &out.RdbFrequency
		*out = new(Persistence_RdbFrequency_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence_STATUS.
func (in *Persistence_STATUS) DeepCopy() *Persistence_STATUS {
	if in == nil {
		return nil
	}
	out := new(Persistence_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence_STATUS_ARM) DeepCopyInto(out *Persistence_STATUS_ARM) {
	*out = *in
	if in.AofEnabled != nil {
		in, out := &in.AofEnabled, &out.AofEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofFrequency != nil {
		in, out := &in.AofFrequency, &out.AofFrequency
		*out = new(Persistence_AofFrequency_STATUS_ARM)
		**out = **in
	}
	if in.RdbEnabled != nil {
		in, out := &in.RdbEnabled, &out.RdbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbFrequency != nil {
		in, out := &in.RdbFrequency, &out.RdbFrequency
		*out = new(Persistence_RdbFrequency_STATUS_ARM)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence_STATUS_ARM.
func (in *Persistence_STATUS_ARM) DeepCopy() *Persistence_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(Persistence_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_STATUS) DeepCopyInto(out *PrivateEndpointConnection_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_STATUS.
func (in *PrivateEndpointConnection_STATUS) DeepCopy() *PrivateEndpointConnection_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_STATUS_ARM) DeepCopyInto(out *PrivateEndpointConnection_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_STATUS_ARM.
func (in *PrivateEndpointConnection_STATUS_ARM) DeepCopy() *PrivateEndpointConnection_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise) DeepCopyInto(out *RedisEnterprise) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise.
func (in *RedisEnterprise) DeepCopy() *RedisEnterprise {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterprise) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabase) DeepCopyInto(out *RedisEnterpriseDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabase.
func (in *RedisEnterpriseDatabase) DeepCopy() *RedisEnterpriseDatabase {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabaseList) DeepCopyInto(out *RedisEnterpriseDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisEnterpriseDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabaseList.
func (in *RedisEnterpriseDatabaseList) DeepCopy() *RedisEnterpriseDatabaseList {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseList) DeepCopyInto(out *RedisEnterpriseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisEnterprise, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseList.
func (in *RedisEnterpriseList) DeepCopy() *RedisEnterpriseList {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_Database_STATUS) DeepCopyInto(out *RedisEnterprise_Database_STATUS) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(DatabaseProperties_ClientProtocol_STATUS)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(DatabaseProperties_ClusteringPolicy_STATUS)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(DatabaseProperties_EvictionPolicy_STATUS)
		**out = **in
	}
	if in.GeoReplication != nil {
		in, out := &in.GeoReplication, &out.GeoReplication
		*out = new(DatabaseProperties_GeoReplication_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS)
		**out = **in
	}
	if in.ResourceState != nil {
		in, out := &in.ResourceState, &out.ResourceState
		*out = new(ResourceState_STATUS)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_Database_STATUS.
func (in *RedisEnterprise_Database_STATUS) DeepCopy() *RedisEnterprise_Database_STATUS {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_Database_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_Database_STATUS_ARM) DeepCopyInto(out *RedisEnterprise_Database_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(DatabaseProperties_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_Database_STATUS_ARM.
func (in *RedisEnterprise_Database_STATUS_ARM) DeepCopy() *RedisEnterprise_Database_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_Database_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_Database_Spec) DeepCopyInto(out *RedisEnterprise_Database_Spec) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(DatabaseProperties_ClientProtocol)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(DatabaseProperties_ClusteringPolicy)
		**out = **in
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(DatabaseProperties_EvictionPolicy)
		**out = **in
	}
	if in.GeoReplication != nil {
		in, out := &in.GeoReplication, &out.GeoReplication
		*out = new(DatabaseProperties_GeoReplication)
		(*in).DeepCopyInto(*out)
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_Database_Spec.
func (in *RedisEnterprise_Database_Spec) DeepCopy() *RedisEnterprise_Database_Spec {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_Database_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_Database_Spec_ARM) DeepCopyInto(out *RedisEnterprise_Database_Spec_ARM) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(DatabaseProperties_ARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_Database_Spec_ARM.
func (in *RedisEnterprise_Database_Spec_ARM) DeepCopy() *RedisEnterprise_Database_Spec_ARM {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_Database_Spec_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_STATUS) DeepCopyInto(out *RedisEnterprise_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(ClusterProperties_MinimumTlsVersion_STATUS)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS)
		**out = **in
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ResourceState != nil {
		in, out := &in.ResourceState, &out.ResourceState
		*out = new(ResourceState_STATUS)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_STATUS.
func (in *RedisEnterprise_STATUS) DeepCopy() *RedisEnterprise_STATUS {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_STATUS_ARM) DeepCopyInto(out *RedisEnterprise_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ClusterProperties_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_STATUS_ARM.
func (in *RedisEnterprise_STATUS_ARM) DeepCopy() *RedisEnterprise_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_Spec) DeepCopyInto(out *RedisEnterprise_Spec) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(ClusterProperties_MinimumTlsVersion)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_Spec.
func (in *RedisEnterprise_Spec) DeepCopy() *RedisEnterprise_Spec {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_Spec_ARM) DeepCopyInto(out *RedisEnterprise_Spec_ARM) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ClusterProperties_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_Spec_ARM.
func (in *RedisEnterprise_Spec_ARM) DeepCopy() *RedisEnterprise_Spec_ARM {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_Spec_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku) DeepCopyInto(out *Sku) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(Sku_Name)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku.
func (in *Sku) DeepCopy() *Sku {
	if in == nil {
		return nil
	}
	out := new(Sku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_ARM) DeepCopyInto(out *Sku_ARM) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(Sku_Name_ARM)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_ARM.
func (in *Sku_ARM) DeepCopy() *Sku_ARM {
	if in == nil {
		return nil
	}
	out := new(Sku_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_STATUS) DeepCopyInto(out *Sku_STATUS) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(Sku_Name_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_STATUS.
func (in *Sku_STATUS) DeepCopy() *Sku_STATUS {
	if in == nil {
		return nil
	}
	out := new(Sku_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_STATUS_ARM) DeepCopyInto(out *Sku_STATUS_ARM) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(Sku_Name_STATUS_ARM)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_STATUS_ARM.
func (in *Sku_STATUS_ARM) DeepCopy() *Sku_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(Sku_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}
