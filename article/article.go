package article

type Article struct {
	Folder string
	Title  string
	Pages  []ArticlePage
}

type ArticlePage struct {
	FileName string
	FullPath string
	Title    string
	Tags     []string
	Page     int
}
