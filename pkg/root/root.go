package root

type BaseRequest struct {
	Url     string
	Headers []Header
}

type Header struct {
	Name  string
	Value string
}
