package v1alpha1

import (
	"context"
	"github.com/mojo-zd/generator-example/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"testing"
)

func TestSampleClient(t *testing.T) {
	conf := rest.Config{} // you should provide the correct config
	cli, err := versioned.NewForConfig(&conf)
	if err != nil {
		t.Fatal("new client failed", err)
		return
	}
	cli.SamplecontrollerV1alpha1().Foos("").List(context.Background(), v1.ListOptions{}) // list crd resources
}
