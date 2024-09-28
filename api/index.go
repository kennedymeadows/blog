package handler

import (
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"fmt"
)

var mds = `
# Simon Lewis

## About Me

I am a software engineer with a passion for building things. I have experience in web development, mobile development, and cloud computing. I am currently working as a software engineer at a startup in Los Angeles.

## Posts

- Post 1
- Post 2
- Post 3

## Contact


`

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}


func Handler(w http.ResponseWriter, r *http.Request) {
	md := []byte(mds)
	html := mdToHTML(md)
    
    fmt.Fprintf(w, "%s", html)
}