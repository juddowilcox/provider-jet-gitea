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

type OrgObservation struct {
	AvatarURL *string `json:"avatarUrl,omitempty" tf:"avatar_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OrgParameters struct {

	// A description of this organisation.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the organisation. Defaults to the value of `name`.
	// +kubebuilder:validation:Optional
	FullName *string `json:"fullName,omitempty" tf:"full_name,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the organisation without spaces.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RepoAdminChangeTeamAccess *bool `json:"repoAdminChangeTeamAccess,omitempty" tf:"repo_admin_change_team_access,omitempty"`

	// Flag is this organisation should be publicly visible or not.
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// A link to a website with more information about this organisation.
	// +kubebuilder:validation:Optional
	Website *string `json:"website,omitempty" tf:"website,omitempty"`
}

// OrgSpec defines the desired state of Org
type OrgSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrgParameters `json:"forProvider"`
}

// OrgStatus defines the observed state of Org.
type OrgStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrgObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Org is the Schema for the Orgs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,giteajet}
type Org struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrgSpec   `json:"spec"`
	Status            OrgStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrgList contains a list of Orgs
type OrgList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Org `json:"items"`
}

// Repository type metadata.
var (
	Org_Kind             = "Org"
	Org_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Org_Kind}.String()
	Org_KindAPIVersion   = Org_Kind + "." + CRDGroupVersion.String()
	Org_GroupVersionKind = CRDGroupVersion.WithKind(Org_Kind)
)

func init() {
	SchemeBuilder.Register(&Org{}, &OrgList{})
}
