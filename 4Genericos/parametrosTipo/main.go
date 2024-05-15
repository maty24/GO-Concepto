package main

func main() {
	sum[int](1, 2, 3, 4, 5)
	sum[float64](1.1, 2.2, 3.3, 4.4, 5.5)
	sum[string]("a", "b", "c", "d", "e")
}

type Number interface {
	~int | ~float64 | string //el ~ significa que acepta int , float64 y string de todas las formas posibles
}

func sum[T Number](nums ...T) T { //el [T acepta int, float64 o string como generico], el ...T es un slice de T generico
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}
