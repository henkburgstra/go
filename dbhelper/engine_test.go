package dbhelper

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func makeEngine() *Engine {
	return &Engine{Driver: "mysql", User: "root", Database: "gestel2"}
}

func TestConnect(t *testing.T) {
	engine := makeEngine()
	db, err := engine.Connect()
	if err != nil {
		t.Fatalf("TestConnect(): engine.Connect(): %s", err.Error())
	}
	db.Close()
}

func TestTableNames(t *testing.T) {
	engine := makeEngine()
	db, err := engine.Connect()
	if err != nil {
		t.Fatalf("TestTableNames(): %s", err.Error())
	}
	defer db.Close()
	names := engine.TableNames()

	if len(names) == 0 {
		t.Errorf("TestTableNames(): length is zero.")
	}
	for _, name := range names {
		fmt.Println(name)
	}
}

func TestTableStructure(t *testing.T) {
	engine := makeEngine()
	db, err := engine.Connect()
	if err != nil {
		t.Fatalf("TestTableStructure(): %s", err.Error())
	}
	defer db.Close()
	names := engine.TableNames()

	if len(names) == 0 {
		t.Errorf("TestTableNames(): length is zero.")
	}
	fmt.Println()
	for _, name := range names {
		ts := engine.TableStructure(name)
		fmt.Println(name)
		fmt.Println("---------------------------------------------")
		for key, value := range ts.Fields() {
			fmt.Println(fmt.Sprintf("%s: %s", key, value.String()))
		}
		fmt.Println()
	}

}
