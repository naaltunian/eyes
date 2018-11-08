package linkgrab

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
)

/*
** Thanks to github.com/vorozhko for his blog post at
** https://vorozhko.net/get-all-links-from-html-page-with-go-lang for
** helping with this solution
 */

func GetLinks(domain string) []string {
	var links2 []string
	res, err := http.Get(domain)
	if err != nil {
		fmt.Println(err)
		return links2
	}
	for _, v := range readLinks(res.Body) {
		links2 = append(links2, v)
	}
	defer res.Body.Close()
	return links2
}

func readLinks(body io.Reader) []string {
	var links1 []string
	t := html.NewTokenizer(body)

	for {
		tt := t.Next()

		switch tt {
		case html.ErrorToken:
			return links1
		case html.StartTagToken, html.EndTagToken:
			token := t.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links1 = append(links1, attr.Val)
					}
				}
			}
		}
	}
}