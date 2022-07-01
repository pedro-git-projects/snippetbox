# Notas

Por enquanto temos uma função *hanlder*, a saber home.

Funções *handler* em Go são aquelas que tem a seguinte assinatura:

```go
func handler(w http.ResponseWriter, r *http.Request) {

}
```

Após criarmos a *handler* criamos um novo *router* ou como é conhecido em Go, um novo *mux*. 

Mux significa **HTTP request multiplexer** e faz exatamente o que o nome sugere, seleciona a função *handler* que tratará a requisição HTTP do padrão correspondente.

O mux pode ser criado através de:

```go
mux := http.NewServeMux()
```

Uma função handler pode ser registrada no multiplexador através da função reciever *HandleFunc*, a qual recebe dois argumentos, sendo o primeiro o padrão e a segunda a função *handler*. 

No nosso caso:

```go
mux.HandleFunc("/", home)
```

Por fim, é essencial que para iniciar o servidor utilizamos a função *reciever* ListenAndServe do pacote HTTP.  
Esta função recebe uma porta e um multiplexador.

```go
err := http.ListenAndServe(":4000", mux)
```
Uma vez iniciada, o fluxo da nossa aplicação consistirá em:

- Iniciar o servidor na porta indicada

- Ouvir requisições HTTP

- Quando uma requisição for feita, o multiplexador a encaminhará para a função handler correspondente

## Notas adicionais

### Endereços de rede

O formato passado para a função ListenAndServe é :4000 porque ela espera "host:port", portanto omitir o host fará com que o servidor ouça em todas as interfaces de rede disponíveis.

Também podemos passar portas nomeadas, como ":http" ou "http-alt" nesse caso o número da porta será buscado em /etc/services e o servidor retornará um erro caso não a encontre.

### Logs

No código foram utilizadas as funções log.Println() e log.Fatal() a diferença dessas funções com aquelas no pacote fmt é que elas vem prefixadas com a data e o horário da chamada. log.Fatal ainda chama os.Exit(1) em sua pilha de execução, terminando o programa.
