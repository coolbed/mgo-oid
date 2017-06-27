package oid

import (
	"testing"
)

func TestNewOID(t *testing.T) {
	objectID := NewOID().String()
	if len(objectID) != 24 {
		t.Errorf("got invalid object id [%s]", objectID)
	}
}

func TestEqual(t *testing.T) {
	o1 := NewOID()
	o2 := NewOID()
	isEqual := Equal(o1, o1)
	if !isEqual {
		t.Errorf("got false expected true")
	}
	isEqual = Equal(o1, o2)
	if isEqual {
		t.Errorf("got true expected false")
	}
}

func BenchmarkNewOID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewOID()
	}
}

func BenchmarkEqual(b *testing.B) {
	o1 := NewOID()
	o2 := NewOID()
	for i := 0; i < b.N; i++ {
		Equal(o1, o2)
	}
}

func BenchmarkBytes(b *testing.B) {
	o1 := NewOID()
	for i := 0; i < b.N; i++ {
		o1.Bytes()
	}
}

func BenchmarkString(b *testing.B) {
	o1 := NewOID()
	for i := 0; i < b.N; i++ {
		_ = o1.String()
	}
}

func BenchmarkTimestamp(b *testing.B) {
	o1 := NewOID()
	for i := 0; i < b.N; i++ {
		o1.Timestamp()
	}
}
