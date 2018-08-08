package engine
type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}
type ParseResult struct {
	Requests []Request
	Item []interface{}
}
func NilParser([]byte) ParseResult{
	return ParseResult{}
}
