package errors

import "fmt"

type Domain struct {
	//error
	Field string
}

func (d Domain) Error() string {
	return fmt.Sprintf("Invalid Domain: [%s] \n", d.Field)
}

type FromTo struct {
	//error
	From string
	To   string
}

func (f FromTo) Error() string {
	return fmt.Sprintf("Unupported Origin Language: [%s] or Target Language: [%s],see 'tl field' for help \n", f.From, f.To)
}
