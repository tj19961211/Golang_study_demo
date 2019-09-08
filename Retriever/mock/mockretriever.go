package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) srting {
	return r.Contents
}
