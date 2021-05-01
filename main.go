// Copyright 2021 yhyzgn md
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-30 23:16
// version: 1.0.0
// desc   :

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/88250/lute"
)

type MK struct {
	Content template.HTML
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	r.PathPrefix("/").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		style := request.URL.Query().Get("style")
		if style == "" {
			style = "solarized-dark"
		}
		mk := read("index", style)
		t, _ := template.ParseFiles("assets/tpl.html")
		t.Execute(writer, mk)
	})
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

func read(name, style string) MK {
	bs, err := ioutil.ReadFile(name + ".md")
	if nil != err {
		panic(err)
	}
	eg := lute.New()
	eg.SetAutoSpace(true)
	eg.SetCodeSyntaxHighlightDetectLang(true)
	eg.SetCodeSyntaxHighlightInlineStyle(true)
	eg.SetCodeSyntaxHighlightLineNum(true)
	eg.SetCodeSyntaxHighlight(true)
	eg.SetVditorCodeBlockPreview(true)
	eg.SetSuperBlock(true)
	eg.SetCodeSyntaxHighlightStyleName(style)
	eg.SetBlockRef(true)
	eg.SetGFMTable(true)
	eg.SetChineseParagraphBeginningSpace(true)
	eg.SetKramdownBlockIAL(true)
	eg.SetKramdownSpanIAL(true)
	eg.SetVditorMathBlockPreview(true)
	eg.SetInlineMathAllowDigitAfterOpenMarker(true)
	eg.SetYamlFrontMatter(true)
	eg.SetRenderListStyle(true)
	bs = eg.Markdown(name, bs)

	content := template.HTML(bs)
	return MK{Content: content}
}
