package lista

type IteradorLista[T any] interface {

	//VerActual devuelve el elemento actual de la iteracion
	VerActual() T

	//HaySiguiente Indica true si existe un elemento a continuacion y si el elemento actual existe
	HaySiguiente() bool

	//Siguiente Proxigue al siguiente elemento de la lista
	Siguiente()

	//Insertar Inserta un elemento en la lista previo al actual de la lista.
	Insertar(T)

	//Borrar Elimina el elemento actual de la lista y lo devuelve por pantalla. Si la lista se encontraba vacia, entra en
	// panico con el mensaje 'La lista esta vacia'.
	Borrar() T
}

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta el dato al inicio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta el dato al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero borra el dato al inicio de la lista. Si la lista se encontraba vacia, entra el panico con el
	// mensaje 'La lista esta vacia'.
	BorrarPrimero() T

	// VerPrimero devuelve el elemento al inicio de la lista (el primero). Si la lista se encontraba vacia, entra en
	// panico con el mensaje 'La lista esta vacia'.
	VerPrimero() T

	// VerUltimo devuelve el elemento al final de la lista (el ultimo). Si la lista se encontraba vacia, entra en
	// panico con el mensaje 'La lista esta vacia'.
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista
	Largo() int

	// Iterar aplica la funcion pasada por parametro a todos los elementos de la lista, hasta que no hayan mas
	// elementos, o la funcion en cuestion devuelva false.
	Iterar(func(T) bool)

	//Iterador genera la interface IteradorLista utilizando la lista
	Iterador() IteradorLista[T]
}
