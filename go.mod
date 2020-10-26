module github.com/mojo-zd/generator-example

go 1.15

require (
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.19.2
	k8s.io/utils v0.0.0-20201015054608-420da100c033 // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.19.3
