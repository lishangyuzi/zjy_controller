package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ZjyController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ZjyControllerSpec `json:"spec,omitempty"`
}

type ZjyControllerSpec struct {
	Replicas *int32                `json:"replicas,omitempty"`
	Template ZjyControllerTemplate `json:"template,omitempty"`
}

type ZjyControllerTemplate struct {
	Name   *string `json:"name,omitempty"`
	Gender *int32  `json:"gender,omitempty"`
}
