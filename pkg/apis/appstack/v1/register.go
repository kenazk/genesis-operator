package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/kenazk/genesis-operator/pkg/apis/appstack"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group: appstack.GroupName,
	Version: "v1",
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme = SchemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&AppStack{},
		&AppStackList{},
	)

	meta_v1.AddToGroupVersion(scheme,SchemeGroupVersion)
	return nil
}
