package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	aux := GeradorUUID()
	aux_type := reflect.TypeOf(aux).String()

	// if aux_type != "uuid.UUID" {
	// 	t.Errorf("GeradorUUID() failed, expected type uuid.UUID, got %v", aux_type)
	// } else {
	// 	t.Logf("GeradorUUID() success, expected type uuid.UUID, got %v", aux_type)
	// }
	assert := assert.New(t)
	assert.Equal(aux_type, "uuid.UUID", "Expected type uuid.UUID, got type uuid.UUID")
}

func TestCodigo(t *testing.T) {
	aux := GeradorCodigo()

	// if len(aux) != 6 {
	// 	t.Errorf("GeradorCodigo() failed, expected length 6, got %v instead", len(aux))
	// } else {
	// 	t.Logf("GeradorCodigo() success, expected length 6, got %v", len(aux))
	// }
	assert := assert.New(t)
	assert.Equal(len(aux), 6, "Expected length 6, got 6")
}

func TestUrl(t *testing.T) {
	aux := GeradorURL("abs67D")

	// if aux != "http://go.io/abs67D" {
	// 	t.Errorf("GeradorURL(\"abs67D\") failed, expected %v, got %v", "http://go.io/abs67D", aux)
	// } else {
	// 	t.Logf("GeradorURL(\"abs67D\") success, expected %v, got %v", "http://go.io/abs67D", aux)
	// }
	assert := assert.New(t)
	assert.Equal(aux, "http://go.io/abs67D", "Expected 'http://go.io/abs67D', got 'http://go.io/abs67D'")
}

func TestTeste(t *testing.T) {
	assert.True(t, true, "True is true!")
}
