package dbhelper

import (
	"fmt"
)

type NoKeyError struct {
	TableName string
}

func (e NoKeyError) Error() string {
	return fmt.Sprintf("No key defined for table '%s'", e.TableName)
}

type EntityField struct {
	Name    string
	Type    string
	Length  int
	Key     bool
	Null    bool
	Default string
}

type TableIndex struct {
}

type EntityRelationship struct {
	ForeignKey       string
	ReferencedTable  string
	ReferencedColumn string
}

type Entity struct {
	Name          string
	registry      *Registry
	Fields        map[string]*EntityField
	Indexes       map[string]*TableIndex
	Relationships map[string]EntityRelationship
}

func NewEntity(name string) *Entity {
	e := new(Entity)
	e.Name = name
	e.Fields = make(map[string]*EntityField)
	e.Indexes = make(map[string]*TableIndex)
	e.Relationships = make(map[string]EntityRelationship)
	return e
}

func (e *Entity) KeyCount() int {
	count := 0
	for _, field := range e.Fields {
		if field.Key {
			count++
		}
	}
	return count
}

func (e *Entity) Key() *EntityField {
	for _, field := range e.Fields {
		if field.Key {
			return field
		}
	}
	return nil
}

func (e *Entity) AddRelationship(relationship EntityRelationship) {
	e.Relationships[relationship.ForeignKey] = relationship
}

func (e *Entity) Relationship(fk string) (relationship EntityRelationship, ok bool) {
	relationship, ok = e.Relationships[fk]
	return
}

func (e *Entity) Keys() []*EntityField {
	keyFields := make([]*EntityField, 0)
	for _, field := range e.Fields {
		if field.Key {
			keyFields = append(keyFields, field)
		}
	}
	return keyFields
}
