package diccionario

type abb[K comparable, V any] struct {
	raiz        *nodo[K, V]
	cantidad    int
	funcion_cmp func(K, K) int
}
type nodo[K comparable, V any] struct {
	der *nodo[K, V]
	izq *nodo[K, V]
	par *parClaveValor[K, V]
}
type iteradorArbol[K comparable, V any] struct {
	arbol  *abb[K, V]
	actual *nodo[K, V]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])

	abb.funcion_cmp = funcion_cmp
	return abb
}
func CrearNodo[K comparable, V any](clave K, valor V) *nodo[K, V] {
	nodo := new(nodo[K, V])
	nodo.par = crearClaveValor(clave, valor)
	return nodo
}
func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	ab.guardarNodo(&ab.raiz, clave, dato)

}
func (ab *abb[K, V]) guardarNodo(nodoR **nodo[K, V], clave K, dato V) {

	if *nodoR == nil {
		*nodoR = CrearNodo(clave, dato)
		ab.cantidad += 1
		return
	} else if ab.funcion_cmp((*nodoR).par.clave, clave) < 0 {

		ab.guardarNodo(&(*nodoR).der, clave, dato)
	} else if ab.funcion_cmp((*nodoR).par.clave, clave) > 0 {

		ab.guardarNodo(&(*nodoR).izq, clave, dato)
	} else {
		(*nodoR).par.dato = dato
	}
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	pertenece, _ := ab.perteneceNodo(ab.raiz, clave)
	return pertenece
}
func (ab *abb[K, V]) perteneceNodo(nodoR *nodo[K, V], clave K) (bool, *nodo[K, V]) {

	if nodoR == nil {
		return false, nil
	} else if ab.funcion_cmp(nodoR.par.clave, clave) < 0 {

		return (ab.perteneceNodo(nodoR.der, clave))

	} else if ab.funcion_cmp(nodoR.par.clave, clave) > 0 {

		return (ab.perteneceNodo(nodoR.izq, clave))

	} else {

		return true, nodoR
	}
}

func (ab *abb[K, V]) Obtener(clave K) V {
	pertenece, nodo := ab.perteneceNodo(ab.raiz, clave)
	if pertenece {
		return nodo.par.dato
	} else {
		panic("La clave no pertenece al diccionario")
	}
}

func (ab *abb[K, V]) Borrar(clave K) V {

	pertenece, nodo := ab.perteneceNodo(ab.raiz, clave)

	if !pertenece {
		panic("La clave no pertenece al diccionario")

	} else {
		dato := nodo.par.dato
		ab.raiz = ab.borrarNodo(&ab.raiz, clave)
		ab.cantidad -= 1
		return dato

	}

}

func (ab *abb[K, V]) borrarNodo(nodo **nodo[K, V], clave K) *nodo[K, V] {

	if *nodo == nil {
		return *nodo
	} else if ab.funcion_cmp((*nodo).par.clave, clave) < 0 {
		(*nodo).der = (ab.borrarNodo(&(*nodo).der, clave))
	} else if ab.funcion_cmp((*nodo).par.clave, clave) > 0 {
		(*nodo).izq = (ab.borrarNodo(&(*nodo).izq, clave))
	} else {

		if (*nodo).izq != nil && (*nodo).der != nil {
			nodoMinimo := ab.encontrarRemplazante((*nodo).izq)
			(*nodo).par = nodoMinimo.par
			(*nodo).izq = ab.borrarNodo(&(*nodo).izq, clave)
		} else if (*nodo).der != nil {

			return (*nodo).der
		} else if (*nodo).izq != nil {
			return (*nodo).izq
		} else {

			*nodo = nil
			return *nodo
		}
	}
	return *nodo
}

// pasar el nodo izq a esta
func (ab *abb[K, V]) encontrarRemplazante(nodoR *nodo[K, V]) *nodo[K, V] {
	if nodoR.der == nil {
		return nodoR
	} else {
		return ab.encontrarRemplazante(nodoR.der)
	}
}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iteradorArbol[K, V])
	return iter
}
func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iteradorArbol[K, V])
	return iter
}
func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {

}

func (i *iteradorArbol[K, V]) HaySiguiente() bool {

	return false
}

func (i *iteradorArbol[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.actual.par.clave, i.actual.par.dato
}

func (i *iteradorArbol[K, V]) Siguiente() {

}
