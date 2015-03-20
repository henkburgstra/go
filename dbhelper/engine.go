package dbhelper

import (
	"database/sql"
	"fmt"
)

type Engine struct {
	db       *sql.DB
	Driver   string
	Host     string
	Port     string
	Database string
	User     string
	Password string
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

func (engine *Engine) mysqlTableStructure(name string, ts *Model) {
	rows, err := engine.db.Query(fmt.Sprintf("DESCRIBE %s", name))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		ts.Scan(rows)
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

func (engine *Engine) TableStructure(name string) *Model {
	ts := NewModel(name)

	if engine.db == nil {
		return ts
	}

	switch engine.Driver {
	case "mysql":
		engine.mysqlTableStructure(name, ts)
	}
	return ts
}
