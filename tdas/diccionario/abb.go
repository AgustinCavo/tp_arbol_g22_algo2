package diccionario

import (
	TDAPila "tdas/pila"
)

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
	arbol *abb[K, V]
	//actual *nodo[K, V]
	pila TDAPila.Pila[*nodo[K, V]]
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
			(*nodo).izq = ab.borrarNodo(&(*nodo).izq, (*nodo).par.clave)
		} else if (*nodo).der != nil {

			return (*nodo).der
		} else if (*nodo).izq != nil {

			return (*nodo).izq
		} else {

			return nil
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
	iter.pila = TDAPila.CrearPilaDinamica[*nodo[K, V]]()
	//Chequear caso arbol vacio
	iter.pila.Apilar(ab.raiz)
	ab.raiz.buscarPrimero(ab.raiz, &iter.pila)
	return iter
}
func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	ab.raiz.iterar(visitar)
}

func (nodoR *nodo[K, V]) iterar(visitar func(clave K, dato V) bool) { //recorrido inorder

	if nodoR == nil {
		return
	}
	nodoR.izq.iterar(visitar)
	if !visitar(nodoR.par.clave, nodoR.par.dato) {
		return
	}
	nodoR.der.iterar(visitar)
}

func (*nodo[K, V]) buscarPrimero(nodoR *nodo[K, V], pila *TDAPila.Pila[*nodo[K, V]]) *TDAPila.Pila[*nodo[K, V]] { //Debo ir buscando el primero y apilando el resto
	if nodoR == nil {
		return pila
	}
	if nodoR.izq != nil {
		(*pila).Apilar(nodoR.izq)
	}
	return nodoR.buscarPrimero(nodoR.izq, pila)
}

func (i *iteradorArbol[K, V]) HaySiguiente() bool {
	return (i.pila).EstaVacia()
}

func (i *iteradorArbol[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	//return i.actual.par.clave, i.actual.par.dato
	nodo := i.pila.VerTope()
	return nodo.par.clave, nodo.par.dato
}

func (i *iteradorArbol[K, V]) Siguiente() {
	//nodoactual := i.pila.Desapilar()
}
