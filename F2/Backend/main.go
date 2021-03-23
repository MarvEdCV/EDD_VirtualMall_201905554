package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"./List"
	"./TreeAVL"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var Vector []NodoLinealizado
var arreglotiendas []List.Tienda
var id int

type Tienda struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
	Logo         string
}
type NodoLinealizado struct {
	Indice       string
	Departamento string
	Calificacion int
	Lista        *List.Lista
}
type Departamento struct {
	Nombre  string
	Tiendas []List.Tienda
}
type Dato struct {
	Indice        string
	Departamentos []Departamento
}
type Sobre struct {
	Datos []Dato
}
type busqueda struct {
	Departamento string
	Nombre       string
	Calificacion int
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Wecome the my GO API!")
}
func rutaInicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Marvin Eduardo Catalan Veliz\nCarnet 201905554")
}
func filex(ruta string) *os.File {
	file, x := os.OpenFile(ruta, os.O_RDWR, 07775)
	if x != nil {
		log.Fatal(x)
	}
	return file
}
func (f *NodoLinealizado) Grafo() {
	os.Create("Grafo.dot")

	grafo := filex("Grafo.dot")
	fmt.Fprintf(grafo, "digraph G{\n")
	fmt.Fprintf(grafo, "rankdir = DR; \n")
	fmt.Fprintf(grafo, "color= black; \n")
	fmt.Fprintf(grafo, "\tnode [shape=cds color=black]; \n")
	fmt.Fprintf(grafo, "label= Linealizacion; \n")

	var componente string = ""

	for i := 0; i < len(Vector); i++ {
		componente = "\t\tnodo" + strconv.Itoa(i) + "[label=\"" + Vector[i].Departamento + "-" + strconv.Itoa(Vector[i].Calificacion) + "-Indice-" + Vector[i].Indice + "\"];\n"
		fmt.Fprintf(grafo, componente)
	}
	for i := 0; i < len(Vector); i++ {
		if i == len(Vector)-1 {
			componente = "nodo" + strconv.Itoa(i) + ";\n}"
			fmt.Fprintf(grafo, componente)
		} else {
			componente = "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(i+1) + ";\n"
			fmt.Fprintf(grafo, componente)
		}
	}
	grafo.Close()
	exec.Command("dot", "-Tpng", "Grafo.dot", "-o", "Grafo.png ").Output()
}
func getArreglo(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(Vector); i++ {
		Vector[i].Grafo()
	}
}
func getSave(w http.ResponseWriter, r *http.Request) {
	fmt.Println("guardar")
}
func metodopost(w http.ResponseWriter, r *http.Request) {
	//Creo una lista para guardar tiendas
	var row, column int
	var countaux int
	body, _ := ioutil.ReadAll(r.Body)
	var matrix Sobre
	json.Unmarshal(body, &matrix)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matrix)
	//fmt.Println(matrix)
	row = len(matrix.Datos)
	column = len(matrix.Datos[0].Departamentos)
	//Numero de posiciones del vector linealizado
	fmt.Println(matrix)
	var posiciones int
	posiciones = row * column * 5
	VectorLinealizado := make([]NodoLinealizado, posiciones)
	countaux = 0
	for x := 0; x < column; x++ {
		for y := 0; y < len(matrix.Datos); y++ {
			Regular := &NodoLinealizado{matrix.Datos[y].Indice, matrix.Datos[y].Departamentos[x].Nombre, 1, List.NewLista()}
			Buena := &NodoLinealizado{matrix.Datos[y].Indice, matrix.Datos[y].Departamentos[x].Nombre, 2, List.NewLista()}
			MuyBuena := &NodoLinealizado{matrix.Datos[y].Indice, matrix.Datos[y].Departamentos[x].Nombre, 3, List.NewLista()}
			Excelente := &NodoLinealizado{matrix.Datos[y].Indice, matrix.Datos[y].Departamentos[x].Nombre, 4, List.NewLista()}
			Magnifica := &NodoLinealizado{matrix.Datos[y].Indice, matrix.Datos[y].Departamentos[x].Nombre, 5, List.NewLista()}
			for z := 0; z < len(matrix.Datos[y].Departamentos[x].Tiendas); z++ {
				var calificacion int
				calificacion = matrix.Datos[y].Departamentos[x].Tiendas[z].Calificacion
				tiendita := matrix.Datos[y].Departamentos[x].Tiendas[z]
				switch calificacion {
				case 1:
					Regular.Lista.Insertar(tiendita)
					//Como en todos estos casos entrara a traer todas las listas en todos los casos guardo en mi lista de tiendas
					arreglotiendas = append(arreglotiendas, tiendita)
				case 2:
					Buena.Lista.Insertar(tiendita)
					//Como en todos estos casos entrara a traer todas las listas en todos los casos guardo en mi lista de tiendas
					arreglotiendas = append(arreglotiendas, tiendita)

				case 3:
					MuyBuena.Lista.Insertar(tiendita)
					//Como en todos estos casos entrara a traer todas las listas en todos los casos guardo en mi lista de tiendas
					arreglotiendas = append(arreglotiendas, tiendita)
				case 4:
					Excelente.Lista.Insertar(tiendita)
					//Como en todos estos casos entrara a traer todas las listas en todos los casos guardo en mi lista de tiendas
					arreglotiendas = append(arreglotiendas, tiendita)
				case 5:
					Magnifica.Lista.Insertar(tiendita)
					//Como en todos estos casos entrara a traer todas las listas en todos los casos guardo en mi lista de tiendas
					arreglotiendas = append(arreglotiendas, tiendita)
				}
			}
			VectorLinealizado[countaux] = *Regular
			countaux++
			VectorLinealizado[countaux] = *Buena
			countaux++
			VectorLinealizado[countaux] = *MuyBuena
			countaux++
			VectorLinealizado[countaux] = *Excelente
			countaux++
			VectorLinealizado[countaux] = *Magnifica
			countaux++
		}
	}
	//fmt.Println(VectorLinealizado)
	//fmt.Println(arreglotiendas)
	Vector = VectorLinealizado
}

type listaTienda struct {
	ListaTienda []List.Tienda `json:"listaTiendas"`
}

func getListaTiendas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&listaTienda{ListaTienda: arreglotiendas})
}
func metodopost1(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var send busqueda
	json.Unmarshal(body, &send)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(send)
	fmt.Println(send)
	fmt.Println(Vector)

	for i := 0; i < len(Vector); i++ {
		if Vector[i].Departamento == send.Departamento && Vector[i].Calificacion == send.Calificacion {
			Vector[i].Lista.Buscar(string(send.Nombre))
		}
	}
}
func Eliminar(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)
	name := data["nombre"]
	categoria := data["categoria"]
	calificacion := data["calificacion"]
	castcat, _ := strconv.ParseInt(calificacion, 10, 64)

	for i := 0; i < len(Vector); i++ {
		if Vector[i].Departamento == categoria && Vector[i].Calificacion == int(castcat) {
			fmt.Println("HOLLLLA")
			Vector[i].Lista.Eliminar(string(name))
			fmt.Println("La tienda :", name, " fue eliminada con exito")
		}
	}
}
func getposicion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	temp := vars["id"]
	tempcast, _ := strconv.ParseInt(temp, 10, 64)
	fmt.Println("\nLA POSICION EN EL VECTOR DE LA LISTA ES: ", tempcast)
	Vector[tempcast].Lista.Imprimir()
}

//Pocesos para Productos e inventario
//Estructura para cargar inventarios
type Inventario struct {
	Invetarios []TiendaEstructura
}

type TiendaEstructura struct {
	Tienda       string
	Departamento string
	Calificacion int
	Productos    []TreeAVL.Productos
}

func CargarInventarios(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var CapaInventario Inventario
	json.Unmarshal(body, &CapaInventario)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CapaInventario)
	var numeroTiendas int
	var numeroProdutos int
	numeroTiendas = len(CapaInventario.Invetarios)

	for x := 0; x < numeroTiendas; x++ {
		numeroProdutos = len(CapaInventario.Invetarios[x].Productos)
		name := CapaInventario.Invetarios[x].Tienda
		Arbol := TreeAVL.NewArbol(name)
		for y := 0; y < numeroProdutos; y++ {
			Arbol.InsertarRaiz(CapaInventario.Invetarios[x].Productos[y])
		}
		//fmt.Println(Arbol.Raiz)
		Arbol.GrafoAVL(name)
	}
}
func main() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", indexRoute)
	myrouter.HandleFunc("/api/Listatiendas", getListaTiendas).Methods("GET")
	myrouter.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	myrouter.HandleFunc("/cargartienda", metodopost).Methods("POST")
	myrouter.HandleFunc("/TiendaEspecifica", metodopost1).Methods("POST")
	myrouter.HandleFunc("/{id}", getposicion).Methods("GET")
	myrouter.HandleFunc("/Eliminar/{categoria}/{nombre}/{calificacion}", Eliminar).Methods("GET")
	myrouter.HandleFunc("/Guardar", getSave).Methods("GET")
	myrouter.HandleFunc("/cargarinventario", CargarInventarios).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myrouter)))
}
