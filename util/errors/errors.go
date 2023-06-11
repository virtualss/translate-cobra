package errors

import "fmt"

type TranslateResp struct {
	Code   string
	Reason string
}

func (t TranslateResp) Error() string {
	return fmt.Sprintf("Translate Fail With Code: [%s] , Reason: [%s]\n", t.Code, t.Reason)
}

type Domain struct {
	//error
	Field string
}

func (d Domain) Error() string {
	return fmt.Sprintf("Invalid Domain: [%s] \n", d.Field)
}

type FromTo struct {
	Field string
	//error
	From string
	To   string
}

func (f FromTo) Error() string {
	return fmt.Sprintf("Filed [%s] unsupport From Language: [%s] to Target Language: [%s] \n", f.Field, f.From, f.To)
}
