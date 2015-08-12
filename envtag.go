package envtag

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"time"
)

var (
	handlers   map[reflect.Type]func(reflect.Value, string) error
	timeFormat string = ""
)

func init() {
	handlers = map[reflect.Type]func(reflect.Value, string) error{
		reflect.TypeOf((*time.Duration)(nil)).Elem(): handleDuration,
		reflect.TypeOf((*time.Time)(nil)).Elem():     handleTime,
	}
}

func Unmarshal(out interface{}) error {
	val := reflect.ValueOf(out)

	if !val.IsValid() || val.Type().Kind() != reflect.Ptr {
		return errors.New("Invalid type (non-pointer)")
	}

	val = val.Elem()

	if val.Type().Kind() != reflect.Struct {
		return errors.New("Invalid type: " + val.Type().Kind().String())
	}

	return unmarshal(val)
}

func unmarshal(val reflect.Value) error {
	n := val.NumField()

	for i := 0; i < n; i++ {
		fieldVal := val.Field(i)

		if fieldVal.Kind() == reflect.Struct {
			unmarshal(fieldVal)
			continue
		}

		if !fieldVal.CanSet() {
			continue
		}

		field := val.Type().Field(i)

		tag := field.Tag.Get("env")
		if tag == "" {
			continue
		}

		envVal := os.Getenv(tag)
		if envVal == "" {
			continue
		}

		if handler, ok := handlers[fieldVal.Type()]; ok {
			err := handler(fieldVal, envVal)
			if err != nil {
				return errors.New(err.Error() + " for env '" + tag + "': " + envVal)
			}
		} else {
			switch fieldVal.Kind() {

			case reflect.String:
				fieldVal.SetString(envVal)

			case reflect.Bool:
				boolValue, err := strconv.ParseBool(envVal)
				if err != nil {
					return errors.New("Invalid boolean for env '" + tag + "': " + envVal)
				}
				fieldVal.SetBool(boolValue)

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intValue, err := strconv.ParseInt(envVal, 0, fieldVal.Type().Bits())
				if err != nil {
					return errors.New("Invalid integer for env '" + tag + "': " + envVal)
				}
				fieldVal.SetInt(intValue)

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				uintValue, err := strconv.ParseUint(envVal, 0, fieldVal.Type().Bits())
				if err != nil {
					return errors.New("Invalid unsigned integer for env '" + tag + "': " + envVal)
				}
				fieldVal.SetUint(uintValue)

			case reflect.Float32, reflect.Float64:
				floatValue, err := strconv.ParseFloat(envVal, fieldVal.Type().Bits())
				if err != nil {
					return errors.New("Invalid float for env '" + tag + "': " + envVal)
				}
				fieldVal.SetFloat(floatValue)

			default:
				return errors.New("Unsupported type '" + fieldVal.Kind().String() + "': " + field.Name)
			}
		}
	}
	return nil
}

func handleDuration(fieldVal reflect.Value, envVal string) error {
	duration, err := time.ParseDuration(envVal)
	if err != nil {
		return errors.New("Invalid time.Duration (" + err.Error() + ")")
	}

	fieldVal.SetInt(duration.Nanoseconds())
	return nil
}

func SetTimeFormat(format string) {
	timeFormat = format
}

func handleTime(fieldVal reflect.Value, envVal string) error {
	if timeFormat == "" {
		return errors.New("time format not provided (envtag.SetTimeFormat)")
	}

	time, err := time.Parse(timeFormat, envVal)
	if err != nil {
		return errors.New("Invalid time.Time (" + err.Error() + ")")
	}

	fieldVal.Set(reflect.ValueOf(time))

	return nil
}
