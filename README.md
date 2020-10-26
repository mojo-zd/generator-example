use https://github.com/kubernetes/code-generator to generator kubernetes types of crd client、informer、listers

#### prepare
1. ``` mkdir -p generator-example/apis/samplecontroller/v1alpha1 && generator-example/apis/samplecontroller/v1alpha1``` 
2. touch doc.go

```
// +k8s:deepcopy-gen=package
// +groupName=samplecontroller.kimo.io

// v1alpha1版本的api包
package v1alpha1
```

3. touch types.go

```
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FooSpec   `json:"spec"`
	Status FooStatus `json:"status"`
}

// FooSpec is the spec for a Foo resource

type FooSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

// FooStatus is the status for a Foo resource
type FooStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type FooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Foo `json:"items"`
}
```

#### how to generate crd client
```
cd generator-example && sh ./hack/update-codegen.sh
```

#### how to use client

```
generator-example/pkg/apis/samplecontroller/v1alpha1/client_test.go
```