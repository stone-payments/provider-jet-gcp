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

type InterconnectAttachmentObservation struct {
	CloudRouterIPAddress *string `json:"cloudRouterIpAddress,omitempty" tf:"cloud_router_ip_address,omitempty"`

	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	CustomerRouterIPAddress *string `json:"customerRouterIpAddress,omitempty" tf:"customer_router_ip_address,omitempty"`

	GoogleReferenceID *string `json:"googleReferenceId,omitempty" tf:"google_reference_id,omitempty"`

	PairingKey *string `json:"pairingKey,omitempty" tf:"pairing_key,omitempty"`

	PartnerAsn *string `json:"partnerAsn,omitempty" tf:"partner_asn,omitempty"`

	PrivateInterconnectInfo []PrivateInterconnectInfoObservation `json:"privateInterconnectInfo,omitempty" tf:"private_interconnect_info,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type InterconnectAttachmentParameters struct {

	// Whether the VLAN attachment is enabled or disabled.  When using
	// PARTNER type this will Pre-Activate the interconnect attachment
	// +kubebuilder:validation:Optional
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`

	// Provisioned bandwidth capacity for the interconnect attachment.
	// For attachments of type DEDICATED, the user can set the bandwidth.
	// For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	// Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	// Defaults to BPS_10G Possible values: ["BPS_50M", "BPS_100M", "BPS_200M", "BPS_300M", "BPS_400M", "BPS_500M", "BPS_1G", "BPS_2G", "BPS_5G", "BPS_10G", "BPS_20G", "BPS_50G"]
	// +kubebuilder:validation:Optional
	Bandwidth *string `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// Up to 16 candidate prefixes that can be used to restrict the allocation
	// of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	// All prefixes must be within link-local address space (169.254.0.0/16)
	// and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	// an unused /29 from the supplied candidate prefix(es). The request will
	// fail if all possible /29s are in use on Google's edge. If not supplied,
	// Google will randomly select an unused /29 from all of link-local space.
	// +kubebuilder:validation:Optional
	CandidateSubnets []*string `json:"candidateSubnets,omitempty" tf:"candidate_subnets,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Desired availability domain for the attachment. Only available for type
	// PARTNER, at creation time. For improved reliability, customers should
	// configure a pair of attachments with one per availability domain. The
	// selected availability domain will be provided to the Partner via the
	// pairing key so that the provisioned circuit will lie in the specified
	// domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	// +kubebuilder:validation:Optional
	EdgeAvailabilityDomain *string `json:"edgeAvailabilityDomain,omitempty" tf:"edge_availability_domain,omitempty"`

	// Indicates the user-supplied encryption option of this interconnect
	// attachment:
	//
	// NONE is the default value, which means that the attachment carries
	// unencrypted traffic. VMs can send traffic to, or receive traffic
	// from, this type of attachment.
	//
	// IPSEC indicates that the attachment carries only traffic encrypted by
	// an IPsec device such as an HA VPN gateway. VMs cannot directly send
	// traffic to, or receive traffic from, such an attachment. To use
	// IPsec-encrypted Cloud Interconnect create the attachment using this
	// option.
	//
	// Not currently available publicly. Default value: "NONE" Possible values: ["NONE", "IPSEC"]
	// +kubebuilder:validation:Optional
	Encryption *string `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// URL of the underlying Interconnect object that this attachment's
	// traffic will traverse through. Required if type is DEDICATED, must not
	// be set if type is PARTNER.
	// +kubebuilder:validation:Optional
	Interconnect *string `json:"interconnect,omitempty" tf:"interconnect,omitempty"`

	// URL of addresses that have been reserved for the interconnect
	// attachment, Used only for interconnect attachment that has the
	// encryption option as IPSEC.
	//
	// The addresses must be RFC 1918 IP address ranges. When creating HA
	// VPN gateway over the interconnect attachment, if the attachment is
	// configured to use an RFC 1918 IP address, then the VPN gateway's IP
	// address will be allocated from the IP address range specified
	// here.
	//
	// For example, if the HA VPN gateway's interface 0 is paired to this
	// interconnect attachment, then an RFC 1918 IP address for the VPN
	// gateway interface 0 will be allocated from the IP address specified
	// for this interconnect attachment.
	//
	// If this field is not specified for interconnect attachment that has
	// encryption option as IPSEC, later on when creating HA VPN gateway on
	// this interconnect attachment, the HA VPN gateway's IP address will be
	// allocated from regional external IP address pool.
	// +kubebuilder:validation:Optional
	IpsecInternalAddresses []*string `json:"ipsecInternalAddresses,omitempty" tf:"ipsec_internal_addresses,omitempty"`

	// Maximum Transmission Unit (MTU), in bytes, of packets passing through
	// this interconnect attachment. Currently, only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
	// +kubebuilder:validation:Optional
	Mtu *string `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// Name of the resource. Provided by the client when the resource is created. The
	// name must be 1-63 characters long, and comply with RFC1035. Specifically, the
	// name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a
	// lowercase letter, and all following characters must be a dash, lowercase
	// letter, or digit, except the last character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the regional interconnect attachment resides.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// URL of the cloud router to be used for dynamic routing. This router must be in
	// the same region as this InterconnectAttachment. The InterconnectAttachment will
	// automatically connect the Interconnect to the network & region within which the
	// Cloud Router is configured.
	// +kubebuilder:validation:Required
	Router *string `json:"router" tf:"router,omitempty"`

	// The type of InterconnectAttachment you wish to create. Defaults to
	// DEDICATED. Possible values: ["DEDICATED", "PARTNER", "PARTNER_PROVIDER"]
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
	// using PARTNER type this will be managed upstream.
	// +kubebuilder:validation:Optional
	VlanTag8021Q *int64 `json:"vlanTag8021Q,omitempty" tf:"vlan_tag8021q,omitempty"`
}

type PrivateInterconnectInfoObservation struct {
	Tag8021Q *int64 `json:"tag8021q,omitempty" tf:"tag8021q,omitempty"`
}

type PrivateInterconnectInfoParameters struct {
}

// InterconnectAttachmentSpec defines the desired state of InterconnectAttachment
type InterconnectAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InterconnectAttachmentParameters `json:"forProvider"`
}

// InterconnectAttachmentStatus defines the observed state of InterconnectAttachment.
type InterconnectAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InterconnectAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InterconnectAttachment is the Schema for the InterconnectAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type InterconnectAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InterconnectAttachmentSpec   `json:"spec"`
	Status            InterconnectAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InterconnectAttachmentList contains a list of InterconnectAttachments
type InterconnectAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InterconnectAttachment `json:"items"`
}

// Repository type metadata.
var (
	InterconnectAttachment_Kind             = "InterconnectAttachment"
	InterconnectAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InterconnectAttachment_Kind}.String()
	InterconnectAttachment_KindAPIVersion   = InterconnectAttachment_Kind + "." + CRDGroupVersion.String()
	InterconnectAttachment_GroupVersionKind = CRDGroupVersion.WithKind(InterconnectAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&InterconnectAttachment{}, &InterconnectAttachmentList{})
}
