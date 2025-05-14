# Blockchain Básico em Go

Este repositório contém um estudo sobre blockchain, implementando um modelo funcional básico na linguagem Go. O objetivo é compreender os conceitos fundamentais, como encadeamento de blocos, hashing com SHA-256 e verificação de integridade.

## Funcionalidades

- Criação de blocos encadeados.
- Geração de hash usando SHA-256.
- Verificação de integridade dos blocos.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação utilizada.
- **SHA-256**: Algoritmo de hash para encadeamento dos blocos.

## Como Executar

1. Certifique-se de ter o [Go](https://go.dev/) instalado.
2. Clone o repositório:

   ```sh
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```

3. Compile e execute o projeto:

   ```sh
   go run main.go
   ```

## Estrutura do Projeto

```
├── main.go         # Implementação principal da blockchain
├── blockchain.go   # Estrutura e funções da blockchain
└── README.md       # Documentação do projeto
```

## Exemplo de Uso

```go
 chain := InitBlockChain()
	chain.AddBlock("Primeiro")
	chain.AddBlock("Segundo")
	chain.AddBlock("Terceiro")
```

 ![image](https://github.com/user-attachments/assets/ba9de503-06a6-4114-83c4-eae919f9743c)


## Contribuição

Sugestões e melhorias são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.


