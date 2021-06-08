## Part of MVC Model

- maincontrollers
- mainmodels
    - Object Modeling
    - Call Database
        - Table Struct Object
        - Mysql
        - Mongodb
        - In Memory Database (Redis, Aerospike)
- mainrouters
    - Setup Router GET, POST, etc
- Middleware
    - Header Checking
    - Authentication / Tokenize
- publics
    - assets
- services
    - core
    - miscs
    - libs
- view
    -  templates


## REST API DESIGNS


- Public
    - Auth
        - POST /auth -- Auth Login , return JWT TOken
        - POST /register -- Register New User
- Users
    - GET /users -- Get user list, return User list
    - GET /users/filter?[params] -- Get User by Params eg, Id, email, name, date
    - GET /user/@id -- Get User by ID
    - PUT /user/@id -- Edit User by ID
    - DELETE /user/@id -- Delete User By ID
    - DELETE /user/@id/confirm -- Confirm Delete user User by ID
- Posts
    - GET /posts -- Get post list, return User list
    - GET /posts/filter?[params] -- Get Post by Params eg, Id, email, name, date
    - GET /post/@id -- Get Post by ID
    - POST /post -- Create New Post
    - PUT /post/@id -- Edit Post by ID
    - DELETE /post/@id -- Delete Post By ID
    - DELETE /post/@id/confirm -- Confirm Delete post by ID

## COMPILE & RUN
build Package ``` go build -o [dist/app] ```
run ```[dist/app]```
build plugin ``` go build -buildmode=plugin -o plugins/[plugin_name/plugin_name.so] plugins/[plugin_name/plugin_name.go] ```
Unit Testing ``` go test ```