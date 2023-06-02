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

type iterador[T any] struct {
	lista    *listaEnlazada[T]
	anterior *nodoLista[T]
	actual   *nodoLista[T]
}

func nodoCrear[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.primero == nil
}

func (l *listaEnlazada[T]) VerPrimero() T {
	l.validarEstaVacia()
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	l.validarEstaVacia()
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) validarEstaVacia() {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := nodoCrear(dato)
	if !l.EstaVacia() {
		nodo.siguiente = l.primero
		l.primero = nodo
	} else {
		l.primero, l.ultimo = nodo, nodo
	}
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := nodoCrear(dato)
	if !l.EstaVacia() {
		l.ultimo.siguiente = nodo
		l.ultimo = nodo
	} else {
		l.primero, l.ultimo = nodo, nodo
	}
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	dato := l.VerPrimero()
	if l.primero.siguiente == nil {
		l.ultimo = nil
	}
	l.primero = l.primero.siguiente
	l.largo--
	return dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodo := &iterador[T]{lista: l, actual: l.primero}
	for i := 0; i <= l.Largo()-1; i++ {
		if !visitar(nodo.actual.dato) {
			break
		}
		nodo.anterior = nodo.actual
		nodo.actual = nodo.actual.siguiente
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterador[T]{lista: l, actual: l.primero}
}

func (l *iterador[T]) VerActual() T {
	l.validarIteradorTermino()
	return l.actual.dato
}

func (l *iterador[T]) validarIteradorTermino() {
	if !l.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (l *iterador[T]) HaySiguiente() bool {
	return l.actual != nil
}

func (l *iterador[T]) Siguiente() {
	l.validarIteradorTermino()
	l.anterior, l.actual = l.actual, l.actual.siguiente
}

func (l *iterador[T]) Insertar(dato_entrante T) {
	nodo := nodoCrear(dato_entrante)
	if l.anterior == nil {
		if l.lista.EstaVacia() {
			l.lista.ultimo = nodo
		}
		l.lista.primero = nodo
		nodo.siguiente = l.actual
		l.actual = l.lista.primero
	} else {
		if l.actual == nil {
			l.lista.ultimo = nodo
		}
		l.anterior.siguiente = nodo
		nodo.siguiente = l.actual
		l.actual = nodo
	}
	l.lista.largo++
}

func (l *iterador[T]) Borrar() T {
	l.validarIteradorTermino()
	dato := l.actual.dato
	if l.anterior == nil {
		if l.actual.siguiente == nil {
			l.lista.ultimo = nil
		}
		l.lista.primero = l.actual.siguiente
	} else {
		if l.actual.siguiente == nil {
			l.lista.ultimo = l.anterior
		}
		l.anterior.siguiente = l.actual.siguiente
	}
	l.actual = l.actual.siguiente
	l.lista.largo--
	return dato
}
