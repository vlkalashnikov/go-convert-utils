package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/xml"
	"fmt"
	"math"
	"reflect"
	"runtime"
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

// B2P returns pointer boolean
func B2P(b bool) *bool {
	return &b
}

func P2B(p *bool) bool {
	if p != nil {
		return *p
	}

	return false
}

func B2F64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func F642B(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func I2B(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func P2S(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func P2T(t *time.Time) time.Time {
	var tt time.Time
	if t != nil {
		return *t
	}

	return tt
}

func P2I(p *int) int {
	if p != nil {
		return *p
	}

	return 0
}

func X2S(in interface{}) string {
	out, _ := xml.MarshalIndent(in, " ", "  ")
	return string(out)
}

func DefaultString(value string, defaultValue ...string) string {
	if len(value) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

func DefaultInt(value int, defaultValue ...int) int {
	if value == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

func DefaultInt64(value int64, defaultValue ...int64) int64 {
	if value == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

func RmElemByIdx[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func Filter[T any](ss []T, fn func(T) bool) (ret []T) {
	for _, s := range ss {
		if fn(s) {
			ret = append(ret, s)
		}
	}
	return
}

func FnName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
