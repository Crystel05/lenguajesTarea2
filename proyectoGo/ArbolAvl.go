package main

import (
//	"container/list"
	"fmt"
)

var comparacionesPromedio int
var sumatoriaNiveles int
var totalInsercionesBin int
var totalBusquBin int

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
	totalInsercionesBin = 0

	for i := 0; i < len(listaNums); i++ {
		comps := arbol.Insertar(listaNums[i])
		//fmt.Println(comps)
		totalComparaciones = totalComparaciones + comps
	}
	totalInsercionesBin = totalComparaciones
	return totalComparaciones

}


func (arbol *arbolAvl) buscarValorAvl(lista []int) int {
	totalComp := 0
	totalBusquBin = 0
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
	totalBusquBin = totalComp
	return totalComp
}


func (arbol *arbolAvl) alturaPromedio() float64 {
	recorridoAmplitudAvl(arbol.raiz)
	div := float64(sumatoriaNiveles) / float64(arbol.longitudArbol())
	return div
}


func (arbol *arbolAvl) totalComparaciones() int {
	return totalBusquBin + totalInsercionesBin
}


func (arbol *arbolAvl) cantidadPromedioComparaciones() float64 {
	comparacionesPromedio = 0
	recorridoAmplitudAvl(arbol.raiz)
	return float64(comparacionesPromedio) / float64(arbol.longitudArbol())
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
			comparacionesPromedio = comparacionesPromedio + mult
			sumatoriaNiveles = sumatoriaNiveles + nivelAct
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
func  (nodoActual *nodoAvl) InsertarAux(pValor int,nod *arbolAvl) {
	nodoActual.Hh=false

	if nodoActual == nil {
		nodoActual = &nodoAvl{pValor,0,0,nil,nil,nil,0,false}
		return
	}
	if pValor < nodoActual.valor {
		if nodoActual.nIzq == nil {
			nodoActual.nIzq = &nodoAvl{pValor,0,0,nil,nil,nil,0,false}
			nodoActual.peso=nodoActual.peso-1
		}else{
			nodoActual.nIzq.InsertarAux(pValor,nod)
			nodoActual.peso=nodoActual.peso-1

		}
		nodoActual.nIzq = nodoActual.nIzq.adjust()
		if nodoActual.nIzq.Hh{
			nodoActual.nIzq.Hh=false
			nodoActual.Hh=true
			nodoActual.peso=nodoActual.nDer.alturaArbolAux(0)-nodoActual.nIzq.alturaArbolAux(0)
		}
		return
	}else {
		if pValor > nodoActual.valor{
			if nodoActual.nDer == nil {
				nodoActual.nDer = &nodoAvl{pValor, 0, 0, nil, nil, nil, 0, false}
				nodoActual.peso = nodoActual.peso + 1
			} else {
				nodoActual.nDer.InsertarAux(pValor,nod)
				nodoActual.peso = nodoActual.peso + 1
			}
			nodoActual.nDer = nodoActual.nDer.adjust()
			if nodoActual.nDer.Hh {
				nodoActual.nDer.Hh = false
				nodoActual.Hh = true
				nodoActual.peso = nodoActual.nDer.alturaArbolAux(0) - nodoActual.nIzq.alturaArbolAux(0)
			}
		}else{
			fmt.Printf( "ponga aqui la repeticion")
			nodoActual.reparapesos(pValor,nod.raiz)
		}
	}


}


func (nod *arbolAvl) Insertar(pValor int)  {
	if nod.raiz == nil {
		nod.raiz =&nodoAvl{pValor,0,0,nil,nil,nil,0,false}
		nod.sumarTamano()
		return
	}
	nod.raiz.InsertarAux(pValor,nod)
	nod.raiz=nod.raiz.adjust()
	if nod.raiz.nDer!=nil {
		if nod.raiz.nDer.Hh {
			nod.raiz.nDer.Hh = false
			nod.raiz.Hh = true
			nod.raiz.peso = nod.raiz.nDer.alturaArbolAux(0) - nod.raiz.nIzq.alturaArbolAux(0)
		}
	}
	if nod.raiz.nIzq!=nil {
		if nod.raiz.nIzq.Hh {
			nod.raiz.nIzq.Hh = false
			nod.raiz.Hh = true
			nod.raiz.peso = nod.raiz.nDer.alturaArbolAux(0) - nod.raiz.nIzq.alturaArbolAux(0)
		}
	}
	nod.sumarTamano()
	//fmt.Println(nod.alturaArbol())
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


func (nod * nodoAvl ) RotIzq () * nodoAvl  {// rotar a la izquierda
	headNode := nod.nDer
	nod.nDer = headNode.nIzq
	headNode.nIzq = nod
	if headNode.nDer!= nil {
		headNode.nDer.peso = headNode.nDer.nDer.alturaArbolAux(0) - headNode.nDer.nIzq.alturaArbolAux(0)
	}
	if headNode.nIzq!= nil{
		headNode.nIzq.peso=headNode.nIzq.nDer.alturaArbolAux(0)-headNode.nIzq.nIzq.alturaArbolAux(0)
	}
	headNode.peso=headNode.nDer.alturaArbolAux(0)-headNode.nIzq.alturaArbolAux(0)
	return headNode
}


func (nod * nodoAvl ) RotDer () * nodoAvl  {// rotar a la derecha
	headNode := nod.nIzq
	nod.nIzq = headNode.nDer
	headNode.nDer = nod
	if headNode.nIzq!= nil{
		headNode.nIzq.peso=headNode.nIzq.nIzq.alturaArbolAux(0)-headNode.nIzq.nDer.alturaArbolAux(0)
	}
	if headNode.nDer!= nil{
		headNode.nDer.peso=headNode.nDer.nIzq.alturaArbolAux(0)-headNode.nDer.nDer.alturaArbolAux(0)
	}

	return headNode
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
func (nod *nodoAvl ) adjust() *nodoAvl  {
	if nod.peso== 2 {
		if nod.nDer.peso > 0{
			nod = nod.RotIzq()
		}else {
			nod = nod.rightThenLeftRotate()
		}
		nod.Hh=true
	}else if nod.peso== -2 {
		if nod.nIzq.peso<0 {
			nod = nod.RotDer()
		} else {
			nod = nod.LeftThenRightRotate()
		}
		nod.Hh=true
	}
	return nod
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

func (nod * nodoAvl ) rightThenLeftRotate () * nodoAvl  {

	sonHeadNode := nod.nDer.RotDer()
	nod.nDer = sonHeadNode
	return nod.RotIzq()
}

func (nod * nodoAvl ) LeftThenRightRotate () * nodoAvl  {

	sonHeadNode := nod.nIzq.RotIzq()
	nod.nIzq= sonHeadNode
	return nod.RotDer()
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
		fmt.Println(nod.valor,nod.peso)
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
func main() {
	arbol := arbolAvl{
		raiz: nil,
		tamanio: 0,
	}
	fmt.Println("insertar 10")
	arbol.Insertar(10)
	arbol.raiz.Orden()
	fmt.Println("insertar 60")
	arbol.Insertar(60)
	arbol.raiz.Orden()
	fmt.Println("insertar 70")
	arbol.Insertar(70)
	arbol.raiz.Orden()
	fmt.Println("insertar 20")
	arbol.Insertar(20)
	arbol.raiz.Orden()
	fmt.Println("insertar 130")
	arbol.Insertar(130)
	arbol.raiz.Orden()
	fmt.Println("insertar 230")
	arbol.Insertar(230)
	arbol.raiz.Orden()
	fmt.Println("insertar 330")
	arbol.Insertar(330)
	arbol.raiz.Orden()
	fmt.Println("insertar 430")
	arbol.Insertar(430)
	arbol.raiz.Orden()
	//fmt.Println("insertar 130")
	//arbol.Insertar(130)
	//arbol.raiz.Orden()
	fmt.Println("insertar 320")
	arbol.Insertar(320)
	arbol.raiz.Orden()
	fmt.Println("insertar 120")
	arbol.Insertar(120)
	arbol.raiz.Orden()
	fmt.Println("insertar 220")
	arbol.Insertar(220)
	arbol.raiz.Orden()
	fmt.Println("insertar 15")
	arbol.Insertar(15)
	arbol.raiz.Orden()
	fmt.Println("insertar 150")
	arbol.Insertar(150)
	arbol.raiz.Orden()
	//fmt.Println("insertar 510")
	//arbol.Insertar(10)
	//arbol.raiz.Orden()
	fmt.Println("insertar 1021")





	arbol.Insertar(21)
	arbol.raiz.Orden()






	fmt.Println(arbol.raiz.valor,"raiz",arbol.raiz.nIzq.valor,arbol.raiz.nDer.valor,arbol.raiz.nDer.peso)
	arbol.raiz.postOrden()
	fmt.Println(arbol.raiz.nIzq.peso)
	//arbol.agregarNodoArbol(10)
	//arbol.agregarNodoArbol(20)
	//arbol.agregarNodoArbol(5)
	//arbol.agregarNodoArbol(12)


	//inOrden(arbol.raiz)
	//fmt.Println("\n")
	//lis := list.New()
	//lis = arbol.buscarLLave(5)
	//fmt.Println("Se encotró: ", lis.Front().Value)
	//fmt.Println("Cant de comparacionesObt: ", lis.Back().Value)
}
