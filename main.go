// reptile project main.go
package main

import (
	"fmt"
	"time"
	"vava6/vatools"

	"github.com/gocolly/colly"
)

func main() {
	for {
		jokeList := NewParserJokeList()
		for i := 1; i <= 13; i++ {
			RandSleep()
			jokeList.Do(fmt.Sprintf("https://www.qiushibaike.com/text/page/%d/", i))
		}
		obLog.Println()
	}
}

// 构造colly.Collector并且设定它的属性
func NewCollector() *colly.Collector {
	// 实例化爬虫
	c := colly.NewCollector()
	// 伪装成Chrome浏览器
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"

	return c
}

func RandSleep() {
	rndTime := vatools.CRnd(1600, 3100)
	time.Sleep(time.Duration(rndTime) * time.Millisecond)
}
