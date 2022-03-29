package fetch

import (
	"fmt"
	"reflect"

	"github.com/tidwall/gjson"
)

// DataSource defines the interface for loading data from a data source.
type DataSource interface {
	Get(key string) string
}

func decodeWithTagFromDataSource(ptr interface{}, tagName string, dataSource gjson.Result) error {
	t := reflect.TypeOf(ptr).Elem()
	v := reflect.ValueOf(ptr).Elem()

	for i := 0; i < t.NumField(); i++ {
		typ := t.Field(i)
		val := v.Field(i)

		kind := val.Kind()

		tagValueName := typ.Tag.Get(tagName)
		tabValueDfeault := typ.Tag.Get("default")
		tagValueRequired := typ.Tag.Get("required")

		switch kind {
		case reflect.String:
			tagValue := dataSource.Get(tagValueName).String()
			if tagValue == "" && tabValueDfeault != "" {
				tagValue = tabValueDfeault
			}

			if tagValueRequired == "true" && tagValue == "" {
				return fmt.Errorf("%s is required", tagValueName)
			}

			val.SetString(tagValue)
		case reflect.Bool:
			tagValue := dataSource.Get(tagValueName).Bool()
			val.SetBool(tagValue)
		case reflect.Int, reflect.Int64:
			tagValue := dataSource.Get(tagValueName).Int()
			if tagValueRequired == "true" && tagValue == 0 {
				return fmt.Errorf("%s is required", tagValueName)
			}

			val.SetInt(tagValue)
		case reflect.Float64:
			tagValue := dataSource.Get(tagValueName).Float()
			if tagValueRequired == "true" && tagValue == 0.0 {
				return fmt.Errorf("%s is required", tagValueName)
			}

			val.SetFloat(tagValue)
		}
	}

	return nil
}

// Decode decodes the given struct pointer from the environment.
func decode(ptr interface{}, response *Response) error {
	return decodeWithTagFromDataSource(ptr, "alias", response.Value())
}
