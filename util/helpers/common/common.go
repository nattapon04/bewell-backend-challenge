package common

import (
	"time"
)

func GetBoolPointer(b bool) *bool {
	return &b
}

func GetStringPointer(s string) *string {
	return &s
}

func GetIntPointer(i int) *int {
	return &i
}

func GetTimePointer(d time.Time) *time.Time {
	return &d
}

func GetInt64Pointer(i int64) *int64 {
	return &i
}

func GetFloat64Pointer(f float64) *float64 {
	return &f
}

func GetFloat32Pointer(f float32) *float32 {
	return &f
}
