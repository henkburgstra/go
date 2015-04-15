package dbhelper

import (
	"fmt"
	"strings"
)

type Query struct {
	modelName string
	registry  *Registry
	sql       string
	sqlParams []interface{}
}

func NewQuery(modelName string, registry *Registry) *Query {
	q := new(Query)
	q.modelName = modelName
	q.registry = registry
	q.sqlParams = make([]interface{}, 0)
	return q
}

func (q *Query) Join() *Query {
	return q
}

func (q *Query) Get(keyValue interface{}) IModel {
	entity := q.registry.Entity(q.modelName)
	if entity == nil {
		fmt.Println("entity == nil")
		return nil
	}
	key := entity.Key()
	if key == nil {
		fmt.Println("key == nil")
		return nil
	}
	model := q.registry.Model(q.modelName)(q.modelName)
	model.SetOwner(model)
	model.SetRegistry(q.registry)

	sql := fmt.Sprintf(`SELECT * 
		FROM %s
		WHERE %s = ?`, entity.Name, key.Name)

	db, err := q.registry.Db()
	if err != nil {
		// TODO: log err
		return nil
	}
	rows, err := db.Query(sql, keyValue)
	if err != nil {
		// TODO: log err
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		model.Scan(rows)
		break
	}

	return model
}

func (q *Query) Columns(cols ...string) *Query {
	return q
}

func (q *Query) FromSql(sql string, params ...interface{}) *Query {
	q.sql = sql
	q.sqlParams = append(q.sqlParams, params...)
	return q
}

func (q *Query) Sql() string {
	if q.sql != "" {
		return q.sql
	}
	e := q.registry.Entity(q.modelName)
	if e == nil {
		return ""
	}
	sql := fmt.Sprintf(`SELECT * 
		FROM %s`, e.Name)

	return sql
}

func (q *Query) All() []IModel {
	models := make([]IModel, 0)
	entity := q.registry.Entity(q.modelName)
	if entity == nil {
		fmt.Println("entity == nil")
		return models
	}
	key := entity.Key()
	if key == nil {
		fmt.Println("key == nil")
		return models
	}

	fieldPrefix := strings.Replace(q.registry.FieldPrefix(), "{model}", q.modelName, 1)
	keyName := strings.TrimPrefix(key.Name, fieldPrefix)

	db, err := q.registry.Db()
	if err != nil {
		// TODO: log err
		return models
	}
	rows, err := db.Query(q.Sql(), q.sqlParams...)
	if err != nil {
		// TODO: log err
		return models
	}
	defer rows.Close()

	var lastKey string
	modelConstructor := q.registry.Model(q.modelName)

	for rows.Next() {
		model := modelConstructor(q.modelName)
		model.SetOwner(model)
		model.SetRegistry(q.registry)
		model.Scan(rows)
		keyValue := model.Field(keyName).String()
		if lastKey != keyValue {
			models = append(models, model)
			lastKey = keyValue
		}
	}

	return models
}
