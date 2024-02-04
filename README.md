# localgit - Manipulando Contribuições Locais do Git com Go

---

A inspiração para este projeto surgiu de uma necessidade pessoal: eu queria uma maneira eficiente de verificar quantos commits eu havia feito em um dia, considerando que meus repositórios não se limitam apenas ao GitHub, mas também incluem o GitLab e outros.

Assim, embarquei na jornada de manipular o .git dos repositórios, utilizando Go para simplificar e aprimorar esse processo.

---

## 🚀 Começando

Essas instruções permitirão que você obtenha uma cópia do projeto em operação na sua máquina local para fins de desenvolvimento e teste.

### 📋 Pré-requisitos

Certifique-se de ter o Go na versão go1.21.6 instalado. Além disso, você precisará das seguintes bibliotecas:

```bash
go get -u github.com/jedib0t/go-pretty/v6/table
go get -u github.com/jedib0t/go-pretty/v6/text
go get -u gopkg.in/src-d/go-git.v4
go get -u gopkg.in/src-d/go-git.v4/plumbing/object
```

### 🔧 Instalação

Clone o repositório e siga os passos abaixo:

```bash
git clone https://github.com/seu-usuario/localgit.git
cd localgit
go build
./localgit -add repositorio -weeks 8
```

Lembrando que este projeto existem duas flags `add` que define o caminho para o diretório em que o repositório esta e `weeks` que define a quantidade de tempo em semanas que será analisado.

### 🛠️ Construído com

- github.com/jedib0t/go-pretty/v6/table
- github.com/jedib0t/go-pretty/v6/text
- gopkg.in/src-d/go-git.v4
- gopkg.in/src-d/go-git.v4/plumbing/object

---

## ✒️ Autores

    Darlan Guimarães - Trabalho Inicial e Documentação - darlangui

## 📄 Licença 

Este projeto está sob a licença MIT - veja o arquivo LICENSE.md para detalhes.

---

## 🎁 Expressões de gratidão

Inspiração: https://flaviocopes.com/go-git-contributions/

  
