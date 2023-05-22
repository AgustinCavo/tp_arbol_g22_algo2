package lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}
type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}
type nodoMovil[T any] struct {
	lista    *listaEnlazada[T]
	anterior *nodoLista[T]
	actual   *nodoLista[T]
}

func nodoCrear[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	return nodo
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.primero == nil
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}
func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := nodoCrear(dato)
	if l.EstaVacia() {
		l.primero = nodo
		l.ultimo = nodo
	} else {
		nodo.siguiente = l.primero
		l.primero = nodo
	}
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := nodoCrear(dato)
	if l.EstaVacia() {
		l.primero = nodo
		l.ultimo = nodo
	} else {
		l.ultimo.siguiente = nodo
		l.ultimo = nodo
	}
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	dato := l.VerPrimero()
	if l.primero.siguiente == nil {
		l.ultimo = l.ultimo.siguiente
	}
	l.primero = l.primero.siguiente
	l.largo--
	return dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodo := new(nodoMovil[T])
	nodo.lista = l
	nodo.actual = l.primero

	for i := 0; i <= l.Largo()-1; i++ {
		if !visitar(nodo.actual.dato) {
			return
		} else {
			nodo.anterior = nodo.actual
			nodo.actual = nodo.actual.siguiente
		}
	}
}
func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {

	iterador := new(nodoMovil[T])

	iterador.lista = l
	iterador.actual = l.primero
	iterador.anterior = nil

	return iterador
}
func (l *nodoMovil[T]) VerActual() T {
	if l.HaySiguiente() {
		return l.actual.dato
	} else {
		panic("El iterador termino de iterar")
	}
}
func (l *nodoMovil[T]) HaySiguiente() bool {
	return l.actual != nil
}
func (l *nodoMovil[T]) Siguiente() {
	if l.HaySiguiente() {
		l.anterior = l.actual
		l.actual = l.actual.siguiente
	} else {
		panic("El iterador termino de iterar")
	}
}
func (l *nodoMovil[T]) Insertar(dato_entrante T) {
	if l.lista.EstaVacia() || l.anterior == nil {

		nodo := nodoCrear(dato_entrante)
		if l.lista.EstaVacia() {
			l.lista.primero = nodo
			l.lista.ultimo = nodo
		} else {
			nodo.siguiente = l.lista.primero
			l.lista.primero = nodo
		}
		l.lista.largo++
		l.actual = l.lista.primero

	} else if l.HaySiguiente() {

		nodo := nodoCrear(dato_entrante)
		nodo.siguiente = l.actual
		l.anterior.siguiente = nodo
		l.actual = nodo
		l.lista.largo++

	} else {

		nodo := nodoCrear(dato_entrante)
		if l.lista.EstaVacia() {
			l.lista.primero = nodo
			l.lista.ultimo = nodo
		} else {
			l.lista.ultimo.siguiente = nodo
			l.lista.ultimo = nodo
		}
		l.lista.largo++
		l.actual = l.lista.ultimo

	}
}
func (l *nodoMovil[T]) Borrar() T {
	if !l.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	dato := l.actual.dato

	if l.HaySiguiente() {
		if l.anterior == nil {
			if l.actual.siguiente == nil {
				l.actual = nil
				l.lista.primero = nil
				l.lista.ultimo = nil
			} else {
				l.lista.primero = l.actual.siguiente
				l.actual = l.actual.siguiente
			}
		} else {
			if l.actual.siguiente == nil {
				l.anterior.siguiente = nil
				l.lista.ultimo = l.anterior
				l.actual = nil
			} else {
				l.anterior.siguiente = l.actual.siguiente
				l.actual = l.actual.siguiente
			}
		}
	} else {
		if l.anterior != nil {
			l.anterior.siguiente = nil
			l.lista.ultimo = l.anterior
			l.actual = l.anterior
		}
	}
	l.lista.largo--
	return dato
}
