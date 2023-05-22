package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()

	require.EqualValues(t, true, cola.EstaVacia())

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	cola.Encolar(1)
	require.EqualValues(t, false, cola.EstaVacia())

}
func TestVariosElementos(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(1)
	require.EqualValues(t, 1, cola.VerPrimero())
	cola.Encolar(2)
	require.EqualValues(t, 1, cola.VerPrimero())
	cola.Encolar(3)
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerPrimero())
	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 3, cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.EqualValues(t, true, cola.EstaVacia())

}

func TestVolumen(t *testing.T) {
	cargas := []int{10, 1000, 100000}

	for _, carga := range cargas {

		cola := TDACola.CrearColaEnlazada[int]()

		for i := 0; i <= carga; i++ {
			cola.Encolar(i)
			require.EqualValues(t, 0, cola.VerPrimero())
		}
		for i := 0; i <= carga; i++ {
			require.EqualValues(t, i, cola.Desencolar())
		}
		require.EqualValues(t, true, cola.EstaVacia())
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	}
}

func TestStrings(t *testing.T) {
	src := []string{"Sublime", "Vscode", "Intellij", "Eclipse"}

	cola := TDACola.CrearColaEnlazada[string]()

	for _, v := range src {
		cola.Encolar(v)
		require.EqualValues(t, src[0], cola.VerPrimero())
	}
	for i := 0; i <= len(src)-1; i++ {
		require.EqualValues(t, src[i], cola.Desencolar())
	}
}
func TestEncolarDesconlar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(1)
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerPrimero())
	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, true, cola.EstaVacia())
}
