package dbhelper

type Query struct {
	model IModel
}

func NewQuery(model IModel) *Query {
	q := new(Query)
	q.model = model
	return q
}

func (q *Query) Get(key interface{}) IModel {
	return nil
}
