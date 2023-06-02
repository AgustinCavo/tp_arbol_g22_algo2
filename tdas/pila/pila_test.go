package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con pila vacia")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilar(t *testing.T) {
	t.Log("Hacemos pruebas apilando algunos elementos")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.Equal(t, 1, pila.VerTope())
	pila.Apilar(2)
	require.Equal(t, 2, pila.VerTope())
	pila.Apilar(3)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 3, pila.VerTope())
	pila.Apilar(4)
	pila.Apilar(5)
	pila.Apilar(7)
	require.Equal(t, 7, pila.VerTope())
}

func TestDesapilar(t *testing.T) {
	t.Log("Hacemos pruebas desapilando algunos elementos")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	pila.Apilar(5)
	pila.Desapilar()
	require.EqualValues(t, 4, pila.VerTope())
	pila.Desapilar()
	require.EqualValues(t, 3, pila.VerTope())
	pila.Desapilar()
	pila.Desapilar()
	require.EqualValues(t, 1, pila.VerTope())
	require.EqualValues(t, 1, pila.VerTope())
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestVerTope(t *testing.T) {
	t.Log("Hacemos pruebas para ver el tope")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	require.Equal(t, 4, pila.VerTope())
	pila.Desapilar()
	require.EqualValues(t, 3, pila.VerTope())
}

func TestApilaryDesapilar(t *testing.T) {
	t.Log("Hacemos pruebas para ver si se mantiene la invariante de pila")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	pila.Apilar(5)
	pila.Desapilar()
	require.EqualValues(t, 4, pila.VerTope())
	pila.Desapilar()
	require.EqualValues(t, 3, pila.VerTope())
	pila.Desapilar()
	require.EqualValues(t, 2, pila.VerTope())
	pila.Desapilar()
	require.EqualValues(t, 1, pila.VerTope())
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilaryDesapilarv2(t *testing.T) {
	t.Log("Hacemos pruebas para ver si se mantiene la invariante de pila")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	pila.Apilar(6)
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilaryDesapilarv3(t *testing.T) {
	t.Log("Hacemos pruebas con elementos diferentes")
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Hola")
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	pila.Apilar("Como")
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPruebaVolumen(t *testing.T) {
	tam := 1000
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < tam; i++ {
		pila.Apilar(i)
	}
	for i := 0; i < tam; i++ {
		pila.Desapilar()
	}
}

func TestApilaryDesapilarv4(t *testing.T) {
	t.Log("primero se apila y desapila constantemente. Luego se apilan muchos elementos, desapilando a medida que se van apilando (validando que el desapilado sea correcto), y luego se desapilantodos los restantes validando que sean correctos. Al final la pila debe quedar vacía.")
	pila := TDAPila.CrearPilaDinamica[int]()
	tam := 100
	for i := 0; i < tam; i++ {
		pila.Apilar(i)
		pila.Desapilar()
	}
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	for j := 0; j < tam; j++ {
		for i := 0; i < 10; i++ {
			pila.Apilar(i)
		}
		pila.Desapilar()
	}
	for j := 0; j < 900; j++ {
		pila.Desapilar()
	}
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}
