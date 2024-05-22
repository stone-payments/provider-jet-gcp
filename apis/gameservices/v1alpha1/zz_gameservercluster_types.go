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

type ConnectionInfoObservation struct {
}

type ConnectionInfoParameters struct {

	// Reference of the GKE cluster where the game servers are installed.
	// +kubebuilder:validation:Required
	GkeClusterReference []GkeClusterReferenceParameters `json:"gkeClusterReference" tf:"gke_cluster_reference,omitempty"`

	// Namespace designated on the game server cluster where the game server
	// instances will be created. The namespace existence will be validated
	// during creation.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type GameServerClusterObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GameServerClusterParameters struct {

	// Required. The resource name of the game server cluster
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// +kubebuilder:validation:Required
	ConnectionInfo []ConnectionInfoParameters `json:"connectionInfo" tf:"connection_info,omitempty"`

	// Human readable description of the cluster.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Location of the Cluster.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The realm id of the game server realm.
	// +kubebuilder:validation:Required
	RealmID *string `json:"realmId" tf:"realm_id,omitempty"`
}

type GkeClusterReferenceObservation struct {
}

type GkeClusterReferenceParameters struct {

	// The full or partial name of a GKE cluster, using one of the following
	// forms:
	//
	// * 'projects/{project_id}/locations/{location}/clusters/{cluster_id}'
	// * 'locations/{location}/clusters/{cluster_id}'
	// * '{cluster_id}'
	//
	// If project and location are not specified, the project and location of the
	// GameServerCluster resource are used to generate the full name of the
	// GKE cluster.
	// +kubebuilder:validation:Required
	Cluster *string `json:"cluster" tf:"cluster,omitempty"`
}

// GameServerClusterSpec defines the desired state of GameServerCluster
type GameServerClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GameServerClusterParameters `json:"forProvider"`
}

// GameServerClusterStatus defines the observed state of GameServerCluster.
type GameServerClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GameServerClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GameServerCluster is the Schema for the GameServerClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type GameServerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameServerClusterSpec   `json:"spec"`
	Status            GameServerClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameServerClusterList contains a list of GameServerClusters
type GameServerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameServerCluster `json:"items"`
}

// Repository type metadata.
var (
	GameServerCluster_Kind             = "GameServerCluster"
	GameServerCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GameServerCluster_Kind}.String()
	GameServerCluster_KindAPIVersion   = GameServerCluster_Kind + "." + CRDGroupVersion.String()
	GameServerCluster_GroupVersionKind = CRDGroupVersion.WithKind(GameServerCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&GameServerCluster{}, &GameServerClusterList{})
}