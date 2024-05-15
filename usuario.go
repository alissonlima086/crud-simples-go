package main

import "fmt"

type usuario struct {
	id int
	nome string
	cpf  string
}

var usuarios []usuario

func cadastrarUsuario(){

	novoId := len(usuarios)

	var novoNome string
	fmt.Println("Digite o nome: ")
	fmt.Scanln(&novoNome)

	var novoCpf string
	fmt.Println("Digite o cpf: ")
	fmt.Scanln(&novoCpf)

	novoUsuario := usuario{id: novoId, nome: novoNome, cpf: novoCpf}
	usuarios = append(usuarios, novoUsuario)

	fmt.Println("Usuario cadastrado com sucesso")
}

func obterUsuario(){
	fmt.Println("----------Todos os usuarios----------")
	for _, u := range usuarios{
		fmt.Printf("Id: %d\n, Nome: %s\n, Cpf: %s\n", u.id, u.nome, u.cpf)
	}
}

func atualizarUsuario(){
	var idUser int
	fmt.Println("Qual o id do usuario que deseja alterar?")
	fmt.Scanln(&idUser)

	for i, u := range usuarios{
		if u.id == idUser{
			var novoNome string
			fmt.Println("Digite o nome: ")
			fmt.Scanln(&novoNome)

			var novoCpf string
			fmt.Println("Digite o cpf: ")
			fmt.Scanln(&novoCpf)

			usuarios[i].nome = novoNome
			usuarios[i].cpf = novoCpf

			fmt.Println("Usuário atualizado com sucesso!")
			return
		}
	}
}

func deletarUsuario() {
	var idUser int
	fmt.Println("Qual o id do usuario que deseja deletar?")
	fmt.Scanln(&idUser)

	var indiceRemocao int

	for i, u := range usuarios {
		if u.id == idUser {
			indiceRemocao = i

			copy(usuarios[i:], usuarios[i+1:])
			usuarios = usuarios[:len(usuarios)-1]

			fmt.Println("Usuário deletado com sucesso!")
			break
		}
	}

	for j := indiceRemocao; j < len(usuarios); j++ {
		usuarios[j].id = j
	}
}


