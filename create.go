package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func ByteCount(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func DownloadAssetFile(url string, file string) {
	basePath, err := os.Getwd()
	HandleError(err)

	fileObj, err := os.Create(filepath.Join(basePath, "assets", file))
	HandleError(err)
	defer fileObj.Close()

	// Download file
	resp, err := http.Get(url)
	HandleError(err)
	defer resp.Body.Close()

	io.Copy(fileObj, resp.Body)
	HandleError(err)

	fi, err := fileObj.Stat()
	HandleError(err)

	size := fi.Size()

	fmt.Printf("Downloaded '%s' %s\n", url, ByteCount(size))
}

func MakeDirectory(directory string) {
	basePath, err := os.Getwd()
	HandleError(err)

	err = os.Mkdir(filepath.Join(basePath, directory), os.ModePerm)
	HandleError(err)
}

func CreateFile(path string, contents string) {
	basePath, err := os.Getwd()
	HandleError(err)

	err = os.WriteFile(filepath.Join(basePath, path), []byte(contents), 0644)
	HandleError(err)
}

func SetupProjectFiles(name string) {

	MakeDirectory("assets")
	MakeDirectory("handlers")
	MakeDirectory("routes")
	MakeDirectory("templates")
	MakeDirectory("templates/errors")

	DownloadAssetFile("https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css", "bulma.min.css")
	DownloadAssetFile("https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js", "alpine.min.js")
	DownloadAssetFile("https://unpkg.com/htmx.org/dist/htmx.min.js", "htmx.min.js")

	CreateFile("handlers/error.lua", errorLua)
	CreateFile("templates/base.tpl", baseTemplate)
	CreateFile("templates/errors/404.html", error404Template)
	CreateFile("main.go", mainGoFile)

	CreateFile("go.mod", CreateGoModFile(name))

}
