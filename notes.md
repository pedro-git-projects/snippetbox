# Notas - 2.6
## Queries de URL 

|Método 		 | Padrão          | Função        | Ação                 |
|----------------|------------------|---------------|---------------------|
|Qualquer um     | /               | home          | Exibe a página home  |
|Qualquer um     | /snippet?id=1   | showSnippet   | Exibe o snippet de id correspondente |
|POST            | /snippet/create | createSnippet | Cria um snippet      |


Estamos simulando o acesso de dados que são buscados através de parâmetros interpolados na URL. Isso é feito utilizando

```go
r.URL.Query().Get(<parâmetro>)
```

O e isso corresponderá no URL a forma http://endereco/sub?***parametro***=***passado***

