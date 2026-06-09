package main

//panic function use kora hoy jeta program ke immediately stop kore dey ebong ekta error message print kore dey. Panic function ke call korle, program er execution immediately stop hoye jay ebong panic message print hoy.

//func main() {
//	panic("something went wrong")
//}


//recovery function use kora hoy jeta panic er poro program ke recover korte dey. Recovery function ke call korle, program er execution abar chalu hoye jay ebong panic message print hoyna.

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered from panic:", r)
		}
	}()

	panic("something went wrong")
}

//ei code e, defer function er modhye ekta anonymous function define kora hoyeche, jeta recover function ke call kore panic message ke recover kore. Jodi panic hoy, tahole recover function ke call kora hobe ebong panic message print hobe.