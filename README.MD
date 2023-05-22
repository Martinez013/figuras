#Modelo de figuras
#Este modulo saca el area y perimetro de un cuadrado 
#Ademas hace lo mismo en el circulo.

Puedes probar con el siguiente codigo como ejemplo:

package main

import "github.com/Martinez013/figuras"

func main() {

	cuadrado1 := figuras.Cuadrado{Lado: 10}
	figuras.Medidas(&cuadrado1)

	circulo1 := figuras.Circulo{Radio: 24}
	figuras.Medidas(&circulo1)
