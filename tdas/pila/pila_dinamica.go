package pila

/* Definición del struct pila proporcionado por la cátedra. */
const tamanio_ini = 2
const num = 4

//Recordatorio: Si  cantidad*num  es igual a la capacidad, nos estamos asegurando que
// tengamos 1/4 de la capacidad usada para luego reducir a la estructura a la mitad.

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
	cap_var  int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, tamanio_ini)
	pila.cap_var = cap(pila.datos) //inicializo con una capacidad inicial igual a la capacidad total
	return pila                    //,esta cap_var va cambiando a medida que apilamos y desapilamos elementos
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if p.cantidad*num == cap(p.datos) {
		p.redimensionar(cap(p.datos) / 2)
		p.cap_var = cap(p.datos) - p.cantidad
		p.datos = p.datos[:cap(p.datos)]
	}
	tope := p.VerTope()
	p.cantidad--
	p.cap_var++
	return tope
}

func (p *pilaDinamica[T]) Apilar(dato T) {
	if p.cap_var == 0 && p.cantidad > 0 {
		p.redimensionar(cap(p.datos) * 2)
		p.cap_var = cap(p.datos) - p.cantidad
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
	p.cap_var--
}

func (p *pilaDinamica[T]) redimensionar(cap int) {
	nuevo := make([]T, cap)
	copy(nuevo, p.datos)
	p.datos = nuevo
}
