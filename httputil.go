// httputil
package htmlutil

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetDoc(url string) (dc *goquery.Document, er error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		er = err
	}

	//doc, err := goquery.NewDocumentFromResponse(resp)
	if doc, err := goquery.NewDocumentFromResponse(resp); err != nil {
		dc = doc
	} else {
		er = err
	}

	return
}

func main() {
	fmt.Println("Hello World!")
}
