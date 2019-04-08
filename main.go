package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "io"
	"log"
	"os"
	_ "os"
	"strings"
)

type siteRoad struct{
	status bool
	url string
}

func linkScrape(url1 string)  string{

	doc, err := goquery.NewDocument(url1)

	if err != nil {
		log.Fatal(err)
	}

	var link1 string

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()

		//fmt.Printf("Link #%d: '%s' - '%s'\n", index, linkText, link)
		//if withoutSpaces(strings.ToUpper(linkText)) == "КОНТАКТЫ" || withoutSpaces(strings.ToUpper(linkText))  == "CONTACT" ||withoutSpaces(strings.ToUpper(linkText)) == "KONTAKT" {
		if (strings.Contains(strings.ToUpper(linkText), "КОНТАКТ") && !strings.Contains(strings.ToUpper(linkText), "ВКОНТАКТ")) ||
			strings.Contains(strings.ToUpper(linkText), "CONTACT") || strings.Contains(strings.ToUpper(linkText), "JOINDRE") ||
			strings.Contains(strings.ToUpper(linkText), "KONTAKT") || strings.Contains(strings.ToUpper(linkText), "CONTACTO")  {//link1 = cut(link, 4)

			link1 = condition(link, url1);

		}else if strings.Contains(strings.ToUpper(link), "CONTACT")||strings.Contains(strings.ToUpper(link), "КОНТАКТ")||
			strings.Contains(strings.ToUpper(link), "KONTAKT") ||strings.Contains(strings.ToUpper(link), "ADDRESS") ||
			strings.Contains(strings.ToUpper(link), "JOINDRE") || strings.Contains(strings.ToUpper(link), "KONTAKTFORMULAR") ||
			strings.Contains(strings.ToUpper(linkText), "CONTACTO"){

			link1 = condition(link, url1);
		}

		//if strings.Contains(strings.ToUpper(link), "CONTACT"){
		//
		//	link1 = condition(link, url1);
		//}
		})

	return link1

}

func main() {

	file, err := os.Open("Site.txt")
	if err != nil {
		// здесь перехватывается ошибка
		return
	}
	defer file.Close()

	// получить размер файла
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// чтение файла
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	s := strings.Split(str, "\r\n")

	var siteRoad1 = siteRoad{url: "", status: true}

	for ii := 0; ii < len(s); ii++ {
		rEmpty  := linkScrape(s[ii])
		if rEmpty != ""{
			siteRoad1.url = rEmpty
			siteRoad1.status = true
			fmt.Println(siteRoad1)
		}else{
			fmt.Println("false")
		}
	}
	}

func condition(link string, url1 string) string{

	if cut(link, 4) == "http" {
		return  link
	}else {
		link = url1 + link
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


