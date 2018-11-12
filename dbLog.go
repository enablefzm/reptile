package main

import (
	"fmt"
)

var obLog *dbLog = &dbLog{}

type dbLog struct {
	countVisit int // 总共请求次数
	success    int // 成功保存数量
	repeat     int // 重复数量
	errs       int // 失败数量
}

func (this *dbLog) clear() {
	this.countVisit = 0
	this.success = 0
	this.repeat = 0
	this.errs = 0
}

func (this *dbLog) Println() {
	fmt.Println(fmt.Sprintf("共请求[%d]次，成功【%d】次数，重复[%d]次数, 错误[%d]次数。", this.countVisit, this.success, this.repeat, this.errs))
}
