package main

import "testing"

func TestSoma(t *testing.T) {
	x := 15
	y := 15

	total := 30
	totalV := Soma(x, y)

	if total != 30 {
		t.Errorf(
			"Resultado da soma é inválido: Resultado: %d. Esperado: %d",
			total,
			totalV,
		)
	}
}
