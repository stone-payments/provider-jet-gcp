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

type DocumentObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type DocumentParameters struct {

	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	// +kubebuilder:validation:Required
	Collection *string `json:"collection" tf:"collection,omitempty"`

	// The Firestore database id. Defaults to '"(default)"'.
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The client-assigned document ID to use for this document during creation.
	// +kubebuilder:validation:Required
	DocumentID *string `json:"documentId" tf:"document_id,omitempty"`

	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
	// +kubebuilder:validation:Required
	Fields *string `json:"fields" tf:"fields,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// DocumentSpec defines the desired state of Document
type DocumentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DocumentParameters `json:"forProvider"`
}

// DocumentStatus defines the observed state of Document.
type DocumentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DocumentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Document is the Schema for the Documents API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Document struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocumentSpec   `json:"spec"`
	Status            DocumentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentList contains a list of Documents
type DocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Document `json:"items"`
}

// Repository type metadata.
var (
	Document_Kind             = "Document"
	Document_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Document_Kind}.String()
	Document_KindAPIVersion   = Document_Kind + "." + CRDGroupVersion.String()
	Document_GroupVersionKind = CRDGroupVersion.WithKind(Document_Kind)
)

func init() {
	SchemeBuilder.Register(&Document{}, &DocumentList{})
}
