package figuras

import "fmt"

type Geometrica interface {
	area() float64
	perimetro() float64
}

func Medidas(g Geometrica) {
	fmt.Println("MEDIDAS: ", g)
	fmt.Println("AREA: ", g.area())
	fmt.Println("PERIMETRO: ", g.perimetro())
}
