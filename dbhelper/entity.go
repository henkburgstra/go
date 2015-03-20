package dbhelper

type Field struct {
	Name string
}

type Entity struct {
	Name     string
	registry *Registry
}
