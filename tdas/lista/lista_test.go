package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con lista vacia")
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsertarPrimero(t *testing.T) {
	t.Log("Hacemos pruebas insertando algunos elementos")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	lista.InsertarPrimero(2)
	require.Equal(t, 2, lista.VerPrimero())
	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	require.Equal(t, 5, lista.Largo())
}

func TestBorrarPrimero(t *testing.T) {
	t.Log("Hacemos pruebas Borrando algunos elementos")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	lista.BorrarPrimero()
	require.EqualValues(t, 4, lista.VerPrimero())
	lista.BorrarPrimero()
	require.EqualValues(t, 3, lista.VerPrimero())
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.False(t, lista.EstaVacia())
	lista.BorrarPrimero()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.True(t, lista.EstaVacia())
}

func TestVerPrimero(t *testing.T) {
	t.Log("Hacemos pruebas para ver el primero de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	require.Equal(t, 2, lista.VerPrimero())
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	require.Equal(t, 4, lista.VerPrimero())
}

func TestVerUltimo(t *testing.T) {
	t.Log("Hacemos pruebas para ver el ultimo de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.VerUltimo())
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	require.Equal(t, 4, lista.VerUltimo())
}

func TestInsertarPrimero_InsertaUltimo(t *testing.T) {
	t.Log("Hacemos pruebas insertando al principio y al final de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.Largo())
	lista.InsertarPrimero(0)
	lista.InsertarUltimo(3)
	lista.InsertarPrimero(-1)
	lista.InsertarUltimo(4)
	lista.InsertarPrimero(-2)
	lista.InsertarUltimo(5)
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, -2, lista.VerPrimero())
}

func TestBorrareInsertar(t *testing.T) {
	t.Log("Hacemos pruebas insertando al principio y al final de la lista, tambien borrando para ver si la estructura mantiene sus invariantes")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 1, lista.VerPrimero())
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	lista.InsertarPrimero(0)
	lista.InsertarUltimo(3)
	lista.InsertarPrimero(-1)
	lista.InsertarUltimo(4)
	lista.InsertarPrimero(-2)
	lista.InsertarUltimo(5)
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 0, lista.VerPrimero())
}
func TestVerLargoLista(t *testing.T) {
	t.Log("Hacemos pruebas insertando elementos y borrando para ver si se mantiene el largo de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarPrimero(0)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(1)
	require.Equal(t, 6, lista.Largo())
	lista.BorrarPrimero()
	require.Equal(t, 5, lista.Largo())
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.Equal(t, 1, lista.Largo())
	lista.BorrarPrimero()
	require.Equal(t, 0, lista.Largo())
}

func TestLista_strings(t *testing.T) {
	t.Log("Hacemos pruebas insertando elementos de otro tipo")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("Hola")
	lista.InsertarPrimero("Como")
	lista.InsertarPrimero("Estas")
	lista.BorrarPrimero()
	require.Equal(t, "Como", lista.VerPrimero())
	require.Equal(t, "Hola", lista.VerUltimo())
	lista.BorrarPrimero()
	require.Equal(t, "Hola", lista.VerPrimero())
	require.Equal(t, "Hola", lista.VerUltimo())
	lista.BorrarPrimero()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.True(t, lista.EstaVacia())
}

func TestPruebaVolumenInsertarPrimero(t *testing.T) {
	const TAMANIO int = 100000
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < TAMANIO; i++ {
		lista.InsertarPrimero(i)
	}
	require.Equal(t, TAMANIO-1, lista.VerPrimero())
	for i := 0; i < TAMANIO; i++ {
		lista.BorrarPrimero()
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestIteradorVerActual(t *testing.T) {
	t.Log("Hacemos pruebas viendo los elementos actuales de la lista y se recorre")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)

	iterador := lista.Iterador()

	require.EqualValues(t, 5, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 4, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
	iterador.Siguiente()
	iterador.Siguiente()
	require.EqualValues(t, 1, iterador.VerActual())
	iterador.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
}

func TestIteradorVerActualVolumen(t *testing.T) {
	t.Log("Hacemos pruebas viendo los elementos actuales de la lista")
	const TAMANIO int = 100000
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= TAMANIO; i++ {
		lista.InsertarPrimero(i)
	}

	iterador := lista.Iterador()

	for i := TAMANIO; i >= 0; i-- {
		require.EqualValues(t, i, iterador.VerActual())
		iterador.Siguiente()
	}
}

func TestIteradorInsertarVacio(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador con vacio")
	lista := TDALista.CrearListaEnlazada[int]()

	iterador := lista.Iterador()

	require.EqualValues(t, true, lista.EstaVacia())
	iterador.Insertar(1)
	require.EqualValues(t, 1, iterador.VerActual())

}

func TestIteradorInsertarPrincipio(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador con vacio")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)

	iterador := lista.Iterador()

	require.EqualValues(t, 2, iterador.VerActual())
	iterador.Insertar(1)
	require.EqualValues(t, 1, iterador.VerActual())
	require.EqualValues(t, true, iterador.HaySiguiente())
	iterador.Siguiente()
	require.EqualValues(t, 2, iterador.VerActual())

}

func TestIteradorInsertarBasico(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador insertando")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(1)

	iterador := lista.Iterador()

	require.EqualValues(t, 1, iterador.VerActual())
	require.EqualValues(t, true, iterador.HaySiguiente())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
	iterador.Insertar(2)
	require.EqualValues(t, 2, iterador.VerActual())
	require.EqualValues(t, true, iterador.HaySiguiente())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
}

func TestIteradorInsertarBorde(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador insertando en el borde")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iterador := lista.Iterador()

	iterador.Siguiente()
	require.EqualValues(t, 2, iterador.VerActual())
	iterador.Siguiente()
	iterador.Insertar(3)
	require.EqualValues(t, 3, iterador.VerActual())
}

func TestIteradorBorrar3Casos(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador borrando")
	lista := TDALista.CrearListaEnlazada[int]()

	iterador := lista.Iterador()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	lista.InsertarPrimero(1)

	iterador2 := lista.Iterador()

	require.EqualValues(t, 1, iterador2.Borrar())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.EqualValues(t, false, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador3 := lista.Iterador()

	iterador3.Siguiente()
	require.EqualValues(t, 2, iterador3.VerActual())
	require.EqualValues(t, 2, iterador3.Borrar())
	require.EqualValues(t, true, iterador3.HaySiguiente())
	require.EqualValues(t, 3, iterador3.VerActual())
	require.EqualValues(t, true, iterador3.HaySiguiente())
	require.EqualValues(t, 3, iterador3.Borrar())
}

func TestIteradorBorrar2Casos(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador borrando")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)

	iterador := lista.Iterador()

	require.EqualValues(t, 1, iterador.VerActual())
	require.EqualValues(t, true, iterador.HaySiguiente())
	require.EqualValues(t, 1, iterador.Borrar())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 2, iterador.VerActual())
	require.EqualValues(t, 2, iterador.Borrar())
	lista.InsertarPrimero(1)

	iterador2 := lista.Iterador()

	require.EqualValues(t, 1, iterador2.VerActual())
	require.EqualValues(t, true, iterador2.HaySiguiente())
	iterador2.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador2.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador2.Borrar() })
}

func TestIteradorBorrarCaso(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador borrando")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iterador := lista.Iterador()

	require.EqualValues(t, 1, iterador.VerActual())
	require.EqualValues(t, true, iterador.HaySiguiente())
	iterador.Siguiente()
	require.EqualValues(t, 2, iterador.VerActual())
	require.EqualValues(t, true, iterador.HaySiguiente())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
	require.EqualValues(t, true, iterador.HaySiguiente()) //va bien
	require.EqualValues(t, 3, iterador.Borrar())
	require.EqualValues(t, false, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
}

func TestIteradorVolumen(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador insertando y borrando en volumen")
	const TAMANIO int = 100000
	lista := TDALista.CrearListaEnlazada[int]()

	iterador := lista.Iterador()

	for i := 1; i <= TAMANIO; i++ {
		iterador.Insertar(i)
		require.EqualValues(t, i, iterador.VerActual())
	}
	require.EqualValues(t, TAMANIO, lista.Largo())
	for i := lista.Largo(); i >= 1; i-- {
		require.EqualValues(t, i, iterador.VerActual())
		require.EqualValues(t, i, iterador.Borrar())
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
}

func TestIterarNumerosCompleto(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador interno con funciones de numeros")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(6)
	lista.InsertarPrimero(50)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(1)

	acum := 0
	acumulador := &acum
	const num int = 10
	lista.Iterar(func(numero int) bool {
		*acumulador = *acumulador + numero
		return true
	})
	require.EqualValues(t, 65, acum)

	cont := 0
	contador := &cont
	lista.Iterar(func(numero int) bool {
		if numero > num {
			*contador++
		}
		return true
	})
	require.EqualValues(t, 1, cont)

	cont_pares := 0
	contador_pares := &cont_pares
	lista.Iterar(func(numero int) bool {
		if numero%2 == 0 {
			*contador_pares++
		}
		return true
	})
	require.EqualValues(t, 2, cont_pares)
}

func TestIterarNumerosCorta(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador interno con funciones de numeros")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(6)
	lista.InsertarPrimero(50)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(1)

	cont := 0
	contador := &cont
	lista.Iterar(func(numero int) bool {
		if numero < 10 {
			*contador++
		} else {
			return false
		}
		return true
	})
	require.EqualValues(t, 3, cont)
}

func TestIterarLetras(t *testing.T) {
	t.Log("Hacemos pruebas viendo como se comporta el iterador interno con funciones de strings")
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero(" Juan")
	lista.InsertarPrimero(" Homero")
	lista.InsertarPrimero(" Roberto")
	lista.InsertarPrimero(" Pedro")
	lista.InsertarPrimero("Ramiro")

	var cadena string
	cadena_ptr := &cadena

	lista.Iterar(func(v string) bool {
		if v == " Homero" {
			return false
		} else {
			*cadena_ptr = *cadena_ptr + v
			return true
		}
	})
	require.EqualValues(t, "Ramiro Pedro Roberto", cadena)
}
