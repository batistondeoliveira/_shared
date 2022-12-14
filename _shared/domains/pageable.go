package domains

type Pageable struct {
	Page         int
	LinesPerPage int
	Orders       []Order
}

type Order struct {
	OrderBy   string
	Direction Direction
}

type Direction string

const (
	ASC  Direction = "ASC"
	DESC           = "DESC"
)
