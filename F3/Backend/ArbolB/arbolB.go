package ArbolB

import (
	"fmt"
)

//ARREGLO DE NODOS
var ArregloUsuarios []Usuario

const Max int = 4
const Min int = 2

type Usuario struct {
	Dpi      string
	Nombre   string
	Correo   string
	Password string
	Cuenta   string
}

type btreeNode struct {
	Val   [Max + 1]Usuario
	Count int
	Link  [Max + 1]*btreeNode
}

var Root *btreeNode

func obtenernodo() *btreeNode {
	return Root
}

//metodo para crear un nuevo nodo
func createNode(User Usuario, Child *btreeNode) *btreeNode {
	Newnode := new(btreeNode)
	Newnode.Val[1] = User
	Newnode.Count = 1
	Newnode.Link[0] = Root
	Newnode.Link[1] = Child
	fmt.Println("se creo un nuevo nodo")
	return Newnode
}

//metodo para colocar el nodo en la posicion adecuada
func addValToNode(User Usuario, Pos int, Node *btreeNode, Child *btreeNode) {
	j := Node.Count
	for j > Pos {
		Node.Val[j+1] = Node.Val[j]
		Node.Link[j+1] = Node.Link[j]
		j--
	}
	Node.Val[j+1] = User
	Node.Link[j+1] = Child
	Node.Count++

}

// Metodo para dividir el nodo split
func splitNode(User Usuario, Puser *Usuario, Pos int, Node *btreeNode, Child *btreeNode, Newnode **btreeNode) {
	var Median, j int
	if Pos > Min {
		Median = Min + 1
	} else {
		Median = Min
	}
	*Newnode = new(btreeNode)
	j = Median + 1
	for j <= Max {
		(*Newnode).Val[j-Median] = Node.Val[j]
		(*Newnode).Link[j-Median] = Node.Link[j]
		j++
	}
	Node.Count = Median
	(*Newnode).Count = Max - Median
	if Pos <= Min {
		addValToNode(User, Pos, Node, Child)
	} else {
		addValToNode(User, Pos-Median, *Newnode, Child)
	}
	*Puser = Node.Val[Node.Count]
	(*Newnode).Link[0] = Node.Link[Node.Count]
	Node.Count--
}
