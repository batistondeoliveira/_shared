package domains

type Filter map[string]FilterProperties

type FilterProperties struct {
	Value      string
	Type       Type
	Expression FilterExpression
}

type Type string

const (
	INTEGER Type = "int"
	BOOLEAN      = "bool"
	STRING       = "string"
)

type FilterExpression string

const (
	ANYWHERE      FilterExpression = "ANYWHERE"
	EQUAL_TO                       = "EQUAL_TO"
	START_WITH                     = "START_WITH"
	ENDS_WITH                      = "ENDS_WITH"
	IN_THE_MIDDLE                  = "IN_THE_MIDDLE"
)
