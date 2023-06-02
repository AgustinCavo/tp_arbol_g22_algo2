package cola

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}
type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func nodoCrear[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	return nodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nodo := nodoCrear(dato)
	if c.EstaVacia() {
		c.primero = nodo
	} else {
		c.ultimo.prox = nodo
	}
	c.ultimo = nodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	dato := c.VerPrimero()
	c.primero = c.primero.prox
	if c.EstaVacia() {
		c.ultimo = nil
	}
	return dato
}
