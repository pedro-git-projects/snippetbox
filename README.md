| Directory | Content                                                               |
|-----------|-----------------------------------------------------------------------|
| cmd       | application specific code                                             |
| internal  | non-application specific code, such as validators and database models |
| ui        | user interface assets                                                 |



| Method | Pattern           | Handler         | Action                                           |
|--------|-------------------|-----------------|--------------------------------------------------|
| GET    | /                 | home                | Display the home page                        |
| GET    | /snippet/:id      | snippetView         | Display a specific snippet                   |
| GET    | /snippet/create   | snippetCreate       | Display HTML form for creating a new snippet |
| POST   | /snippet/create   | snippetCreatePost   | Create a new snippet                         |
| GET    | /static           | http.FileServer     | Serve a specific static file                 |
