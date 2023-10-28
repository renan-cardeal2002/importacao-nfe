# Minha Aplicação

Bem-vindo à documentação da ImportaNFe. Este README contém informações sobre como configurar, usar e contribuir para o projeto.

## Sumário

- [Visão Geral](#visão-geral)
- [Requisitos](#requisitos)

## Visão Geral

Essa alicação, recebe um XML de uma NFe, reconhece os dados da mesma, e insere os dados em uma base local. Simples e objetiva. Criei algumas rotas, que não inserem na base, que servem somente para consuta dos dados da NFe "/produtos", "/emitente" e "/destinatario"

## Requisitos

Para rodar essa aplicação, tentei utilizar o máximo bibliotecas nativas, porém tive que instalar duas externas, para rodar a api e obter o JWT de login

```bash
# Clone o repositório
git clone https://github.com/renan-cardeal2002/importacao-nfe

# Inicie a aplicação
go run .
