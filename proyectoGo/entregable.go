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
type nodoAvl struct {
	valor	int
	cantidadRepetidas  int
	CantComparaciones	int
	nIzq	*nodoAvl
	nDer	*nodoAvl
	padre	*nodoAvl
	peso int
	Hh bool

}

type arbolAvl struct {
	raiz    *nodoAvl
	tamanio int

}

func (arbol *arbolAvl) llenarArbolAvl(listaNums []int) int {
	totalComparaciones := 0
	totalInsercionesAVL = 0

	for i := 0; i < len(listaNums); i++ {
		comps := i
		arbol.Insertar(listaNums[i])
		//fmt.Println(comps)
		totalComparaciones = totalComparaciones + comps
	}
	totalInsercionesAVL = totalComparaciones
	return totalComparaciones
}

func (arbol *arbolAvl) buscarValorAvl(lista []int) int {
	totalComp := 0
	totalBusquAVL = 0
	for i := 0; i < len(lista); i++ {
		var lis [2]int
		lis = arbol.Buscar(lista[i])
		totalComp = totalComp + lis[1]
	}
	totalBusquAVL = totalComp
	return totalComp
}


func (arbol *arbolAvl) alturaPromedio() float64 {
	recorridoAmplitudAvl(arbol.raiz)
	div := float64(sumatoriaNivelesavl) / float64(arbol.longitudArbol())
	return div
}

func (arbol *arbolAvl) totalComparaciones() int {
	return totalBusquAVL + totalInsercionesAVL
}

func (arbol *arbolAvl) cantidadPromedioComparaciones() float64 {
	comparacionesPromedioavl = 0
	recorridoAmplitudAvl(arbol.raiz)
	return float64(comparacionesPromedioavl) / float64(arbol.longitudArbol())
}

func recorridoAmplitudAvl(raiz *nodoAvl) {
	alt := alturaArbolAvl(raiz)

	for i := 1; i <= alt; i++ {
		verNivelAvl(raiz, i, i, 0)
	}
}

func verNivelAvl(raiz *nodoAvl, nivel int, nivelAct int, cantNodos int) {
	if raiz == nil {
		return
	} else {
		if nivel == 1 {
			mult := nivelAct * raiz.cantidadRepetidas
			comparacionesPromedioavl = comparacionesPromedioavl + mult
			sumatoriaNivelesavl = sumatoriaNivelesavl + nivelAct
			cantNodos += 1
		} else if nivel > 1 {
			verNivelAvl(raiz.nIzq, nivel-1, nivelAct, cantNodos)
			verNivelAvl(raiz.nDer, nivel-1, nivelAct, cantNodos)
		}
	}
}

func (arbol *arbolAvl) densidad() float64 {
	div := float64(arbol.longitudArbol()) / float64(alturaArbolAvl(arbol.raiz))
	return div
}

func alturaArbolAvl(nod *nodoAvl) int { //con solo esta me devuelve la altura del árbol

	if nod == nil {
		return 0
	} else {
		alturaIzqu := alturaArbolAvl(nod.nIzq)
		alturaDer := alturaArbolAvl(nod.nDer)

		if alturaIzqu > alturaDer {
			return alturaIzqu + 1
		} else {
			return alturaDer + 1
		}
	}
}

func (arbol *arbolAvl) agregarTamanio() {
	arbol.tamanio += 1
}

func (arbol *arbolAvl) longitudArbol() int {
	return arbol.tamanio
}

func  (nodoActual *nodoAvl) reparapesos(pValor int,nod *nodoAvl) {
	if pValor < nod.valor {
		nod.peso=nod.peso+1
		nodoActual.reparapesos(pValor,nod.nIzq)
	}else{
		if pValor > nod.valor {
			nod.peso=nod.peso-1
			nodoActual.reparapesos(pValor,nod.nDer)
		}

	}
}

func  (nodoActual *nodoAvl) InsertarAux(pValor int,nod *arbolAvl) int{
	comp:=0
	nodoActual.Hh=false
	comp++
	if nodoActual == nil {
		nodoActual = &nodoAvl{pValor,0,0,nil,nil,nil,0,false}
		return comp
	}
	comp++
	if pValor < nodoActual.valor {
		comp++
		if nodoActual.nIzq == nil {
			nodoActual.nIzq = &nodoAvl{pValor,0,0,nil,nil,nil,0,false}
			nodoActual.peso=nodoActual.nDer.alturaArbolAux(0) - nodoActual.nIzq.alturaArbolAux(0)
		}else{
			comp=comp+nodoActual.nIzq.InsertarAux(pValor,nod)
			nodoActual.peso=nodoActual.nDer.alturaArbolAux(0) - nodoActual.nIzq.alturaArbolAux(0)

		}
		var l [2]*nodoAvl
		l= nodoActual.nIzq.adjust()
		nodoActual.nIzq = l[0]
		comp++
		if nodoActual.nIzq.Hh{
			nodoActual.nIzq.Hh=false
			nodoActual.Hh=true
			nodoActual.peso=nodoActual.nDer.alturaArbolAux(0)-nodoActual.nIzq.alturaArbolAux(0)
		}
		return comp
	}else {
		comp++
		if pValor > nodoActual.valor{
			comp++
			if nodoActual.nDer == nil {
				nodoActual.nDer = &nodoAvl{pValor, 0, 0, nil, nil, nil, 0, false}
				nodoActual.peso = nodoActual.nDer.alturaArbolAux(0) - nodoActual.nIzq.alturaArbolAux(0)
			} else {
				comp=comp+nodoActual.nDer.InsertarAux(pValor,nod)
				nodoActual.peso = nodoActual.nDer.alturaArbolAux(0) - nodoActual.nIzq.alturaArbolAux(0)
			}
			var l [2]*nodoAvl
			l= nodoActual.nDer.adjust()
			nodoActual.nDer = l[0]
			if nodoActual.nDer.Hh {
				nodoActual.nDer.Hh = false
				nodoActual.Hh = true
				nodoActual.peso = nodoActual.nDer.alturaArbolAux(0) - nodoActual.nIzq.alturaArbolAux(0)
			}
		}else{
			//fmt.Printf( "ponga aqui la repeticion")
			nodoActual.cantidadRepetidas=nodoActual.cantidadRepetidas+1
			nodoActual.reparapesos(pValor,nod.raiz)
		}
	}

	return comp
}

func (nod *arbolAvl) getNodo(pValor int)  *nodoAvl{
	return nod.raiz.getNodoAux(pValor)
}

func (nod *nodoAvl) getNodoAux(pValor int)  *nodoAvl{
	if(nod.valor>pValor){
		return nod.nIzq.getNodoAux(pValor)
	}else {
		if(nod.valor<pValor){
			return nod.nDer.getNodoAux(pValor)
		}
	}
	return nod
}

func (nod *arbolAvl) Insertar(pValor int)  int{
	comp:=0
	comp++
	if nod.raiz == nil {
		nod.raiz =&nodoAvl{pValor,0,0,nil,nil,nil,0,false}
		nod.sumarTamano()
		return comp
	}
	comp=comp+nod.raiz.InsertarAux(pValor,nod)
	var l [2]*nodoAvl
	l=nod.raiz.adjust()
	nod.raiz=l[0]
	comp=comp+l[1].valor
	comp++
	if nod.raiz.nDer!=nil {
		comp++
		if nod.raiz.nDer.Hh {
			nod.raiz.nDer.Hh = false
			nod.raiz.Hh = true
			nod.raiz.peso = nod.raiz.nDer.alturaArbolAux(0) - nod.raiz.nIzq.alturaArbolAux(0)
		}
	}
	comp++
	if nod.raiz.nIzq!=nil {
		comp++
		if nod.raiz.nIzq.Hh {
			nod.raiz.nIzq.Hh = false
			nod.raiz.Hh = true
			nod.raiz.peso = nod.raiz.nDer.alturaArbolAux(0) - nod.raiz.nIzq.alturaArbolAux(0)
		}
	}
	nod.sumarTamano()
	//fmt.Println(nod.alturaArbol())

	nod.getNodo(pValor).CantComparaciones=comp
	return comp
}

func (nod *arbolAvl) Busqueda(pNodo int) bool {
	return nod.BusquedaAux(pNodo,nod.raiz)
}

func (nod *arbolAvl) BusquedaAux(pNodo int,nodoActual *nodoAvl) bool {
	if nodoActual == nil {
		return false
	}
	compare := pNodo - nodoActual.valor
	if compare < 0 {
		return nod.BusquedaAux(pNodo,nodoActual.nIzq)
	}else if compare > 0 {
		return nod.BusquedaAux(pNodo,nodoActual.nDer)
	}else {
		return true
	}
}

func (nod * nodoAvl ) RotIzq () [2]*nodoAvl   {// rotar a la izquierda
	var l [2]*nodoAvl
	comp:=0
	headNode := nod.nDer
	nod.nDer = headNode.nIzq
	headNode.nIzq = nod
	comp++
	if headNode.nDer!= nil {
		headNode.nDer.peso = headNode.nDer.nDer.alturaArbolAux(0) - headNode.nDer.nIzq.alturaArbolAux(0)
	}
	comp++
	if headNode.nIzq!= nil{
		headNode.nIzq.peso=headNode.nIzq.nDer.alturaArbolAux(0)-headNode.nIzq.nIzq.alturaArbolAux(0)
	}
	headNode.peso=headNode.nDer.alturaArbolAux(0)-headNode.nIzq.alturaArbolAux(0)
	l[0]=headNode
	nodoAux:=new(nodoAvl)
	nodoAux.valor=comp
	l[1]=nodoAux
	return l
}

func (nod * nodoAvl ) RotDer () [2]*nodoAvl {// rotar a la derecha
	var l [2]*nodoAvl
	comp:=0
	headNode := nod.nIzq
	nod.nIzq = headNode.nDer
	headNode.nDer = nod
	comp++
	if headNode.nIzq!= nil{
		headNode.nIzq.peso=headNode.nIzq.nIzq.alturaArbolAux(0)-headNode.nIzq.nDer.alturaArbolAux(0)
	}
	comp++
	if headNode.nDer!= nil{
		headNode.nDer.peso=headNode.nDer.nIzq.alturaArbolAux(0)-headNode.nDer.nDer.alturaArbolAux(0)
	}
	l[0]=headNode
	nodoAux:=new(nodoAvl)
	nodoAux.valor=comp
	l[1]=nodoAux
	return l
}

func (nod *nodoAvl ) adjusttodo() {
	if nod == nil {
		return
	}
	nod.nIzq.adjusttodo()
	nod.nDer.adjusttodo()
	if nod!= nil{
		nod.peso=nod.nDer.alturaArbolAux(0)-nod.nIzq.alturaArbolAux(0)
	}
}

func (nod *nodoAvl ) adjust() [2]*nodoAvl  {
	var l [2]*nodoAvl
	comp:=0
	comp++
	if nod.peso== 2 {
		comp++
		if nod.nDer.peso > 0{
			l=nod.RotIzq()
			comp=comp+l[1].valor
			nod=l[0]
		}else {
			l = nod.rightThenLeftRotate()
			comp=comp+l[1].valor
			nod=l[0]
		}
		nod.Hh=true
	}else if nod.peso== -2 {
		comp++
		if nod.nIzq.peso<0 {
			l= nod.RotDer()
			comp=comp+l[1].valor
			nod=l[0]
		} else {
			l= nod.LeftThenRightRotate()
			comp=comp+l[1].valor
			nod=l[0]
		}
		nod.Hh=true
	}
	comp++
	l[0]=nod
	nodoAux:=new(nodoAvl)
	nodoAux.valor=comp
	l[1]=nodoAux
	return l
}

func (arbol *arbolAvl) sumarTamano() {
	arbol.tamanio += 1
}

func (nod *nodoAvl ) getAll() []int {
	valores  := []int{}
	return addValues(valores,nod)
}

func addValues(valores []int,nod *nodoAvl ) []int {
	if nod != nil {
		valores = addValues(valores,nod.nIzq)
		valores = append(valores,nod.valor)

		valores = addValues(valores,nod.nDer)
	}
	return valores
}

func (nod *nodoAvl ) getAltura() int {
	if nod == nil {
		return 0
	}
	return 0
}

func maximo(a int,b int) int {
	if a > b {
		return a
	} else {
		return b

	}
}

func (nod * nodoAvl ) rightThenLeftRotate () [2]*nodoAvl  {//
	comp:=0
	sonHeadNode := nod.nDer.RotDer()
	comp=sonHeadNode[1].valor
	nod.nDer = sonHeadNode[0]
	sonHeadNode=nod.RotIzq()
	sonHeadNode[1].valor=comp+sonHeadNode[1].valor
	sonHeadNode[0].peso=sonHeadNode[0].nDer.alturaArbolAux(0)-sonHeadNode[0].nIzq.alturaArbolAux(0)
	return sonHeadNode
}

func (nod * nodoAvl ) LeftThenRightRotate () [2]*nodoAvl  {//[2]*nodoAvl
	comp:=0
	sonHeadNode := nod.nIzq.RotIzq()
	comp=sonHeadNode[1].valor
	nod.nIzq= sonHeadNode[0]
	sonHeadNode=nod.RotDer()
	sonHeadNode[1].valor=comp+sonHeadNode[1].valor
	sonHeadNode[0].peso=sonHeadNode[0].nDer.alturaArbolAux(0)-sonHeadNode[0].nIzq.alturaArbolAux(0)
	return sonHeadNode
}

func (nod * nodoAvl )postOrden(){
	if nod != nil{
		nod.nIzq.postOrden()
		fmt.Println(nod.valor)
		nod.nDer.postOrden()
	}
}
func (nod * nodoAvl )Orden(){
	if nod != nil{
		fmt.Println(nod.valor,nod.peso,nod.cantidadRepetidas)
		nod.nIzq.Orden()

		nod.nDer.Orden()
	}
}
func (nod *nodoAvl) alturaArbolAux( max int) int{

	if nod != nil{
		return 1+maximo(nod.nDer.alturaArbolAux(max),nod.nIzq.alturaArbolAux(max))
	}
	return 0

}

func (arbol *arbolAvl) alturaArbol() int {
	return arbol.raiz.alturaArbolAux(0)

}

func (nod *nodoAvl) BuscarAux(valor int,cont int) [2]int {
	var l [2]int
	if(nod==nil){
		l[0]=0
		l[1]=cont+1
		return l
	}
	if(nod.valor<valor){
		nod.nDer.BuscarAux(valor,cont+1)
	}else{
		if(nod.valor>valor){
			nod.nIzq.BuscarAux(valor,cont+1)
		}else{
			l[0]=1
			l[1]=cont+1

		}
	}
	return l
}
func (arbol *arbolAvl) Buscar(valor int) [2]int {
	return arbol.raiz.BuscarAux(valor,0)

}

func generarNumerosAvl(semilla int, cantidad int) []int {

	var numeros []int

	i := 0

	for s := range generarNumerosPseudoAleatoriosAvl(semilla, cantidad) {
		s = s % 200
		numeros = append(numeros, s)
		i++
	}

	return numeros
}

func generarNumerosPseudoAleatoriosAvl(x int, n int) <-chan int {

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

func (nod * nodoAvl)contarIzqAvl()  int{
	if(nod==nil){
		return  0
	}
	return 1+nod.nIzq.contarIzqAvl()
}

//árbol Rojo Negro

const (
	red   byte = byte(0)
	black byte = byte(1)
)

type RBNode struct {
	value               int
	cont                int
	color               byte
	nivel               int
	comparaciones       int
	left, right, parent *RBNode
}

type RBTree struct {
	root                       *RBNode
	tamanio                    int
	comparaciones              int
	encontrado                 bool
	cantComparaciones          int
	totalComparaciones         int
	totalComparacionesBusqueda int
	totalNodos                 int
}

func (arbolRN *RBTree) printTree(root *RBNode) {

	fmt.Println(fmt.Sprint("Valor ", root.value))
	if root.left != nil {
		arbolRN.printTree(root.left)
	}

	if root.right != nil {
		arbolRN.printTree(root.right)
	}
}

func (arbolRN *RBTree) insert(root *RBNode, newNode *RBNode) *RBNode {
	arbolRN.comparaciones = arbolRN.comparaciones + 1
	if root == nil {
		arbolRN.totalNodos += 1
		newNode.comparaciones = arbolRN.comparaciones
		return newNode
	}
	if newNode.value < root.value {
		root.left = arbolRN.insert(root.left, newNode)
		root.left.parent = root
	} else if newNode.value > root.value {
		root.right = arbolRN.insert(root.right, newNode)
		root.right.parent = root
	} else if newNode.value == root.value {
		root.cont += 1
		newNode.cont = 2
	}
	return root
}

func (arbolRN *RBTree) fixViolation(root *RBNode, pt *RBNode) {
	var parent_pt *RBNode = nil
	var grand_parent_pt *RBNode = nil

	for (pt != root) && (pt.color != black) && (pt.parent.color == red) {

		parent_pt = pt.parent
		grand_parent_pt = pt.parent.parent

		if parent_pt == grand_parent_pt.left {

			var uncle_pt *RBNode = grand_parent_pt.right

			if uncle_pt != nil && uncle_pt.color == red {
				grand_parent_pt.color = red
				parent_pt.color = black
				uncle_pt.color = black
				pt = grand_parent_pt
			} else {
				if pt == parent_pt.right {
					arbolRN.rotateLeft(root, parent_pt)
					pt = parent_pt
					parent_pt = pt.parent
				}
				arbolRN.rotateRight(root, grand_parent_pt)
				//Intercambiar colores
				var temp byte
				temp = parent_pt.color
				parent_pt.color = grand_parent_pt.color
				grand_parent_pt.color = temp

				pt = parent_pt
			}
		} else {

			var uncle_pt *RBNode = grand_parent_pt.left

			if (uncle_pt != nil) && (uncle_pt.color == red) {
				grand_parent_pt.color = red
				parent_pt.color = black
				uncle_pt.color = black
				pt = grand_parent_pt
			} else {
				if pt == parent_pt.left {
					arbolRN.rotateRight(root, parent_pt)
					pt = parent_pt
					parent_pt = pt.parent
				}

				arbolRN.rotateLeft(root, grand_parent_pt)

				//Intercambiar colores
				var temp byte
				temp = parent_pt.color
				parent_pt.color = grand_parent_pt.color
				grand_parent_pt.color = temp

				pt = parent_pt
			}
		}
	}

	root.color = black
}

func (arbolRN *RBTree) addValue(value int) int {

	Node := RBNode{
		value:  value,
		cont:   1,
		color:  red,
		left:   nil,
		right:  nil,
		parent: nil,
	}

	arbolRN.comparaciones = 0
	// Do a normal BST insert
	arbolRN.root = arbolRN.insert(arbolRN.root, &Node)

	// fix red black Tree violations
	if Node.cont == 1 {
		arbolRN.fixViolation(arbolRN.root, &Node)
	}


	return arbolRN.comparaciones

}

func (arbolRN *RBTree) rotateLeft(root *RBNode, pt *RBNode) {
	var pt_right *RBNode = pt.right
	pt.right = pt_right.left
	if pt.right != nil {
		pt.right.parent = pt
	}
	pt_right.parent = pt.parent
	if pt.parent == nil {
		root = pt_right
	} else if pt == pt.parent.left {
		pt.parent.left = pt_right
	} else {
		pt.parent.right = pt_right
	}
	pt_right.left = pt
	pt.parent = pt_right
}

func (arbolRN *RBTree) rotateRight(root *RBNode, pt *RBNode) {
	var pt_left *RBNode = pt.left

	pt.left = pt_left.right

	if pt.left != nil {
		pt.left.parent = pt
	}

	pt_left.parent = pt.parent

	if pt.parent == nil {
		root = pt_left
	} else if pt == pt.parent.left {
		pt.parent.left = pt_left
	} else {
		pt.parent.right = pt_left
	}
	pt_left.right = pt
	pt.parent = pt_left
}

func (arbolRN *RBTree) getAlturaMaxima(root *RBNode) int {
	a := 1
	b := 1
	if root.left == nil && root.right == nil {
		return 1
	} else {
		if root.left != nil {
			a = arbolRN.getAlturaMaxima(root.left)
		}
		if root.right != nil {
			b = arbolRN.getAlturaMaxima(root.right)
		}

		if a < b {
			return b + 1
		} else {
			return a + 1
		}
	}
}

func (arbolRN *RBTree) getAltPromedio(root *RBNode, nivel int) int {
	a := 1
	b := 1
	if root.left == nil && root.right == nil {
		return nivel
	} else {
		if root.left != nil {
			a = arbolRN.getAltPromedio(root.left, nivel+1)
		}
		if root.right != nil {
			b = arbolRN.getAltPromedio(root.right, nivel+1)
		}
		return nivel + a + b
	}
}

func (arbolRN *RBTree) getCompPromedio(root *RBNode, nivel int) int {
	a := 1
	b := 1
	if root.left == nil && root.right == nil {
		return nivel * root.cont
	} else {
		if root.left != nil {
			a = arbolRN.getCompPromedio(root.left, nivel+1)
		}
		if root.right != nil {
			b = arbolRN.getCompPromedio(root.right, nivel+1)
		}

		return (nivel * root.cont) + a + b
	}
}

func (arbolRN *RBTree) getAlturaPromedio(root *RBNode) float64 {
	return float64(arbolRN.getAltPromedio(root, 1)) / float64(arbolRN.totalNodos)
}

func (arbolRN *RBTree) getComparacionesPromedio(root *RBNode) float64 {
	return float64(arbolRN.getCompPromedio(root, 1)) / float64(arbolRN.totalNodos)
}

func (arbolRN *RBTree) getTotalComparaciones() int {
	return arbolRN.totalComparaciones + arbolRN.totalComparacionesBusqueda
}

func (arbolRN *RBTree) busqueda(root *RBNode, value int) {
	if root != nil {
		arbolRN.comparaciones = arbolRN.comparaciones + 1
		if value == root.value {
			arbolRN.encontrado = true
		} else {
			if value < root.value {
				arbolRN.busqueda(root.left, value)
			} else {
				arbolRN.busqueda(root.right, value)
			}
		}
	}
}

func (arbolRN *RBTree) busquedaRN(value int) (int, bool) {
	arbolRN.comparaciones = 0
	arbolRN.encontrado = false

	arbolRN.busqueda(arbolRN.root, value)

	return arbolRN.comparaciones, arbolRN.encontrado
}

func (arbolRN *RBTree) insercion(array []int) {
	arbolRN.cantComparaciones = 0
	arbolRN.totalComparaciones = 0
	arbolRN.totalNodos = 0

	for i := 0; i < len(array); i++ {
		//fmt.Println(array[i])
		arbolRN.totalComparaciones += arbolRN.addValue(array[i])
	}
}

func (arbolRN *RBTree) busquedas(array []int) {
	arbolRN.totalComparacionesBusqueda = 0
	for i := 0; i < len(array); i++ {
		//fmt.Println(arbolRN.busquedaRN(array[i]))
		arbolRN.busquedaRN(array[i])
		arbolRN.totalComparacionesBusqueda += arbolRN.comparaciones

	}

	println(arbolRN.totalComparacionesBusqueda)
}

func (arbolRN *RBTree) getDensidad() float64 {
	a := float64(arbolRN.totalNodos) / float64(arbolRN.getAlturaMaxima(arbolRN.root))
	return a
}


func main() {

}
