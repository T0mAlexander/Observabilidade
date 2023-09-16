# Prometheus e Go

## Descrição

Este repositório demonstra um exemplo de uso do Prometheus para observabilidade de uma aplicação feita em Go

## Pré-requisitos

- **[Docker](https://docs.docker.com/get-docker/)** instalado na máquina com versão mínima em 24.0.x
- **[Golang](https://go.dev/dl/)** instalado na máquina com versão mínima em 1.21.x

## Instruções de uso

1. Clone o repositório

```bash
git clone -b go-store https://github.com/T0mAlexander/Go-Alura
```

2. Navegue até o diretório

```bash
cd Go-Alura
```

3. Construa e inicialize o container do Docker em sua máquina a partir do arquivo `docker-compose.yml`

```bash
docker-compose up -d
```

4. Compile e instancie o servidor

```bash
go run main.go
```

> Nota: a aplicação será executada no endereço [localhost:3333](http://localhost:3333)
