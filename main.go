package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "io"
	"log"
	"os"
	_ "os"
	"regexp"
	"strings"
)

type siteRoad struct{
	status bool
	url string
}

func linkScrape(url1 string)  string{

	rUrl, err := regexp.MatchString(`контакт|\bcontact|joindre|\bkontakt|contacto|contacta|kontakty|info\.html|location`, strings.ToLower(url1))

	defer func() {
		_check(err)
	}()

	if rUrl == true {
		return url1
	}else {

	doc, err := goquery.NewDocument(url1)

		defer func() {
			_check(err)
		}()

	var link1 string

		doc.Find("body a").Each(func(index int, item *goquery.Selection) {
			linkTag := item
			link, _ := linkTag.Attr("href")
			linkText := linkTag.Text()

			r, err := regexp.MatchString(`контакт|\bcontact|joindre|kontakt|contacto|contacta|kontakty`, strings.ToLower(linkText))
			rLink, err := regexp.MatchString(`контакт|\bcontact|joindre|\bkontakt|contacto|contacta|kontakty`, strings.ToUpper(link))

			defer func() {
				_check(err)
			}()

			switch {
				case r == true:
					link1 = condition(link, url1)
				case rLink == true:
					link1 = condition(link, url1)
				}
					})

		return link1
	}
}

func exampleLink(url1 string) {

	doc, err := goquery.NewDocument(url1)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {

		linkTag := item
		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()


		r, err := regexp.MatchString(`КОНТАКТ|CONTACT|JOINDRE|KONTAKT|CONTACTO`, strings.ToUpper(linkText))

		if err != nil {
			fmt.Println(err)
		}
		if r != false {
			fmt.Printf("'%s' - '%s'\n", linkText, url1+link)
		}else{
			fmt.Println(r)
		}
	})
	}

func main() {

	//exampleLink("https://www.kaliningradartmuseum.ru/")

	file, err := os.Open("Site.txt")
	_check(err)
	defer file.Close()

	// получить размер файла
	stat, err := file.Stat()
	_check(err)
	// чтение файла
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	_check(err)

	str := string(bs)
	s := strings.Split(str, "\r\n")

	var siteRoad1 = siteRoad{url: "", status: true}

	for ii := 0; ii < len(s); ii++ {
		//exampleLink(s[ii])
		res, err := regexp.MatchString(`HTTP|HTTPS`, strings.ToUpper(s[ii]))
		_check(err)
		if res == true {
			rEmpty := linkScrape(s[ii])
			if rEmpty != "" {

				reWiki, err := regexp.MatchString(`wikipedia|contacts.google.com|www.paypal.com|vikidia`, strings.ToLower(rEmpty))

				_check(err)

				if reWiki == true {
					fmt.Println("false")
				}else {
					siteRoad1.url = rEmpty
					siteRoad1.status = true
					fmt.Println(siteRoad1)
				}
			} else {
				fmt.Println("false")
			}
		}
	}
	}

func condition(link string, url1 string) string{

	cutLine := cut(link, 4)
	switch  {
	case cutLine == "http": return link
	case cutLine == "/www": return link
	case cutLine == "//ww": return link
	case cutLine == "www.": return link
	default:
		re := regexp.MustCompile(".*://|/.*")
		cleanLink := re.ReplaceAllString(url1, "")
		//fmt.Println(cleanLink)
		link = "http://"+cleanLink + "/"+link
		return  link
	}

}

func cut(text string, limit int) string {
	runes := []rune(text)
	if len(runes) >= limit {
		return string(runes[:limit])
	}
	return text
}

func _check(err error) {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
