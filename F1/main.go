package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"./List"
	"github.com/gorilla/mux"
)

var Vector []NodoLinealizado

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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor en Go")
}
func getArreglo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[1,2,3,4]")
}
func metodopost(w http.ResponseWriter, r *http.Request) {
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
				case 2:
					Buena.Lista.Insertar(tiendita)
				case 3:
					MuyBuena.Lista.Insertar(tiendita)
				case 4:
					Excelente.Lista.Insertar(tiendita)
				case 5:
					Magnifica.Lista.Insertar(tiendita)
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
	} /*
		for z := 0; z < posiciones; z++ {
			fmt.Println("posicion:", z)
			fmt.Println(VectorLinealizado[z].Indice, VectorLinealizado[z].Departamento, VectorLinealizado[z].Calificacion)
			VectorLinealizado[z].Lista.Imprimir()
			fmt.Println()
		}*/
	fmt.Println(VectorLinealizado)
	Vector = VectorLinealizado
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
			Vector[i].Lista.Eliminar(name)
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

var id int

func request() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homePage)
	myrouter.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	myrouter.HandleFunc("/cargartienda", metodopost).Methods("POST")
	myrouter.HandleFunc("/TiendaEspecifica", metodopost1).Methods("POST")
	myrouter.HandleFunc("/{id}", getposicion).Methods("GET")
	myrouter.HandleFunc("/Eliminar/{categoria}/{nombre}/{calificacion}", Eliminar).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", myrouter))
}

func rutaInicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Marvin Eduardo Catalan Veliz\nCarnet 201905554")
}
func main() {
	request()
}
