package main

type base1C struct {
	Name string
	Path string
	ID   string
}

func (b base1C) String() string {
	return b.Name + "|" + b.Path + "|" + b.ID
}

func (b base1C) Complete() bool {
	return b.Name != "" && b.ID != "" && b.Path != ""
}
