package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// 本物のparser完成する前に使用する
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
