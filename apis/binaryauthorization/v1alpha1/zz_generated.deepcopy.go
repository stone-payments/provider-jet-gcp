// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
func (in *AdmissionWhitelistPatternsObservation) DeepCopyInto(out *AdmissionWhitelistPatternsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionWhitelistPatternsObservation.
func (in *AdmissionWhitelistPatternsObservation) DeepCopy() *AdmissionWhitelistPatternsObservation {
	if in == nil {
		return nil
	}
	out := new(AdmissionWhitelistPatternsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionWhitelistPatternsParameters) DeepCopyInto(out *AdmissionWhitelistPatternsParameters) {
	*out = *in
	if in.NamePattern != nil {
		in, out := &in.NamePattern, &out.NamePattern
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionWhitelistPatternsParameters.
func (in *AdmissionWhitelistPatternsParameters) DeepCopy() *AdmissionWhitelistPatternsParameters {
	if in == nil {
		return nil
	}
	out := new(AdmissionWhitelistPatternsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestationAuthorityNoteObservation) DeepCopyInto(out *AttestationAuthorityNoteObservation) {
	*out = *in
	if in.DelegationServiceAccountEmail != nil {
		in, out := &in.DelegationServiceAccountEmail, &out.DelegationServiceAccountEmail
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestationAuthorityNoteObservation.
func (in *AttestationAuthorityNoteObservation) DeepCopy() *AttestationAuthorityNoteObservation {
	if in == nil {
		return nil
	}
	out := new(AttestationAuthorityNoteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestationAuthorityNoteParameters) DeepCopyInto(out *AttestationAuthorityNoteParameters) {
	*out = *in
	if in.NoteReference != nil {
		in, out := &in.NoteReference, &out.NoteReference
		*out = new(string)
		**out = **in
	}
	if in.PublicKeys != nil {
		in, out := &in.PublicKeys, &out.PublicKeys
		*out = make([]PublicKeysParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestationAuthorityNoteParameters.
func (in *AttestationAuthorityNoteParameters) DeepCopy() *AttestationAuthorityNoteParameters {
	if in == nil {
		return nil
	}
	out := new(AttestationAuthorityNoteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attestor) DeepCopyInto(out *Attestor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attestor.
func (in *Attestor) DeepCopy() *Attestor {
	if in == nil {
		return nil
	}
	out := new(Attestor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Attestor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamBinding) DeepCopyInto(out *AttestorIamBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamBinding.
func (in *AttestorIamBinding) DeepCopy() *AttestorIamBinding {
	if in == nil {
		return nil
	}
	out := new(AttestorIamBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttestorIamBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamBindingList) DeepCopyInto(out *AttestorIamBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AttestorIamBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamBindingList.
func (in *AttestorIamBindingList) DeepCopy() *AttestorIamBindingList {
	if in == nil {
		return nil
	}
	out := new(AttestorIamBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttestorIamBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamBindingObservation) DeepCopyInto(out *AttestorIamBindingObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamBindingObservation.
func (in *AttestorIamBindingObservation) DeepCopy() *AttestorIamBindingObservation {
	if in == nil {
		return nil
	}
	out := new(AttestorIamBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamBindingParameters) DeepCopyInto(out *AttestorIamBindingParameters) {
	*out = *in
	if in.Attestor != nil {
		in, out := &in.Attestor, &out.Attestor
		*out = new(string)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]ConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamBindingParameters.
func (in *AttestorIamBindingParameters) DeepCopy() *AttestorIamBindingParameters {
	if in == nil {
		return nil
	}
	out := new(AttestorIamBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamBindingSpec) DeepCopyInto(out *AttestorIamBindingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamBindingSpec.
func (in *AttestorIamBindingSpec) DeepCopy() *AttestorIamBindingSpec {
	if in == nil {
		return nil
	}
	out := new(AttestorIamBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamBindingStatus) DeepCopyInto(out *AttestorIamBindingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamBindingStatus.
func (in *AttestorIamBindingStatus) DeepCopy() *AttestorIamBindingStatus {
	if in == nil {
		return nil
	}
	out := new(AttestorIamBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamMember) DeepCopyInto(out *AttestorIamMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamMember.
func (in *AttestorIamMember) DeepCopy() *AttestorIamMember {
	if in == nil {
		return nil
	}
	out := new(AttestorIamMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttestorIamMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamMemberConditionObservation) DeepCopyInto(out *AttestorIamMemberConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamMemberConditionObservation.
func (in *AttestorIamMemberConditionObservation) DeepCopy() *AttestorIamMemberConditionObservation {
	if in == nil {
		return nil
	}
	out := new(AttestorIamMemberConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamMemberConditionParameters) DeepCopyInto(out *AttestorIamMemberConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamMemberConditionParameters.
func (in *AttestorIamMemberConditionParameters) DeepCopy() *AttestorIamMemberConditionParameters {
	if in == nil {
		return nil
	}
	out := new(AttestorIamMemberConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamMemberList) DeepCopyInto(out *AttestorIamMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AttestorIamMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamMemberList.
func (in *AttestorIamMemberList) DeepCopy() *AttestorIamMemberList {
	if in == nil {
		return nil
	}
	out := new(AttestorIamMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttestorIamMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamMemberObservation) DeepCopyInto(out *AttestorIamMemberObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamMemberObservation.
func (in *AttestorIamMemberObservation) DeepCopy() *AttestorIamMemberObservation {
	if in == nil {
		return nil
	}
	out := new(AttestorIamMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamMemberParameters) DeepCopyInto(out *AttestorIamMemberParameters) {
	*out = *in
	if in.Attestor != nil {
		in, out := &in.Attestor, &out.Attestor
		*out = new(string)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]AttestorIamMemberConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamMemberParameters.
func (in *AttestorIamMemberParameters) DeepCopy() *AttestorIamMemberParameters {
	if in == nil {
		return nil
	}
	out := new(AttestorIamMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamMemberSpec) DeepCopyInto(out *AttestorIamMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamMemberSpec.
func (in *AttestorIamMemberSpec) DeepCopy() *AttestorIamMemberSpec {
	if in == nil {
		return nil
	}
	out := new(AttestorIamMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamMemberStatus) DeepCopyInto(out *AttestorIamMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamMemberStatus.
func (in *AttestorIamMemberStatus) DeepCopy() *AttestorIamMemberStatus {
	if in == nil {
		return nil
	}
	out := new(AttestorIamMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamPolicy) DeepCopyInto(out *AttestorIamPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamPolicy.
func (in *AttestorIamPolicy) DeepCopy() *AttestorIamPolicy {
	if in == nil {
		return nil
	}
	out := new(AttestorIamPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttestorIamPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamPolicyList) DeepCopyInto(out *AttestorIamPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AttestorIamPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamPolicyList.
func (in *AttestorIamPolicyList) DeepCopy() *AttestorIamPolicyList {
	if in == nil {
		return nil
	}
	out := new(AttestorIamPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttestorIamPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamPolicyObservation) DeepCopyInto(out *AttestorIamPolicyObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamPolicyObservation.
func (in *AttestorIamPolicyObservation) DeepCopy() *AttestorIamPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(AttestorIamPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamPolicyParameters) DeepCopyInto(out *AttestorIamPolicyParameters) {
	*out = *in
	if in.Attestor != nil {
		in, out := &in.Attestor, &out.Attestor
		*out = new(string)
		**out = **in
	}
	if in.PolicyData != nil {
		in, out := &in.PolicyData, &out.PolicyData
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamPolicyParameters.
func (in *AttestorIamPolicyParameters) DeepCopy() *AttestorIamPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(AttestorIamPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamPolicySpec) DeepCopyInto(out *AttestorIamPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamPolicySpec.
func (in *AttestorIamPolicySpec) DeepCopy() *AttestorIamPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AttestorIamPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorIamPolicyStatus) DeepCopyInto(out *AttestorIamPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorIamPolicyStatus.
func (in *AttestorIamPolicyStatus) DeepCopy() *AttestorIamPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AttestorIamPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorList) DeepCopyInto(out *AttestorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Attestor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorList.
func (in *AttestorList) DeepCopy() *AttestorList {
	if in == nil {
		return nil
	}
	out := new(AttestorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttestorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorObservation) DeepCopyInto(out *AttestorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorObservation.
func (in *AttestorObservation) DeepCopy() *AttestorObservation {
	if in == nil {
		return nil
	}
	out := new(AttestorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorParameters) DeepCopyInto(out *AttestorParameters) {
	*out = *in
	if in.AttestationAuthorityNote != nil {
		in, out := &in.AttestationAuthorityNote, &out.AttestationAuthorityNote
		*out = make([]AttestationAuthorityNoteParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorParameters.
func (in *AttestorParameters) DeepCopy() *AttestorParameters {
	if in == nil {
		return nil
	}
	out := new(AttestorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorSpec) DeepCopyInto(out *AttestorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorSpec.
func (in *AttestorSpec) DeepCopy() *AttestorSpec {
	if in == nil {
		return nil
	}
	out := new(AttestorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttestorStatus) DeepCopyInto(out *AttestorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttestorStatus.
func (in *AttestorStatus) DeepCopy() *AttestorStatus {
	if in == nil {
		return nil
	}
	out := new(AttestorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdmissionRulesObservation) DeepCopyInto(out *ClusterAdmissionRulesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdmissionRulesObservation.
func (in *ClusterAdmissionRulesObservation) DeepCopy() *ClusterAdmissionRulesObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterAdmissionRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAdmissionRulesParameters) DeepCopyInto(out *ClusterAdmissionRulesParameters) {
	*out = *in
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.EnforcementMode != nil {
		in, out := &in.EnforcementMode, &out.EnforcementMode
		*out = new(string)
		**out = **in
	}
	if in.EvaluationMode != nil {
		in, out := &in.EvaluationMode, &out.EvaluationMode
		*out = new(string)
		**out = **in
	}
	if in.RequireAttestationsBy != nil {
		in, out := &in.RequireAttestationsBy, &out.RequireAttestationsBy
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAdmissionRulesParameters.
func (in *ClusterAdmissionRulesParameters) DeepCopy() *ClusterAdmissionRulesParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterAdmissionRulesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionObservation) DeepCopyInto(out *ConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionObservation.
func (in *ConditionObservation) DeepCopy() *ConditionObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionParameters) DeepCopyInto(out *ConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionParameters.
func (in *ConditionParameters) DeepCopy() *ConditionParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultAdmissionRuleObservation) DeepCopyInto(out *DefaultAdmissionRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultAdmissionRuleObservation.
func (in *DefaultAdmissionRuleObservation) DeepCopy() *DefaultAdmissionRuleObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultAdmissionRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultAdmissionRuleParameters) DeepCopyInto(out *DefaultAdmissionRuleParameters) {
	*out = *in
	if in.EnforcementMode != nil {
		in, out := &in.EnforcementMode, &out.EnforcementMode
		*out = new(string)
		**out = **in
	}
	if in.EvaluationMode != nil {
		in, out := &in.EvaluationMode, &out.EvaluationMode
		*out = new(string)
		**out = **in
	}
	if in.RequireAttestationsBy != nil {
		in, out := &in.RequireAttestationsBy, &out.RequireAttestationsBy
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultAdmissionRuleParameters.
func (in *DefaultAdmissionRuleParameters) DeepCopy() *DefaultAdmissionRuleParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultAdmissionRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PkixPublicKeyObservation) DeepCopyInto(out *PkixPublicKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PkixPublicKeyObservation.
func (in *PkixPublicKeyObservation) DeepCopy() *PkixPublicKeyObservation {
	if in == nil {
		return nil
	}
	out := new(PkixPublicKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PkixPublicKeyParameters) DeepCopyInto(out *PkixPublicKeyParameters) {
	*out = *in
	if in.PublicKeyPem != nil {
		in, out := &in.PublicKeyPem, &out.PublicKeyPem
		*out = new(string)
		**out = **in
	}
	if in.SignatureAlgorithm != nil {
		in, out := &in.SignatureAlgorithm, &out.SignatureAlgorithm
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PkixPublicKeyParameters.
func (in *PkixPublicKeyParameters) DeepCopy() *PkixPublicKeyParameters {
	if in == nil {
		return nil
	}
	out := new(PkixPublicKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyObservation) DeepCopyInto(out *PolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyObservation.
func (in *PolicyObservation) DeepCopy() *PolicyObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyParameters) DeepCopyInto(out *PolicyParameters) {
	*out = *in
	if in.AdmissionWhitelistPatterns != nil {
		in, out := &in.AdmissionWhitelistPatterns, &out.AdmissionWhitelistPatterns
		*out = make([]AdmissionWhitelistPatternsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterAdmissionRules != nil {
		in, out := &in.ClusterAdmissionRules, &out.ClusterAdmissionRules
		*out = make([]ClusterAdmissionRulesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultAdmissionRule != nil {
		in, out := &in.DefaultAdmissionRule, &out.DefaultAdmissionRule
		*out = make([]DefaultAdmissionRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.GlobalPolicyEvaluationMode != nil {
		in, out := &in.GlobalPolicyEvaluationMode, &out.GlobalPolicyEvaluationMode
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyParameters.
func (in *PolicyParameters) DeepCopy() *PolicyParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicKeysObservation) DeepCopyInto(out *PublicKeysObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicKeysObservation.
func (in *PublicKeysObservation) DeepCopy() *PublicKeysObservation {
	if in == nil {
		return nil
	}
	out := new(PublicKeysObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicKeysParameters) DeepCopyInto(out *PublicKeysParameters) {
	*out = *in
	if in.ASCIIArmoredPgpPublicKey != nil {
		in, out := &in.ASCIIArmoredPgpPublicKey, &out.ASCIIArmoredPgpPublicKey
		*out = new(string)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PkixPublicKey != nil {
		in, out := &in.PkixPublicKey, &out.PkixPublicKey
		*out = make([]PkixPublicKeyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicKeysParameters.
func (in *PublicKeysParameters) DeepCopy() *PublicKeysParameters {
	if in == nil {
		return nil
	}
	out := new(PublicKeysParameters)
	in.DeepCopyInto(out)
	return out
}
