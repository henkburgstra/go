package dbhelper

import (
	"fmt"
	"strconv"
)

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

func StrToInt(s string) int64 {
	if v, err := strconv.ParseInt(s, 0, 64); err == nil {
		return v
	}
	return 0
}

func (v *Value) IsNil() bool {
	switch v.Value.(type) {
	case nil:
		return true
	}
	return false
}

func (v *Value) Int() int64 {
	switch value := v.Value.(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case *int:
		return int64(*value)
	case *int8:
		return int64(*value)
	case *int16:
		return int64(*value)
	case *int32:
		return int64(*value)
	case *int64:
		return *value
	case string:
		return StrToInt(value)
	case *string:
		return StrToInt(*value)
	case []byte:
		return StrToInt(string(value))
	case *[]byte:
		return StrToInt(string(*value))
	}
	return 0
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

type FieldData map[string]*Value
