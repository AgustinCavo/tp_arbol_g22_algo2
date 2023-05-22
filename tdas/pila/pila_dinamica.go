package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	p := pilaDinamica[T]{datos: make([]T, 15, 15), cantidad: 0}
	return &p
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {

	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(any T) {

	p.redimensionar()
	p.datos[p.cantidad] = any

	p.cantidad++

}

func (p *pilaDinamica[T]) Desapilar() T {

	aux := p.VerTope()
	p.cantidad--
	p.redimensionar()

	return aux
}

func (p *pilaDinamica[T]) redimensionar() {

	if cap(p.datos) == p.cantidad {

		aux := make([]T, p.cantidad*2, p.cantidad*2)
		copy(aux, p.datos)
		p.datos = aux
	} else if p.cantidad*4 <= cap(p.datos) && p.cantidad >= 5 {
		aux := make([]T, p.cantidad, cap(p.datos)/2)
		copy(aux, p.datos)
		p.datos = aux
	}
}
