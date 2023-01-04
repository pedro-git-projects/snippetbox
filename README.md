| Directory | Content                                                               |
|-----------|-----------------------------------------------------------------------|
| cmd       | application specific code                                             |
| internal  | non-application specific code, such as validators and database models |
| ui        | user interface assets                                                 |


| Method | Pattern           | Handler         | Action                       |
|--------|-------------------|-----------------|------------------------------|
| Any    | /                 | home            | Display the home page        |
| Any    | /snippet/view?id= | snippetView     | Display a specific snippet   |
| POST   | /snippet/create   | snippetCreate   | Create a new snippet         |
| ANY    | /static           | http.FileServer | Serve a specific static file |
