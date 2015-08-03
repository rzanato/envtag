package envtag

import (
	"math"
	"os"
	"strconv"
	"testing"
)

func init() {
	os.Clearenv()
}

func TestSimple(t *testing.T) {
	type TestStruct struct {
		Field string `env:"TEST_FIELD"`
	}

	testStruct := TestStruct{}

	os.Setenv("TEST_FIELD", "test")

	err := Unmarshal(&testStruct)

	if err != nil || testStruct.Field != "test" {
		t.Error(err.Error())
	}
}

func TestTypes(t *testing.T) {
	type TestStruct struct {
		StringSimple string  `env:"TEST_STRING_SIMPLE"`
		StringEmpty  string  `env:"TEST_STRING_EMPTY"`
		BoolTrue     bool    `env:"TEST_BOOL_TRUE"`
		BoolFalse    bool    `env:"TEST_BOOL_FALSE"`
		BoolZero     bool    `env:"TEST_BOOL_ZERO"`
		BoolOne      bool    `env:"TEST_BOOL_ONE"`
		IntMax       int     `env:"TEST_INT_MAX"`
		IntMin       int     `env:"TEST_INT_MIN"`
		IntMax8      int8    `env:"TEST_INT_MAX8"`
		IntMin8      int8    `env:"TEST_INT_MIN8"`
		IntMax16     int16   `env:"TEST_INT_MAX16"`
		IntMin16     int16   `env:"TEST_INT_MIN16"`
		IntMax32     int32   `env:"TEST_INT_MAX32"`
		IntMin32     int32   `env:"TEST_INT_MIN32"`
		IntMax64     int64   `env:"TEST_INT_MAX64"`
		IntMin64     int64   `env:"TEST_INT_MIN64"`
		UintMax      uint    `env:"TEST_UINT_MAX"`
		UintMax8     uint8   `env:"TEST_UINT_MAX8"`
		UintMax16    uint16  `env:"TEST_UINT_MAX16"`
		UintMax32    uint32  `env:"TEST_UINT_MAX32"`
		UintMax64    uint64  `env:"TEST_UINT_MAX64"`
		FloatMax32   float32 `env:"TEST_FLOAT_MAX32"`
		FloatMin32   float32 `env:"TEST_FLOAT_MIN32"`
		FloatMax64   float64 `env:"TEST_FLOAT_MAX64"`
		FloatMin64   float64 `env:"TEST_FLOAT_MIN64"`
	}

	testStruct := TestStruct{}

	testValues := map[string]string{
		"TEST_STRING_SIMPLE": "Test",
		"TEST_STRING_EMPTY":  "",
		"TEST_BOOL_TRUE":     "true",
		"TEST_BOOL_FALSE":    "false",
		"TEST_BOOL_ZERO":     "0",
		"TEST_BOOL_ONE":      "1",
		"TEST_INT_MAX":       strconv.Itoa(int(math.MaxInt32)),
		"TEST_INT_MIN":       strconv.Itoa(int(math.MinInt32)),
		"TEST_INT_MAX8":      strconv.Itoa(int(math.MaxInt8)),
		"TEST_INT_MIN8":      strconv.Itoa(int(math.MinInt8)),
		"TEST_INT_MAX16":     strconv.Itoa(int(math.MaxInt16)),
		"TEST_INT_MIN16":     strconv.Itoa(int(math.MinInt16)),
		"TEST_INT_MAX32":     strconv.Itoa(int(math.MaxInt32)),
		"TEST_INT_MIN32":     strconv.Itoa(int(math.MinInt32)),
		"TEST_INT_MAX64":     strconv.FormatInt(math.MaxInt64, 10),
		"TEST_INT_MIN64":     strconv.FormatInt(math.MinInt64, 10),
		"TEST_UINT_MAX":      strconv.FormatUint(uint64(math.MaxUint32), 10),
		"TEST_UINT_MAX8":     strconv.FormatUint(uint64(math.MaxUint8), 10),
		"TEST_UINT_MAX16":    strconv.FormatUint(uint64(math.MaxUint16), 10),
		"TEST_UINT_MAX32":    strconv.FormatUint(uint64(math.MaxUint32), 10),
		"TEST_UINT_MAX64":    strconv.FormatUint(math.MaxUint64, 10),
		"TEST_FLOAT_MAX32":   strconv.FormatFloat(float64(math.MaxFloat32), 'f', -1, 32),
		"TEST_FLOAT_MIN32":   strconv.FormatFloat(float64(math.SmallestNonzeroFloat32), 'f', -1, 32),
		"TEST_FLOAT_MAX64":   strconv.FormatFloat(math.MaxFloat64, 'f', -1, 64),
		"TEST_FLOAT_MIN64":   strconv.FormatFloat(math.SmallestNonzeroFloat64, 'f', -1, 64),
	}

	for k, v := range testValues {
		os.Setenv(k, v)
	}

	err := Unmarshal(&testStruct)

	if err != nil {
		t.Error(err.Error())
	}

	if testStruct.StringSimple != "Test" || testStruct.StringEmpty != "" {
		t.Error("String type failed")
	}

	if !testStruct.BoolTrue || testStruct.BoolFalse || !testStruct.BoolOne || testStruct.BoolZero {
		t.Error("Bool type failed")
	}

	if testStruct.IntMax != math.MaxInt32 || testStruct.IntMin != math.MinInt32 {
		t.Error("Int type failed")
	}

	if testStruct.IntMax8 != math.MaxInt8 || testStruct.IntMin8 != math.MinInt8 {
		t.Error("Int8 type failed")
	}

	if testStruct.IntMax16 != math.MaxInt16 || testStruct.IntMin16 != math.MinInt16 {
		t.Error("Int16 type failed")
	}

	if testStruct.IntMax32 != math.MaxInt32 || testStruct.IntMin32 != math.MinInt32 {
		t.Error("Int32 type failed")
	}

	if testStruct.IntMax64 != math.MaxInt64 || testStruct.IntMin64 != math.MinInt64 {
		t.Error("Int64 type failed")
	}

	if testStruct.UintMax != math.MaxUint32 {
		t.Error("Uint type failed")
	}

	if testStruct.UintMax8 != math.MaxUint8 {
		t.Error("Uint8 type failed")
	}

	if testStruct.UintMax16 != math.MaxUint16 {
		t.Error("Uint16 type failed")
	}

	if testStruct.UintMax32 != math.MaxUint32 {
		t.Error("Uint32 type failed")
	}

	if testStruct.UintMax64 != math.MaxUint64 {
		t.Error("Uint64 type failed")
	}

	if testStruct.FloatMax32 != math.MaxFloat32 || testStruct.FloatMin32 != math.SmallestNonzeroFloat32 {
		t.Error("Float32 type failed")
	}

	if testStruct.FloatMax64 != math.MaxFloat64 || testStruct.FloatMin64 != math.SmallestNonzeroFloat64 {
		t.Error("Float64 type failed")
	}
}

func TestInnerStruct(t *testing.T) {
	type TestStruct struct {
		Sub struct {
			Field string `env:"TEST_FIELD"`
		}
	}

	testStruct := TestStruct{}

	os.Setenv("TEST_FIELD", "test")

	err := Unmarshal(&testStruct)

	if err != nil || testStruct.Sub.Field != "test" {
		t.Error(err.Error())
	}
}

func TestInvalidTypes(t *testing.T) {
	type TestStruct struct {
		Field string `env:"TEST_FIELD"`
	}

	testStruct := TestStruct{}

	if err := Unmarshal(nil); err == nil {
		t.Error("cannot accept nil")
	}

	if err := Unmarshal(testStruct); err == nil {
		t.Error("expected struct reference")
	}
}
