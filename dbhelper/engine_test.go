package dbhelper

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func makeEngine() *Engine {
	return &Engine{Driver: "mysql", User: "root", Password: "borland", Database: "amersfoort2"}
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
	//	for _, name := range names {
	//		fmt.Println(name)
	//	}
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
}

func TestRef(t *testing.T) {
	engine := makeEngine()
	db, err := engine.Connect()
	if err != nil {
		t.Fatalf("TestRef(): %s", err.Error())
	}
	defer db.Close()
	registry := makeRegistry(engine)
	patient := registry.Query("patient").Get("PJJG-AA1001")
	huisarts, ok := patient.Ref("huisarts")

	if !ok {
		t.Errorf("TestRef(): huisarts niet gevonden")
	}
	for _, name := range huisarts.FieldNames() {
		fmt.Printf("huisarts veldnaam: %s\n", name)
	}

}
