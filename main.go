// reptile project main.go
package main

import (
	"fmt"
	"time"
	"vava6/vatools"

	"github.com/gocolly/colly"
)

func main() {
	// 执行爬取糗事百科的断子
	doPullJokeList()

	// 执行获取笑话断子
	// doPullJokes8List()

	fmt.Scanln()
}

// 拉取糗事百科的断子
func doPullJokeList() error {
	tk := time.NewTicker(10 * time.Minute)
	doPullJokeList()
	for {
		fmt.Println(vatools.GetNowTimeString(), " waiting 10m ...")
		<-tk.C
		doPullJokeList()
	}
}

func _doPullJokeList() error {
	jokeList := NewParserJokeList()
	for i := 1; i <= 13; i++ {
		RandSleep()
		jokeList.Do(fmt.Sprintf("https://www.qiushibaike.com/text/page/%d/", i))
	}
	fmt.Println("\nComplete")
	obLog.Println()
	return nil
}

// 拉取笑话断子的数据
func doPullJokes8List() error {
	// 构造对象
	c := NewJokes8List()

	for i := 1; i <= 1; i++ {
		RandSleep()
		obLog.countVisit++
		c.Do(fmt.Sprintf("%s%d", DOMAIN_JOKES8_PATH, i))
		fmt.Println("")
		obLog.Println()
	}
	return nil
}

// 构造colly.Collector并且设定它的属性
func NewCollector() *colly.Collector {
	// 实例化爬虫
	c := colly.NewCollector()
	// 随机获取一个IP地址
	//	err := c.SetProxy("http://165.227.97.247:3128")
	//	if err != nil {
	//		fmt.Println("代理设定错误：", err.Error())
	//		// os.Exit(1)
	//	}
	// 伪装成Chrome浏览器
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"

	return c
}

func SetRequestHeadersBase(r *colly.Request) {
	r.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	r.Headers.Add("Accept-Encoding", "gzip, deflate, br")
	r.Headers.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	r.Headers.Add("Connection", "keep-alive")
	r.Headers.Add("Cookie", "ga=GA1.2.1126375876.1541781037; __cur_art_index=2400; Hm_lvt_18a964a3eb14176db6e70f1dd0a3e557=1541815767,1542033687; Hm_lvt_2670efbdd59c7e3ed3749b458cafaa37=1541815639,1541836747,1541906357,1542121813; _gid=GA1.2.835845646.1542121813; _xsrf=2|f5374f7b|5ce53c48554ddbb738e6829355a2dd09|1542124213")
	r.Headers.Add("Upgrade-Insecure-Requests", "1")
}

// 设定糗事百科的Request的Headers
func SetRequestHeaders(r *colly.Request) {
	SetRequestHeadersBase(r)
	r.Headers.Add("Host", "www.qiushibaike.com")
}

func RandSleep() {
	rndTime := vatools.CRnd(1000, 2000)
	time.Sleep(time.Duration(rndTime) * time.Millisecond)
}

func SaveDb(tableName string, db *JokeDB) error {
	state, err := SaveJoke(tableName, db)
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
