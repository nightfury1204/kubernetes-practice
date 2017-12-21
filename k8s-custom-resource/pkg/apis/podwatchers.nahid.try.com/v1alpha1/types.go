
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//PodWatcher is specification for PodWatcher resource
type PodWatcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodWatcherSpec   `json:"spec"`
	Status PodWatcherStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//PodwatcherSpec is spec for Podwatcher resource
type PodWatcherSpec struct {
	metav1.LabelSelector `json:"selector"`
	Replicas *int32 `json:"replicas"`
	Template PodTemplate `json:"template"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//PodTemplate is template of pod for PodWatcher resource
type PodTemplate struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec apiv1.PodSpec `json:"spec"`
}

//PodWatcherStatus is the status for PodWatcher resource
type PodWatcherStatus struct {
	AvailabelReplicas int32 `json:"availabelReplicas"`
}