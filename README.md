
# Mini Aplicação - Golang CLI

Esta é uma **mini aplicação** desenvolvida com o intuito de testar meus conhecimentos na linguagem **Golang**. A aplicação utiliza a biblioteca `urfave/cli` para criar uma interface de linha de comando (CLI) que permite buscar o IP e o nome dos servidores de um determinado host.

## Funcionalidades

- **Buscar IP de um host**: Este comando permite buscar e listar os endereços IP associados a um nome de host (ex: `google.com`).
- **Buscar servidores de nome (NS)**: Este comando permite buscar e listar os servidores de nome (NS) associados a um nome de host.

## Pré-requisitos

Para executar a aplicação, você precisa ter instalado:

- **Golang** (versão 1.16 ou superior): [Instruções de instalação](https://golang.org/doc/install)
- **Git** (opcional, se estiver clonando este repositório)

## Como executar

### 1. Clonar o repositório (opcional)

Se você está clonando o projeto de um repositório remoto, execute:

```bash
git clone https://github.com/seu-usuario/mini-aplicacao-golang.git
```

### 2. Instalar as dependências

As dependências são gerenciadas pelo Go Modules (`go.mod` e `go.sum`).

### 3. Executar a aplicação

Para rodar a aplicação diretamente, utilize o comando:

```bash
go run main.go
```

### 4. Comandos disponíveis

A aplicação fornece dois comandos principais:

#### 1. Buscar IP de um host

Este comando permite buscar os endereços IP de um host específico. Por exemplo:

```bash
go run main.go ip --host google.com
```

Esse comando irá retornar uma lista de endereços IP associados ao host `google.com`.

#### 2. Buscar servidores de nome (NS) de um host

Este comando permite buscar os servidores de nome (NS) de um host específico. Por exemplo:

```bash
go run main.go servidor --host google.com
```

Esse comando irá retornar uma lista de servidores de nome associados ao host `google.com`.

### 5. Ajuda

Você pode visualizar todos os comandos disponíveis utilizando o seguinte comando:

```bash
go run main.go --help
```

Para obter ajuda específica de um comando, por exemplo, o comando `ip`, use:

```bash
go run main.go ip --help
```

## Estrutura do Projeto

```bash
.
├── app/
│   └── app.go        # Código principal da aplicação CLI
├── main.go           # Ponto de entrada da aplicação
├── go.mod            # Arquivo de dependências do Go
├── go.sum            # Checksum das dependências
└── README.md         # Instruções de uso (este arquivo)
```
Este projeto é focado no aprendizado de **Golang**.
