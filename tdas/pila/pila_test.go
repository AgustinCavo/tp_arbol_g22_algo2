package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()
	require.EqualValues(t, true, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pila.Apilar(1)
	require.EqualValues(t, false, pila.EstaVacia())

}
func TestVariosElementos(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(1)
	require.EqualValues(t, 1, pila.VerTope())
	pila.Apilar(2)
	require.EqualValues(t, 2, pila.VerTope())
	pila.Apilar(3)
	require.EqualValues(t, 3, pila.VerTope())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

}
func TestVolumen(t *testing.T) {
	cargas := []int{10, 1000, 100000}

	for _, carga := range cargas {

		pila := TDAPila.CrearPilaDinamica[int]()

		for i := 0; i <= carga; i++ {
			pila.Apilar(i)
			require.EqualValues(t, i, pila.VerTope())
		}
		for i := carga; i >= 0; i-- {
			require.EqualValues(t, i, pila.Desapilar())
		}
		require.EqualValues(t, true, pila.EstaVacia())
		require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
		require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	}
}
func TestStrings(t *testing.T) {
	src := []string{"Sublime", "Vscode", "Intellij", "Eclipse"}

	pila := TDAPila.CrearPilaDinamica[string]()

	for _, v := range src {
		pila.Apilar(v)
		require.EqualValues(t, v, pila.VerTope())
	}
	for i := len(src) - 1; i >= 0; i-- {
		require.EqualValues(t, src[i], pila.Desapilar())
	}
}
