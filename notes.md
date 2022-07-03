# Notas - 2.8
## Template Engine 

A linguagem Go  possui uma template engine já integrada em suas bibliotecas padrão.

Para criar um template criamos um arquivo de nome *nome*.*papel*.tmpl. Tal que os papéis são *page*, *partial* e *layout*.    

Para manter o templates DRY, basta utilizarmos o sistema de ações. Por exemplo, para definir o layout base da página basta:

```go
{{define "base"}}
{{end}}
```

Para invocar uma ação pré-definida, neste caso base, utilizamos:

```go
{{template "base" .}}
```

Para ler um template basta na função handler correspondente:

```go
files := []string{
	// lista de templates 
}
ts, err := template.ParseFiles(files...)
// tratamento de erros
err = ts.Execute(w, nil)
```
