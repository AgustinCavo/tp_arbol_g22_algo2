package cola

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func crearNodo[T any](dato T) *nodo[T] {
	return &nodo[T]{dato: dato, siguiente: nil}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{primero: nil, ultimo: nil}
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

func (c *colaEnlazada[T]) Encolar(any T) {
	if c.EstaVacia() {
		c.primero = crearNodo(any)
		c.ultimo = c.primero
		c.primero.siguiente = c.ultimo
	} else {
		c.ultimo.siguiente = crearNodo(any)
		c.ultimo = c.ultimo.siguiente
	}
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	aux := c.primero.dato

	if c.primero == c.ultimo {
		c.primero = nil
	} else {
		c.primero = c.primero.siguiente
	}
	return aux
}
