package main

import "fmt"

func main() {

	for {
		fmt.Println("\n===== Menu =====")
		fmt.Println("1. Cadastrar usuário")
		fmt.Println("2. Obter todos os usuários")
		fmt.Println("3. Atualizar usuário")
		fmt.Println("4. Deletar usuário")
		fmt.Println("5. Buscar usuario por Id")
		fmt.Println("6. Buscar usuario por email")
		fmt.Println("7. Sair")

		var opcao int
		fmt.Println("- Escolha uma opção: ")
		fmt.Scanln(&opcao)

		switch opcao{
		case 1:
			insertUser()
		case 2:
			getUser()
		case 3:
			updateUser()
		case 4:
			deleteUser()
		case 5:
			getUserById()
		case 6:
			getUserbyEmail()
		case 7:
			fmt.Println("Saindo")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente")
		}
	}

}