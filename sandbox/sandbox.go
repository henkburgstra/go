package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"unsafe"
)

type T1 struct {
	i    int
	s    string
	Naam string
}

func (t *T1) print() {
	fmt.Println("print")
}

type Value struct {
	Value interface{}
}

func (v *Value) Get() interface{} {
	return v.Value
}

func (v *Value) Addr() *interface{} {
	return &v.Value
}

func (v *Value) Set(value interface{}) {
	v.Value = value
}

type Test struct {
	Naam  string
	Naam2 *string
}

func deelGeheugen() {
	// Deel geheugen tussen variabelen en struct velden
	t := new(Test)
	t.Naam = "test"
	s := "test2"
	t.Naam2 = &s

	a := &t.Naam
	fmt.Println(*a)

	*a = "test2"
	fmt.Println(*a)
	fmt.Println(t.Naam)

	t.Naam = "test3"
	fmt.Println(*a)
	fmt.Println(t.Naam)

	var b interface{}
	b = &t.Naam

	fmt.Println(*(b.(*string)))

	c := reflect.ValueOf(t).Elem().Addr().Interface()
	d := c.(*Test)
	d.Naam = "test4"
	fmt.Println(d.Naam)
	fmt.Println(t.Naam)
	fmt.Println(*a)

	e := reflect.ValueOf(t).Elem()
	field := e.FieldByName("Naam")
	field2 := e.FieldByName("Naam2")

	var i interface{}
	up := unsafe.Pointer(field2.UnsafeAddr())

	fmt.Println("--------------------------")
	switch value := field2.Interface().(type) {
	default:
		fmt.Println("huh??")
	case int, int8, int16, int32, int64:
		fmt.Println("int-achtig %d", value)
	case *int:
		fmt.Println("*int %d", *value)
	case *int8:
		fmt.Println("*int8 %d", *value)
	case *int16:
		fmt.Println("*int16 %d", *value)
	case *int32:
		fmt.Println("*int32 %d", *value)
	case *int64:
		fmt.Println("*int64 %d", *value)
	case string:
		fmt.Println("string %s", value)
		//i = (*string)up
	case *string:
		i = (**string)(up)
		fmt.Println("*string %s", **i.(**string))
	case []byte:
		fmt.Println("[]byte %s", string(value))
	case *[]byte:
		fmt.Println("*[]byte %s", string(*value))
	case nil:
		fmt.Println("NULL")
	}
	fmt.Println("--------------------------")

	i = (**string)(up)
	fmt.Println(**i.(**string))
	k := *i.(**string)
	*k = "bla"
	fmt.Println(*k)
	fmt.Println(**i.(**string))
	f := (*string)(unsafe.Pointer(field.UnsafeAddr()))
	*f = "via FieldByName"
	fmt.Println(field.Interface())
	fmt.Println(*f)
	fmt.Println(*a)
	fmt.Println(t.Naam)

}

func instantieer(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println(t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Name)
	}
	instance := reflect.New(t).Elem().Addr().Interface()

	fmt.Println(instance)
}

func main() {
	deelGeheugen()
	var z T1
	instantieer(z)

	v1 := reflect.TypeOf(T1{}) //(*t1)(nil))
	x := reflect.New(v1).Elem().Addr().Interface()
	y := x.(*T1)
	fmt.Println(x.(*T1).i)
	y.print()

	db, err := sql.Open("mysql", "root:@/gestel2")
	if err != nil {
		fmt.Println("fout bij het maken van een database verbinding.")
		return
	}
	rows, err := db.Query(`SELECT *
		FROM behandeldag_verrichtingen
		WHERE behandeldag_verrichtingen_id <= ?`, 2)
	if err != nil {
		fmt.Println("fout bij het uitvoeren van de query.")
		return
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	values := make([]*Value, len(columns))
	pValues := make([]interface{}, len(columns))
	for i, _ := range columns {
		values[i] = new(Value)
		pValues[i] = values[i].Addr()
	}

	for rows.Next() {
		fmt.Println("Resultaat gevonden.")
		rows.Scan(pValues...)

		for i, _ := range columns {
			switch values[i].Get().(type) {
			case int, int8, int16, int32, int64:
				//				if false {
				//					fmt.Println(s)
				//				}
				fmt.Println("Waarde was een integer:", values[i].Value)
			case *int, *int8, *int16, *int32, *int64:
				fmt.Println("Waarde was een pointer naar een integer:", values[i])
			case []byte:
				fmt.Println("Waarde was een array van bytes:", string(values[i].Value.([]byte)))
			case *[]byte:
				fmt.Println("Waarde was een pointer naar een array van bytes:", values[i].Value)
			case nil:
				fmt.Println("Waarde was nil.")
			}

		}

	}

}
