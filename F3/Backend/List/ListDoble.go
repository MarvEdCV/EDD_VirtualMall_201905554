package List

import (
	"fmt"
)

// Lugar donde almacenaremos la informacion
type nodo struct {
	anterior  *nodo
	siguiente *nodo
	tienda    Tienda
}

// Estructura para almacenar nodos de informacion
type Lista struct {
	inicio *nodo
	ultimo *nodo
	tam    int
}

// crear una nueva lista
func NewLista() *Lista {
	return &Lista{nil, nil, 0}
}

//insertar un nodo
func (m *Lista) Insertar(tienda Tienda) {
	nuevo := &nodo{nil, nil, tienda}

	if m.inicio == nil {
		m.inicio = nuevo
		m.ultimo = nuevo
	} else {
		m.ultimo.siguiente = nuevo
		nuevo.anterior = m.ultimo
		m.ultimo = nuevo
	}
	m.tam++
}

// imprimir la lista
func (m *Lista) Imprimir() {
	aux := m.inicio
	fmt.Println("{")
	for aux != nil {
		fmt.Println("\n\"Nombre\":", aux.tienda.Nombre, ",\n\"Descripcion\":", aux.tienda.Descripcion, ",\n\"Contacto\":", aux.tienda.Contacto, ",\n\"Calificacion\":", aux.tienda.Calificacion, ",\n\"Link\":", aux.tienda.Logo)
		fmt.Println()
		aux = aux.siguiente
	}
	fmt.Print("}")
	fmt.Println()
	fmt.Println("\nTOTAL TIENDAS = ", m.tam)
	if m.tam == 0 {
		fmt.Println("****LISTA VACIA*****")
	}
}

type Tienda struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
	Logo         string `json:"Logo"`
}

//Buscar Elemento dentro de lista
func (m *Lista) Buscar(nombre string) *nodo {
	variable := 0
	aux := m.inicio
	for aux != nil {
		if aux.tienda.Nombre == nombre {
			variable++
			fmt.Println("Si se encontro el nodo")
			fmt.Print("{\n\"Nombre\": \""+aux.tienda.Nombre+"\"\n"+"\"Descripcion\": \""+aux.tienda.Descripcion+"\"\n\"Contacto\": \""+aux.tienda.Contacto+"\"\n\"Calificacion\": \"", aux.tienda.Calificacion)
			fmt.Print("\"\n}\n")
			return aux
		}
		aux = aux.siguiente
	}
	return aux
}

//Eliminar nodo de la lista
func (m *Lista) Eliminar(nombre string) {
	aux := m.Buscar(nombre)
	fmt.Print(aux)
	if m.inicio == aux {
		m.inicio = aux.siguiente
		aux.siguiente.anterior = nil
		aux.siguiente = nil
	} else if m.ultimo == aux {
		m.ultimo = aux.anterior
		aux.anterior.siguiente = nil
		aux.anterior = nil
	} else {
		aux.anterior.siguiente = aux.siguiente
		aux.siguiente.anterior = aux.anterior
		aux.anterior = nil
		aux.siguiente = nil
	}
	m.tam--
}
