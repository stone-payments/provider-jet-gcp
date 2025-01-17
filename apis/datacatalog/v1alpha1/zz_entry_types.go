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

type BigqueryDateShardedSpecObservation struct {
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	ShardCount *int64 `json:"shardCount,omitempty" tf:"shard_count,omitempty"`

	TablePrefix *string `json:"tablePrefix,omitempty" tf:"table_prefix,omitempty"`
}

type BigqueryDateShardedSpecParameters struct {
}

type BigqueryTableSpecObservation struct {
	TableSourceType *string `json:"tableSourceType,omitempty" tf:"table_source_type,omitempty"`

	TableSpec []TableSpecObservation `json:"tableSpec,omitempty" tf:"table_spec,omitempty"`

	ViewSpec []ViewSpecObservation `json:"viewSpec,omitempty" tf:"view_spec,omitempty"`
}

type BigqueryTableSpecParameters struct {
}

type EntryObservation struct {
	BigqueryDateShardedSpec []BigqueryDateShardedSpecObservation `json:"bigqueryDateShardedSpec,omitempty" tf:"bigquery_date_sharded_spec,omitempty"`

	BigqueryTableSpec []BigqueryTableSpecObservation `json:"bigqueryTableSpec,omitempty" tf:"bigquery_table_spec,omitempty"`

	IntegratedSystem *string `json:"integratedSystem,omitempty" tf:"integrated_system,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type EntryParameters struct {

	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display information such as title and description. A short name to identify the entry,
	// for example, "Analytics Data - Jan 2011".
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the entry group this entry is in.
	// +kubebuilder:validation:Required
	EntryGroup *string `json:"entryGroup" tf:"entry_group,omitempty"`

	// The id of the entry to create.
	// +kubebuilder:validation:Required
	EntryID *string `json:"entryId" tf:"entry_id,omitempty"`

	// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
	// +kubebuilder:validation:Optional
	GcsFilesetSpec []GcsFilesetSpecParameters `json:"gcsFilesetSpec,omitempty" tf:"gcs_fileset_spec,omitempty"`

	// The resource this metadata entry refers to.
	// For Google Cloud Platform resources, linkedResource is the full name of the resource.
	// For example, the linkedResource for a table resource from BigQuery is:
	// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	// this field is optional and defaults to an empty string.
	// +kubebuilder:validation:Optional
	LinkedResource *string `json:"linkedResource,omitempty" tf:"linked_resource,omitempty"`

	// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
	// attached to it. See
	// https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
	// for what fields this schema can contain.
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// The type of the entry. Only used for Entries with types in the EntryType enum.
	// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType. Possible values: ["FILESET"]
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// This field indicates the entry's source system that Data Catalog does not integrate with.
	// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	// +kubebuilder:validation:Optional
	UserSpecifiedSystem *string `json:"userSpecifiedSystem,omitempty" tf:"user_specified_system,omitempty"`

	// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	// When creating an entry, users should check the enum values first, if nothing matches the entry
	// to be created, then provide a custom value, for example "my_special_type".
	// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	// +kubebuilder:validation:Optional
	UserSpecifiedType *string `json:"userSpecifiedType,omitempty" tf:"user_specified_type,omitempty"`
}

type GcsFilesetSpecObservation struct {
	SampleGcsFileSpecs []SampleGcsFileSpecsObservation `json:"sampleGcsFileSpecs,omitempty" tf:"sample_gcs_file_specs,omitempty"`
}

type GcsFilesetSpecParameters struct {

	// Patterns to identify a set of files in Google Cloud Storage.
	// See [Cloud Storage documentation](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames)
	// for more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:
	//
	// * gs://bucket_name/dir/*: matches all files within bucket_name/dir directory.
	// * gs://bucket_name/dir/**: matches all files in bucket_name/dir spanning all subdirectories.
	// * gs://bucket_name/file*: matches files prefixed by file in bucket_name
	// * gs://bucket_name/??.txt: matches files with two characters followed by .txt in bucket_name
	// * gs://bucket_name/[aeiou].txt: matches files that contain a single vowel character followed by .txt in bucket_name
	// * gs://bucket_name/[a-m].txt: matches files that contain a, b, ... or m followed by .txt in bucket_name
	// * gs://bucket_name/a/*/b: matches all files in bucket_name that match a/*/b pattern, such as a/c/b, a/d/b
	// * gs://another_bucket/a.txt: matches gs://another_bucket/a.txt
	// +kubebuilder:validation:Required
	FilePatterns []*string `json:"filePatterns" tf:"file_patterns,omitempty"`
}

type SampleGcsFileSpecsObservation struct {
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	SizeBytes *int64 `json:"sizeBytes,omitempty" tf:"size_bytes,omitempty"`
}

type SampleGcsFileSpecsParameters struct {
}

type TableSpecObservation struct {
	GroupedEntry *string `json:"groupedEntry,omitempty" tf:"grouped_entry,omitempty"`
}

type TableSpecParameters struct {
}

type ViewSpecObservation struct {
	ViewQuery *string `json:"viewQuery,omitempty" tf:"view_query,omitempty"`
}

type ViewSpecParameters struct {
}

// EntrySpec defines the desired state of Entry
type EntrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EntryParameters `json:"forProvider"`
}

// EntryStatus defines the observed state of Entry.
type EntryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EntryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Entry is the Schema for the Entrys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Entry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EntrySpec   `json:"spec"`
	Status            EntryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EntryList contains a list of Entrys
type EntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Entry `json:"items"`
}

// Repository type metadata.
var (
	Entry_Kind             = "Entry"
	Entry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Entry_Kind}.String()
	Entry_KindAPIVersion   = Entry_Kind + "." + CRDGroupVersion.String()
	Entry_GroupVersionKind = CRDGroupVersion.WithKind(Entry_Kind)
)

func init() {
	SchemeBuilder.Register(&Entry{}, &EntryList{})
}
