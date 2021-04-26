package grafo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

// Nodo grafo
type nodo struct {
	Departamento string
	siguiente    *nodo
	lista        *ListaAdyasentes
	Distancia    int
}

//Lista nodos del grafo
type Grafo struct {
	inicio *nodo
	ultimo *nodo
}

//lista de nodos adyasentes
//los nodos en la lista de adyasentes solo guardaran la informacion no la lista de adyasentes propia
type ListaAdyasentes struct {
	inicio *nodo
	ultimo *nodo
}

func NewGrafo() *Grafo {
	fmt.Println("se creo un nuevo grafo")
	return &Grafo{nil, nil}
}

//Funcion que agrega un nuevo nodo al grafo
func (g *Grafo) InsertarNuevo(Depa string) {
	//inicializo la lista
	listaAd := &ListaAdyasentes{nil, nil}
	nuevo := &nodo{Depa, nil, listaAd, 0} //se le envia la lista inicializada

	//validaciones para insertar
	//se insertara siempre de ultimo
	if g.inicio == nil {
		g.inicio = nuevo
		g.ultimo = nuevo
	} else {
		g.ultimo.siguiente = nuevo
		g.ultimo = nuevo
	}
}

var listayautilizados []string

//funcion para agregar nodos adyasenes a los nodos
func (g *Grafo) AgregarAdyasente(Depa string, Depaadya string, Dist int) {
	//dato es el nodo al cual se le agregara un adyasente
	//ad sera el adyasente que se agrege
	// la lista de adyasentes de add siempre tiene que estar vacia

	//*** ambos nodos deben existir en el grafo antes de agregarse como adyasentes**
	nodografo := g.Buscar(Depa)

	if nodografo == nil {
		fmt.Println("no se encontro el nodo dato en el grafo")
	} else {
		//buscar el nodo ad
		nosoAd := g.Buscar(Depaadya)
		if nosoAd == nil {
			fmt.Println("no se encontro el nodo adyacente en el grafo")
		} else {
			//se crea un nodo nuevo con la info del nodo ad para que no modifique el siguiente del nodo original
			nodoad := &nodo{Depaadya, nil, nil, Dist}
			//insertamos el nuevo nodo en la lista de adyasentes
			nodografo.lista.InsertarNuevoAd(nodoad)

			//** al no se dirigido se agregega el nodo grafo como adyadente del nodo ad
			nodog := &nodo{Depa, nil, nil, Dist}
			nosoAd.lista.InsertarNuevoAd(nodog)
		}
	}

}

//Buscar nodo
func (g *Grafo) Buscar(Depa string) *nodo {
	//verificar si el grafo esta vacio
	if g.inicio != nil {
		pivote := g.inicio
		for pivote != nil {
			if pivote.Departamento == Depa {
				return pivote
			} else {
				pivote = pivote.siguiente
			}
		}
	}
	return nil
}

var retornarcamino string

//Ruta mas corta
func (g Grafo) RutaMasCorta(Inicio string, Fin string, tamanio int) string {
	terminado := false
	rr := Inicio
	for terminado == false {
		var Camino string
		var listatemp []int
		var NumMenor int
		var Anterior string
		Camino = "***" + Inicio + "-->"
		listayautilizados = append(listayautilizados, "inicio")
		if g.inicio != nil {
			temp := g.inicio
			if g.BuscarNodoyaUtilizado(temp.Departamento) == true {
				temp = temp.siguiente
			}
			for temp != nil && Inicio != Fin {
				tempad := temp.lista.inicio
				tempad1 := temp.lista.inicio
				if temp.Departamento == Inicio {
					for tempad1 != nil {
						dist := tempad1.Distancia
						if Anterior == tempad1.Departamento {
							dist = dist * 10000000000
						}
						listatemp = append(listatemp, dist)
						tempad1 = tempad1.siguiente
					}
					NumMenor = ordenarMenor(listatemp, len(listatemp))
					listatemp = nil
				}
				if temp.Departamento == Inicio {

					for tempad != nil {
						if tempad.Distancia == NumMenor {
							Camino = Camino + tempad.Departamento + "-->"
							if g.BuscarNodoyaUtilizado(tempad.Departamento) == false {
								listayautilizados = append(listayautilizados, tempad.Departamento)
							}
							Anterior = Inicio
							Inicio = tempad.Departamento
						}
						tempad = tempad.siguiente
					}
				}
				if temp.siguiente == nil {
					temp = g.inicio
				} else {
					temp = temp.siguiente
				}

			}
			fmt.Println(Camino + rr)
			retornarcamino = retornarcamino + Camino + rr

		}
		if len(listayautilizados) == tamanio {
			terminado = true
		} else {
			terminado = false
		}
	}
	return retornarcamino
}
func (g Grafo) BuscarNodoyaUtilizado(Nodo string) bool {
	for i := 0; i < len(listayautilizados); i++ {
		if Nodo == listayautilizados[i] {
			return true
		}
	}
	return false
}

//Funcion para ordenar de menor a mayor
func ordenarMenor(listNum []int, Cant int) int {
	tmp := 0
	for x := 0; x < Cant; x++ {
		for y := 0; y < Cant; y++ {
			if listNum[x] < listNum[y] {
				tmp = listNum[y]
				listNum[y] = listNum[x]
				listNum[x] = tmp
			}
		}
	}
	fmt.Print("\nArray dinamico ordenado: ")
	for i := 0; i < Cant; i++ {
		fmt.Print("[", listNum[i], "]")
	}
	fmt.Println()
	return listNum[0]
}

//Funcion que agrega un nuevo nodo adyasente a un nodo
func (l *ListaAdyasentes) InsertarNuevoAd(nuevo *nodo) {

	//validaciones para insertar
	//se insertara siempre de ultimo
	if l.inicio == nil {
		l.inicio = nuevo
		l.ultimo = nuevo
	} else {
		l.ultimo.siguiente = nuevo
		l.ultimo = nuevo
	}
}

//graficar
func (g *Grafo) Graficar() {
	os.Create("Grafo.dot")

	graphdot := getFile("Grafo.dot")

	fmt.Fprintf(graphdot, "graph G {\n")
	fmt.Fprintf(graphdot, "rankdir = LR; \n")
	fmt.Fprintf(graphdot, "subgraph cluster_1 { \n")
	fmt.Fprintf(graphdot, "node [color=seagreen1, style=filled, shape=egg]; \n")

	//recorrer el grafo
	//crear nodos
	pivote := g.inicio
	var text_aux string = ""
	for pivote != nil {
		text_aux = "n" + pivote.Departamento + "[label=\"" + pivote.Departamento + " \"] \n"
		fmt.Fprintf(graphdot, text_aux)
		pivote = pivote.siguiente
	}
	//enlazar adyasentes
	pivote = g.inicio
	for pivote != nil {
		pivoteAd := pivote.lista.inicio
		for pivoteAd != nil {
			text_aux = "n" + pivote.Departamento + " -- n" + pivoteAd.Departamento + "[label=" + strconv.Itoa(pivoteAd.Distancia) + "]\n"
			fmt.Fprintf(graphdot, text_aux)
			pivoteAd = pivoteAd.siguiente
		}
		pivote = pivote.siguiente
	}

	fmt.Fprintf(graphdot, "label = \"Grafo de Pedido\";\n")
	fmt.Fprintf(graphdot, "}\n")
	fmt.Fprintf(graphdot, "}\n")
	exec.Command("dot", "-Tpng", "Grafo.dot", "-o", "Grafo.png ").Output()

	graphdot.Close()
}

//obtener el archivo
func getFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDWR, 0775)

	if err != nil {
		log.Fatal(err)
	}
	return file
}
