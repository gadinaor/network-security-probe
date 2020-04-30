// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressRule) DeepCopyInto(out *EgressRule) {
	*out = *in
	if in.FQDNs != nil {
		in, out := &in.FQDNs, &out.FQDNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CIDR4 != nil {
		in, out := &in.CIDR4, &out.CIDR4
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CIDR6 != nil {
		in, out := &in.CIDR6, &out.CIDR6
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.L3.DeepCopyInto(&out.L3)
	in.L4.DeepCopyInto(&out.L4)
	in.L7.DeepCopyInto(&out.L7)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressRule.
func (in *EgressRule) DeepCopy() *EgressRule {
	if in == nil {
		return nil
	}
	out := new(EgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRule) DeepCopyInto(out *HTTPRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRule.
func (in *HTTPRule) DeepCopy() *HTTPRule {
	if in == nil {
		return nil
	}
	out := new(HTTPRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
	*out = *in
	if in.CIDR4 != nil {
		in, out := &in.CIDR4, &out.CIDR4
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CIDR6 != nil {
		in, out := &in.CIDR6, &out.CIDR6
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.L3.DeepCopyInto(&out.L3)
	in.L4.DeepCopyInto(&out.L4)
	in.L7.DeepCopyInto(&out.L7)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRule.
func (in *IngressRule) DeepCopy() *IngressRule {
	if in == nil {
		return nil
	}
	out := new(IngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L3Rule) DeepCopyInto(out *L3Rule) {
	*out = *in
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L3Rule.
func (in *L3Rule) DeepCopy() *L3Rule {
	if in == nil {
		return nil
	}
	out := new(L3Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4ProtocolPortRule) DeepCopyInto(out *L4ProtocolPortRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4ProtocolPortRule.
func (in *L4ProtocolPortRule) DeepCopy() *L4ProtocolPortRule {
	if in == nil {
		return nil
	}
	out := new(L4ProtocolPortRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4Rule) DeepCopyInto(out *L4Rule) {
	*out = *in
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProtocolPorts != nil {
		in, out := &in.ProtocolPorts, &out.ProtocolPorts
		*out = make([]L4ProtocolPortRule, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4Rule.
func (in *L4Rule) DeepCopy() *L4Rule {
	if in == nil {
		return nil
	}
	out := new(L4Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L7Rule) DeepCopyInto(out *L7Rule) {
	*out = *in
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = make([]HTTPRule, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L7Rule.
func (in *L7Rule) DeepCopy() *L7Rule {
	if in == nil {
		return nil
	}
	out := new(L7Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicy) DeepCopyInto(out *NetworkPolicy) {
	*out = *in
	in.Egress.DeepCopyInto(&out.Egress)
	in.Ingress.DeepCopyInto(&out.Ingress)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicy.
func (in *NetworkPolicy) DeepCopy() *NetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessProfile) DeepCopyInto(out *ProcessProfile) {
	*out = *in
	in.NetworkPolicy.DeepCopyInto(&out.NetworkPolicy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessProfile.
func (in *ProcessProfile) DeepCopy() *ProcessProfile {
	if in == nil {
		return nil
	}
	out := new(ProcessProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityProfile) DeepCopyInto(out *SecurityProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityProfile.
func (in *SecurityProfile) DeepCopy() *SecurityProfile {
	if in == nil {
		return nil
	}
	out := new(SecurityProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityProfileList) DeepCopyInto(out *SecurityProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityProfileList.
func (in *SecurityProfileList) DeepCopy() *SecurityProfileList {
	if in == nil {
		return nil
	}
	out := new(SecurityProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityProfileSpec) DeepCopyInto(out *SecurityProfileSpec) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NetworkAttacks != nil {
		in, out := &in.NetworkAttacks, &out.NetworkAttacks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.DefaultNetworkPolicy.DeepCopyInto(&out.DefaultNetworkPolicy)
	if in.ProcessProfiles != nil {
		in, out := &in.ProcessProfiles, &out.ProcessProfiles
		*out = make([]*ProcessProfile, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ProcessProfile)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityProfileSpec.
func (in *SecurityProfileSpec) DeepCopy() *SecurityProfileSpec {
	if in == nil {
		return nil
	}
	out := new(SecurityProfileSpec)
	in.DeepCopyInto(out)
	return out
}