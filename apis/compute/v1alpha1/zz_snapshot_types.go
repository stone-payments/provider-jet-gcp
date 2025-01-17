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

type SnapshotEncryptionKeyObservation struct {
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type SnapshotEncryptionKeyParameters struct {

	// The name of the encryption key that is stored in Google Cloud KMS.
	// +kubebuilder:validation:Optional
	KmsKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	// +kubebuilder:validation:Optional
	KmsKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	// +kubebuilder:validation:Optional
	RawKeySecretRef *v1.SecretKeySelector `json:"rawKeySecretRef,omitempty" tf:"-"`
}

type SnapshotObservation struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	Licenses []*string `json:"licenses,omitempty" tf:"licenses,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	SnapshotID *int64 `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	StorageBytes *int64 `json:"storageBytes,omitempty" tf:"storage_bytes,omitempty"`
}

type SnapshotParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Labels to apply to this Snapshot.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The customer-supplied encryption key of the snapshot. Required if the
	// source snapshot is protected by a customer-supplied encryption key.
	// +kubebuilder:validation:Optional
	SnapshotEncryptionKey []SnapshotEncryptionKeyParameters `json:"snapshotEncryptionKey,omitempty" tf:"snapshot_encryption_key,omitempty"`

	// A reference to the disk used to create this snapshot.
	// +kubebuilder:validation:Required
	SourceDisk *string `json:"sourceDisk" tf:"source_disk,omitempty"`

	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// +kubebuilder:validation:Optional
	SourceDiskEncryptionKey []SourceDiskEncryptionKeyParameters `json:"sourceDiskEncryptionKey,omitempty" tf:"source_disk_encryption_key,omitempty"`

	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	// +kubebuilder:validation:Optional
	StorageLocations []*string `json:"storageLocations,omitempty" tf:"storage_locations,omitempty"`

	// A reference to the zone where the disk is hosted.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SourceDiskEncryptionKeyObservation struct {
}

type SourceDiskEncryptionKeyParameters struct {

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	// +kubebuilder:validation:Optional
	KmsKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	// +kubebuilder:validation:Optional
	RawKeySecretRef *v1.SecretKeySelector `json:"rawKeySecretRef,omitempty" tf:"-"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Snapshot is the Schema for the Snapshots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
