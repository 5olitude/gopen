package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cavaliercoder/grab"
	"github.com/tidwall/gjson"
)

func main() {
	fmt.Println("Choose the Content you like \n [1] All \n [2] Books \n [3] Conferences \n [4] Journals \n [5] Courses \n [6] Standards")
	fmt.Println("******ENTER OPTION BELOW**********")
	types := []string{"All", "Books", "Conferences", "Journals", "Courses", "Standards"}
	var useropt int
	fmt.Scanln(&useropt)
	fmt.Println("Enter the subject you wanna Explore freely!")
	var searchterm string
	fmt.Scanln(&searchterm)
	fmt.Println("Enter the no of results you wanna fetch budddy!")
	var rows int
	fmt.Scanln(&rows)
	postBody, _ := json.Marshal(map[string]interface{}{
		"newsearch":    true,
		"queryText":    searchterm,
		"highlight":    true,
		"returnFacets": [1]string{"All"},
		"returnType":   "SEARCH",
		"contentType":  types[useropt-1],
		"rowsPerPage":  rows,
	})
	req, err := http.NewRequest("POST", "https://ieeexplore.ieee.org/rest/search", bytes.NewBuffer(postBody))
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://ieeexplore.ieee.org")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	kresp := fmt.Sprintln("response Body:", string(body))
	articletitle := gjson.Get(kresp, "records.#.articleTitle")
	for key1, name := range articletitle.Array() {
		fmt.Println("              ****************************************************..............")
		key1 = key1 + 1
		println("[", +key1, "]", "  ", name.String())
		fmt.Println("              ***************************************************..............")
		fmt.Println("                                     ")

	}
	fmt.Println("                                                ")
	fmt.Println("Enter the index no of paper you wanna download?")
	var papindex int
	fmt.Scanln(&papindex)
	option := gjson.Get(kresp, "records.#.documentLink")
	doclinker := option.Array()
	optioncom := fmt.Sprintln(doclinker[papindex-1])
	linkappender := fmt.Sprintf("https://sci-hub.se/https://ieeexplore.ieee.org" + optioncom)
	linkcleaner := strings.TrimSuffix(linkappender, "\n")
	repeat := fmt.Sprintf(linkcleaner)
	doc, err := goquery.NewDocument(repeat)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("a[onclick]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("onclick")
		linkclean := fmt.Sprintln(href)
		cleaner := strings.NewReplacer("location.href=", "", "'", "", "hideme(this)", "", "\n", "")
		final := cleaner.Replace(linkclean)
		resp, _ := grab.Get(".", final)
		fmt.Println("saved to " + resp.Filename)
	})

}
