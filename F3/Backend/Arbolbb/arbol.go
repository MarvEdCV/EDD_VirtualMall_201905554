package Arbolbb

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/fernet/fernet-go"
)

type Usuario struct {
	Dpi      int    `json:"DPI"`
	Nombre   string `json:"Nombre"`
	Correo   string `json:"Correo"`
	Password string `json:"Password"`
	Cuenta   string `json:"Cuenta"` //Esta puede ser Admin o Usuario
}

type Key struct {
	Usuario   Usuario
	Izquierdo *Nodo
	Derecho   *Nodo
}

func NewKey(usuario Usuario) *Key {
	k := Key{
		Usuario:   usuario,
		Izquierdo: nil,
		Derecho:   nil,
	}
	return &k
}

type Nodo struct {
	Max       int
	NodoPadre *Nodo
	Keys      []*Key
}

func NewNodo(max int) *Nodo {
	keys := make([]*Key, max)
	n := Nodo{max, nil, keys}
	return &n
}

func (this *Nodo) Colocar(i int, llave *Key) {
	this.Keys[i] = llave
}

func (this *Nodo) CountKeys() int {
	count := 0
	for this.Keys[count] != nil {
		count++
		if count == this.Max {
			break
		}
	}
	return count
}

func (this *Nodo) getMin() int {
	return (this.Max - 1) / 2
}

func (this *Nodo) getLast() *Key {
	for i := 0; i < this.Max; i++ {
		if this.Keys[i] == nil {
			return this.Keys[i-1]
		}
	}
	return nil
}

type Arbol struct {
	k    int
	Raiz *Nodo
}

func NewArbol(nivel int) *Arbol {
	a := Arbol{
		k:    nivel,
		Raiz: nil,
	}
	nodoraiz := NewNodo(nivel)
	a.Raiz = nodoraiz
	return &a
}

func (this *Arbol) Insertar(newKey *Key) {

	if this.Raiz.Keys[0] == nil {
		this.Raiz.Colocar(0, newKey)
	} else if this.Raiz.Keys[0].Izquierdo == nil {
		lugarinsertado := -1
		node := this.Raiz
		lugarinsertado = this.colocarNodo(node, newKey)
		if lugarinsertado != -1 {
			if lugarinsertado == node.Max-1 {
				middle := node.Max / 2
				llavecentral := node.Keys[middle]
				derecho := NewNodo(this.k)
				izquierdo := NewNodo(this.k)
				indiceIzquierdo := 0
				indiceDerecho := 0
				for j := 0; j < node.Max; j++ {
					if node.Keys[j].Usuario.Dpi < llavecentral.Usuario.Dpi {
						izquierdo.Colocar(indiceIzquierdo, node.Keys[j])
						indiceIzquierdo++
						node.Colocar(j, nil)
					} else if node.Keys[j].Usuario.Dpi > llavecentral.Usuario.Dpi {
						derecho.Colocar(indiceDerecho, node.Keys[j])
						indiceDerecho++
						node.Colocar(j, nil)
					}
				}
				node.Colocar(middle, nil)
				this.Raiz = node
				this.Raiz.Colocar(0, llavecentral)
				izquierdo.NodoPadre = this.Raiz
				derecho.NodoPadre = this.Raiz
				llavecentral.Izquierdo = izquierdo
				llavecentral.Derecho = derecho
			}
		}
	} else if this.Raiz.Keys[0].Izquierdo != nil {
		node := this.Raiz
		for node.Keys[0].Izquierdo != nil {
			loop := 0
			for i := 0; i < node.Max; i, loop = i+1, loop+1 {

				if node.Keys[i] != nil {
					if node.Keys[i].Usuario.Dpi > newKey.Usuario.Dpi {
						node = node.Keys[i].Izquierdo
						break
					}
				} else {
					node = node.Keys[i-1].Derecho
					break
				}
			}
			if loop == node.Max {
				node = node.Keys[loop-1].Derecho
			}
		}
		indicecolocado := this.colocarNodo(node, newKey)
		if indicecolocado == node.Max-1 {
			for node.NodoPadre != nil {
				indicemedio := node.Max / 2
				llavecentral := node.Keys[indicemedio]
				izquierdo := NewNodo(this.k)
				derecho := NewNodo(this.k)
				indiceizquierdo, indicederecho := 0, 0
				for i := 0; i < node.Max; i++ {
					if node.Keys[i].Usuario.Dpi < llavecentral.Usuario.Dpi {
						izquierdo.Colocar(indiceizquierdo, node.Keys[i])
						indiceizquierdo++
						node.Colocar(i, nil)
					} else if node.Keys[i].Usuario.Dpi > llavecentral.Usuario.Dpi {
						derecho.Colocar(indicederecho, node.Keys[i])
						indicederecho++
						node.Colocar(i, nil)
					}
				}
				node.Colocar(indicemedio, nil)
				llavecentral.Izquierdo = izquierdo
				llavecentral.Derecho = derecho
				node = node.NodoPadre
				izquierdo.NodoPadre = node
				derecho.NodoPadre = node
				for i := 0; i < izquierdo.Max; i++ {
					if izquierdo.Keys[i] != nil {
						if izquierdo.Keys[i].Izquierdo != nil {
							izquierdo.Keys[i].Izquierdo.NodoPadre = izquierdo
						}
						if izquierdo.Keys[i].Derecho != nil {
							izquierdo.Keys[i].Derecho.NodoPadre = izquierdo
						}
					}
				}
				for i := 0; i < derecho.Max; i++ {
					if derecho.Keys[i] != nil {
						if derecho.Keys[i].Izquierdo != nil {
							derecho.Keys[i].Izquierdo.NodoPadre = derecho
						}
						if derecho.Keys[i].Derecho != nil {
							derecho.Keys[i].Derecho.NodoPadre = derecho
						}

					}
				}
				lugarcolocado := this.colocarNodo(node, llavecentral)
				if lugarcolocado == node.Max-1 {
					if node.NodoPadre == nil {
						indicecentralraiz := node.Max / 2
						llavecentralraiz := node.Keys[indicecentralraiz]
						izquierdoraiz := NewNodo(this.k)
						derechoraiz := NewNodo(this.k)
						indicederechoraiz, indiceizquierdoraiz := 0, 0

						for i := 0; i < node.Max; i++ {
							if node.Keys[i].Usuario.Dpi < llavecentralraiz.Usuario.Dpi {
								izquierdoraiz.Colocar(indiceizquierdoraiz, node.Keys[i])
								indiceizquierdoraiz++
								node.Colocar(i, nil)
							} else if node.Keys[i].Usuario.Dpi > llavecentralraiz.Usuario.Dpi {
								derechoraiz.Colocar(indicederechoraiz, node.Keys[i])
								indicederechoraiz++
								node.Colocar(i, nil)
							}
						}

						node.Colocar(indicecentralraiz, nil)
						node.Colocar(0, llavecentralraiz)
						for i := 0; i < this.k; i++ {
							if izquierdoraiz.Keys[i] != nil {
								izquierdoraiz.Keys[i].Izquierdo.NodoPadre = izquierdoraiz
								izquierdoraiz.Keys[i].Derecho.NodoPadre = izquierdoraiz
							}
						}
						for i := 0; i < this.k; i++ {
							if derechoraiz.Keys[i] != nil {
								derechoraiz.Keys[i].Izquierdo.NodoPadre = derechoraiz
								derechoraiz.Keys[i].Derecho.NodoPadre = derechoraiz
							}
						}
						llavecentralraiz.Izquierdo = izquierdoraiz
						llavecentralraiz.Derecho = derechoraiz
						izquierdoraiz.NodoPadre = node
						derechoraiz.NodoPadre = node
						this.Raiz = node
					}
					continue

				} else {
					break
				}
			}

		}
	}

}

func (this *Arbol) colocarNodo(node *Nodo, newkey *Key) int {
	index := -1
	for i := 0; i < node.Max; i++ {
		if node.Keys[i] == nil {
			placed := false
			for j := i - 1; j >= 0; j-- {
				if node.Keys[j].Usuario.Dpi > newkey.Usuario.Dpi {
					node.Colocar(j+1, node.Keys[j])
				} else {
					node.Colocar(j+1, newkey)
					node.Keys[j].Derecho = newkey.Izquierdo
					if (j+2) < this.k && node.Keys[j+2] != nil {
						node.Keys[j+2].Izquierdo = newkey.Derecho
					}
					placed = true
					break
				}
			}
			if placed == false {
				node.Colocar(0, newkey)
				node.Keys[1].Izquierdo = newkey.Derecho
			}
			index = i
			break
		}
	}
	return index
}

func graficar(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	asciiValue1 := 92
	asciiValue2 := 110
	character := rune(asciiValue1)
	character2 := rune(asciiValue2)
	separador := string(character) + string(character2)

	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++
			} else {
				fmt.Fprintf(cad, "<f%d>DPI: %d %s Nombre: %s %s Correo: %s |", j, actual.Keys[i].Usuario.Dpi, separador, actual.Keys[i].Usuario.Nombre, separador, actual.Keys[i].Usuario.Correo)
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficar(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficar(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual)) // cada %s o %d es un indicador de que elemnto va ahi %s para strings, %d para numeros, pensemos que estamos sustituyendos esos simbolos por los valores que deseamos
	}
}

func graficarCifradoSensible(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	asciiValue1 := 92
	asciiValue2 := 110
	character := rune(asciiValue1)
	character2 := rune(asciiValue2)
	separador := string(character) + string(character2)

	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++
			} else {
				fmt.Fprintf(cad, "<f%d>DPI: %s %s Nombre: %s %s Correo: %s |", j, encriptarFernetDpi(actual.Keys[i].Usuario.Dpi), separador, actual.Keys[i].Usuario.Nombre, separador, encriptarFernetCorreo(actual.Keys[i].Usuario.Correo))
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficarCifradoSensible(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficarCifradoSensible(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual)) // cada %s o %d es un indicador de que elemnto va ahi %s para strings, %d para numeros, pensemos que estamos sustituyendos esos simbolos por los valores que deseamos
	}
}

func (this *Arbol) Graficar() {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\n")
	m := make(map[string]*Nodo)
	graficar(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	guardarArchivo(builder.String(), "arbolNoCifrado")
	generarImagen("arbolNoCifrado.png", "arbolNoCifrado")
}

func (this *Arbol) GraficarSensible() {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\n")
	m := make(map[string]*Nodo)
	graficarCifradoSensible(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	guardarArchivo(builder.String(), "arbolSensible")
	generarImagen("arbolSensible.png", "arbolSensible")
}

func guardarArchivo(cadena string, name string) {
	f, err := os.Create(name + "diagrama.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(cadena)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written succesfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func generarImagen(nombre string, name string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./"+name+"diagrama.dot").Output()
	mode := int(0777)
	ioutil.WriteFile(nombre, cmd, os.FileMode(mode))
}

func (this *Arbol) Eliminar(key int) *Key {
	return this.eliminarB(this.Raiz, key)
}

func (this *Arbol) eliminarB(raiz *Nodo, key int) *Key {
	llaves := raiz.CountKeys()
	for i := 0; i < llaves; i++ {
		if key < raiz.Keys[i].Usuario.Dpi {
			if raiz.Keys[i].Izquierdo != nil {
				return this.eliminarB(raiz.Keys[i].Izquierdo, key)
			} else {
				return nil
			}
		} else if key == raiz.Keys[i].Usuario.Dpi {
			elim := raiz.Keys[i]
			if raiz.Keys[i+1] != nil {
				if raiz.Keys[i+1] == nil {
					raiz.Keys[i] = nil
					fmt.Println(raiz.Keys[i].Usuario.Dpi, " eliminado")
				} else {
					for i := 0; i < llaves-1; i++ {
						raiz.Keys[i] = raiz.Keys[i+1]
						if i == llaves-1-1 {
							raiz.Keys[i+1] = nil
						}
					}
				}
				if raiz.CountKeys() < raiz.getMin() {
					this.rebalancear(raiz)
				}
			}
			return elim
		}
		if i == raiz.CountKeys()-1 {
			if raiz.Keys[i].Derecho != nil {
				return this.eliminarB(raiz.Keys[i].Derecho, key)
			} else {
				return nil
			}
		}
	}
	return nil
}

func (this *Arbol) rebalancear(nodo *Nodo) {
	if nodo.NodoPadre != nil {
		llaves := nodo.NodoPadre.CountKeys()
		pos := 0
		for pos = 0; pos < llaves; pos++ {
			if nodo.NodoPadre.Keys[pos].Izquierdo == nodo {
				break
			}
		}
		if nodo.NodoPadre.Keys[pos].Derecho.CountKeys() > nodo.getMin() {
			mover := this.eliminarB(nodo, nodo.NodoPadre.Keys[pos].Derecho.Keys[0].Usuario.Dpi)
			mover.Derecho = nodo.NodoPadre.Keys[pos].Derecho
			mover.Izquierdo = nodo.NodoPadre.Keys[pos].Izquierdo
			nodo.NodoPadre.Keys[pos].Izquierdo = nil
			nodo.NodoPadre.Keys[pos].Derecho = nil
			nodo.insertKey(nodo.NodoPadre.Keys[pos])
			nodo.NodoPadre.Keys[pos] = mover
		} else if pos > 0 && nodo.NodoPadre.Keys[pos-1].Izquierdo.CountKeys() > nodo.getMin() {
			mover := this.eliminarB(nodo, nodo.NodoPadre.Keys[pos-1].Izquierdo.getLast().Usuario.Dpi)
			nodo.insertKey(mover)
		}
	}
}

func (nodo *Nodo) insertKey(key *Key) {
	for pos := 0; pos < nodo.Max; pos++ {
		if nodo.Keys[pos] == nil {
			nodo.Keys[pos] = key
			nodo.Keys[pos-1].Derecho = nodo.Keys[pos].Izquierdo
			break
		} else if nodo.Keys[pos].Usuario.Dpi > key.Usuario.Dpi {
			for i := nodo.CountKeys(); i > pos; i-- {
				nodo.Keys[i] = nodo.Keys[i-1]
			}
			nodo.Keys[pos] = key
			if pos > 0 {
				nodo.Keys[pos-1].Derecho = nodo.Keys[pos].Izquierdo
			}
			if pos < nodo.Max-1 && nodo.Keys[pos+1] != nil {
				nodo.Keys[pos+1].Izquierdo = nodo.Keys[pos].Derecho
			}
			break
		}
	}
	if nodo.CountKeys() == nodo.Max {
		mid := (nodo.Max - 1) / 2
		separador := nodo.Keys[mid]
		separador.Izquierdo = NewNodo(nodo.Max)
		separador.Derecho = NewNodo(nodo.Max)
		for i := 0; i < mid; i++ {
			if nodo.Keys[i].Izquierdo != nil {
				nodo.Keys[i].Izquierdo.NodoPadre = separador.Izquierdo
			}
			if nodo.Keys[i].Derecho != nil {
				nodo.Keys[i].Derecho.NodoPadre = separador.Derecho
			}
			separador.Izquierdo.Keys[i] = nodo.Keys[i]
			if nodo.Keys[mid+1+i].Izquierdo != nil {
				nodo.Keys[mid+1+i].Izquierdo.NodoPadre = separador.Izquierdo
			}
			if nodo.Keys[mid+1+i].Derecho != nil {
				nodo.Keys[mid+1+i].Derecho.NodoPadre = separador.Derecho
			}
			separador.Derecho.Keys[i] = nodo.Keys[mid+1+i]
		}
		if nodo.NodoPadre == nil {
			separador.Izquierdo.NodoPadre = nodo
			separador.Derecho.NodoPadre = nodo
			newKeys := make([]*Key, 5)
			newKeys[0] = separador
			nodo.Keys = newKeys
		} else {
			separador.Izquierdo.NodoPadre = nodo.NodoPadre
			separador.Derecho.NodoPadre = nodo.NodoPadre
			nodo.NodoPadre.insertKey(separador)
		}
	}
}

func encriptarFernetCorreo(correo string) string {
	k := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	tok, err := fernet.EncryptAndSign([]byte(correo), k[0])
	if err != nil {
		panic(err)
	}
	msg := fernet.VerifyAndDecrypt(tok, 60*time.Second, k)
	fmt.Println(string(msg), string(tok))
	return string(tok)
}

func encriptarFernetDpi(Dpi int) string {
	k := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	tok, err := fernet.EncryptAndSign([]byte(strconv.Itoa(Dpi)), k[0])
	if err != nil {
		panic(err)
	}
	msg := fernet.VerifyAndDecrypt(tok, 60*time.Second, k)
	fmt.Println(string(msg), string(tok))
	return string(tok)
}
