package main

func main() {
	a := true
	if a {
		println("hello world")
	}

	b := 1
	if b == 1 {
		println("неявное преобразование ( if b ) не работает")
	}

	mm := map[string]string{"firstName": "Dmitriy", "lastName": "Pashkov"}
	if firstName, ok := mm["firstName"]; ok {
		println("firstName key exist, = ", firstName)
	} else {
		println("no firstName")
	}

	if firstName, ok := mm["firstName"]; !ok {
		println("no firstName")
	} else if firstName == "Dmitriy" {
		println("firstName is Dmitriy")
	} else {
		println("firstName is not Dmitriy")
	}

	for {
		println("бесконечный цикл")
		break
	}

	sl := []int{3, 4, 5, 6, 7, 8}
	value := 0
	idx := 0

	// операции по slice
	for idx < 4 {
		if idx < 2 {
			idx++
			continue
		}
		value = sl[idx]
		idx++
		println("while-stype loop, idx:", idx, "value:", value)
	}

	for i := 0; i < len(sl); i++ {
		println("c-style loop", i, sl[i])
	}

	for idx := range sl {
		println("range slice by index", idx)
	}

	for idx, val := range sl {
		println("range slice by idx-value", idx, val)
	}

	// операции по map
	for key := range mm {
		println("range map by key", key)
	}

	for key, val := range mm {
		println("range map by key-val", key, val)
	}

	for _, val := range mm {
		println("range map by val", val)
	}

	mm["firstName"] = "Dmitriy"
	mm["flag"] = "Ok"
	switch mm["firstName"] {
	case "Dmitriy", "Evgeny":
		println("switch - name is Dmitriy")
		// в отличии от других языков - не переходим в другой вариант по-умолчанию
	case "Petr":
		if mm["flag"] == "Ok" {
			break // выходим из switch, чтобы не выполнять переход в другой вариант
		}
		println("switch - name is Pert")
		fallthrough // переходим в следующий вариант
	default:
		println("switch - some other name")
	}

	// как замена множественным if else
	switch {
	case mm["firstName"] == "Dmitriy":
		println("switch2 - Dmitriy")
	case mm["lastName"] == "Pashkov":
		println("switch2 - Pashkov")
	}

	// выход из цикла будучи внутри switch
Loop:
	for key, val := range mm {
		println("switch in loop", key, val)
		switch {
		case key == "firstName" && val == "Dmitriy":
			println("switch - break loop here")
			break Loop
		}
	}

}
