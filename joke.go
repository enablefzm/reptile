package main

import (
	"fmt"
	"strings"

	"vava6/vatools"

	"github.com/gocolly/colly"
)

// ×√+-
type JokeDB struct {
	ID       int    // 笑话ID
	Keywords string // 关键字
	Vote     int    // 好笑的投票数
	Comment  int    // 评论数
	Content  string // 内容
}

func (this *JokeDB) Clear() {
	this.ID = 0
	this.Keywords = ""
	this.Vote = 0
	this.Comment = 0
	this.Content = ""
}

func (this *JokeDB) Save() error {
	state, err := SaveJoke(this)
	switch state {
	case IS_OK:
		// fmt.Println(this.ID, "保存成功")
		obLog.success++
		fmt.Print("+")
	case IS_ERROR:
		// fmt.Println(this.ID, "保存错误：", err.Error())
		obLog.errs++
		fmt.Print("×")
	case IS_EXISTENCE:
		// fmt.Println(this.ID, "已存在无法再次保存")
		obLog.repeat++
		fmt.Print("-")
	}
	return err
}

func NewJoke() *Joke {
	p := &Joke{
		c:  NewCollector(),
		db: &JokeDB{},
	}
	// 获取KeyWords时触发
	p.c.OnHTML("meta[name|='keywords']", p.parserKeywords)
	// 获得正文
	p.c.OnHTML("div[class$='content']", p.parserContent)
	// 获得Vote和Comment
	p.c.OnHTML("div[class$='stats']", p.parserVoteAndComment)
	// 发起请求时的操作
	p.c.OnRequest(p.onRequest)
	return p
}

type Joke struct {
	c  *colly.Collector
	db *JokeDB
}

func (this *Joke) Do(sPathID string) {
	obLog.countVisit++
	// 清除数据
	this.db.Clear()
	// 获得ID
	arr := strings.Split(sPathID, "/")
	if len(arr) == 3 {
		this.db.ID = vatools.SInt(arr[2])
	}
	// 判断当前库里是否有这个信息存在有则退出
	res, err := CheckJoke(this.db.ID)
	if err != nil {
		obLog.errs++
		fmt.Print("×")
		return
	}
	if res != IS_NULL {
		obLog.repeat++
		fmt.Print("-")
		return
	}
	// 生成网址
	sPath := fmt.Sprint(DOMAIN_PATH, sPathID)
	if err := this.c.Visit(sPath); err != nil {
		obLog.errs++
		fmt.Print("≠")
		return
	}

	// 完成进行保存
	this.db.Save()
}

func (this *Joke) onRequest(r *colly.Request) {
	// fmt.Println("内容获取网址：", r.URL.String())
	SetRequestHeaders(r)
}

func (this *Joke) parserKeywords(e *colly.HTMLElement) {
	this.db.Keywords = e.Attr("content")
}

func (this *Joke) parserContent(e *colly.HTMLElement) {
	strContent, err := e.DOM.Html()
	if err != nil {
		strContent = e.Text
	}
	strContent = strings.Trim(strContent, "\n")
	this.db.Content = strContent
}

func (this *Joke) parserVoteAndComment(e *colly.HTMLElement) {
	e.ForEach("i[class$='number']", this.opeateVoteAndComment)
}

func (this *Joke) opeateVoteAndComment(i int, e *colly.HTMLElement) {
	val := vatools.SInt(e.Text)
	if i == 0 {
		// 保存好评
		this.db.Vote = val
	} else {
		this.db.Comment = val
	}
}
