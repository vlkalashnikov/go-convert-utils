package utils

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

// I2S converting any value to a string
func I2S(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// I2I64 converting any value to a int64
func I2I64(value interface{}) int64 {
	if i, err := strconv.ParseInt(I2S(value), 10, 64); err == nil {
		return i
	}

	return 0
}

// B2I converting byte slice to a int
func B2I(b []byte) int {
	if i, err := strconv.Atoi(B2S(b)); err == nil {
		return i
	}

	return 0
}

// B2I64 converting byte slice to a int64
func B2I64(b []byte) int64 {
	if i, err := strconv.ParseInt(B2S(b), 10, 64); err == nil {
		return i
	}

	return 0
}

// B2T converts byte slice to a time.Time
func B2T(b []byte, format string) *time.Time {
	t, err := time.Parse(format, B2S(b))
	if err != nil {
		return nil
	}

	return &t
}

// B2B converts byte slice to a bool
func B2B(b []byte) bool {
	p, err := strconv.ParseBool(B2S(b))
	if err != nil {
		return false
	}

	return p
}

// B2S converts byte slice to a string without memory allocation.
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// S2B converts string to a byte slice without memory allocation.
func S2B(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len

	return
}

// DS returns the value or a default value if it is set
func DS(value string, defaultValue ...string) string {
	if len(value) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

//B2P returns pointer boolean
func B2P(b bool) *bool {
	return &b
}

func BF64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func F64B(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}
