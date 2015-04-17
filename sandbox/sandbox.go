package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
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

func (v *Value) SetAddr(value interface{}) {
	switch value.(type) {
	default:
		fmt.Println("huh??")
	case *string, *int, *int8, *int16, *int32, *int64:
		v.Value = value
	}
}

func (v *Value) Set(value interface{}) {
	switch value.(type) {
	default:
		fmt.Println("huh??")
	case int, int8, int16, int32, int64:
		fmt.Println("int-achtig %d", value)
	case *int:
		fmt.Println("*int %d", *value.(*int))
	case *int8:
		fmt.Println("*int8 %d", *value.(*int8))
	case *int16:
		fmt.Println("*int16 %d", *value.(*int16))
	case *int32:
		fmt.Println("*int32 %d", *value.(*int32))
	case *int64:
		fmt.Println("*int64 %d", *value.(*int64))
	case string:
		*v.Value.(*string) = value.(string)
		fmt.Println("--- string:", value)
	case *string:
		*v.Value.(*string) = *value.(*string)
		fmt.Println("--- *string:", *value.(*string))
	case []byte:
		fmt.Println("[]byte %s", string(value.([]byte)))
	case *[]byte:
		fmt.Println("*[]byte %s", string(*value.(*[]byte)))
	case *interface{}:
		*v.Value.(*interface{}) = *value.(*interface{})
	case nil:
		fmt.Println("NULL")
	}
}

func (v *Value) String() string {
	switch value := v.Value.(type) {
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", value)
	case *int:
		return fmt.Sprintf("%d", *value)
	case *int8:
		return fmt.Sprintf("%d", *value)
	case *int16:
		return fmt.Sprintf("%d", *value)
	case *int32:
		return fmt.Sprintf("%d", *value)
	case *int64:
		return fmt.Sprintf("%d", *value)
	case string:
		return value
	case *string:
		return *value
	case []byte:
		return string(value)
	case *[]byte:
		return string(*value)
	case nil:
		return "NULL"
	}
	return ""
}

type Test struct {
	Naam  string
	Naam2 *string
}

func deelGeheugen1() {
	fmt.Println("t := new(Test)")
	t := new(Test)
	fmt.Println(`t.Naam = "Naam"`)
	t.Naam = "Naam"
	fmt.Println("a := &t.Naam")
	a := &t.Naam
	fmt.Println("*a:", *a)
	fmt.Println("t.Naam:", t.Naam)
	fmt.Println(`*a = "Andere naam"`)
	*a = "Andere naam"
	fmt.Println("t.Naam:", t.Naam)
	fmt.Println(`t.Naam = "Nog een andere naam"`)
	t.Naam = "Nog een andere naam"
	fmt.Println("*a:", *a)
	fmt.Println("--------------------------")
}

func deelGeheugen2() {
	fmt.Println("t := new(Test)")
	t := new(Test)
	fmt.Println(`t.Naam = "Naam"`)
	t.Naam = "Naam"
	fmt.Println(`v = new(Value)`)
	v := new(Value)
	fmt.Println(`v.Value := &t.Naam`)
	v.Value = &t.Naam
	fmt.Println("*v.Value.(*string):", *v.Value.(*string))
	fmt.Println("t.Naam:", t.Naam)
	fmt.Println(`*v.Value.(*string) = "Andere naam"`)
	*v.Value.(*string) = "Andere naam"
	fmt.Println("t.Naam:", t.Naam)
	fmt.Println(`t.Naam = "Nog een andere naam"`)
	t.Naam = "Nog een andere naam"
	fmt.Println("*v.Value.(*string):", *v.Value.(*string))
	fmt.Println("--------------------------")
}

func deelGeheugen3() {
	fmt.Println("t := new(Test)")
	t := new(Test)
	fmt.Println(`t.Naam = "Naam"`)
	t.Naam = "Naam"
	fmt.Println(`v = new(Value)`)
	v := new(Value)
	fmt.Println(`v.Set(&t.Naam)`)
	v.SetAddr(&t.Naam)
	fmt.Println("*v.Value.(*string):", *v.Value.(*string))
	fmt.Println("t.Naam:", t.Naam)
	fmt.Println(`v.Set("Andere naam")`)
	v.Set("Andere naam")
	fmt.Println("t.Naam:", t.Naam)
	fmt.Println(`t.Naam = "Nog een andere naam"`)
	t.Naam = "Nog een andere naam"
	fmt.Println("*v.Value.(*string):", *v.Value.(*string))
	fmt.Println("--------------------------")
}

func deelGeheugen4() {
	fmt.Println("t := new(Test)")
	t := new(Test)

	e := reflect.ValueOf(t).Elem()
	field := e.FieldByName("Naam")
	//	up := (*string)(unsafe.Pointer(field.UnsafeAddr()))

	v := new(Value)
	//	v.SetAddr(up)
	v.SetAddr((*string)(unsafe.Pointer(field.UnsafeAddr())))

	t.Naam = "xxx"
	fmt.Println(v.String())
	//fmt.Println(*(*string)(v.Get().(*string)))

	v.Set("yyy")
	fmt.Println(t.Naam)

	x := "blub"
	v.Set(&x)

	//var p int64 = 256
	//	var q *int32
	//	r := (int32)(p)
	//q = &r

	fmt.Println(t.Naam)
}

func deelGeheugen() {
	deelGeheugen1()
	deelGeheugen2()
	deelGeheugen3()
	deelGeheugen4()
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

type Cond map[string]map[string]interface{}

type Connective struct {
	Operator string
	Operands []interface{}
}

func And(ops ...interface{}) Connective {
	return Connective{"AND", ops}
}

func Or(ops ...interface{}) Connective {
	return Connective{"OR", ops}
}

func processConnective(con Connective) string {
	args := make([]string, 0)
	//	values := make([]interface{}, 0)

	for _, op := range con.Operands {
		switch c := op.(type) {
		case Connective:
			r := processConnective(c)
			if r != "" {
				args = append(args, r)
			}
		case Cond:
			r := processCondition(c, con.Operator)
			if r != "" {
				args = append(args, r)
			}
		}
	}
	return fmt.Sprintf("(%s)", strings.Join(args, fmt.Sprintf(" %s ", con.Operator)))
}

func processCondition(c Cond, op string) string {
	args := make([]string, 0)
	values := make([]interface{}, 0)

	for model, conds := range c {
		for field, value := range conds {
			args = append(args, fmt.Sprintf("%s.%s = ?", model, field))
			values = append(values, value)
		}
	}

	return strings.Join(args, fmt.Sprintf(" %s ", op))
}

func Filter(f interface{}) string {
	// Filter verwacht een Cond of een Connective
	switch c := f.(type) {
	default:
		return ""
	case Cond:
		return processCondition(c, "AND")
	case Connective:
		return processConnective(c)
	}
}

func testSlices(s *[]string, c int) {
	for i := 1; i < 5; i++ {
		*s = append(*s, fmt.Sprintf("mySlice %d", c*i))
	}
}

func main() {
	mySlice := make([]string, 0)
	for i := 1; i < 5; i++ {
		testSlices(&mySlice, i)
	}
	for _, sl := range mySlice {
		fmt.Println(sl)
	}
	cond1 := Cond{
		"patient": {"key": "ACTB-09D0034"},
		"relatie": {"verwijsfunctie": "huisarts"}}
	cond2 := Cond{
		"patient": {"key": "ACTB-09D0034"},
		"relatie": {"verwijsfunctie": "huisarts"}}
	fmt.Println(Filter(Or(And(cond1), And(cond2))))
	deelGeheugen()
	var z T1
	instantieer(z)

	v1 := reflect.TypeOf(T1{}) //(*t1)(nil))
	x := reflect.New(v1).Elem().Addr().Interface()
	y := x.(*T1)
	fmt.Println(x.(*T1).i)
	y.print()

	db, err := sql.Open("mysql", "root:borland@/amersfoort2")
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
