package main
//
//import (
//	"container/list"
//	"fmt"
//)
//
//type Nodo struct {
//	clave             int
//	cantidadRepetidas int
//	comparacionesObt     int
//	hijoIzquierdo     *Nodo
//	hijoDerecho       *Nodo
//	padre             *Nodo
//	contadorQC int
//}
//
//func (nod *Nodo) sumarRepetidas() {
//	nod.cantidadRepetidas += 1
//}
//
//func (nod *Nodo) comparacionesCantidad(comp int) {
//	nod.comparacionesObt = comp
//}
//
////***********************************************************************
//
//type arbolBinarioBusqueda struct {
//	raiz    *Nodo
//	tamanno int
//}
//
//func (arbol *arbolBinarioBusqueda) agregarTamanno() {
//	arbol.tamanno += 1
//}
//
//func (arbol *arbolBinarioBusqueda) longitudArbol() int {
//	return arbol.tamanno
//}
//
//func (arbol *arbolBinarioBusqueda) insertarP(llave int, nodoActual *Nodo) {
//	if llave < nodoActual.clave {
//		if nodoActual.hijoIzquierdo != nil {
//			arbol.insertarP(llave, nodoActual.hijoIzquierdo)
//		} else {
//			nodoActual.hijoIzquierdo = &Nodo{clave: llave, hijoIzquierdo: nil, hijoDerecho: nil, padre: nodoActual}
//			nodoActual.hijoIzquierdo.sumarRepetidas()
//			arbol.agregarTamanno()
//		}
//
//	} else if llave > nodoActual.clave {
//		if nodoActual.hijoDerecho != nil {
//			arbol.insertarP(llave, nodoActual.hijoDerecho)
//		} else {
//			nodoActual.hijoDerecho = &Nodo{clave: llave, hijoIzquierdo: nil, hijoDerecho: nil, padre: nodoActual}
//			nodoActual.hijoDerecho.sumarRepetidas()
//			arbol.agregarTamanno()
//		}
//	} else {
//		nodoActual.sumarRepetidas()
//	}
//}
//
//func (arbol *arbolBinarioBusqueda) insertar(llave int) {
//	if arbol.raiz != nil {
//		arbol.insertarP(llave, arbol.raiz)
//	} else {
//		arbol.raiz = &Nodo{clave: llave, hijoIzquierdo: nil, hijoDerecho: nil, padre: nil}
//		arbol.raiz.sumarRepetidas()
//		arbol.agregarTamanno()
//	}
//}
//
//func (arbol *arbolBinarioBusqueda) agregarNodoArbol(llave int) {
//	arbol.insertar(llave)
//}
//
//func (arbol *arbolBinarioBusqueda) obtenerP(llave int, nodoActual Nodo, comparacionesObt int) *Nodo {
//	if &nodoActual == nil {
//		return nil
//	} else if nodoActual.clave == llave {
//		comparacionesObt = comparacionesObt + 1
//		nodoActual.comparacionesCantidad(comparacionesObt)
//		return &nodoActual
//	} else if llave < nodoActual.clave {
//		return arbol.obtenerP(llave, *nodoActual.hijoIzquierdo, comparacionesObt+1)
//	} else {
//		return arbol.obtenerP(llave, *nodoActual.hijoDerecho, comparacionesObt+1)
//	}
//}
//
//func (arbol *arbolBinarioBusqueda) obtener(llave int) *Nodo {
//	if &arbol.raiz != nil {
//		res := arbol.obtenerP(llave, *arbol.raiz, 0)
//		if res != nil {
//			return res
//		} else {
//			return nil
//		}
//	} else {
//		return nil
//	}
//}
//
//func (arbol *arbolBinarioBusqueda) buscarLLave(llave int) *list.List {
//	nod := arbol.obtener(llave)
//	lista := list.New()
//	if nod != nil {
//		lista.PushFront(true)
//	} else {
//		lista.PushFront(false)
//	}
//	lista.PushBack(nod.comparacionesObt)
//	return lista
//
//	//lista.Front().Value para saber si se insertó o no, lista.Back().Value para la cantidad de comparacionesObt
//}
//
//func inOrden(nodoAct *Nodo) {
//	if nodoAct != nil {
//		inOrden(nodoAct.hijoIzquierdo)
//		fmt.Println(nodoAct.clave, " ", nodoAct.cantidadRepetidas)
//		inOrden(nodoAct.hijoDerecho)
//	}
//}
//
//func postOrden(nodoAct *Nodo){
//	if nodoAct != nil{
//		postOrden(nodoAct.hijoIzquierdo)
//		postOrden(nodoAct.hijoDerecho)
//		fmt.Println(nodoAct)
//	}
//}

//func main() {
//	var nod *Nodo
//	nod = nil
//	arbol := arbolBinarioBusqueda{
//		raiz:    nod,
//		tamanno: 0,
//	}
//
//	arbol.agregarNodoArbol(10)
//	arbol.agregarNodoArbol(10)
//	arbol.agregarNodoArbol(20)
//	arbol.agregarNodoArbol(5)
//	arbol.agregarNodoArbol(12)
//
//	fmt.Println("\n")
//
//	inOrden(arbol.raiz)
//
//	fmt.Println("\n")
//
//	lis := list.New()
//	lis = arbol.buscarLLave(5)
//
//	fmt.Println("Se encotró: ", lis.Front().Value)
//	fmt.Println("Cant de comparacionesObt: ", lis.Back().Value)
//}
