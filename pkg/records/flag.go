package records

import "io"

type Flags struct {
	Action   string
	Category string
	Domain   string

	Args []string

	Rdr io.Reader `json:"-"`
}

func (f Flags) Arg(i int) string {
	if len(f.Args) > i {
		return f.Args[i]
	}
	return ""
}
