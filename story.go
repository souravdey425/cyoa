package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template
var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`

func NewRequest(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

// ServeHTTP implements http.Handler.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		fmt.Println(err)
	}
}

func JsonStory(f io.Reader) (Story, error) {
	d := json.NewDecoder(f)
	var story Story
	d.Decode(&story)
	return story, nil
}

type Story map[string]AutoGenerated
type AutoGenerated struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"paragraphs"`
	Options    []Options `json:"options"`
}
type Options struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}
