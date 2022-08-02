go-render
=========
[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/mrusme/go-render) [![license](http://img.shields.io/badge/license-GPLv3-red.svg?style=flat)](https://raw.githubusercontent.com/mrusme/go-render/master/LICENSE)


Tiny Go library for the 
[Render API](https://api-docs.render.com/reference/introduction).


## Installation

```sh
go get -u github.com/mrusme/go-render
```


## Getting Started


### Listing Services

```go
package main

import (
  "log"
  "encoding/json"
  "github.com/mrusme/go-render"
)

func main() {
  r := render.New("apiKeyHere")

  services, err := r.ListServices()
  if err != nil {
    log.Panic(err)
  }

  for _, service := range services {
    b, _ := json.MarshalIndent(service, "", "    ")
    log.Printf("%s\n\n", string(b))
  }
}
```

