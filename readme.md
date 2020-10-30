# lodash

The package provides some functional programming style tools that can help you handle slice easily.

It is a implement of [lodash](https://lodash.com/) using Golang

## Function

You can find document at this [link](https://pkg.go.dev/github.com/ken00535/lodash/pkg/handler)

it provides functions as following

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
go get -u github.com/ken00535/lodash/pkg/handler
```

## Usage

You can use this package like

```go
import (
	"fmt"

	lo "github.com/ken00535/lodash/pkg/handler"
)

type personType struct {
	Name   string
	Gender string
}

func main() {
	people := []personType{{Name: "Ken", Gender: "Male"}, {Name: "Cythia", Gender: "Female"}}
	actual := []personType{}
	lo.Start(people).Filter(func(person interface{}) bool {
		return person.(personType).Name == "Ken"
	}).End(&actual)
	fmt.Println(actual)
    // -> expect is []personType{{Name: "Ken", Gender: "Male"}}

	team := make(map[string]personType)
	team["US"] = personType{Name: "Ken", Gender: "Male"}
	team["JP"] = personType{Name: "Akemi", Gender: "Female"}
    Start(team).Values().End(&actual)
    /* -> expect is []{
		{Name: "Ken", Gender: "Male"},
		{Name: "Akemi", Gender: "Female"},
	}*/
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
go test ./pkg/lodash/handler
```
