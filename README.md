## Visão Geral
MS responsável por listar ips e servidores onde estão hospedados determinados dominio

## Observação
Caso nem um valor seja especificado em --host o valor padrão é google.com.br


## Comandos
Lista IPs

```bash
go run main.go ip --host dominio
```

## Comandos
Lista servidores

```bash
go run main.go server --host dominio
```