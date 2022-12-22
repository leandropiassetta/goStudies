package main

import "fmt"

// LINK BOM SOBRE MAPA :

// https://www.digitalocean.com/community/tutorials/understanding-maps-in-go-pt

func main() {
	// Normalmente, os mapas são usados em Go para reter dados relacionados, como as informações contidas em um ID. Um mapa com dados se parece com este:

	// map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	// tipo metodo .keys():

	// for key, value := range sammy {
	// fmt.Printf("%q is the key for the value %q\n", key, value)
	// }

	// para obter uma lista apenadas das chaves :

	// 	keys := []string{}

	// for key := range sammy {
	// 	keys = append(keys, key)
	// }
	// fmt.Printf("%q", keys)

	// Você pode usar o mesmo padrão para recuperar apenas os valores em um mapa. No próximo exemplo, você fará a alocação prévia da fatia para evitar as alocações e, dessa forma, tornará o programa mais eficiente:

	// 	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	// items := make([]string, len(sammy))

	// var i int

	// for _, v := range sammy {
	// 	items[i] = v
	// 	i++
	// }
	// fmt.Printf("%q", items)

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	usuario["nome"] = "Leandro"

	fmt.Println(usuario["nome"])

	armazenandoValorMap := usuario["nome"]
	usuario["nome"] = armazenandoValorMap

	fmt.Println(usuario["nome"])
	fmt.Println(armazenandoValorMap)
	// map aninhado

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joão",
			"ultimo":   "pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "CTC",
		},
	}

	fmt.Println(usuario2)
	// delete(usuario2, "nome")
	// fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "virgem",
	}

	fmt.Println(usuario2)
}

// if resultProximity.PassProximities {
// 	mapFindings := *command.Findings

// 	finding = mapFindings[resultProximity.FindingKey]

// 	finding.Likelihood = resultProximity.Likelihood

// 	mapFindings[resultProximity.FindingKey] = finding

// }
