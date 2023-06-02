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
	arbol  *abb[K, V]
	actual *nodo[K, V]
	pila   TDAPila.Pila[*nodo[K, V]]
	desde  *K
	hasta  *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.funcion_cmp = funcion_cmp
	return abb
}

func crearNodo[K comparable, V any](clave K, valor V) *nodo[K, V] {
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
		*nodoR = crearNodo(clave, dato)
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

func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	corte := false
	corte1 := &corte
	ab.raiz.iterar(visitar, corte1)
}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if (desde == hasta && hasta == nil) || ab.raiz == nil {
		ab.Iterar(visitar)
		return
	}
	iter := new(iteradorArbol[K, V])
	iter.actual = ab.raiz
	iter.pila = TDAPila.CrearPilaDinamica[*nodo[K, V]]()
	corte := false
	corte1 := &corte

	if desde == nil { //busco el elemento mas chico
		iter.actual.buscar_Primero(iter.actual, &iter.pila)
		desde = &iter.pila.VerTope().par.clave
	} else if hasta == nil { //busco el elemento mas grande
		iter.actual.buscar_Ultimo(iter.actual, &iter.pila)
		hasta = &iter.pila.VerTope().par.clave
	}

	//itero por rango
	iter.actual.iterarRango(iter.actual, visitar, desde, hasta, ab.funcion_cmp, corte1)
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iteradorArbol[K, V])
	iter.pila = TDAPila.CrearPilaDinamica[*nodo[K, V]]()
	iter.actual = ab.raiz

	iter.actual.buscar_Primero(iter.actual, &iter.pila)
	return iter
}

func (i *iteradorArbol[K, V]) HaySiguiente() bool {
	if (i.pila).EstaVacia() {
		return false
	}
	if i.desde == i.hasta && i.hasta == nil {
		return true
	}
	nodo_act := i.pila.VerTope()
	clave_act := nodo_act.par.clave
	funcion := i.arbol.funcion_cmp
	hasta := *i.hasta
	if funcion(clave_act, hasta) > 0 {
		return false
	}
	return true
}

func (i *iteradorArbol[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodo := i.pila.VerTope()
	return nodo.par.clave, nodo.par.dato
}

func (i *iteradorArbol[K, V]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	i.actual = i.pila.Desapilar()
	if i.actual.der != nil {
		i.actual = i.actual.der
		if i.desde == i.hasta && i.hasta == nil {
			i.actual.buscar_Primero(i.actual, &i.pila)
		} else {
			i.buscarPrimeroIterRango(i.actual, &i.pila, i.desde, i.arbol.funcion_cmp)
		}
	}
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

	iter := new(iteradorArbol[K, V])
	iter.arbol = ab
	iter.pila = TDAPila.CrearPilaDinamica[*nodo[K, V]]()
	iter.actual = ab.raiz
	iter.desde = desde
	iter.hasta = hasta
	if iter.actual == nil || (iter.desde == iter.hasta && iter.hasta == nil) { //Caso arbol vacio o no hay iteracion con rango
		return ab.Iterador()
	}
	//Busco un elemento mayor o igual a desde
	pila_aux2 := TDAPila.CrearPilaDinamica[*nodo[K, V]]()
	if iter.desde == nil {
		iter.actual.buscar_Primero(iter.actual, &pila_aux2) //Busco el primero de todos los elementos
		iter.desde = &pila_aux2.VerTope().par.clave

	} else if iter.hasta == nil {
		iter.actual.buscar_Ultimo(iter.actual, &pila_aux2) //busco el ultimo de todos los elementos
		iter.hasta = &pila_aux2.VerTope().par.clave

	}
	pila_aux := TDAPila.CrearPilaDinamica[*nodo[K, V]]()
	pila_con_el_primero_como_tope := iter.actual.buscarPrimerMayor_mas_chicoIterRango(iter.actual, &pila_aux, iter.desde, ab.funcion_cmp)
	if (*pila_con_el_primero_como_tope).EstaVacia() { //No encontramos un mayor elemento que "desde"---estamos en un rango invalido
		return iter
	}
	//Si encontre
	iter.actual = (*pila_con_el_primero_como_tope).VerTope()

	iter.buscarPrimeroIterRango(iter.actual, &iter.pila, iter.desde, ab.funcion_cmp)
	return iter
}

func (nodoR *nodo[K, V]) iterar(visitar func(clave K, dato V) bool, corte *bool) { //recorrido inorder

	if nodoR == nil {
		return
	}
	if !*corte {
		nodoR.izq.iterar(visitar, corte)
	}
	if !*corte {
		if !visitar(nodoR.par.clave, nodoR.par.dato) {
			*corte = true
			return
		}
	}
	if !*corte {
		nodoR.der.iterar(visitar, corte)
	}
}

func (*nodo[K, V]) buscar_Primero(nodoR *nodo[K, V], pila *TDAPila.Pila[*nodo[K, V]]) *TDAPila.Pila[*nodo[K, V]] { //Debo ir buscando el primero y apilando el resto
	if nodoR == nil {
		return pila
	}
	(*pila).Apilar(nodoR)
	return nodoR.buscar_Primero(nodoR.izq, pila)
}

func (*nodo[K, V]) buscar_Ultimo(nodoR *nodo[K, V], pila *TDAPila.Pila[*nodo[K, V]]) *TDAPila.Pila[*nodo[K, V]] { //Debo ir buscando el primero y apilando el resto
	if nodoR == nil {
		return pila
	}
	(*pila).Apilar(nodoR)
	return nodoR.buscar_Ultimo(nodoR.der, pila)
}

/////Iterador Externo por rango

func (nodo *nodo[K, V]) buscarPrimerMayor_mas_chicoIterRango(nodoR *nodo[K, V], pila *TDAPila.Pila[*nodo[K, V]], desde *K, compar func(K, K) int) *TDAPila.Pila[*nodo[K, V]] { //Debo ir buscando el primero y apilando el resto
	if nodoR == nil {
		return pila
	}
	if compar(nodoR.par.clave, *desde) >= 0 {
		(*pila).Apilar(nodoR)
		return pila
	}
	return nodoR.buscarPrimerMayor_mas_chicoIterRango(nodoR.der, pila, desde, compar)
}

func (i *iteradorArbol[K, V]) buscarPrimeroIterRango(nodoR *nodo[K, V], pila *TDAPila.Pila[*nodo[K, V]], desde *K, compar func(K, K) int) *TDAPila.Pila[*nodo[K, V]] { //Debo ir buscando el primero y apilando el resto
	if i.actual == nil {
		return pila
	}
	if compar(nodoR.par.clave, *desde) < 0 {
		i.actual = i.actual.der
		return i.buscarPrimeroIterRango(i.actual, pila, desde, compar)
	} else if compar(nodoR.par.clave, *desde) >= 0 {
		(*pila).Apilar(nodoR)
		i.actual = i.actual.izq
	}
	return i.buscarPrimeroIterRango(i.actual, pila, desde, compar)
}

func (nodo *nodo[K, V]) iterarRango(nodoR *nodo[K, V], visitar func(clave K, dato V) bool, desde *K, hasta *K, funcion_cmp func(K, K) int, corte *bool) { //recorrido inorder

	if nodoR == nil {
		return
	}
	if funcion_cmp(nodoR.par.clave, *desde) >= 0 && !*corte {
		nodoR.iterarRango(nodoR.izq, visitar, desde, hasta, funcion_cmp, corte)
	}
	if funcion_cmp(nodoR.par.clave, *desde) >= 0 && funcion_cmp(nodoR.par.clave, *hasta) <= 0 && !*corte {
		if !visitar(nodoR.par.clave, nodoR.par.dato) {
			*corte = true
			return
		}
	}
	if funcion_cmp(nodoR.par.clave, *hasta) <= 0 && !*corte {
		nodoR.iterarRango(nodoR.der, visitar, desde, hasta, funcion_cmp, corte)
	}
}
