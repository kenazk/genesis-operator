package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// appstack describes a AppStack resource
type AppStack struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	Spec AppStackSpec `json:"spec"`
}

type AppStackSpec struct {
	Region string `json:"message"`
	Environment string `json:"message"`
	SomeValue *int32 `json:"someValue"`

}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppStackList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []AppStack `json:"items"`
}
