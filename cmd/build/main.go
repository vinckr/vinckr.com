package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// struct for data object including content and global config

type data struct {
	Body       string
	SiteTitle  string
	Year       string
	Author     string
	Pagematter pagematter
}

type pagematter struct {
	PageTitle string
}

var matter pagematter

// get files in directory
func getFilesFromDirectory(path string) []fs.DirEntry {
	files, readDirErr := os.ReadDir(path)
	if readDirErr != nil {
		log.Fatalf("Error getting filelist: %s", readDirErr)
	}
	return files
}

// read markdown file from directory
func readMarkdownFileFromDirectory(path string, filename string) []byte {
	md, readErr := os.ReadFile(path + filename)
	if readErr != nil {
		log.Fatalf("Error reading markdown file: %s", readErr)
	}
	return md
}

// split body and frontmatter
func splitBodyAndFrontmatter(md []byte) []byte {
	bodyOnly, err := frontmatter.Parse(strings.NewReader(string(md)), &matter)
	if err != nil {
		log.Fatalf("Error parsing frontmatter: %s", err)
	}
	return bodyOnly
}

// insert body in template
func buildTemplate(data data, templates ...string) string {
	var t = template.Must(template.ParseFiles(templates...))
	build := new(strings.Builder)
	templateErr := t.ExecuteTemplate(build, "Page", data)
	if templateErr != nil {
		log.Fatalf("Error building the template %s", templateErr)
	}
	return build.String()
}

// write html file
func writeHTMLFile(fileName string, outpath string, page string) {
	outPath := outpath + strings.TrimSuffix(fileName, ".md") + ".html"
	writeErr := os.WriteFile(outPath, []byte(page), 0644)
	if writeErr != nil {
		log.Fatalf("Error writing file: %s", writeErr)
	}
	fmt.Printf("\n" + fileName + " written to " + outPath)
}

func buildPage(fileName string, dir string, outpath string, templates ...string) {

	// global config
	author := "vinckr"
	sitetitle := "vinckr.com"
	currentyear := time.Now().Format("2006")
	// get markdown body
	md := readMarkdownFileFromDirectory(dir, fileName)
	bodyOnly := splitBodyAndFrontmatter(md)
	// convert markdown to html body
	extensions := parser.CommonExtensions | parser.Footnotes
	parser := parser.NewWithExtensions(extensions)
	body := markdown.ToHTML(bodyOnly, parser, nil)
	// build page object with html body and frontmatter
	page := data{string(body), sitetitle, currentyear, author, matter}
	fmt.Printf("\nBuilding page %s:", page.Pagematter.PageTitle)
	// build page with template and write to file
	build := buildTemplate(page, templates...)
	writeHTMLFile(fileName, outpath, build)
}

func buildPages(dir string, outpath string, templates ...string) {

	files := getFilesFromDirectory(dir)
	// build pages from files in directory
	for _, file := range files {
		fileName := file.Name()
		buildPage(fileName, dir, outpath, templates...)
	}
}

func main() {

	//build blogindex
	buildPage("blog.md", "./markdown/", "./public/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/body.tmpl")

	//build home page
	buildPage("index.md", "./markdown/", "./public/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/body.tmpl")

	// build all pages
	buildPages("./markdown/blog/", "./public/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/body.tmpl")

}