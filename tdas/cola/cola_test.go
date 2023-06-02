package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con cola vacia")
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolar(t *testing.T) {
	t.Log("Hacemos pruebas encolando algunos elementos")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.Equal(t, 1, cola.VerPrimero())
	cola.Encolar(2)
	require.Equal(t, 1, cola.VerPrimero())
	cola.Encolar(3)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.VerPrimero())
	cola.Encolar(4)
	cola.Encolar(5)
	cola.Encolar(7)
	require.Equal(t, 1, cola.VerPrimero())
}

func TestDesencolar(t *testing.T) {
	t.Log("Hacemos pruebas desencolando algunos elementos")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	cola.Encolar(5) //
	cola.Desencolar()
	require.EqualValues(t, 2, cola.VerPrimero())
	cola.Desencolar()
	require.EqualValues(t, 3, cola.VerPrimero())
	cola.Desencolar()
	cola.Desencolar()
	require.False(t, cola.EstaVacia())
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestVerPrimero(t *testing.T) {
	t.Log("Hacemos pruebas para ver el primero de la cola")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	require.Equal(t, 1, cola.VerPrimero())
	cola.Desencolar()
	require.EqualValues(t, 2, cola.VerPrimero())
}

func TestEncolaryDesencolar(t *testing.T) {
	t.Log("Hacemos pruebas para ver si se mantiene la invariante de cola")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	cola.Encolar(5)
	cola.Desencolar()
	require.EqualValues(t, 2, cola.VerPrimero())
	cola.Desencolar()
	require.EqualValues(t, 3, cola.VerPrimero())
	cola.Desencolar()
	require.EqualValues(t, 4, cola.VerPrimero())
	cola.Desencolar()
	require.EqualValues(t, 5, cola.VerPrimero())
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolaryDesencolarv2(t *testing.T) {
	t.Log("Hacemos pruebas para ver si se mantiene la invariante de cola")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	cola.Encolar(6)
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolaryDesencolarv3(t *testing.T) {
	t.Log("Hacemos pruebas con elementos diferentes")
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("Hola")
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	cola.Encolar("Como")
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestPruebaVolumen(t *testing.T) {
	tam := 1000
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < tam; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < tam; i++ {
		cola.Desencolar()
	}
}

func TestEncolaryDesencolarv4(t *testing.T) {
	t.Log("primero se encola y desencola constantemente. Luego se encolan muchos elementos, desencolando a medida que se van encolando (validando que el desencolado sea correcto), y luego se desencolan todos los restantes validando que sean correctos. Al final la cola debe quedar vacía.")
	cola := TDACola.CrearColaEnlazada[int]()
	tam := 100
	for i := 0; i < tam; i++ {
		cola.Encolar(i)
		cola.Desencolar()
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	for j := 0; j < tam; j++ {
		for i := 0; i < 10; i++ {
			cola.Encolar(i)
		}
		cola.Desencolar()
	}
	for j := 0; j < 900; j++ {
		cola.Desencolar()
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}
