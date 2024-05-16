package main

import (
	"fmt"
)

type usuario struct {
	id    int
	nome  string
	cpf   string
	email string
}

var usuarios []usuario

func insertUser() {

	novoId := len(usuarios)

	var novoNome string
	fmt.Println("Digite o nome: ")
	fmt.Scanln(&novoNome)

	var novoEmail string
	fmt.Println("Digite o email: ")
	fmt.Scanln(&novoEmail)

	var novoCpf string
	fmt.Println("Digite o cpf: ")
	fmt.Scanln(&novoCpf)

	novoUsuario := usuario{id: novoId, nome: novoNome, cpf: novoCpf, email: novoEmail}
	usuarios = append(usuarios, novoUsuario)

	fmt.Println("Usuario cadastrado com sucesso")
}

func getUser() {
	fmt.Println("----------Todos os usuarios----------")
	for _, u := range usuarios {
		fmt.Printf("Id: %d\n, Nome: %s\n, Cpf: %s\n, Email: %s\n", u.id, u.nome, u.cpf, u.email)
	}
}

func updateUser() {
	var idUser int
	fmt.Println("Qual o id do usuario que deseja alterar?")
	fmt.Scanln(&idUser)

	for i, u := range usuarios {
		if u.id == idUser {
			var novoNome string
			fmt.Println("Digite o nome: ")
			fmt.Scanln(&novoNome)

			var novoEmail string
			fmt.Println("Digite o email: ")
			fmt.Scanln(&novoEmail)

			var novoCpf string
			fmt.Println("Digite o cpf: ")
			fmt.Scanln(&novoCpf)

			usuarios[i].nome = novoNome
			usuarios[i].cpf = novoCpf
			usuarios[i].email = novoEmail

			fmt.Println("Usuário atualizado com sucesso!")
			return
		}
	}
}

func deleteUser() {
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

func getUserById() {

	var idUser int
	fmt.Println("Digite o id do usuario que deseja obter")
	fmt.Scanln(&idUser)

	fmt.Println("----------Busca por Id----------")

	for _, u := range usuarios {
		if u.id == idUser {
			fmt.Printf("Id: %d\n, Nome: %s\n, Cpf: %s\n, Email: %s\n", u.id, u.nome, u.cpf, u.email)
		}
	}
}

func getUserbyEmail() {
	var emailUser string
	fmt.Println("Digite o email do usuario que deseja obter: ")
	fmt.Scanln(&emailUser)

	fmt.Println("----------Busca por Email----------")

	for _, u := range usuarios {
		if u.email == emailUser{
			fmt.Printf("Id: %d\n, Nome: %s\n, Cpf: %s\n, Email: %s\n", u.id, u.nome, u.cpf, u.email)
		}
	}
}
