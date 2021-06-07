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

- POST /auth -- Auth Login , return JWT TOken
- GET /users -- Get user list, return User list
- GET /users/filter?[params] -- Get User by Params eg, Id, email, name, date
- GET /user/@id -- Get User by ID
- POST /user -- Create New User
- PUT /user/@id -- Edit User by ID
- DELETE /user/@id -- Delete User By ID
- DELETE /user/@id/confirm -- Confirm Delete user User by ID


## COMPILE & RUN
compile ``` go build -o dist/app ```
run ```dist/app```
