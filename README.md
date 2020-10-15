# gormgen

gormgen is a code generation tool to generate a better API to query and update [gorm](https://gorm.io) structs without having to deal with `interface{}`s or with database column names.

**Note** : gormgen is still is still in early development phase. It may contain bugs and the API is not yet stable. Your suggestions for improving gormgen are welcome through issues/PRs.

## Why to use gormgen

```go

// Querying

// The gorm way:
users := []User{}
err := db.Where("age > ?", 20).Order("age ASC").Limit(10).Find(&users).Error

// gormgen way
users, err := (&UserQueryBuilder{}).
  WhereAge(gormgen.GreaterThanPredicate, 20).
  OrderByAge(true).
  Limit(10).
  QueryAll(db)


// Creating Object
user := &User{
  Name: "Bla",
  Age: 20,
}

// The gorm way
err := db.Create(user).Error

// The gormgen way
err := user.Save(db)
```

- No more ugly `interface{}`s when doing in the `Where` function. Using gormgen, the passed values will be type checked.
- No more ugly strings for column names for `Where` and `Order` functions. By this, you won't need to convert the field name to the column name yourself, gormgen will do it for you. Also, you won't forget to change a column name when you change the field name because your code won't compile until you fix it everywhere.
- Query results are returned in a more intuitive way instead of passing them as a param. Also the errors are returned the "Go" way instead of explicitly accessing them.
- It doesn't alter your struct, so it's still compatible with gorm and you can still use the gorm way whenever you want (or for missing features in gormgen).

## How it works

If you have the following :

```go
//go:generate gormgen -structs User,Admin -inputDir . -importPkgs gorm.io/gorm -transformErr true
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int
	Email string
}

type Admin struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int
	Email string
}
```

Run `go generate` and gormgen will generate go file for you :
`
gen_admin.go
gen_user.go
`

For the actual generated code, check the examples folder.

## install
 - `git clone github.com/olongfen/gormgen`
 - `cd cmd/gormgen/ `
 - `go build -o gormgen main.go`
 - mv gormgen to GOPATH/bin/
 - `gormgen -structs User,Admin -inputDir ./example -importPkgs gorm.io/gorm -transformErr true`


## How to use it

- `go get -u github.com/olongfen/gormgen/...`
- Add the `//go:generate` comment mentioned above anywhere in your code.
- Add `go generate` to your build steps.

## Not yet supported features
- [X] Inferring database column name from gorm convention or gorm struct tag.
- [X] Support for anonymous structs (IMPORTANT for gorm.Model).

## Contributing

Your contributions and ideas are welcomed through issues and pull requests.

*Note for development* : Make sure to have `gormgen` in your path to be able to run the tests. Also, always run the tests with `make test` to regenerate the test structs.

## Note

The parser of this package is heavily inspired from the source code of `https://godoc.org/golang.org/x/tools/cmd/stringer`. That's where I learned how to parse and type check a go package.
