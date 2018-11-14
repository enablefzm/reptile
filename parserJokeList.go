// 笑话页面索引器
//	Authon Jimmy 2018.11.10
//	将索引页面索引到的信息转到其分析器去处理
package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

const (
	DOMAIN_PATH string = "https://www.qiushibaike.com"
)

func NewParserJokeList() *ParserJokeList {
	pt := &ParserJokeList{
		c:            NewCollector(),
		contentPaths: make([]string, 0, 20),
	}
	// 设定OnHtml时需要执行的函数
	pt.c.OnHTML("div[class*=\"article block untagged mb15\"]", pt.ParserOnHtml)
	// 设定发请求时处理
	pt.c.OnRequest(pt.ParserOnRequest)
	// 返回对象
	return pt
}

type ParserJokeList struct {
	c            *colly.Collector
	path         string
	contentPaths []string
}

func (this *ParserJokeList) ParserOnHtml(e *colly.HTMLElement) {
	strPath := e.ChildAttr("a[class|='contentHerf']", "href")
	this.contentPaths = append(this.contentPaths, strPath)
}

func (this *ParserJokeList) ParserOnRequest(r *colly.Request) {
	// fmt.Print(r.URL.String())
	// fmt.Print("○")
	fmt.Print("|")
	SetRequestHeaders(r)
}

func (this *ParserJokeList) Do(sPath string) error {
	this.contentPaths = make([]string, 0, 20)
	// fmt.Print("开始抓取")
	err := this.c.Visit(sPath)
	// fmt.Println("，分析页面数据完成，抓取详细页面信息")
	if err == nil {
		for _, jokePath := range this.contentPaths {
			// 随机等待时间
			RandSleep()
			this.GetJokeContent(jokePath)
		}
	} else {
		fmt.Println(err.Error())
	}
	// fmt.Println("完成!")
	return err
}

func (this *ParserJokeList) GetJokeContent(jokePath string) {
	// 获取详细页面抓取对象
	ptJoke := NewJoke()
	ptJoke.Do(jokePath)
}
