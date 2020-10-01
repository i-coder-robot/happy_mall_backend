package resp

type Category struct {
	CategoryID string                `json:"categoryID"`
	Name       string                `json:"name"`
	Desc       string                `json:"desc"`
	Order      int                   `json:"order"`
	ParentID   string                `json:"parentId"`
	Children   map[string]*Category2 `json:"children"`
}

type Category2 struct {
	CategoryID string                `json:"categoryID"`
	Name       string                `json:"name"`
	Desc       string                `json:"desc"`
	Order      int                   `json:"order"`
	ParentID   string                `json:"parentId"`
	Children   map[string]*Category3 `json:"children"`
}

type Category3 struct {
	Key        string `json:"key"`
	Id         string `json:"id"`
	CategoryID string `json:"categoryID"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Order      int    `json:"order"`
	ParentID   string `json:"parentId"`
	IsDeleted  bool   `json:"isDeleted"`
}
