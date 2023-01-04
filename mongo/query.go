package mongo

type Query struct {
}

func (q *Query) EqualTo(fieldName string, value interface{}) {}

func (q *Query) NotEqualTo(fieldName string, value interface{}) {}

func (q *Query) Exists(fieldName string) {}

func (q *Query) NotExists(fieldName string) {}

func (q *Query) LessThan(fieldName string, value interface{})        {}
func (q *Query) LessThanOrEqual(fieldName string, value interface{}) {}

func (q *Query) GreaterThanOrEqual(fieldName string, value interface{}) {}
func (q *Query) GreaterThan(fieldName string, value interface{})        {}

func (q *Query) In(fieldName string, value interface{})    {}
func (q *Query) NotIn(fieldName string, value interface{}) {}

func (q *Query) All(fieldName string, value interface{}) {}

func (q *Query) Text(fieldName string, value interface{}) {}

func (q *Query) Regex(fieldName string, value interface{}) {}
