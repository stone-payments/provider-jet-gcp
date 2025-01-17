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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GlobalAddressObservation struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type GlobalAddressParameters struct {

	// The IP address or beginning of the address range represented by this
	// resource. This can be supplied as an input to reserve a specific
	// address or omitted to allow GCP to choose a valid one for you.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of the address to reserve.
	//
	// * EXTERNAL indicates public/external single IP address.
	// * INTERNAL indicates internal IP ranges belonging to some network. Default value: "EXTERNAL" Possible values: ["EXTERNAL", "INTERNAL"]
	// +kubebuilder:validation:Optional
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"]
	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The URL of the network in which to reserve the IP range. The IP range
	// must be in RFC1918 space. The network cannot be deleted if there are
	// any reserved IP ranges referring to it.
	//
	// This should only be set when using an Internal address.
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The prefix length of the IP range. If not present, it means the
	// address field is a single IP address.
	//
	// This field is not applicable to addresses with addressType=EXTERNAL,
	// or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
	// +kubebuilder:validation:Optional
	PrefixLength *int64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of the resource. Possible values include:
	//
	// * VPC_PEERING - for peer networks
	//
	// * PRIVATE_SERVICE_CONNECT - for ([Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html) only) Private Service Connect networks
	// +kubebuilder:validation:Optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`
}

// GlobalAddressSpec defines the desired state of GlobalAddress
type GlobalAddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalAddressParameters `json:"forProvider"`
}

// GlobalAddressStatus defines the observed state of GlobalAddress.
type GlobalAddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalAddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalAddress is the Schema for the GlobalAddresss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type GlobalAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalAddressSpec   `json:"spec"`
	Status            GlobalAddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalAddressList contains a list of GlobalAddresss
type GlobalAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalAddress `json:"items"`
}

// Repository type metadata.
var (
	GlobalAddress_Kind             = "GlobalAddress"
	GlobalAddress_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalAddress_Kind}.String()
	GlobalAddress_KindAPIVersion   = GlobalAddress_Kind + "." + CRDGroupVersion.String()
	GlobalAddress_GroupVersionKind = CRDGroupVersion.WithKind(GlobalAddress_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalAddress{}, &GlobalAddressList{})
}
