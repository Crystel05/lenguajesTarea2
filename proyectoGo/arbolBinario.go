package main

import "fmt"

type Nodo struct {

	clave int
	cantidadRepetidas int
	comparaciones int
	hijoIzquierdo *Nodo
	hijoDerecho *Nodo
	padre *Nodo

}

func (nod *Nodo)sumarRepetidas(){
	nod.cantidadRepetidas += 1
}

func (nod *Nodo) comparacionesCantidad(comp int){
	nod.comparaciones = comp
}

//***********************************************************************

type arbolBinarioBusqueda struct {
	raiz *Nodo
	tamanno int
}

func (arbol *arbolBinarioBusqueda) agregarTamanno() {
	arbol.tamanno += 1
}

func (arbol *arbolBinarioBusqueda) longitudArbol () int {
	return arbol.tamanno
}

func (arbol *arbolBinarioBusqueda) insertarP( llave int, nodoActual *Nodo) {
	if llave < nodoActual.clave{
		if nodoActual.hijoIzquierdo != nil {
			arbol.insertarP(llave, nodoActual.hijoIzquierdo)
		} else {
			nodoActual.hijoIzquierdo = &Nodo{clave: llave, hijoIzquierdo: nil, hijoDerecho: nil, padre: nodoActual}
			nodoActual.hijoIzquierdo.sumarRepetidas()
			arbol.agregarTamanno()
		}

	} else if llave > nodoActual.clave{
		if nodoActual.hijoDerecho != nil{
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

func (arbol *arbolBinarioBusqueda) agregarNodoArbol(llave int) {
	arbol.insertar(llave)
}

func (arbol *arbolBinarioBusqueda) obtenerP(llave int, nodoActual Nodo, comparaciones int) *Nodo {
	if &nodoActual == nil{
		return nil
	}else if nodoActual.clave == llave{
		comparaciones = comparaciones + 1
		nodoActual.comparacionesCantidad(comparaciones)
		return &nodoActual
	} else if llave < nodoActual.clave{
		return arbol.obtenerP(llave, *nodoActual.hijoIzquierdo, comparaciones+1)
	} else {
		return arbol.obtenerP(llave, *nodoActual.hijoDerecho, comparaciones+1)
	}
}

func (arbol *arbolBinarioBusqueda) obtener(llave int) *Nodo {
	if &arbol.raiz != nil{
		res := arbol.obtenerP(llave, *arbol.raiz, 0)
		if res != nil{
			return res
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (arbol *arbolBinarioBusqueda) getItem(llave int) {
	nod := arbol.obtener(llave)
	if nod != nil{
		fmt.Println("Se encontró realizando ", nod.comparaciones, " comparaciones")
	} else {
		fmt.Println("No se encontró, se realizaron: ", arbol.longitudArbol(), " comparaciones")
	}
}

func inOrden(nodoAct *Nodo){
	if nodoAct != nil{
		inOrden(nodoAct.hijoIzquierdo)
		fmt.Println(nodoAct.clave, " ", nodoAct.cantidadRepetidas)
		inOrden(nodoAct.hijoDerecho)
	}
}

func main() {
	var nod *Nodo
	nod = nil
	arbol := arbolBinarioBusqueda{
		raiz: nod,
		tamanno: 0,
	}

	arbol.agregarNodoArbol(10)
	arbol.agregarNodoArbol(10)
	arbol.agregarNodoArbol(20)
	arbol.agregarNodoArbol(5)
	arbol.agregarNodoArbol(12)

	fmt.Println("\n")

	inOrden(arbol.raiz)

	fmt.Println("\n")

	arbol.getItem(5)
}