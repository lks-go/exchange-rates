package currency

var List map[Symbol]struct{}

type Symbol string

const (
	USD Symbol = "USD"
	RUB Symbol = "RUB"
	EUR Symbol = "EUR"
)

func init() {
	List = map[Symbol]struct{}{
		USD: {},
		RUB: {},
		EUR: {},
	}
}

func Available(s Symbol) bool {
	_, ok := List[s]

	return ok
}
