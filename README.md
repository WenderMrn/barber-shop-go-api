# Barber Shop Go API

## Dependencies

Install Dep
- go get -u github.com/golang/dep/cmd/dep

Init Dep (create vendor folder, Gopkg.lock and Gopkg.toml)
- dep init

Add new pkg
- dep ensure -add <pkg_path>

Synchronize vendor folder
- dep ensure -v

### Get started
 - Instalar o dep no link acima
 - Configurar a variável `GOPATH`
 - Criar a pasta raiz do projeto `$ mkdir -p $GOPATH/src/github.com/wendermrn`
 - Baixar o projeto`$ cd $GOPATH/src/github.com/wendermrn && git clone URL_DO_PROJETO`
 - Baixar dependências `$ dep ensure -update`
 - Sincronizar dependências `$ dep ensure -v`
 - Executar o projeto `$ go run main.go`

## Show Data from sqlite
Install [sqlitebrowser](https://sqlitebrowser.org/dl/)
