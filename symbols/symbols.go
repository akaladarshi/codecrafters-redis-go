package symbols

type Symbol string

const (
	Plus     = Symbol('+')
	Asterisk = Symbol('*')
	Dollar   = Symbol('$')
	Colon    = Symbol(':')
)

func (s Symbol) String() string {
	return string(s)
}
