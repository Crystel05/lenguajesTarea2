package main

import (
	"fmt"
)

/////Numeros random//////

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

////////ARBOL///////
//Variables Globales

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
	//fmt.Println(fmt.Sprint("Color ", root.color))
	//fmt.Println(fmt.Sprint("Contador ", root.cont))
	//fmt.Println(fmt.Sprint("Nivel ", root.nivel))
	if root.left != nil {
		//fmt.Println("Izquierda")
		arbolRN.printTree(root.left)
	}
	if root.right != nil {
		//fmt.Println("Derecha")
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

// This function fixes violations
// caused by BST insertion

func (arbolRN *RBTree) fixViolation(root *RBNode, pt *RBNode) {
	var parent_pt *RBNode = nil
	var grand_parent_pt *RBNode = nil

	for (pt != root) && (pt.color != black) && (pt.parent.color == red) {

		parent_pt = pt.parent
		grand_parent_pt = pt.parent.parent

		/*  Case : A
		    Parent of pt is left child
		    of Grand-parent of pt */
		if parent_pt == grand_parent_pt.left {

			var uncle_pt *RBNode = grand_parent_pt.right

			/* Case : 1
			   The uncle of pt is also red
			   Only Recoloring required */
			if uncle_pt != nil && uncle_pt.color == red {
				grand_parent_pt.color = red
				parent_pt.color = black
				uncle_pt.color = black
				pt = grand_parent_pt
			} else {
				/* Case : 2
				   pt is right child of its parent
				   Left-rotation required */
				if pt == parent_pt.right {
					arbolRN.rotateLeft(root, parent_pt)
					pt = parent_pt
					parent_pt = pt.parent
				}

				/* Case : 3
				   pt is left child of its parent
				   Right-rotation required */
				arbolRN.rotateRight(root, grand_parent_pt)
				//Intercambiar colores
				var temp byte
				temp = parent_pt.color
				parent_pt.color = grand_parent_pt.color
				grand_parent_pt.color = temp

				pt = parent_pt
			}
		} else {
			/* Case : B
			   Parent of pt is right child
			   of Grand-parent of pt */

			var uncle_pt *RBNode = grand_parent_pt.left

			/*  Case : 1
			    The uncle of pt is also red
			    Only Recoloring required */
			if (uncle_pt != nil) && (uncle_pt.color == red) {
				grand_parent_pt.color = red
				parent_pt.color = black
				uncle_pt.color = black
				pt = grand_parent_pt
			} else {
				/* Case : 2
				   pt is left child of its parent
				   Right-rotation required */
				if pt == parent_pt.left {
					arbolRN.rotateRight(root, parent_pt)
					pt = parent_pt
					parent_pt = pt.parent
				}

				/* Case : 3
				   pt is right child of its parent
				   Left-rotation required */
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

	//fmt.Println(fmt.Sprint("Comparaciones ", arbolRN.comparaciones))

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

	var node *RBNode
	node = nil

	arbolRN := RBTree{
		root:    node,
		tamanio: 0,
	}

	array := generarNumeros(13, 200)
	//array := generarNumeros(61, 400)
	//array := generarNumeros(97, 600)
	//array := generarNumeros(89, 800)
	//array := generarNumeros(43, 1000)

	arbolRN.insercion(array)
	//arbolRN.printTree(arbolRN.root)
	array = generarNumeros(101, 10000)

	arbolRN.busquedas(array)
	println(arbolRN.totalComparaciones)
	//fmt.Println(arbolRN.busquedaRN(199))
	//fmt.Println(arbolRN.totalNodos)
	fmt.Println("Altura maxima: ", arbolRN.getAlturaMaxima(arbolRN.root))
	fmt.Println("Altura promedio: ", arbolRN.getAlturaPromedio(arbolRN.root))
	fmt.Println("Densidad: ", arbolRN.getDensidad())
	fmt.Println("comparacones totales: ", arbolRN.getTotalComparaciones())
	fmt.Println("comparaciones promedio: ", arbolRN.getComparacionesPromedio(arbolRN.root))

	//arbolRN.busquedas(array)

}
