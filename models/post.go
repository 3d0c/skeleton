package models

type PostScheme struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Post struct {
}

func (*Post) Construct(args ...interface{}) interface{} {
	this := &Post{}
	return this
}

func (this *Post) Find(id string) *PostScheme {
	return &PostScheme{id, "yyy"}
}

func (this *Post) FindAll() []*PostScheme {
	return []*PostScheme{
		&PostScheme{"1", "yyy"},
		&PostScheme{"2", "ccc"},
	}
}
