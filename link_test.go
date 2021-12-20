package main

import "testing"

func TestCodigo(t *testing.T) {
	aux := GeradorCodigo()

	if len(aux) != 6 {
		t.Errorf("GeradorCodigo() failed, expected length 6, got %v instead", len(aux))
	} else {
		t.Logf("GeradorCodigo() success, expected length 6, got %v", len(aux))
	}
}

func TestUrl(t *testing.T) {
	aux := GeradorURL("abs67D")

	if aux != "http://go.io/abs67D" {
		t.Errorf("GeradorURL(\"abs67D\") failed, expected %v, got %v", "http://go.io/abs67D", aux)
	} else {
		t.Logf("GeradorURL(\"abs67D\") success, expected %v, got %v", "http://go.io/abs67D", aux)
	}
}
