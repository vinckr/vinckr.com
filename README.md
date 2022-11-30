# 🍑 - "A Static Site" Builder

## Build

Run this command to build the blog HTML from the files in `/markdown`

```bash
make build
```

Run this command to build the blog HTML from the files in `/markdown`
and commit all changes.

```bash
make u
```

## TODO Launch

- [x] build all markdown files from folder
- [x] inject data
  - [ ] from config.yaml
- [x] basic CSS
- [x] Navigation
- [x] basic welcome
- [x] other pages
- [x] Publish & deploy
- [x] build welcome page from template
  - [x] modify build func to take arguments (folder and template)
- [x] fancy welcome
- [x] frontmatter per blogpost
  - [x] Title
- [ ] dynamically build navigation
- [ ] mobile/desktop view
- [ ] proper struct for global and site config
- [ ] more pages
  - [ ] art (+pic-host)
  - [ ] music (+ mixcloud embed + 5 tracks blog)
- [ ] new sitemap/navigation
- [ ] page category?

## Links

unsorted list of URLS in url-unsorted.txt
run commands

```bash
cat url-unsorted.txt | sort -u > url.txt
go run cmd/links/title.go
```

## Structure:

welcome page -> hand-written HTML

blog -> generated from markdown using mdtohtml

content-pages -> generated from markdown using mdtohtml

galleries?

[![](https://mermaid.ink/img/pako:eNpVzz0LwjAQBuC_Es4lBS24ZlCs9QtURAWHxiG0pw1tkpKmioj_3Vg62JvunnuX9w2pyRAY3K2ocnKOuSZ-ZskFy9QovJLRaEIiGpXmHnS_luZ0K3VR92xJZ9YFv_tf1_SID4nPfnZFd00t02mHUYtxcjC1I4Px9V8XSRiGPdl0uX2rMASFVgmZ-R7vn3BwOSrkwPyaCVtw4Prjc6Jx5vTSKTBnGxxCU2XCYSyFr6-A3URZ4-cLtzBNKw)](https://mermaid.live/edit#pako:eNpVzz0LwjAQBuC_Es4lBS24ZlCs9QtURAWHxiG0pw1tkpKmioj_3Vg62JvunnuX9w2pyRAY3K2ocnKOuSZ-ZskFy9QovJLRaEIiGpXmHnS_luZ0K3VR92xJZ9YFv_tf1_SID4nPfnZFd00t02mHUYtxcjC1I4Px9V8XSRiGPdl0uX2rMASFVgmZ-R7vn3BwOSrkwPyaCVtw4Prjc6Jx5vTSKTBnGxxCU2XCYSyFr6-A3URZ4-cLtzBNKw)
