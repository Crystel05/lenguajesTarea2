package main

import (
	"fmt"
	"math"
)

//Globales
var comparacionesPromedio int
var sumatoriaNiveles int
var totalInsercionesBin int
var totalBusquBin int
var comparacionesPromedioavl int
var sumatoriaNivelesavl int
var totalInsercionesAVL int
var totalBusquAVL int

//Números pseudoaleatorios

func generarNumerosPseudoAleatorios(x int, n int) <-chan int {

	p := 2048
	mult := 109
	i := 853
	k := 0
	sal := x
	out := make(chan int)
	go func() {
		for k < n {
			salida := (mult*sal + i) % p
			sal = salida
			k++
			out <- salida
		}
		close(out)
	}()
	return out
}

func generarNumeros(semilla int, cantidad int) []int {

	var numeros []int

	i := 0

	for s := range generarNumerosPseudoAleatorios(semilla, cantidad) {
		s = s % 200
		numeros = append(numeros, s)
		i++
	}

	return numeros
}

//árbol binario de búsqueda

type Nodo struct {
	clave             int
	cantidadRepetidas int
	comparacionesObt  int
	hijoIzquierdo     *Nodo
	hijoDerecho       *Nodo
	padre             *Nodo
	contadorQC        bool
}

func (nod *Nodo) sumarRepetidas() {
	nod.cantidadRepetidas += 1
}

func (nod *Nodo) comparacionesCantidad(comp int) {
	nod.comparacionesObt = comp
}

type arbolBinarioBusqueda struct {
	raiz    *Nodo
	tamanno int
}

func (arbol *arbolBinarioBusqueda) agregarTamanno() {
	arbol.tamanno += 1
}

func (arbol *arbolBinarioBusqueda) longitudArbol() int {
	return arbol.tamanno
}

func alturaArbol(nod *Nodo) int {
	if nod == nil {
		return 0
	} else {
		alturaIzqu := alturaArbol(nod.hijoIzquierdo)
		alturaDer := alturaArbol(nod.hijoDerecho)

		if alturaIzqu > alturaDer {
			return alturaIzqu + 1
		} else {
			return alturaDer + 1
		}
	}
}

func (arbol *arbolBinarioBusqueda) alturaPromedio() float64 {
	sumatoriaNiveles = 0
	comparacionesPromedio = 0
	recorridoAmplitud(arbol.raiz)
	div := float64(sumatoriaNiveles) / float64(arbol.longitudArbol())
	return div
}

func (arbol *arbolBinarioBusqueda) densidad() float64 {
	div := float64(arbol.longitudArbol()) / float64(alturaArbol(arbol.raiz))
	return div
}

func (arbol *arbolBinarioBusqueda) totalComparaciones() int {
	return totalBusquBin + totalInsercionesBin
}

func (arbol *arbolBinarioBusqueda) cantidadPromedioComparaciones() float64 {
	comparacionesPromedio = 0
	sumatoriaNiveles = 0
	recorridoAmplitud(arbol.raiz)
	return float64(comparacionesPromedio) / float64(arbol.longitudArbol())
}

func (arbol *arbolBinarioBusqueda) insertarP(llave int, nodoActual *Nodo, rep int) {
	if llave < nodoActual.clave {
		if nodoActual.hijoIzquierdo != nil {
			arbol.insertarP(llave, nodoActual.hijoIzquierdo, rep)
		} else {
			nodoActual.hijoIzquierdo = &Nodo{clave: llave, hijoIzquierdo: nil, hijoDerecho: nil, padre: nodoActual}
			if rep != 0{
				nodoActual.hijoIzquierdo.cantidadRepetidas = rep
			}else {
				nodoActual.hijoIzquierdo.sumarRepetidas()
			}
			arbol.agregarTamanno()
		}

	} else if llave > nodoActual.clave {
		if nodoActual.hijoDerecho != nil {
			arbol.insertarP(llave, nodoActual.hijoDerecho, rep)
		} else {
			nodoActual.hijoDerecho = &Nodo{clave: llave, hijoIzquierdo: nil, hijoDerecho: nil, padre: nodoActual}
			if rep != 0{
				nodoActual.hijoDerecho.cantidadRepetidas = rep
			}else {
				nodoActual.hijoDerecho.sumarRepetidas()
			}
			arbol.agregarTamanno()
		}
	} else {
		nodoActual.sumarRepetidas()
	}
}

func (arbol *arbolBinarioBusqueda) insertar(llave int, rep int) {
	if arbol.raiz != nil {
		arbol.insertarP(llave, arbol.raiz, rep)
	} else {
		arbol.raiz = &Nodo{clave: llave, hijoIzquierdo: nil, hijoDerecho: nil, padre: nil}
		if rep != 0{
			arbol.raiz.cantidadRepetidas = rep
		} else{
			arbol.raiz.sumarRepetidas()
		}
		arbol.agregarTamanno()

	}
}

func (arbol *arbolBinarioBusqueda) agregarNodoArbol(llave int) int {
	arbol.insertar(llave, 0)
	nod := arbol.obtener(llave)
	return nod.comparacionesObt
}

func (arbol *arbolBinarioBusqueda) obtenerP(llave int, nodoActual *Nodo, comparaciones int) *Nodo {
	if nodoActual == nil {
		nod := Nodo{clave: -1,
			comparacionesObt: comparaciones + 1}
		return &nod
	} else if nodoActual.clave == llave {
		comparaciones = comparaciones + 1
		nodoActual.comparacionesCantidad(comparaciones)
		return nodoActual
	} else if llave < nodoActual.clave {
		return arbol.obtenerP(llave, nodoActual.hijoIzquierdo, comparaciones+1)
	} else {
		return arbol.obtenerP(llave, nodoActual.hijoDerecho, comparaciones+1)
	}
}

func (arbol *arbolBinarioBusqueda) obtener(llave int) *Nodo {
	return arbol.obtenerP(llave, arbol.raiz, 0)
}

func (arbol *arbolBinarioBusqueda) buscarLLave(llave int) [2]int {
	nod := arbol.obtener(llave)
	var l [2]int

	if nod.clave != -1 {
		l[0] = 1
	} else {
		l[0] = 0
	}
	l[1] = nod.comparacionesObt
	nod.comparacionesObt = 0
	return l
}

//Recorridos

func (arbol *arbolBinarioBusqueda) preOrden(root *Nodo) {

	//fmt.Println(fmt.Sprint("Valor ", root.clave))
	if root.hijoIzquierdo != nil {
		//fmt.Println("<<Izquierda")
		arbol.preOrden(root.hijoIzquierdo)
	}
	if root.hijoDerecho != nil {
		//fmt.Println("Derecha>>")
		arbol.preOrden(root.hijoDerecho)
	}
}

func postOrden(nodoAct *Nodo) {
	if nodoAct != nil {
		postOrden(nodoAct.hijoIzquierdo)
		postOrden(nodoAct.hijoDerecho)
		fmt.Println(nodoAct.clave, " ", nodoAct.cantidadRepetidas)
	}
}

func recorridoAmplitud(raiz *Nodo) {
	alt := alturaArbol(raiz)

	for i := 1; i <= alt; i++ {
		verNivel(raiz, i, i)
	}
}

func verNivel(raiz *Nodo, nivel int, nivelAct int) {

	if raiz == nil {
		return
	} else {

		if nivel == 1 {

			mult := nivelAct * raiz.cantidadRepetidas
			comparacionesPromedio = comparacionesPromedio + mult
			sumatoriaNiveles = sumatoriaNiveles + nivelAct

		} else if nivel > 1 {
			verNivel(raiz.hijoIzquierdo, nivel-1, nivelAct)
			verNivel(raiz.hijoDerecho, nivel-1, nivelAct)
		}
	}

}

func (arbol *arbolBinarioBusqueda) llenarArbol(listaNums []int) int {
	totalComparaciones := 0
	totalInsercionesBin = 0

	for i := 0; i < len(listaNums); i++ {
		comps := arbol.agregarNodoArbol(listaNums[i])
		totalComparaciones = totalComparaciones + comps
	}
	totalInsercionesBin = totalComparaciones
	return totalComparaciones

}

func (arbol *arbolBinarioBusqueda) buscarLlaves(lista []int) int {
	totalComp := 0
	totalBusquBin = 0
	for i := 0; i < len(lista); i++ {
		var lis [2]int
		lis = arbol.buscarLLave(lista[i])
		if lis[0] == 0 {
			fmt.Println("No encontrada ", lis[1])
		} else {
			fmt.Println("Encontrada ", lis[1])
		}

		totalComp = totalComp + lis[1]
	}
	totalBusquBin = totalComp
	return totalComp
}

func (arbol *arbolBinarioBusqueda) creardsw(nodo *Nodo) {

	if nodo == nil {
		return
	}
	arbol.creardsw(nodo.hijoIzquierdo)
	arbol.insertar(nodo.clave, nodo.cantidadRepetidas)
	arbol.creardsw(nodo.hijoDerecho)
	arbol.longitudArbol()

}

func (nod *Nodo) RotIzq() *Nodo { // rotar a la izquierda
	headNode := nod.hijoDerecho
	nod.hijoDerecho = headNode.hijoIzquierdo
	headNode.hijoIzquierdo = nod
	nod = headNode
	return nod
}

func (nod *Nodo) Rotnodo() *Nodo {
	//println("Clave: ", nod.clave)
	if nod.hijoDerecho == nil {
		//println("Entra al primer if")
		return nod
	}

	if nod.contadorQC {
		//println("Entra al segundo if")
		nod.contadorQC = false
		nod = nod.RotIzq()
	}

	//println(nod.hijoIzquierdo.clave)
	if nod.hijoDerecho != nil {
		nod.hijoDerecho = nod.hijoDerecho.Rotnodo()
	}
	return nod
}

func (arbol *arbolBinarioBusqueda) calculoNiveles(tamano int) int {
	variable := int((math.Pow(float64(2), float64(tamano))))
	if variable > arbol.longitudArbol() {
		return tamano
	}
	return arbol.calculoNiveles(tamano + 1)
}

func (arbol *arbolBinarioBusqueda) calculosobrantes(tamano int) int {
	variable := int((math.Pow(float64(2), float64(tamano))))
	if variable-1 > arbol.longitudArbol() {
		return arbol.longitudArbol() - int((math.Pow(float64(2), float64(tamano-1)))-1)
	}
	//println(variable,tamano)
	return arbol.calculosobrantes(tamano + 1)
}
func (nod *Nodo) contarIzq() int {
	if nod == nil {
		return 0
	}
	return 1 + nod.hijoIzquierdo.contarIzq()
}
func (nod *Nodo) contarDer() int {
	if nod == nil {
		return 0
	}
	return 1 + nod.hijoDerecho.contarDer()
}
func (arbol *arbolBinarioBusqueda) dsw() {
	dsw := new(arbolBinarioBusqueda)

	dsw.creardsw(arbol.raiz)

	nodoaux := dsw.raiz
	variable := dsw.calculosobrantes(0)


	for i := 0; i < variable; i++ {
		nodoaux.contadorQC = true
		nodoaux = nodoaux.hijoDerecho.hijoDerecho
	}

	dsw.preOrden(dsw.raiz)
	dsw.raiz = dsw.raiz.Rotnodo()
	dsw.preOrden(dsw.raiz)

	for j := 1; j < 2; j++ {
		dswNodo := dsw.raiz
		for i := 0; i < dsw.raiz.contarDer()/2; i++ {
			dswNodo.contadorQC = true
			if dswNodo.hijoDerecho.hijoDerecho == nil {
				dswNodo.contadorQC = false
			}
			dswNodo = dswNodo.hijoDerecho.hijoDerecho

		}
		dsw.raiz = dsw.raiz.Rotnodo()
		if arbol.calculoNiveles(0) > dsw.raiz.contarIzq() {
			j--
		}
	}

	arbol.raiz = dsw.raiz
}



//árbol AVL




func main() {

}
