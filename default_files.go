package main

import (
	"fmt"
	"runtime"
	"strings"
)

const errorLua string = `local baghl = require("baghl")

function Error(code)
    return baghl.HTMLresponse({}, "errors/"..code..".html", {}, code)
end`

const baseTemplate string = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{block "title" .}}BAGHL Site{{end}}</title>
    <script src="/asset/htmx.min.js"></script>
    <script defer src="/asset/alpine.min.js"></script>
    <link rel="stylesheet" href="/asset/bulma.min.css">
  </head>
  <body>
  {{block "content" .}}
  {{end}}
  </body>
</html>`

const error404Template string = `{{define "title"}}404 Not Found{{end}}

{{define "content"}}
<h1 class="is-size-1">404 NOT FOUND</h1>
{{end}}`

const mainGoFile string = `
package main

import (
    "github.com/grqphical07/baghl"
)

func Main() {
    router := baghl.CreateRouter()

    router.Run("127.0.0.1:8000")
}
`

func CreateGoModFile(name string) string {
	return fmt.Sprintf(`module example/%s
go %s

require github.com/grqphical07/baghl`, name, strings.Replace(runtime.Version(), "go", "", 1))
}
