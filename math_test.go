package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 40 {
		t.Errorf("Soma incorreta, obtido: %d, esperado: %d.", total, 30)
	}
}
