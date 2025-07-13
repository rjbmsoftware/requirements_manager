# Requirements manager

A way of grouping of and tracking requirements between systems

## TODO

- [ ] APIs
    - [ ] requirements
    - [ ] products


### Some day

* swagger docs to use generated DTOs instead of the ent structs directly


### Swagger docs

- https://goswagger.io/go-swagger/reference/annotations/params/
- https://github.com/swaggo/swag
- go install github.com/swaggo/swag/cmd/swag@latest
- go get -u github.com/swaggo/echo-swagger
- go get -u github.com/swaggo/swag // otherwise error unknown field LeftDelim in struct literal of type "github.com/swaggo/swag".Spec

- swag fmt // formats the annotations
- swag init // generates docs
