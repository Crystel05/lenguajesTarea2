package main

import "fmt"

//Variables Globales
var comparaciones int
var encontrado bool
var cantComparaciones int
var totalComparaciones int

const (
	red   byte = byte(0)
	black byte = byte(1)
)

type RBNode struct {
	value               int
	cont                int
	color               byte
	nivel               int
	left, right, parent *RBNode
}

type RBTree struct {
	root    *RBNode
	tamanio int
}

func (arbolRN *RBTree) printTree(root *RBNode) {

	fmt.Println(fmt.Sprint("Valor ", root.value))
	fmt.Println(fmt.Sprint("Color ", root.color))
	fmt.Println(fmt.Sprint("Contador ", root.cont))
	fmt.Println(fmt.Sprint("Nivel ", root.nivel))
	if root.left != nil {
		fmt.Println("Izquierda")
		arbolRN.printTree(root.left)
	}
	if root.right != nil {
		fmt.Println("Derecha")
		arbolRN.printTree(root.right)
	}
}

func (arbolRN *RBTree) insert(root *RBNode, newNode *RBNode) *RBNode {
	comparaciones = comparaciones + 1
	/* If the tree is empty, return a new node */
	if root == nil {
		//newNode.nivel = comparaciones - 1
		return newNode
	}
	/* Otherwise, recur down the tree */
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
	/* return the (unchanged) node pointer */
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

	comparaciones = 0
	// Do a normal BST insert
	arbolRN.root = arbolRN.insert(arbolRN.root, &Node)

	// fix red black Tree violations
	if Node.cont == 1 {
		arbolRN.fixViolation(arbolRN.root, &Node)
	}

	fmt.Println(fmt.Sprint("Comparaciones ", comparaciones))

	return comparaciones

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

func (arbolRN *RBTree) busqueda(root *RBNode, value int) {
	if root != nil {
		comparaciones = comparaciones + 1
		if value == root.value {
			encontrado = true
			totalComparaciones = totalComparaciones + comparaciones
		} else {
			if value < root.value {
				arbolRN.busqueda(root.left, value)
			} else {
				arbolRN.busqueda(root.right, value)
			}
		}
	}

}

func (arbolRN *RBTree) getAltura(root *RBNode, altura int) {
	//if root.getAltura
}

func (arbolRN *RBTree) busquedaRN(value int) (int, bool) {
	comparaciones = 0
	encontrado = false
	cantComparaciones = cantComparaciones + 1

	arbolRN.busqueda(arbolRN.root, value)

	return comparaciones, encontrado

}

func (arbolRN *RBTree) insercion(array [8]int) {
	cantComparaciones = 0
	totalComparaciones = 0

	for i := 0; i < 8; i++ {
		fmt.Println(array[i])
		arbolRN.addValue(array[i])
	}

}

func main() {

	var node *RBNode
	node = nil

	arbolRN := RBTree{
		root:    node,
		tamanio: 0,
	}

	array := [8]int{10, 52, 6, 60, 70, 80, 65, 60}

	arbolRN.insercion(array)

	//	arbolRN.addValue(10)
	//	arbolRN.addValue(52)
	//	arbolRN.addValue(6)
	//	arbolRN.addValue(60)
	//	arbolRN.addValue(70)
	//	arbolRN.addValue(80)
	//	arbolRN.addValue(65)
	//	arbolRN.addValue(60)
	arbolRN.printTree(arbolRN.root)

	//	fmt.Println(arbolRN.busquedaRN(90))

}
