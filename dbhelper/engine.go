package dbhelper

import (
	"database/sql"
	"fmt"
)

type Engine struct {
	db        *sql.DB
	Driver    string
	Host      string
	Port      string
	Database  string
	User      string
	Password  string
	Connected bool
}

func (engine *Engine) ConnectionString() string {
	switch engine.Driver {
	case "sqlite3":
		return engine.Database
	case "mysql":
		return fmt.Sprintf("%s:%s@/%s", engine.User, engine.Password, engine.Database)
		//return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		//	engine.User, engine.Password, engine.Host, engine.Port, engine.Database)
	case "mssql":
		return fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
			engine.Host, engine.User, engine.Password, engine.Database)
	default:
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			engine.User, engine.Password, engine.Host, engine.Port, engine.Database)
	}
}

func (engine *Engine) Connect() (*sql.DB, error) {
	db, err := sql.Open(engine.Driver, engine.ConnectionString())
	if err == nil {
		engine.db = db
		engine.Connected = true
	}
	return db, err
}

func (engine *Engine) Db() *sql.DB {
	return engine.db
}

func (engine *Engine) mysqlTableNames() []string {
	names := make([]string, 0)
	rows, err := engine.db.Query(`SHOW TABLES`)

	if err != nil {
		return names
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return names
		}
		names = append(names, name)
	}
	return names
}

func (engine *Engine) mysqlTableStructure(name string, entity *Entity) {
	rows, err := engine.db.Query(fmt.Sprintf("DESCRIBE %s", name))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		ts := NewModel("fieldstructure")
		ts.Scan(rows)
		f := &EntityField{
			Name:    ts.Field("Field").String(),
			Type:    ts.Field("Type").String(),
			Length:  0, // TODO: lengte distilleren uit het type (varchar(8))
			Key:     ts.Field("Key").String() == "PRI",
			Null:    ts.Field("Null").String() == "YES",
			Default: ts.Field("Default").String(),
		}
		entity.Fields[f.Name] = f
	}
}

func (engine *Engine) mysqlRelationships(registry *Registry) {
	rows, err := engine.db.Query(`
		SELECT rc.CONSTRAINT_NAME AS ConstraintName, 
		rc.TABLE_NAME AS TableName, kc.COLUMN_NAME AS ColumnName, 
		rc.REFERENCED_TABLE_NAME AS ReferencedTableName, 
		kc.REFERENCED_COLUMN_NAME AS ReferencedColumnName, 
		rc.UPDATE_RULE AS UpdateRule, rc.DELETE_RULE AS DeleteRule 
		FROM INFORMATION_SCHEMA.REFERENTIAL_CONSTRAINTS AS rc
		JOIN INFORMATION_SCHEMA.KEY_COLUMN_USAGE AS kc
		ON rc.CONSTRAINT_NAME = kc.CONSTRAINT_NAME
		WHERE rc.CONSTRAINT_SCHEMA = ?
		ORDER BY rc.TABLE_NAME, kc.COLUMN_NAME`, engine.Database)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		m := NewModel("relationship")
		m.Scan(rows)
		r := EntityRelationship{
			ForeignKey:       m.Field("ColumnName").String(),
			ReferencedTable:  m.Field("ReferencedTableName").String(),
			ReferencedColumn: m.Field("ReferencedColumnName").String(),
		}
		entity := registry.Entity(registry.TrimTableAffixes(m.Field("TableName").String()))
		if entity != nil {
			entity.AddRelationship(r)
		}
	}
}

func (engine *Engine) LoadRelationships(registry *Registry) {
	switch engine.Driver {
	case "mysql":
		engine.mysqlRelationships(registry)
	}
}

func (engine *Engine) TableNames() []string {
	names := make([]string, 0)
	if engine.db == nil {
		return names
	}

	switch engine.Driver {
	case "mysql":
		names = engine.mysqlTableNames()
	}

	return names
}

func (engine *Engine) TableStructure(name string) *Entity {
	entity := NewEntity(name)

	if engine.db == nil {
		return entity
	}

	switch engine.Driver {
	case "mysql":
		engine.mysqlTableStructure(name, entity)
	}
	return entity
}
