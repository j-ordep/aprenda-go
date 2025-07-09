package main

import "fmt"

func main() {

	x := struct {
		nome 	  	string
		sabor 	  	string
		ondeTem   	[]string
		vaiBemCom	map[string]string
	}{
		nome: "Yout",
		sabor: "doce",
		ondeTem: []string{"EUA", "Mexico"},
		vaiBemCom: map[string]string {
			"no café da manhã": "chá",
			"na janta": "coca cola",
		},
	}

	fmt.Println(x)

	fmt.Println("")
	
	for _, v := range x.ondeTem {
		fmt.Println(v)
	}

}