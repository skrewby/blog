package article

type Article struct {
	Folder string
	Pages  []ArticlePage
}

type ArticlePage struct {
	FileName string
	FullPath string
	Title    string
	Tags     []string
}
