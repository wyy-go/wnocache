# wnocache 

![GitHub Repo stars](https://img.shields.io/github/stars/wyy-go/wnocache?style=social)
![GitHub](https://img.shields.io/github/license/wyy-go/wnocache)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wyy-go/wnocache)
![GitHub CI Status](https://img.shields.io/github/workflow/status/wyy-go/wnocache/ci?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/wyy-go/wnocache)](https://goreportcard.com/report/github.com/wyy-go/wnocache)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/wyy-go/wnocache?tab=doc)
[![codecov](https://codecov.io/gh/wyy-go/wnocache/branch/main/graph/badge.svg)](https://codecov.io/gh/wyy-go/wnocache)

gin middleware/handler for removing ETag related headers and adding "no cache" headers

**NOTE: This middleware will work only in development mode**

## Usage

~~~ go
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/wyy-go/wnocache"
)

func main() {
	g := gin.Default()
	g.Use(wnocache.New())
	g.Run(":8080")
}

~~~

The middleware will remove the following headers:

  * ETag
  * If-Modified-Since
  * If-Match
  * If-None-Match
  * If-Range
  * If-Unmodified-Since

and add the following headers:

  * Cache-Control: no-cache, no-store, max-age=0, must-revalidate, no-transform
  * Pragma:        no-cache
  * Expires:       Fri, 29 Aug 1997 02:14:00 EST