package main

import (
	"container/list"
	"fmt"
	"math"
)
//Globales
var suma int
var listaNodos []Nodo

//Structs
//Nodo

type Nodo struct {
	clave             int
	cantidadRepetidas int
	comparacionesObt  int
	comparacionesInser  int
	hijoIzquierdo     *Nodo
	hijoDerecho       *Nodo
	padre             *Nodo
}

func (nod *Nodo) sumarRepetidas() {
	nod.cantidadRepetidas += 1
}

func (nod *Nodo) comparacionesCantidad(comp int) {
	nod.comparacionesObt = comp
}

func (nod *Nodo) sumarCompInser(comparaciones int){
	nod.comparacionesInser = comparaciones
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

func (arbol *arbolBinarioBusqueda) alturaMax() int{
	return arbol.tamanno -1
}

func (arbol *arbolBinarioBusqueda) alturaPromedio() float64 { //cambiar, nada que ver
	return 2 * math.Sqrt(math.Pi*float64(arbol.tamanno))
}

func (arbol *arbolBinarioBusqueda) densidad() float64{
	div := float64(arbol.longitudArbol())/float64(alturaSubArbol(arbol.raiz))
	return div
}

func (arbol *arbolBinarioBusqueda)totalComparaciones() int{
	sumaComparaciones := 0
	for i, s:= range cantidadTotalComparaciones(arbol.raiz, 0){
		i = i
		sumaComparaciones = sumaComparaciones + s.comparacionesInser
	}

	return sumaComparaciones
}

func (arbol *arbolBinarioBusqueda) cantidadPromedioComparaciones() float64{
	return float64(recorridoAmplitud(arbol.raiz)) / float64(arbol.longitudArbol())
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
	nod.comparacionesInser = nod.comparacionesObt
	return nod.comparacionesInser
}


func (arbol *arbolBinarioBusqueda) obtenerP(llave int, nodoActual *Nodo, comparaciones int) *Nodo {
	if nodoActual == nil {
		nod := Nodo{clave: -1,
			comparacionesObt: comparaciones +1}
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

func (arbol *arbolBinarioBusqueda) buscarLLave(llave int) *list.List {
	nod := arbol.obtener(llave)
	l := list.New()

	if nod.clave != -1{
		l.PushFront(true)
	} else {
		l.PushFront(false)
	}
	l.PushBack(nod.comparacionesObt)
	return l
}

func cantidadTotalComparaciones(nodoAct *Nodo, compar int) []Nodo{
	if nodoAct != nil {
		cantidadTotalComparaciones(nodoAct.hijoIzquierdo, compar)
		listaNodos = append(listaNodos, *nodoAct)
		cantidadTotalComparaciones(nodoAct.hijoDerecho, compar)
	}
	return listaNodos
}


//Recorridos

func inOrden(nodoAct *Nodo) {
	if nodoAct != nil {
		inOrden(nodoAct.hijoIzquierdo)
		fmt.Println(nodoAct.clave, " ", nodoAct.cantidadRepetidas, " ", nodoAct.comparacionesInser)
		inOrden(nodoAct.hijoDerecho)
	}
}

func postOrden(nodoAct *Nodo){
	if nodoAct != nil{
		postOrden(nodoAct.hijoIzquierdo)
		postOrden(nodoAct.hijoDerecho)
		fmt.Println(nodoAct.clave, " ", nodoAct.cantidadRepetidas)
	}
}

func recorridoAmplitud(raiz *Nodo) int{
	alt := alturaSubArbol(raiz)

	//sumaComparacionesNivel := 0

	for i:=1; i<=alt; i++ {
		verNivel(raiz, i, i)
	}
	return suma
}

func verNivel(raiz *Nodo, nivel int, nivelAct int) {

	if raiz == nil {
		return
	}else{

		if nivel == 1 {
			mult := nivelAct * raiz.comparacionesInser
			suma = suma + mult

		} else if nivel > 1 {
			verNivel(raiz.hijoIzquierdo, nivel-1, nivelAct)
			verNivel(raiz.hijoDerecho, nivel-1, nivelAct)
		}
	}

}

func alturaSubArbol(nod *Nodo) int{ //con solo esta me devuelve la altura del árbol
	if nod == nil{
		return 0
	} else {
		alturaIzqu := alturaSubArbol(nod.hijoIzquierdo)
		alturaDer := alturaSubArbol(nod.hijoDerecho)

		if alturaIzqu > alturaDer{
			return alturaIzqu + 1
		} else{
			return alturaDer + 1
		}
	}
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

func main() {
	//números
	//var lista200 []int
	//var listaBusqueda []int

	//lista200 = generarNumeros(13, 200)
	//listaBusqueda = generarNumeros(53, 10000)


	arbBinarioBus := arbolBinarioBusqueda{
		raiz:    nil,
		tamanno: 0,
	}

	//fmt.Println("Agregar nodos al árbol\n")
	//
	//for i, n:= range lista200{
	//	i = i
	//	fmt.Println(arbBinarioBus.agregarNodoArbol(n))
	//}
	//
	//fmt.Println("\n\nBuscar nodos: \n")
	//
	//for k, m := range listaBusqueda{
	//	k = k
	//	nodo := arbBinarioBus.buscarLLave(m)
	//	fmt.Println(nodo.Front().Value, " ", nodo.Back().Value)
	//}

	fmt.Println("Agregar nodos al árbol\n")

	fmt.Println(10, " ", arbBinarioBus.agregarNodoArbol(10))
	fmt.Println(5, arbBinarioBus.agregarNodoArbol(5))
	fmt.Println(45, arbBinarioBus.agregarNodoArbol(45))
	fmt.Println(95, arbBinarioBus.agregarNodoArbol(95))
	fmt.Println(2, arbBinarioBus.agregarNodoArbol(2))
	fmt.Println(2, arbBinarioBus.agregarNodoArbol(2))
	fmt.Println(7, " ", arbBinarioBus.agregarNodoArbol(7))
	fmt.Println(32, " ",arbBinarioBus.agregarNodoArbol(32))
	fmt.Println(332, " ", arbBinarioBus.agregarNodoArbol(332))
	fmt.Println(95, " ", arbBinarioBus.agregarNodoArbol(95))

	fmt.Println("\n\nBuscar nodos\n")

	var lis []int
	lis = append(lis, 5)
	lis = append(lis, 10)
	lis = append(lis, 15)
	lis = append(lis, 3)

	for i, s:= range lis{
		i=i
		l := arbBinarioBus.buscarLLave(s)
		fmt.Println(l.Front().Value, " ", l.Back().Value)
	}

	fmt.Println("\nAltura máxima: ", arbBinarioBus.alturaMax())

	fmt.Println("\nDensidad árbol: ", arbBinarioBus.densidad())

	fmt.Println("\nCantidad total de comparaciones: ", arbBinarioBus.totalComparaciones())














}
