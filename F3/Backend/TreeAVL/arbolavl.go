package TreeAVL

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

var arbol string

//Variables globales
var Ok bool //Variable para verificar si existe o no el nodo
//Estructuras a utilizar
//Estructura para mis productos
type Productos struct {
	Nombre      string  `json:"Nombre"`
	Codigo      int     `json:"Codigo"`
	Descripcion string  `json:"Descripcion"`
	Precio      float64 `json:"Precio"`
	Cantidad    int     `json:"Cantidad"`
	Imagen      string  `json:"Imagen"`
}

//Estructura para cada nodo del arbol(en este caso productos)
type NodoAVL struct {
	Producto Productos
	Izq      *NodoAVL
	Der      *NodoAVL
	Feq      int
}

//Estructura para el arbol
type Arbol struct {
	Raiz           *NodoAVL
	Nombre         string
	Tamanio        int
	ListaProductos []Productos
}

//Funciones a utilizar
//Funcion para crear un nuevo arbol
func NewArbol(name string) *Arbol {
	return &Arbol{nil, name, 0, nil}
}

//Funcion para insertar nuevo nodo
func (tree *Arbol) InsertarRaiz(Productonuevo Productos) bool {
	nuevo := &NodoAVL{Productonuevo, nil, nil, 0}
	Ok = false
	//Si la raiz no existe es porque el arbol esta vacio y lo inserto en la raiz
	if tree.Raiz == nil {
		tree.Raiz = nuevo
		tree.Tamanio++
	} else {
		//Llamamos al metodo que inserta con recursividad
		InsertarAVL(tree, tree.Raiz, nuevo)
		return Ok
	}
	return Ok
}
func InsertarAVL(tree *Arbol, raiz *NodoAVL, nuevo *NodoAVL) {
	//Comparamos si el codigo del nuevo a insertar es mayor al de la raiz
	if nuevo.Producto.Codigo > raiz.Producto.Codigo {
		//Si es mayor y el apuntador de la derecha esta vacio inserto a la derecha
		if raiz.Der == nil {
			raiz.Der = nuevo
		} else {
			//Si este no es nulo enonces el nodo nuevo vuelve a entrar al metodo para insertarlo donde corresponde pero ya con la raiz eliminada y verificada
			InsertarAVL(tree, raiz.Der, nuevo)
		}
	} else if nuevo.Producto.Codigo < raiz.Producto.Codigo { //Comparamos si el codigo del nuevo es menor al de la raiz
		//Si el de la izquierda es nulo lo insertamos a la izquierda de la raiz
		if raiz.Izq == nil {
			raiz.Izq = nuevo
		} else {
			//Si no es nulo entonces apuntamos la nueva raiz al izquierdo para que entre recursivamente al metodo y compare de nuevo
			InsertarAVL(tree, raiz.Izq, nuevo)
		}
	} else if nuevo.Producto.Codigo == raiz.Producto.Codigo { //Comparamos si el producto es igual osea ya existe
		Ok = true
	}
	//Al finalizar cualquiert tipo de insersion equilibramos el arbol
	MantenerEquilibrio(tree, raiz)
}
func (tree *Arbol) TamanioArbol() int { //Obtenemos cuantos productos tenemos en el arbol
	return tree.Tamanio
}
func profundidad(raiz *NodoAVL) int {
	if raiz == nil {
		return 0
	} else {
		var profizq = profundidad(raiz.Izq)
		var profder = profundidad(raiz.Der)

		if profizq > profder {

			return profizq + 1
		} else {
			return profder + 1
		}
	}
}

// Metodo para retornar el valor de la profundidad
func (tree *Arbol) RetornarProf() int {
	profundidad := profundidad(tree.Raiz)
	fmt.Println("La profundidad es: ", profundidad)
	return profundidad
}

//Metodos para buscar y obtener nodo
func BuscarNodo(raiz *NodoAVL, ProductoBuscar Productos) *NodoAVL {
	if raiz == nil {
		//La lista esta vacia no retorna nada
		fmt.Println("No se encontro el producto")
		return nil
	} else if raiz.Producto.Codigo == ProductoBuscar.Codigo {
		//Se encuentra el nodo y se retorna es la misma raiz
		fmt.Println("Producto encontrado con éxito")
		return raiz
	} else {
		var auxiliar *NodoAVL
		if ProductoBuscar.Codigo > raiz.Producto.Codigo { //Si es mayor buscaremos en el apuntador derecho de la raiz recursivamente con un nuevo sub arbol mas bajo
			auxiliar = BuscarNodo(raiz.Der, ProductoBuscar)
		} else if ProductoBuscar.Codigo < raiz.Producto.Codigo { //Si es mayor buscaremos en el apuntador izquierd de la raiz recursivamente con un nuevo sub arbol mas bajo
			auxiliar = BuscarNodo(raiz.Izq, ProductoBuscar)
		}
		return auxiliar
	}
}
func (tree *Arbol) ObtenerNodo(ProductoBuscar Productos) *NodoAVL {
	var ProductoEncontrado = BuscarNodo(tree.Raiz, ProductoBuscar)
	fmt.Println("El producto encontrado es:")
	fmt.Println(ProductoEncontrado)
	return ProductoEncontrado
}

//funciones para equilibrar el arbol
//Metodo para obtener el padre
func ObtenerPadreAVL(raiz *NodoAVL, ProductoObtener Productos) *NodoAVL {
	if ProductoObtener.Codigo > raiz.Producto.Codigo {
		if ProductoObtener.Codigo == raiz.Der.Producto.Codigo { //Si es igual es porque la raiz de ese momento es el padre
			return raiz
		} else {
			return ObtenerPadreAVL(raiz.Der, ProductoObtener) //Si no es igual cambiamos de raiz y entramos recursivamente al metodo
		}
	} else if ProductoObtener.Codigo < raiz.Producto.Codigo {
		if ProductoObtener.Codigo == raiz.Izq.Producto.Codigo {
			return raiz
		} else {
			return ObtenerPadreAVL(raiz.Izq, ProductoObtener)
		}
	} else {
		return nil
	}
} //Rotaciones
func RotacionII(tree *Arbol, n *NodoAVL, n1 *NodoAVL) {
	n.Izq = n1.Der
	n1.Der = n
	if n1.Feq == -1 {
		n.Feq = 0
		n1.Feq = 0
	} else {
		n.Feq = -1
		n1.Feq = 0
	}
	if tree.Raiz == n {
		n = n1
		tree.Raiz = n1
	} else {
		temp := ObtenerPadreAVL(tree.Raiz, n.Producto)
		if temp.Izq == n {
			temp.Izq = n1
		} else {
			temp.Der = n1
		}
	}
	//fmt.Println("Se realizo rotacion Izquierda Izquierda")
}
func RotacionDD(tree *Arbol, n *NodoAVL, n1 *NodoAVL) {
	n.Der = n1.Izq
	n.Izq = n

	if n1.Feq == 1 {
		n.Feq = 0
		n1.Feq = 0
	} else {
		n.Feq = 1
		n1.Feq = 0
	}
	if tree.Raiz == n {
		n = n1
		tree.Raiz = n1
	} else {
		temp := ObtenerPadreAVL(tree.Raiz, n.Producto)
		if temp.Izq == n {
			temp.Izq = n1
		} else {
			temp.Der = n1
		}
	}
	//fmt.Println("Se realizo rotacion Derecha Derecha")
}
func RotacionID(tree *Arbol, n *NodoAVL, n1 *NodoAVL, n2 *NodoAVL) {
	n.Izq = n2.Der
	n2.Der = n
	n1.Der = n2.Izq
	n2.Izq = n1

	if n2.Feq == 1 {
		n1.Feq = -1
	} else {
		n1.Feq = 0
	}
	if n2.Feq == -1 {
		n.Feq = 1
	} else {
		n.Feq = 0
	}
	n2.Feq = 0

	if tree.Raiz == n {
		n = n2
		tree.Raiz = n2
	} else {
		temp := ObtenerPadreAVL(tree.Raiz, n.Producto)
		if temp.Izq == n {
			temp.Izq = n2
		} else {
			temp.Der = n2
		}
	}

	//fmt.Println("Se realizo rotacion Izquierda Derecha ")
}
func RotacionDI(tree *Arbol, n *NodoAVL, n1 *NodoAVL, n2 *NodoAVL) {
	n.Der = n2.Izq
	n2.Izq = n
	n1.Izq = n2.Der
	n2.Der = n1
	if n2.Feq == 1 {
		n.Feq = -1
	} else {
		n.Feq = 0
	}
	if n2.Feq == -1 {
		n1.Feq = 1
	} else {
		n1.Feq = 0
	}
	n2.Feq = 0

	if tree.Raiz == n {
		n = n2
		tree.Raiz = n2
	} else {
		temp := ObtenerPadreAVL(tree.Raiz, n.Producto)
		if temp.Izq == n {
			temp.Izq = n2
		} else {
			temp.Der = n2
		}
	}

	//fmt.Println("Se realizo rotacion Derecha Izquierda ")
}

// Metodo que sirve para equilibrar el arbol
func MantenerEquilibrio(tree *Arbol, raiz *NodoAVL) {
	izq := profundidad(raiz.Izq)
	der := profundidad(raiz.Der)

	raiz.Feq = der - izq

	if raiz.Feq == -2 {
		if raiz.Izq.Feq > 0 {
			RotacionID(tree, raiz, raiz.Izq, raiz.Izq.Der)
		} else {
			RotacionII(tree, raiz, raiz.Izq)
		}
	} else if raiz.Feq == 2 {
		if raiz.Der.Feq < 0 {
			RotacionDI(tree, raiz, raiz.Der, raiz.Der.Izq)
		} else {
			RotacionDD(tree, raiz, raiz.Der)
		}
	}
}

//RECORRIDOS
// IZQ - RAIZ - DERECHA
func inorden(raia *NodoAVL) {
	if raia.Izq != nil {
		inorden(raia.Izq)
		arbol = arbol + strconv.Itoa(raia.Producto.Codigo) + "->" + strconv.Itoa(raia.Izq.Producto.Codigo) + ";\n"
	}

	//fmt.Print("nodo: ", raia.Producto, "    ")

	if raia.Der != nil {
		inorden(raia.Der)
		arbol = arbol + strconv.Itoa(raia.Producto.Codigo) + "->" + strconv.Itoa(raia.Der.Producto.Codigo) + ";\n"
	}
}

//metodo que retorna el recorrido
func (ar *Arbol) RecorridoInorden() {
	inorden(ar.Raiz)
	//fmt.Println("Termino el recorrido")
}

//Metodo recursivo para recorre Preorder
func Preorder(raiz *NodoAVL) {
	//fmt.Print("nodo: ", raiz.Producto, "    ")
	arbol = arbol + "node[label=\"Producto: " + raiz.Producto.Nombre + " Codigo: " + strconv.Itoa(raiz.Producto.Codigo) + "\"]" + strconv.Itoa(raiz.Producto.Codigo) + ";\n"
	if raiz.Izq != nil {
		Preorder(raiz.Izq)
	}

	if raiz.Der != nil {
		Preorder(raiz.Der)
	}
}

//metodo que retorna el recorrido
func (ar *Arbol) RecorridoPreorden() {
	Preorder(ar.Raiz)
	fmt.Println("Termino el recorrido")
}
func filex(ruta string) *os.File {
	file, x := os.OpenFile(ruta, os.O_RDWR, 07775)
	if x != nil {
		log.Fatal(x)
	}
	return file
}
func (tree *Arbol) GrafoAVL(name string) {
	os.Create("Arboles/" + name + "AVL.dot")
	grafo := filex("Arboles/" + name + "AVL.dot")
	fmt.Fprintf(grafo, "digraph Hash{\n")
	fmt.Fprintf(grafo, "node [color =\"turquoise\"];\n")
	//fmt.Fprintf(grafo, "color= black; \n")
	fmt.Fprintf(grafo, "subgraph clusterMarco {label=\"Arbol AVL de "+name+"\";color=black;\n")

	var code string = ""

	tree.RecorridoPreorden()
	tree.RecorridoInorden()

	code = arbol
	fmt.Fprintf(grafo, code)
	fmt.Fprintf(grafo, "}\n}")
	grafo.Close()
	exec.Command("dot", "-Tpng", "Arboles/"+name+"AVL.dot", "-o", "Arboles/"+name+"AVL.png ").Output()
	arbol = ""

}
