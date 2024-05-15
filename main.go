package main

import "fmt"

func main() {

	for {
		fmt.Println("\n===== Menu =====")
		fmt.Println("1. Cadastrar usuário")
		fmt.Println("2. Obter todos os usuários")
		fmt.Println("3. Atualizar usuário")
		fmt.Println("4. Deletar usuário")
		fmt.Println("5. Sair")

		var opcao int
		fmt.Println("- Escolha uma opção: ")
		fmt.Scanln(&opcao)

		switch opcao{
		case 1:
			cadastrarUsuario()
		case 2:
			obterUsuario()
		case 3:
			atualizarUsuario()
		case 4:
			deletarUsuario()
		case 5:
			fmt.Println("Saindo")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente")
		}
	}

}