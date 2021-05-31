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

//Structs
//Nodo

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

//Árbol binario

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
	recorridoAmplitud(arbol.raiz)
	return float64(comparacionesPromedio) / float64(arbol.longitudArbol())
}

func (arbol *arbolBinarioBusqueda) insertarP(llave int, nodoActual *Nodo) {
	if llave < nodoActual.clave {
		if nodoActual.hijoIzquierdo != nil {
			arbol.insertarP(llave, nodoActual.hijoIzquierdo)
		} else {
			nodoActual.hijoIzquierdo = &Nodo{clave: llave, hijoIzquierdo: nil, hijoDerecho: nil, padre: nodoActual}
			nodoActual.hijoIzquierdo.sumarRepetidas()
			arbol.agregarTamanno()
		}

	} else if llave > nodoActual.clave {
		if nodoActual.hijoDerecho != nil {
			arbol.insertarP(llave, nodoActual.hijoDerecho)
		} else {
			nodoActual.hijoDerecho = &Nodo{clave: llave, hijoIzquierdo: nil, hijoDerecho: nil, padre: nodoActual}
			nodoActual.hijoDerecho.sumarRepetidas()
			arbol.agregarTamanno()
		}
	} else {
		nodoActual.sumarRepetidas()
	}
}

func (arbol *arbolBinarioBusqueda) insertar(llave int) {
	if arbol.raiz != nil {
		arbol.insertarP(llave, arbol.raiz)
	} else {
		arbol.raiz = &Nodo{clave: llave, hijoIzquierdo: nil, hijoDerecho: nil, padre: nil}
		arbol.raiz.sumarRepetidas()
		arbol.agregarTamanno()
	}
}

func (arbol *arbolBinarioBusqueda) agregarNodoArbol(llave int) int {
	arbol.insertar(llave)
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

	fmt.Println(fmt.Sprint("Valor ", root.clave))
	if root.hijoIzquierdo != nil {
		fmt.Println("<<Izquierda")
		arbol.preOrden(root.hijoIzquierdo)
	}
	if root.hijoDerecho != nil {
		fmt.Println("Derecha>>")
		arbol.preOrden(root.hijoDerecho)
	}
}

func inOrden(nodoAct *Nodo) {
	if nodoAct != nil {
		inOrden(nodoAct.hijoIzquierdo)
		fmt.Println(nodoAct.clave, " ", nodoAct.cantidadRepetidas)
		inOrden(nodoAct.hijoDerecho)
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
		verNivel(raiz, i, i, 0)
	}
}

func verNivel(raiz *Nodo, nivel int, nivelAct int, cantNodos int) {

	if raiz == nil {
		return
	} else {

		if nivel == 1 {

			mult := nivelAct * raiz.cantidadRepetidas
			comparacionesPromedio = comparacionesPromedio + mult
			sumatoriaNiveles = sumatoriaNiveles + nivelAct
			cantNodos += 1

		} else if nivel > 1 {
			verNivel(raiz.hijoIzquierdo, nivel-1, nivelAct, cantNodos)
			verNivel(raiz.hijoDerecho, nivel-1, nivelAct, cantNodos)
		}
	}

}

//Funciones pedidas
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

//Números aleatorios

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

func (arbol *arbolBinarioBusqueda) creardsw(nodo *Nodo) {

	if nodo == nil {
		return
	}
	arbol.creardsw(nodo.hijoIzquierdo)
	arbol.insertar(nodo.clave)
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
	//println("Aun no se cae1")
	dsw := new(arbolBinarioBusqueda)

	dsw.creardsw(arbol.raiz)
	//println("Aun no se cae2")

	nodoaux := dsw.raiz
	variable := dsw.calculosobrantes(0)

	//println("Aun no se cae3", variable)

	for i := 0; i < variable; i++ {
		nodoaux.contadorQC = true
		nodoaux = nodoaux.hijoDerecho.hijoDerecho
	}
	//println("Aun no se cae4")

	dsw.preOrden(dsw.raiz)
	dsw.raiz = dsw.raiz.Rotnodo()
	println("Recorrido dentro dsw")
	dsw.preOrden(dsw.raiz)
	println("Final del recorrido")

	for j := 1; j < 2; j++ {
		dswNodo := dsw.raiz
		for i := 0; i < dsw.raiz.contarDer()/2; i++ {
			//	println(dsw.raiz.contarDer()/2)
			dswNodo.contadorQC = true
			if dswNodo.hijoDerecho.hijoDerecho == nil {
				dswNodo.contadorQC = false
			}
			dswNodo = dswNodo.hijoDerecho.hijoDerecho

		}
		//println("finalizo4",j)
		dsw.raiz = dsw.raiz.Rotnodo()
		if arbol.calculoNiveles(0) > dsw.raiz.contarIzq() {
			j--
		}
	}

	arbol.raiz = dsw.raiz
	//println("Recorrido dentro dsw")
	//arbol.preOrden(arbol.raiz)
	//println("Final del recorrido")
}

func main() {
	//números
	//var listaInsertar []int
	prueba := []int{1, 2, 8, 5, 63, 101, 102, 150, 140, 130}
	//var listaBusqueda []int

	//listaInsertar = generarNumeros(13, 1000)
	//listaBusqueda = generarNumeros(53, 10000)

	arbol := new(arbolBinarioBusqueda)

	fmt.Println(arbol.llenarArbol(prueba))
	fmt.Println(alturaArbol(arbol.raiz))
	arbol.dsw()
	println("Recorrido")
	arbol.preOrden(arbol.raiz)
	//fmt.Println(alturaArbol(arbol.raiz))

	//fmt.Println(arbol.llenarArbol(listaInsertar))
	//
	//fmt.Println(arbol.buscarLlaves(listaBuscar))
	//
	//fmt.Println(alturaArbol(arbol.raiz))
	//
	//fmt.Println(arbol.alturaPromedio())
	//
	//fmt.Println(arbol.totalComparaciones())
	//
	//fmt.Println(arbol.cantidadPromedioComparaciones())
	//
	//fmt.Println(arbol.densidad())

	//arbol.dsw()

}
