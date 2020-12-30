package root

type Header struct {
	Header  string
	Content string
}

type Request struct {
	Name    string
	Method  string
	Headers []Header
	Body    string
}
