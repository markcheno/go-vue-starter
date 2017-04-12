# go-vue-starter

Copyright 2017 Mark Chenoweth

## Golang Starter project with Vue.js single page client

### Work in progress...

### Features:
- Middleware: [Negroni](https://github.com/urfave/negroni)

- Router: [Gorilla](https://github.com/gorilla/mux)

- Orm: [Gorm](https://github.com/jinzhu/gorm) (sqlite or postgres)

- Jwt authentication: [jwt-go](https://github.com/dgrijalva/jwt-go) and [go-jwt-middleware](https://github.com/auth0/go-jwt-middleware)

- [Vue.js](https://vuejs.org/) spa client with webpack

- User management

### TODO:
- config from file

- email confirmation

- logrus

- letsencrypt tls

### To get started:

``` bash
# clone repositordy
go get github.com/markcheno/go-vue-starter
cd $GOPATH/src/github.com/markcheno/go-vue-starter

# install Go depenancies (and make sure ports 3000/8080 are open)
go get -u ./... 
go run server.go

# open a new terminal and change to the client dir
cd client

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build
```

### License

MIT License  - see LICENSE for more details