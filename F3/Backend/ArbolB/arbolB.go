package ArbolB

import (
	"fmt"
)

//ARREGLO DE NODOS
var ArregloUsuarios []Usuario

const Max int = 4
const Min int = 2

type Usuario struct {
	Dpi      int
	Nombre   string
	Correo   string
	Password string
	Cuenta   string
}

type btreeNode struct {
	Val   [Max + 1]*Usuario
	Count int
	Link  [Max + 1]*btreeNode
}

var Root *btreeNode

func obtenernodo() *btreeNode {
	return Root
}

//metodo para crear un nuevo nodo
func createNode(User *Usuario, Child *btreeNode) *btreeNode {
	Newnode := new(btreeNode)
	Newnode.Val[1] = User
	Newnode.Count = 1
	Newnode.Link[0] = Root
	Newnode.Link[1] = Child
	fmt.Println("se creo un nuevo nodo")
	return Newnode
}

//metodo para colocar el nodo en la posicion adecuada
func addValToNode(User *Usuario, Pos int, Node *btreeNode, Child *btreeNode) {
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
func splitNode(User *Usuario, Puser **Usuario, Pos int, Node *btreeNode, Child *btreeNode, Newnode **btreeNode) {
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

/* establece el valor DPI en el nodo */
func setValueInNode(User *Usuario, puser **Usuario, Node *btreeNode, Child **btreeNode) bool {
	var Pos int
	if Node == nil {
		*puser = User
		*Child = nil
		return true

	}
	if User.Dpi < Node.Val[1].Dpi {
		Pos = 0

	} else {
		for Pos = Node.Count; User.Dpi < Node.Val[Pos].Dpi && Pos > 1; Pos-- {
			if User.Dpi == Node.Val[Pos].Dpi {
				fmt.Println("NODO DUPLICADO")
				return false
			}
		}
	}
	if setValueInNode(User, puser, Node.Link[Pos], Child) {
		if Node.Count < Max {
			addValToNode(*puser, Pos, Node, *Child)
		} else {
			splitNode(*puser, puser, Pos, Node, *Child, Child)
			return true
		}
	}
	return false

}
func Insertion(User *Usuario) {
	var Flag bool
	var i *Usuario
	var Child *btreeNode

	Flag = setValueInNode(User, &i, Root, &Child)
	if Flag == true {
		Root = createNode(i, Child)
	}
}

/* copia sucesor del valor a eliminar */
func CopySuccessor(MyNode *btreeNode, Pos int) {
	var Dummy *btreeNode
	Dummy = MyNode.Link[Pos]
	for Dummy.Link[0] != nil {
		Dummy = Dummy.Link[0]
	}
	MyNode.Val[Pos] = Dummy.Val[1]
}

/* elimina el valor del nodo dado y reorganiza los valores */

func RemoveVal(MyNode *btreeNode, Pos int) {
	var i int
	i = Pos + 1
	for i <= MyNode.Count {
		MyNode.Val[i-1] = MyNode.Val[i]
		MyNode.Link[i-1] = MyNode.Link[i]
		i++
	}
	MyNode.Count--
}

/* cambia el valor de padre a hijo derecho */
func DoRightShift(MyNode *btreeNode, Pos int) {
	var x *btreeNode
	x = MyNode.Link[Pos]
	j := x.Count

	for j > 0 {
		x.Val[j+1] = x.Val[j]
		x.Link[j+1] = x.Link[j]
	}
	x.Val[1] = MyNode.Val[Pos]
	x.Link[1] = x.Link[0]
	x.Count++

	x = MyNode.Link[Pos-1]
	MyNode.Val[Pos] = x.Val[x.Count]
	MyNode.Link[Pos] = x.Link[x.Count]
	x.Count--
	return
}
func DoLeftShift(Mynode *btreeNode, Pos int) {
	var j int = 1
	var x *btreeNode = Mynode.Link[Pos-1]

	x.Count++
	x.Val[x.Count] = Mynode.Val[Pos]
	x.Link[x.Count] = Mynode.Link[Pos].Link[0]

	x = Mynode.Link[0]
	Mynode.Val[Pos] = x.Val[1]
	x.Count--

	for j <= x.Count {
		x.Val[j] = x.Val[j+1]
		x.Link[j] = x.Link[j+1]
		j++
	}
	return

}

//fusionar nodos 176
func MergeNodes(Mynode *btreeNode, Pos int) {
	var j int = 1
	var x1 *btreeNode = Mynode.Link[Pos]
	var x2 *btreeNode = Mynode.Link[Pos-1]

	x2.Count++
	x2.Val[x2.Count] = Mynode.Val[Pos]
	x2.Link[x2.Count] = Mynode.Link[0]

	for j <= x1.Count {
		x2.Count++
		x2.Val[x2.Count] = x1.Val[j]
		x2.Link[x2.Count] = x1.Link[j]
		j++
	}
	j = Pos
	for j < Mynode.Count {
		Mynode.Val[j] = Mynode.Val[j+1]
		Mynode.Link[j] = Mynode.Link[j+1]
		j++

	}
	Mynode.Count--
	x1 = nil
}

//metodo para ajustar el nodo dado
func AdjustNode(Mynode *btreeNode, Pos int) {
	if Pos == 0 { /////////////////revisar
		if Mynode.Link[1].Count > Min {
			DoLeftShift(Mynode, 1)

		} else {
			MergeNodes(Mynode, 1)
		}
	} else {
		if Mynode.Count != Pos {
			if Mynode.Link[Pos-1].Count > Min {
				DoRightShift(Mynode, Pos)
			} else {
				if Mynode.Link[Pos+1].Count > Min {
					DoLeftShift(Mynode, Pos+1)
				} else {
					MergeNodes(Mynode, Pos)
				}
			}
		} else {
			if Mynode.Link[Pos-1].Count > Min {
				DoRightShift(Mynode, Pos)
			} else {
				MergeNodes(Mynode, Pos)
			}
		}
	}

}

//Metodo para eliminar el usuario del nodo
func DelValFromNode(User *Usuario, Mynode *btreeNode) bool {
	var Pos int
	var flag bool = false
	if Mynode != nil {
		if User.Dpi < Mynode.Val[1].Dpi {
			Pos = 0
			flag = false
		} else {
			for Pos = Mynode.Count; User.Dpi < Mynode.Val[Pos].Dpi && Pos > 1; Pos-- {
				if User.Dpi == Mynode.Val[Pos].Dpi {
					flag = true
				} else {
					flag = false
				}
			}
		}
		if flag == true {
			if Mynode.Link[Pos-1] != nil {
				CopySuccessor(Mynode, Pos)
				flag = DelValFromNode(Mynode.Val[Pos], Mynode.Link[Pos])
				if flag == false {
					fmt.Println("EL USUARIO NO ESTA EN EL ARBOL")
				}
			} else {
				RemoveVal(Mynode, Pos)
			}
		} else {
			flag = DelValFromNode(User, Mynode.Link[Pos])
		}
		if Mynode.Link[Pos] != nil {
			if Mynode.Link[Pos].Count < Min {
				AdjustNode(Mynode, Pos)
			}

		}

	}
	return flag
}

//Metodo para eliminar el usuario del arbol
func Deletion(User *Usuario, Mynode *btreeNode) {
	var tmp *btreeNode
	if DelValFromNode(User, Mynode) == false {
		fmt.Println("El usuario no se encuentra en el arbol")
		return
	} else {
		if Mynode.Count == 0 {
			tmp = Mynode
			fmt.Println(tmp)
			Mynode = Mynode.Link[0]
			tmp = nil
		}
	}
	Root = Mynode
	return
}

//Metodo para buscar un usuario en el nodo
func Searching(User *Usuario, Pos *int, Mynode *btreeNode) {
	if Mynode == nil {
		return
	}
	if User.Dpi < Mynode.Val[1].Dpi {
		*Pos = 0
	} else {
		for *Pos = Mynode.Count; User.Dpi < Mynode.Val[*Pos].Dpi && *Pos > 1; (*Pos)-- {
			if User.Dpi == Mynode.Val[*Pos].Dpi {
				fmt.Println("Usuario encontrado")
				return
			}
		}
	}
	Searching(User, Pos, Mynode.Link[*Pos])
	return
}

// cruce del arbol B
func Traversal(Mynode *btreeNode) {
	var i int
	if Mynode != nil {
		for i = 0; i < Mynode.Count; i++ {
			Traversal(Mynode.Link[i])
			fmt.Println(Mynode.Val[i+1])

		}
		Traversal(Mynode.Link[i])
	}
}
