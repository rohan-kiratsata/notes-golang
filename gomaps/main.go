package main

import "fmt"

func main() {

	lang := make(map[string]string)

	lang["JS"] = "javascript"
	lang["PY"] = "python"
	lang["RB"] = "ruby"

	fmt.Println("langs:", lang["JS"])

	delete(lang, "RB")
	fmt.Println("langs after delete:", lang)

	for key, value := range lang {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
