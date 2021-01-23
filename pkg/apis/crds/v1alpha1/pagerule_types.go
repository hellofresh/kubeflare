/*
Copyright 2019 Replicated, Inc.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ForwardingURLPageRule struct {
	RequestURL  string `json:"requestUrl"`
	StatusCode  int    `json:"statusCode"`
	RedirectURL string `json:"redirectUrl"`
}

type Rule struct {
	ForwardingURL *ForwardingURLPageRule `json:"forwardingUrl,omitempty"`
	Priority      *int                   `json:"priority,omitempty"`
	Status        *string                `json:"status,omitempty"`
}

// PageRuleSpec defines the desired state of PageRule
type PageRuleSpec struct {
	Zone string `json:"zone"`
	Rule *Rule  `json:"pageRule,omitempty"`
}

// PageRuleStatus defines the observed state of PageRule
type PageRuleStatus struct {
	ID string `json:"id,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PageRule is the Schema for the pagerules API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type PageRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PageRuleSpec   `json:"spec,omitempty"`
	Status PageRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PageRuleList contains a list of PageRule
type PageRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PageRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PageRule{}, &PageRuleList{})
}
