package main

/*
func esPrimo(num int) bool {
	i := 2
	esPrim := true

	for esPrim && i < num {
		if num%i == 0 {
			esPrim = false
		} else {
			i += 1
		}
	}

	if esPrim {
		return true
	} else {
		return false
	}
}

func generarNumerosPseudoAleatorios(x int, n int) <-chan int {

	p := 2048
	mult := 109
	i := 853
	k := 0
	sal := x
	out := make(chan int)
	go func() {
		for k < n {
			salida := (mult*sal + i) % p
			sal = salida
			k++
			out <- salida
		}
		close(out)
	}()
	return out
}

func generarNumeros(cantidad int) []int {

	var numeros []int

	var numeroPrimo int
	es := true
	i := 0

	for es {
		numeroPrimo = rand.Intn(101-11) + 11
		if esPrimo(numeroPrimo) {
			es = false
		}
	}
	for s := range generarNumerosPseudoAleatorios(numeroPrimo, cantidad) {
		s = s % 200
		numeros = append(numeros, s)
		i++
	}

	return numeros
}

func main() {

	fmt.Println(generarNumeros(200)) //solo le entra la cantidad

}*/
