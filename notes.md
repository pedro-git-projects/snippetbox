# Notas - 2.4
## Multiplexação de Requisições 

| Padrão          | Função        | Ação                |
|-----------------|---------------|---------------------|
| /               | home          | Exibe a página home |
| /snippet        | showSnippet   | Exibe um snippet    |
| /snippet/create | createSnippet | Cria um snippet     |

Para criar novas rotas utilizamos o mesmo procedimento:

```go
createSnippet(w http.ResponseWriter, r *http.Request) {

}

showSnippet(w http.ResponseWriter, r *http.Request) {

}
```

```go
mux.HandleFunc("/snippet", showSnippet)
mux.HandleFunc("/snippet/create", createSnippet)
```

Além disso adicionamos uma modificação na função home, porque no mux padrão da linguagem Go, todo padrão terminado em / corresponde a qualquer caminho que sufixe / e não corresponda exatamente a outra rota.

Podemos lidar com isso através de uma checagem simples:

```go
home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// ...
}
