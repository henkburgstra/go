package dbhelper

import (
	"database/sql"
	//	"fmt"
	//"reflect"
	"strings"
	"unicode"
)

func Underscore2Camel(underscores string) string {
	els := strings.Split(underscores, "_")
	for i, el := range els {
		r := []rune(el)
		r[0] = unicode.ToUpper(r[0])
		els[i] = string(r)
	}
	return strings.Join(els, "")
}

type IModel interface {
	Name() string
	Fields() FieldData
	Registry() *Registry
	SetRegistry(*Registry)
	Scan(*sql.Rows)
}

type Model struct {
	name     string
	registry *Registry
	fields   FieldData
}

func NewModel(name string) *Model {
	m := new(Model)
	m.name = name
	m.fields = make(FieldData)
	return m
}

func (m *Model) Name() string {
	return m.name
}

func (m *Model) Fields() FieldData {
	return m.fields
}

func (m *Model) Registry() *Registry {
	return m.registry
}

func (m *Model) SetRegistry(r *Registry) {
	m.registry = r
}

func (m *Model) Scan(rows *sql.Rows) {
	columns, _ := rows.Columns()
	values := make([]*Value, len(columns))
	pValues := make([]interface{}, len(columns))
	for i, _ := range columns {
		values[i] = new(Value)

		//		s := reflect.ValueOf(m).Elem()
		//		// Kijk of het model een veld heeft dat overeenkomt met de veldnaam uit de database
		//		// waarbij veldnamen met undersores worden vertaald naar CamelCase:
		//		// de_naam -> DeNaam
		//		structField := s.FieldByName(Underscore2Camel(name))
		//		if structField.IsValid() {
		//			pValues[i] =
		//		}

		pValues[i] = values[i].Addr()
	}

	rows.Scan(pValues...)

	for i, _ := range columns {
		m.Fields()[columns[i]] = values[i]
	}

}
