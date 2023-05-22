package diccionario

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	TDALista "tdas/lista"
)

const TAM = 100
const FACTORCARGA float32 = 2
const FACTORDESCARGA float32 = 0.5

//FactorCarga--->achicar 0,5

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[*parClaveValor[K, V]]
	tam      int
	cantidad int
}
type iterador[K comparable, V any] struct {
	hash   *hashAbierto[K, V]
	actual *parClaveValor[K, V]
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
func hashSha256[K comparable](clave K) int {
	hash := sha256.Sum256(convertirABytes(clave))
	return int(binary.LittleEndian.Uint32(hash[:]))
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashAbierto[K, V])
	hash.tabla = crearTabla[K, V](TAM)
	hash.tam = TAM
	return hash
}

func (d *hashAbierto[K, V]) Cantidad() int {
	return d.cantidad
}
func crearClaveValor[K comparable, V any](clave K, dato V) *parClaveValor[K, V] {
	par := new(parClaveValor[K, V])
	par.clave = clave
	par.dato = dato
	return par
}
func (d *hashAbierto[K, V]) Guardar(clave K, dato V) {
	par := crearClaveValor(clave, dato)

	claveHasheada := d.claveHasheadaModulo(clave)

	if d.tabla[claveHasheada] == nil {
		d.tabla[claveHasheada] = TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
	}

	d.insertar(par, claveHasheada)

	d.agrandar()
}
func (d *hashAbierto[K, V]) insertar(par *parClaveValor[K, V], posTabla int) {
	iter := d.tabla[posTabla].Iterador()

	for iter.HaySiguiente() {

		if iter.VerActual().clave == par.clave {
			iter.Borrar()
			iter.Insertar(par)
			return
		}
		iter.Siguiente()
	}
	d.cantidad += 1
	iter.Insertar(par)
}

func (d *hashAbierto[K, V]) Pertenece(clave K) bool {
	pertenece, _ := d.encontrarPar(clave)
	return pertenece
}

func (d *hashAbierto[K, V]) Obtener(clave K) V {

	pertenece, iter := d.encontrarPar(clave)

	if !pertenece {
		panic("La clave no pertenece al diccionario")
	}

	return iter.VerActual().dato
}

func (d *hashAbierto[K, V]) Borrar(clave K) V {

	pertenece, iter := d.encontrarPar(clave)

	if !pertenece {
		panic("La clave no pertenece al diccionario")
	}

	dato := iter.VerActual().dato

	iter.Borrar()

	if d.tabla[d.claveHasheadaModulo(clave)].EstaVacia() {
		d.tabla[d.claveHasheadaModulo(clave)] = nil
	}
	d.cantidad -= 1
	d.achicar()
	return dato
}

func (d *hashAbierto[K, V]) agrandar() {
	aux := *d

	if float32((d.cantidad / d.tam)) >= FACTORCARGA {

		d.tam = d.tam * 2
		d.tabla = crearTabla[K, V](d.tam)
		d.copiar(aux)

	}
}
func (d *hashAbierto[K, V]) achicar() {
	aux := *d
	if float32((d.cantidad / d.tam)) <= FACTORDESCARGA {

		d.tam = d.tam / 2
		d.tabla = crearTabla[K, V](d.tam)

		d.copiar(aux)

	}
}
func (d *hashAbierto[K, V]) copiar(aux hashAbierto[K, V]) {
	d.cantidad = 0
	for index, value := range aux.tabla {
		if aux.tabla[index] != nil {
			iter := value.Iterador()
			for iter.HaySiguiente() {
				d.Guardar(iter.VerActual().clave, iter.VerActual().dato)
				iter.Siguiente()
			}
		}
	}
}

func (d *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	_, par := d.encontrarSiguienteHash(-1)
	return &iterador[K, V]{hash: d, actual: par}
}
func (d *hashAbierto[K, V]) claveHasheadaModulo(clave K) int {
	return hashSha256(clave) % d.tam
}
func (d *hashAbierto[K, V]) encontrarSiguienteHash(posicion int) (bool, *parClaveValor[K, V]) {
	inicio := posicion + 1

	for i := inicio; i < d.tam; i++ {
		if d.tabla[i] != nil {
			primero := d.tabla[i].VerPrimero()
			return true, primero
		}
	}
	return false, nil
}
func (i *iterador[K, V]) HaySiguiente() bool {
	if i.actual != nil {

		_, iter := i.hash.encontrarPar(i.actual.clave)
		if iter.HaySiguiente() {
			return true
		} else {
			claveHasheada := i.hash.claveHasheadaModulo(i.actual.clave)
			siguiente, _ := i.hash.encontrarSiguienteHash(claveHasheada)
			return siguiente
		}
	}
	return false
}

func (i *iterador[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.actual.clave, i.actual.dato
}

func (i *iterador[K, V]) Siguiente() {
	if i.HaySiguiente() {

		_, iter := i.hash.encontrarPar(i.actual.clave)
		if iter.HaySiguiente() {
			iter.Siguiente()
		}
		if iter.HaySiguiente() {
			i.actual = iter.VerActual()

		} else {

			claveHasheada := i.hash.claveHasheadaModulo(i.actual.clave)
			siguiente, par := i.hash.encontrarSiguienteHash(claveHasheada)
			if siguiente {
				i.actual = par
			} else {
				i.actual = nil

			}
		}
	} else {
		panic("El iterador termino de iterar")
	}

}

func (d *hashAbierto[K, V]) encontrarPar(clave K) (bool, TDALista.IteradorLista[*parClaveValor[K, V]]) {

	claveHasheada := d.claveHasheadaModulo(clave)

	if d.tabla[claveHasheada] == nil {
		return false, nil
	}

	iter := d.tabla[claveHasheada].Iterador()

	for iter.HaySiguiente() {
		if iter.VerActual().clave == clave {
			return true, iter
		}

		iter.Siguiente()
	}
	return false, nil
}
func (d *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {

	for index, value := range d.tabla {
		if d.tabla[index] != nil {
			iter := value.Iterador()
			for iter.HaySiguiente() {
				clave := iter.VerActual().clave
				dato := iter.VerActual().dato
				if !visitar(clave, dato) {
					return
				}
				iter.Siguiente()
			}
		}
	}
}
func crearTabla[K comparable, V any](tam int) []TDALista.Lista[*parClaveValor[K, V]] {
	return make([]TDALista.Lista[*parClaveValor[K, V]], tam)
}
