package main

import (
//	"container/list"
	"fmt"
)

var comparacionesPromedioavl int
var sumatoriaNivelesavl int
var totalInsercionesAVL int
var totalBusquAVL int

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
		//if lis[0] == 0{
		//	fmt.Println("No encontrada ", lis[1])
		//} else {
		//	fmt.Println("Encontrada ", lis[1])
		//}

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
	//if nod == nil {
	//	return 0
	//} else {
	//	return maximo(alturaArbolAvl(nod.nIzq),alturaArbolAvl(nod.nDer))+1
	//
	//}

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
		//fmt.Println(nod.valor,nod.peso,"uno")
		nod.peso=nod.nDer.alturaArbolAux(0)-nod.nIzq.alturaArbolAux(0)
		//fmt.Println(nod.valor,nod.peso,"dos")

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
	return 0 /////////////////////////////////////////////////////////////////////////////////////////
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

func main() {
	arbol := arbolAvl{
		raiz: nil,
		tamanio: 0,
	}
	var numeros []int
	numeros=generarNumerosAvl(13,200)

	fmt.Println("insertasr ",arbol.llenarArbolAvl(numeros))



	//
	fmt.Println("altura ",alturaArbolAvl(arbol.raiz))
	//
	fmt.Println("altura promedio ",arbol.alturaPromedio())
	//
	fmt.Println("comparaciones  ",arbol.totalComparaciones())
	//
	fmt.Println("comparaciones promedio ",arbol.cantidadPromedioComparaciones())

	fmt.Println("densidad ",arbol.densidad())
}
