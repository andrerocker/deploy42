package config

type Namespace struct {
	Endpoint string
	Chaining []string
	Commands CommandList
}

type NamespaceList []Namespace
