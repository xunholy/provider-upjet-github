/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RunnerGroupInitParameters struct {

	// Whether public repositories can be added to the runner group. Defaults to false.
	// Whether public repositories can be added to the runner group.
	AllowsPublicRepositories *bool `json:"allowsPublicRepositories,omitempty" tf:"allows_public_repositories,omitempty"`

	// Name of the runner group
	// Name of the runner group.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/repo/v1alpha1.Repository
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Repository in repo to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
	// If 'true', the runner group will be restricted to running only the workflows specified in the 'selected_workflows' array. Defaults to 'false'.
	RestrictedToWorkflows *bool `json:"restrictedToWorkflows,omitempty" tf:"restricted_to_workflows,omitempty"`

	// IDs of the repositories which should be added to the runner group
	// List of repository IDs that can access the runner group.
	// +listType=set
	SelectedRepositoryIds []*int64 `json:"selectedRepositoryIds,omitempty" tf:"selected_repository_ids,omitempty"`

	// List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
	// List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to 'true'.
	SelectedWorkflows []*string `json:"selectedWorkflows,omitempty" tf:"selected_workflows,omitempty"`

	// Visibility of a runner group. Whether the runner group can include all, selected, or private repositories. A value of private is not currently supported due to limitations in the GitHub API.
	// The visibility of the runner group.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type RunnerGroupObservation struct {

	// Whether public repositories can be added to the runner group. Defaults to false.
	// Whether public repositories can be added to the runner group.
	AllowsPublicRepositories *bool `json:"allowsPublicRepositories,omitempty" tf:"allows_public_repositories,omitempty"`

	// Whether this is the default runner group
	// Whether this is the default runner group.
	Default *bool `json:"default,omitempty" tf:"default,omitempty"`

	// An etag representing the runner group object
	// An etag representing the runner group object
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the runner group is inherited from the enterprise level
	// Whether the runner group is inherited from the enterprise level
	Inherited *bool `json:"inherited,omitempty" tf:"inherited,omitempty"`

	// Name of the runner group
	// Name of the runner group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
	// If 'true', the runner group will be restricted to running only the workflows specified in the 'selected_workflows' array. Defaults to 'false'.
	RestrictedToWorkflows *bool `json:"restrictedToWorkflows,omitempty" tf:"restricted_to_workflows,omitempty"`

	// The GitHub API URL for the runner group's runners
	// The GitHub API URL for the runner group's runners.
	RunnersURL *string `json:"runnersUrl,omitempty" tf:"runners_url,omitempty"`

	// GitHub API URL for the runner group's repositories
	// GitHub API URL for the runner group's repositories.
	SelectedRepositoriesURL *string `json:"selectedRepositoriesUrl,omitempty" tf:"selected_repositories_url,omitempty"`

	// IDs of the repositories which should be added to the runner group
	// List of repository IDs that can access the runner group.
	// +listType=set
	SelectedRepositoryIds []*int64 `json:"selectedRepositoryIds,omitempty" tf:"selected_repository_ids,omitempty"`

	// List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
	// List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to 'true'.
	SelectedWorkflows []*string `json:"selectedWorkflows,omitempty" tf:"selected_workflows,omitempty"`

	// Visibility of a runner group. Whether the runner group can include all, selected, or private repositories. A value of private is not currently supported due to limitations in the GitHub API.
	// The visibility of the runner group.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type RunnerGroupParameters struct {

	// Whether public repositories can be added to the runner group. Defaults to false.
	// Whether public repositories can be added to the runner group.
	// +kubebuilder:validation:Optional
	AllowsPublicRepositories *bool `json:"allowsPublicRepositories,omitempty" tf:"allows_public_repositories,omitempty"`

	// Name of the runner group
	// Name of the runner group.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/repo/v1alpha1.Repository
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Repository in repo to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
	// If 'true', the runner group will be restricted to running only the workflows specified in the 'selected_workflows' array. Defaults to 'false'.
	// +kubebuilder:validation:Optional
	RestrictedToWorkflows *bool `json:"restrictedToWorkflows,omitempty" tf:"restricted_to_workflows,omitempty"`

	// IDs of the repositories which should be added to the runner group
	// List of repository IDs that can access the runner group.
	// +kubebuilder:validation:Optional
	// +listType=set
	SelectedRepositoryIds []*int64 `json:"selectedRepositoryIds,omitempty" tf:"selected_repository_ids,omitempty"`

	// List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
	// List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to 'true'.
	// +kubebuilder:validation:Optional
	SelectedWorkflows []*string `json:"selectedWorkflows,omitempty" tf:"selected_workflows,omitempty"`

	// Visibility of a runner group. Whether the runner group can include all, selected, or private repositories. A value of private is not currently supported due to limitations in the GitHub API.
	// The visibility of the runner group.
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

// RunnerGroupSpec defines the desired state of RunnerGroup
type RunnerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RunnerGroupParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RunnerGroupInitParameters `json:"initProvider,omitempty"`
}

// RunnerGroupStatus defines the observed state of RunnerGroup.
type RunnerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RunnerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RunnerGroup is the Schema for the RunnerGroups API. Creates and manages an Actions Runner Group within a GitHub organization
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type RunnerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.visibility) || (has(self.initProvider) && has(self.initProvider.visibility))",message="spec.forProvider.visibility is a required parameter"
	Spec   RunnerGroupSpec   `json:"spec"`
	Status RunnerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RunnerGroupList contains a list of RunnerGroups
type RunnerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunnerGroup `json:"items"`
}

// Repository type metadata.
var (
	RunnerGroup_Kind             = "RunnerGroup"
	RunnerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RunnerGroup_Kind}.String()
	RunnerGroup_KindAPIVersion   = RunnerGroup_Kind + "." + CRDGroupVersion.String()
	RunnerGroup_GroupVersionKind = CRDGroupVersion.WithKind(RunnerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&RunnerGroup{}, &RunnerGroupList{})
}