package main

import (
	"fmt"
	"net/http"
	"strings"
	"golang.org/x/net/html"
)

func main()  {
	
}

func title(url string) error{
	resp , err := http.Get(url)
	if err != nil {
		return err
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s,not text/html", url, ct)
	}
	doc,err := html.Parse(resp.Body)
	resp.Body.Close()
	if  err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func (n *html.Node)  {
		if n.Type ==  html.ElementNode{
			
		}
	}
}