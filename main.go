package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	//"strconv"
	"strings"
)

type nodo struct {
	nombre, apellido, apodo, favoritos string
	Siguiente, Anterior                *nodo
}

type lista struct {
	cabeza *nodo
	cola   *nodo
}

func (this *lista) Insertar(nuevo *nodo) {
	if this.cabeza == nil {
		this.cabeza = nuevo
		this.cola = nuevo
	} else {
		this.cola.Siguiente = nuevo
		nuevo.Anterior = this.cola
		this.cola = nuevo
	}
}

func (this *lista) GraficarLista() {
	var cadena strings.Builder
	aux := this.cabeza
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"box\" shape=\"record\"];\n")
	for aux != nil {
		fmt.Fprintf(&cadena, " node%p[label=\"{%v|%v|%v|%v}\" style=\"radial\" fillcolor= \"yellow;0.1:red\"];\n", &(*aux), "Nombre: "+aux.nombre, "Apellido: "+aux.apellido, "Apodo: "+aux.apodo, "Favorito: "+aux.favoritos)
		aux = aux.Siguiente
	}
	aux2 := this.cabeza
	for aux2 != nil {
		if aux2.Siguiente != nil {
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(*aux2), &(*aux2.Siguiente))
			fmt.Fprintf(&cadena, "node%p->node%p;\n", &(*aux2.Siguiente), &(*aux2))
		}
		aux2 = aux2.Siguiente
	}
	fmt.Fprintf(&cadena, "}")
	generarDot(cadena.String())
}
func generarDot(cadena string) {
	f, err := os.Create("lista.dot")
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
	fmt.Println(l, "Dot creado satisfactoriamente")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//pasar a svg
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tsvg", "./lista.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("lista.svg", cmd, os.FileMode(mode))
}

func main() {
	li := lista{nil, nil}
	a := nodo{"Marvin", "Martinez", "Marvin25ronal", "Jugar apex", nil, nil}
	b := nodo{"Yaiza", "Pineda", "Bambi", "Patinar", nil, nil}
	c := nodo{"Jonathan", "Lopez", "Pancho", "Comer", nil, nil}
	d := nodo{"usuario1", "bla", "bla", "Jugar apex", nil, nil}
	e := nodo{"usuario2", "bla", "bla", "Jugar apex", nil, nil}
	f := nodo{"usuario3", "sale edd", "vamos con todo", "100 en la fase 1", nil, nil}
	li.Insertar(&a)
	li.Insertar(&b)
	li.Insertar(&c)
	li.Insertar(&d)
	li.Insertar(&e)
	li.Insertar(&f)
	li.GraficarLista()
}
