# Minha Aplicação

Bem-vindo à documentação da ImportaNFe. Este README contém informações sobre como configurar, usar e contribuir para o projeto.

## Sumário

- [Visão Geral](#visão-geral)
- [Requisitos](#requisitos)

## Visão Geral

Essa alicação, recebe um XML de uma NFe, reconhece os dados da mesma, e insere os dados em uma base local. Simples e objetiva. Criei algumas rotas, que não inserem na base, que servem somente para consuta dos dados da NFe "/produtos", "/emitente" e "/destinatario"

A rota "/insereNFe" pega o documento informado e importa a mesma na base de dados.

A minha intenção era fazer uma tela para poder testar, incluindo vários XML, porém o prazo ficou curto e acabei não completando.

Outra ideia era salvar o token no storage do navegador, e verificar a cada requisicao se está logado, por meio de um middleware.

A questão dos valores dos produtos, não achei nem no manual nem no XML o campo vAdicional e vMargem, para mim não ficou muito claro se eu devia calcular os mesmos, ou de onde pegar. Porém criei os campos na tabela e estou inserindo-os zerados.

O banco de dados foi feito em mysql como solicitado, e coloquei os scripts na pasta scripts.

Acho que é isso, foi uma ótima esperiencia utilizar o Golang, e creio que tenho muito à aprender, e a melhorar. Espero fazer parte da equipe em breve!

## Requisitos

Para rodar essa aplicação, tentei utilizar o máximo bibliotecas nativas, porém tive que instalar duas externas, para rodar a api e obter o JWT de login.
- "github.com/dgrijalva/jwt-go"
- "github.com/gin-gonic/gin"

- Coloque o arquivo XML na pasta docs (foi o que foi possível, melhoria* receber o XML base64 da requisição e salva-lo já na pasta docs)
- Informe o nome do arquivo pelo parametro da url (nfe)
- Informe o cnpj logado pelo parametro da url (cnpj)
    exemplo: http://localhost:8080/inserirNfe?nfe=41230910541434000152550010000012411749316397-nfe.xml&cnpj=10541434000152

```bash
# Clone o repositório
git clone https://github.com/renan-cardeal2002/importacao-nfe

# Inicie a aplicação
go run .
