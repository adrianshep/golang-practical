package main

import (
  "fmt"
  "net/http"
  "html/template"
  "io/ioutil"
  "encoding/xml"
)

type Sitemapindex struct {
  Locations []string `xml:"sitemap>loc"`
}

type News struct {
  Titles []string `xml:"url>news>title"`
  Keywords []string `xml:"url>news>keywords"`
  Locations []string `xml:"url>loc"`
}

type NewsMap struct {
  Keyword string
  Location string
}

type NewsAggPage struct {
  Title string
  News map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
  p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
  t, _ := template.ParseFiles("basictemplating.html")
  t.Execute(w,p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/agg/", newsAggHandler)
  http.ListenAndServe(":8000", nil)
}
