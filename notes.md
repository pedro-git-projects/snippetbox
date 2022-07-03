# Notas - 2.5
## Customizando Headers HTTP 

|Método 		 | Padrão          | Função        | Ação                 |
|----------------|------------------|---------------|---------------------|
|Qualquer um     | /               | home          | Exibe a página home  |
|Qualquer um     | /snippet        | showSnippet   | Exibe um snippet     |
|POST            | /snippet/create | createSnippet | Cria um snippet      |

Neste ponto do desenvolvimento restringimos os métodos HTTP que a função para criar snippets aceita. 

Para tanto basta fazer uso do próprio pacote http da linguagem:

```go
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST") // indica o método permitido 
		http.Error(w, "Method Not Allowed", 405) 
		return
	}
	w.Write([]byte("creating snippet"))
}
```
