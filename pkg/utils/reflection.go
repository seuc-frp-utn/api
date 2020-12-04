package utils

import "reflect"

// FillStruct receives a struct represented by a reflect.Value and a JSON object called body.
// It returns the populated struct with the values from the JSON object.
// FillStruct needs that the reflected Value has `json:"name"` labels for each field of the struct.
// The body needs to have properties called by the same name as the JSON labels.
//
// USE CASE:
//
// 		t := reflect.TypeOf(typeOf)
//
//		if t.Kind() != reflect.Struct {
//			ctx.AbortWithError(http.StatusInternalServerError, errors.New("input body is not a struct"))
//			return
//		}
//
//		var body map[string]interface{}
//		if err := ctx.BindJSON(&body); err != nil {
//			ctx.AbortWithError(http.StatusBadRequest, err)
//			return
//		}
//
//		entity := reflect.Indirect(reflect.New(t))
//		if !entity.CanAddr() {
//			ctx.AbortWithError(http.StatusInternalServerError, errors.New("not addressable"))
//		}
//
//		entity = utils.FillStruct(entity, body)
//
// NOTE:
//		entity needs to be addressable (that's the main reason of using Indirect) in order to be populated.
// 		entity will be of type typeOf, and it will be populated with the values from body.
//
// DISCLAIMER:
// 		This function doesn't support structs nor slices as JSON values.
//
func FillStruct(v reflect.Value, payload map[string]interface{}) reflect.Value {
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag

		key, ok := tag.Lookup("json")
		if !ok {
			continue
		}
		value := payload[key]
		if value == nil {
			continue
		}

		field := v.Field(i)
		if !field.IsValid() {
			continue
		}

		if !field.CanSet() {
			continue
		}

		if !field.CanAddr() {
			continue
		}

		switch field.Kind() {
		case reflect.Slice, reflect.Array:
			converted, ok := value.([]interface{})
			if ! ok {
				break
			}
			v := reflect.ValueOf(converted)
			field.Set(v)
		case reflect.String:
			field.SetString(value.(string))
		case reflect.Bool:
			field.SetBool(value.(bool))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			converted, ok := value.(float64)
			parsed := int64(converted)
			if !ok {
				break
			}
			field.SetInt(parsed)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			converted, ok := value.(float64)
			parsed := uint64(converted)
			if !ok {
				break
			}
			field.SetUint(parsed)
		case reflect.Float32, reflect.Float64:
			converted, ok := value.(float64)
			if !ok {
				break
			}
			field.SetFloat(converted)
		default:
			converted, ok := value.(map[string]interface{})
			if !ok {
				break
			}
			v := reflect.ValueOf(converted)
			field.Set(v)
		}
	}
	return v
}
