# go-lodash

The package provides some functional programming style tools that can help you handle slice easily.

It is inspired by [lodash](https://lodash.com/)

## Function

You can find document at this [link](https://pkg.go.dev/github.com/ken00535/dao/pkg/handler)

dao provides functions as following

### slice

- Difference
- Filter
- Find
- ForEach
- Map
- Omit
- Reduce
- Uniq

### map

- FindKey
- Keys
- Merge
- Pick

## Install

```
go get -u github.com/ken00535/dao/pkg/handler
```

## Usage

You can use this package like

```go
import (
	"fmt"

	dao "github.com/ken00535/dao/pkg/handler"
)

type personType struct {
	Name   string
	Gender string
}

func main() {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	actual := []personType{}
	dao.Filter(people, &actual, func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	})
	fmt.Println(actual)
}
```

## Build

You can build example by

```
make
```

And run it

```
./bin/main
```

## Test

If your want to test

```bash
go test ./pkg/dao/handler
```
