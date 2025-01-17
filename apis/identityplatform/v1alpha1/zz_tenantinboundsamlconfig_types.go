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

type IdpConfigIdpCertificatesObservation struct {
}

type IdpConfigIdpCertificatesParameters struct {

	// The x509 certificate
	// +kubebuilder:validation:Optional
	X509Certificate *string `json:"x509Certificate,omitempty" tf:"x509_certificate,omitempty"`
}

type SpConfigSpCertificatesObservation struct {
	X509Certificate *string `json:"x509Certificate,omitempty" tf:"x509_certificate,omitempty"`
}

type SpConfigSpCertificatesParameters struct {
}

type TenantInboundSamlConfigIdpConfigObservation struct {
}

type TenantInboundSamlConfigIdpConfigParameters struct {

	// The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
	// +kubebuilder:validation:Required
	IdpCertificates []IdpConfigIdpCertificatesParameters `json:"idpCertificates" tf:"idp_certificates,omitempty"`

	// Unique identifier for all SAML entities
	// +kubebuilder:validation:Required
	IdpEntityID *string `json:"idpEntityId" tf:"idp_entity_id,omitempty"`

	// Indicates if outbounding SAMLRequest should be signed.
	// +kubebuilder:validation:Optional
	SignRequest *bool `json:"signRequest,omitempty" tf:"sign_request,omitempty"`

	// URL to send Authentication request to.
	// +kubebuilder:validation:Required
	SsoURL *string `json:"ssoUrl" tf:"sso_url,omitempty"`
}

type TenantInboundSamlConfigObservation struct {
}

type TenantInboundSamlConfigParameters struct {

	// Human friendly display name.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// SAML IdP configuration when the project acts as the relying party
	// +kubebuilder:validation:Required
	IdpConfig []TenantInboundSamlConfigIdpConfigParameters `json:"idpConfig" tf:"idp_config,omitempty"`

	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// +kubebuilder:validation:Required
	SpConfig []TenantInboundSamlConfigSpConfigParameters `json:"spConfig" tf:"sp_config,omitempty"`

	// The name of the tenant where this inbound SAML config resource exists
	// +kubebuilder:validation:Required
	Tenant *string `json:"tenant" tf:"tenant,omitempty"`
}

type TenantInboundSamlConfigSpConfigObservation struct {
	SpCertificates []SpConfigSpCertificatesObservation `json:"spCertificates,omitempty" tf:"sp_certificates,omitempty"`
}

type TenantInboundSamlConfigSpConfigParameters struct {

	// Callback URI where responses from IDP are handled. Must start with 'https://'.
	// +kubebuilder:validation:Required
	CallbackURI *string `json:"callbackUri" tf:"callback_uri,omitempty"`

	// Unique identifier for all SAML entities.
	// +kubebuilder:validation:Required
	SpEntityID *string `json:"spEntityId" tf:"sp_entity_id,omitempty"`
}

// TenantInboundSamlConfigSpec defines the desired state of TenantInboundSamlConfig
type TenantInboundSamlConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TenantInboundSamlConfigParameters `json:"forProvider"`
}

// TenantInboundSamlConfigStatus defines the observed state of TenantInboundSamlConfig.
type TenantInboundSamlConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TenantInboundSamlConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TenantInboundSamlConfig is the Schema for the TenantInboundSamlConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type TenantInboundSamlConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TenantInboundSamlConfigSpec   `json:"spec"`
	Status            TenantInboundSamlConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantInboundSamlConfigList contains a list of TenantInboundSamlConfigs
type TenantInboundSamlConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TenantInboundSamlConfig `json:"items"`
}

// Repository type metadata.
var (
	TenantInboundSamlConfig_Kind             = "TenantInboundSamlConfig"
	TenantInboundSamlConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TenantInboundSamlConfig_Kind}.String()
	TenantInboundSamlConfig_KindAPIVersion   = TenantInboundSamlConfig_Kind + "." + CRDGroupVersion.String()
	TenantInboundSamlConfig_GroupVersionKind = CRDGroupVersion.WithKind(TenantInboundSamlConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&TenantInboundSamlConfig{}, &TenantInboundSamlConfigList{})
}
