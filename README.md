# TaskManager-Go 📝

Um gerenciador de tarefas via linha de comando (CLI), desenvolvido em Go como projeto de portfólio durante meus estudos da linguagem.

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=flat&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-em%20desenvolvimento-yellow)

## 📖 Sobre o projeto

O **TaskManager** é uma ferramenta de linha de comando para gerenciar tarefas do dia a dia — adicionar, listar, marcar como concluídas e remover — com persistência em arquivo JSON. Foi construído com foco em praticar conceitos fundamentais de Go: structs, slices, manipulação de arquivos, tratamento de erros e organização de código em pacotes.

## ✨ Funcionalidades

- ✅ Adicionar novas tarefas
- 📋 Listar todas as tarefas
- ☑️ Marcar tarefas como concluídas
- 🗑️ Remover tarefas
- 💾 Persistência automática em arquivo JSON (`tasks.json`)

## 🚀 Como usar

### Pré-requisitos

- [Go](https://go.dev/dl/) 1.22 ou superior instalado

### Instalação

```bash
git clone https://github.com/fshenrique6/TaskManager-Go.git
cd TaskManager-Go
go build -o todo
```

### Comandos disponíveis

```bash
# Adicionar uma nova tarefa
./todo add "Estudar Go"

# Listar todas as tarefas
./todo list

# Marcar uma tarefa como concluída (usando o ID)
./todo done 1

# Remover uma tarefa (usando o ID)
./todo remove 1
```

### Exemplo de uso

```bash
$ ./todo add "Comprar leite"
Tarefa adicionada: Comprar leite

$ ./todo add "Estudar Go"
Tarefa adicionada: Estudar Go

$ ./todo list
[ ] 1 - Comprar leite
[ ] 2 - Estudar Go

$ ./todo done 1
Tarefa marcada como concluída: 1

$ ./todo list
[x] 1 - Comprar leite
[ ] 2 - Estudar Go
```

## 🏗️ Estrutura do projeto

```
taskmanager/
├── go.mod
├── main.go              # Ponto de entrada e parsing de comandos
├── task/
│   └── task.go          # Struct Task e regras de negócio
└── storage/
    └── storage.go       # Persistência (leitura/escrita em JSON)
```

## 🛠️ Tecnologias e conceitos aplicados

- Linguagem [Go](https://go.dev/)
- Manipulação de arquivos (`os`, `encoding/json`)
- Organização de código em pacotes (`task`, `storage`)
- Slices e manipulação de índices
- Tratamento de erros idiomático em Go

## 🔮 Próximos passos

- [ ] Adicionar suporte a prioridades e tags
- [ ] Filtros na listagem (por status, tag, etc.)
- [ ] Migrar o parsing de comandos para a biblioteca [Cobra](https://github.com/spf13/cobra)
- [ ] Testes unitários

## 👤 Autor

**Henrique Souza**
Projeto desenvolvido para fins de aprendizado e portfólio.

- GitHub: [@fshenrique6](https://github.com/fshenrique6)
