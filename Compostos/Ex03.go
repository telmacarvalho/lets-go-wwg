package main

import "fmt"

func main() {
	cidades := make(map[string]string)
	cidades["Rio de Janeiro"] = "Brasil"
	cidades["Paris"] = "França"
	cidades["Londres"] = "Inglaterra"
	cidades["São Paulo"] = "Brasil"
	cidades["Chicago"] = "EUA"
	cidades["Nova York"] = "EUA"
	cidades["Toronto"] = "Canadá"
	cidades["Barra Mansa"] = "Brasil"
	fmt.Println(cidades)
	fmt.Println(cidades["Rio de Janeiro"])

	contagem := make(map[string]int)
	contagem["Brasil"] = 0
	contagem["França"] = 0
	contagem["Inglaterra"] = 0
	contagem["EUA"] = 0
	contagem["Canadá"] = 0

	if cidades["Rio de Janeiro"] == "Brasil" {
		contagem["Brasil"] = contagem["Brasil"] + 1
	}
	if cidades["Rio de Janeiro"] == "França" {
		contagem["França"] = contagem["França"] + 1
	}
	if cidades["Rio de Janeiro"] == "Inglaterra" {
		contagem["Inglaterra"] = contagem["Inglaterra"] + 1
	}
	if cidades["Rio de Janeiro"] == "EUA" {
		contagem["EUA"] = contagem["EUA"] + 1
	}
	if cidades["Rio de Janeiro"] == "Canadá" {
		contagem["Canadá"] = contagem["Canadá"] + 1
	}

	if cidades["Paris"] == "Brasil" {
		contagem["Brasil"] = contagem["Brasil"] + 1
	}
	if cidades["Paris"] == "França" {
		contagem["França"] = contagem["França"] + 1
	}
	if cidades["Paris"] == "Inglaterra" {
		contagem["Inglaterra"] = contagem["Inglaterra"] + 1
	}
	if cidades["Paris"] == "EUA" {
		contagem["EUA"] = contagem["EUA"] + 1
	}
	if cidades["Paris"] == "Canadá" {
		contagem["Canadá"] = contagem["Canadá"] + 1
	}
	if cidades["Londres"] == "Brasil" {
		contagem["Brasil"] = contagem["Brasil"] + 1
	}
	if cidades["Londres"] == "França" {
		contagem["França"] = contagem["França"] + 1
	}
	if cidades["Londres"] == "Inglaterra" {
		contagem["Inglaterra"] = contagem["Inglaterra"] + 1
	}
	if cidades["Londres"] == "EUA" {
		contagem["EUA"] = contagem["EUA"] + 1
	}
	if cidades["Londres"] == "Canadá" {
		contagem["Canadá"] = contagem["Canadá"] + 1
	}
	if cidades["São Paulo"] == "Brasil" {
		contagem["Brasil"] = contagem["Brasil"] + 1
	}
	if cidades["São Paulo"] == "França" {
		contagem["França"] = contagem["França"] + 1
	}
	if cidades["São Paulo"] == "Inglaterra" {
		contagem["Inglaterra"] = contagem["Inglaterra"] + 1
	}
	if cidades["São Paulo"] == "EUA" {
		contagem["EUA"] = contagem["EUA"] + 1
	}
	if cidades["São Paulo"] == "Canadá" {
		contagem["Canadá"] = contagem["Canadá"] + 1
	}

	if cidades["Chicago"] == "Brasil" {
		contagem["Brasil"] = contagem["Brasil"] + 1
	}
	if cidades["Chicago"] == "França" {
		contagem["França"] = contagem["França"] + 1
	}
	if cidades["Chicago"] == "Inglaterra" {
		contagem["Inglaterra"] = contagem["Inglaterra"] + 1
	}
	if cidades["Chicago"] == "EUA" {
		contagem["EUA"] = contagem["EUA"] + 1
	}
	if cidades["Chicago"] == "Canadá" {
		contagem["Canadá"] = contagem["Canadá"] + 1
	}

	if cidades["Nova York"] == "Brasil" {
		contagem["Brasil"] = contagem["Brasil"] + 1
	}
	if cidades["Nova York"] == "França" {
		contagem["França"] = contagem["França"] + 1
	}
	if cidades["Nova York"] == "Inglaterra" {
		contagem["Inglaterra"] = contagem["Inglaterra"] + 1
	}
	if cidades["Nova York"] == "EUA" {
		contagem["EUA"] = contagem["EUA"] + 1
	}
	if cidades["Nova York"] == "Canadá" {
		contagem["Canadá"] = contagem["Canadá"] + 1
	}

	if cidades["Toronto"] == "Brasil" {
		contagem["Brasil"] = contagem["Brasil"] + 1
	}
	if cidades["Toronto"] == "França" {
		contagem["França"] = contagem["França"] + 1
	}
	if cidades["Toronto"] == "Inglaterra" {
		contagem["Inglaterra"] = contagem["Inglaterra"] + 1
	}
	if cidades["Toronto"] == "EUA" {
		contagem["EUA"] = contagem["EUA"] + 1
	}
	if cidades["Toronto"] == "Canadá" {
		contagem["Canadá"] = contagem["Canadá"] + 1
	}

	if cidades["Barra Mansa"] == "Brasil" {
		contagem["Brasil"] = contagem["Brasil"] + 1
	}
	if cidades["Barra Mansa"] == "França" {
		contagem["França"] = contagem["França"] + 1
	}
	if cidades["Barra Mansa"] == "Inglaterra" {
		contagem["Inglaterra"] = contagem["Inglaterra"] + 1
	}
	if cidades["Barra Mansa"] == "EUA" {
		contagem["EUA"] = contagem["EUA"] + 1
	}
	if cidades["Barra Mansa"] == "Canadá" {
		contagem["Canadá"] = contagem["Canadá"] + 1
	}

	fmt.Println(contagem)
}
