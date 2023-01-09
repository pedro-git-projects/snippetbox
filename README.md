| Directory | Content                                                               |
|-----------|-----------------------------------------------------------------------|
| cmd       | application specific code                                             |
| internal  | non-application specific code, such as validators and database models |
| ui        | user interface assets                                                 |



| Method | Pattern           | Handler         | Action                                            |
|--------|-------------------|-----------------|---------------------------------------------------|
| GET    | /                 | home                | Displays the home page                        |
| GET    | /snippet/view/:id | snippetView         | Displays a specific snippet                   |
| GET    | /snippet/create   | snippetCreate       | Displays HTML form for creating a new snippet |
| POST   | /snippet/create   | snippetCreatePost   | Creates a new snippet                         |
| GET    | /user/signup      | userSignup          | Displays HTML form for signing up a new user  |
| POST   | /user/signup      | userSignupPost      | Creates a new user                            |
| GET    | /user/login       | userLogin           | Displays HTML form for signing in a  user     |
| POST   | /user/login       | userLoginPost       | Authenticates and logs in a  user             |
| POST   | /user/logout      | userLogoutPost      | Logs out a user                               |
| GET    | /static           | http.FileServer     | Serves a specific static file                 |
