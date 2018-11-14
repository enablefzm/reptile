package main

import (
	"fmt"
	"strings"
	"vava6/vatools"

	"github.com/gocolly/colly"
)

const (
	DOMAIN_JOKES8_PATH = "http://www.jokes8.com/home/getList?jokeType=3&pageIndex=" // 1500
	JOKES8_TABLE_NAME  = "jokes8"
)

func NewJokes8List() *Jokes8List {
	pt := &Jokes8List{
		c:  NewCollector(),
		db: &JokeDB{},
	}
	pt.c.OnHTML("article", pt.ParserOnHtml)
	pt.c.OnRequest(pt.ParserOnRequest)
	return pt
}

type Jokes8List struct {
	c  *colly.Collector
	db *JokeDB
}

func (this *Jokes8List) Do(sPath string) error {
	err := this.c.Visit(sPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (this *Jokes8List) ParserOnHtml(e *colly.HTMLElement) {
	// 获取正文内容
	strHtml, err := e.DOM.Find("ul > p").Html()
	if err != nil {
		return
	}
	strContent := strings.Replace(strHtml, " ", "", -1)
	strContent = strings.Trim(strContent, "\n")
	// 获取ID和赞
	pNode := e.DOM.Find("a[class$='btn-icon-good']")
	href, _ := pNode.Attr("href")
	arr := strings.Split(href, "=")
	if len(arr) < 2 {
		return
	}
	this.db.Clear()
	this.db.ID = vatools.SInt(arr[1])
	this.db.Vote = vatools.SInt(pNode.Text())
	this.db.Content = strContent
	// 判断当前表里是否有这个对象
	res, err := CheckJoke(JOKES8_TABLE_NAME, this.db.ID)
	if err != nil {
		obLog.errs++
		fmt.Print("x")
		return
	}
	if res != IS_NULL {
		obLog.repeat++
		fmt.Print("-")
		return
	}
	this.db.Save(JOKES8_TABLE_NAME)
}

func (this *Jokes8List) ParserOnRequest(r *colly.Request) {
	fmt.Print(r.URL.String(), " | ")
}
