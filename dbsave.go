package main

import (
	"fmt"
	"os"
	"vava6/mysql"
)

const (
	IS_OK        int8 = 0
	IS_EXISTENCE int8 = 1
	IS_ERROR     int8 = 2
	IS_NULL      int8 = 3
)

var DBSave *mysql.DBs

func init() {
	err := LinkDBServer()
	if err != nil {
		os.Exit(1)
	}
}

func LinkDBServer() error {
	var err error

	// 加载配置文件
	cfg := mysql.NewCfg()
	cfg.DBName = "joke"
	cfg.Address = "119.23.235.220"
	cfg.User = "joke"
	cfg.Pass = "jokeUser&2018"

	fmt.Print("【正在连接数据库...")
	DBSave, err = mysql.NewDBs(
		cfg.DBName,
		cfg.Address,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.MaxConn,
		cfg.MinConn,
	)
	if err != nil {
		fmt.Print("连接数据库失败-", err.Error())
	} else {
		fmt.Print("连接数据库成功")
	}
	fmt.Println("】")
	return err
}

func SaveJoke(tableName string, db *JokeDB) (int8, error) {
	_, err := DBSave.Insert(tableName, map[string]interface{}{
		"id":       db.ID,
		"content":  db.Content,
		"keywords": db.Keywords,
		"vote":     db.Vote,
		"comment":  db.Comment,
	})
	if err != nil {
		return IS_ERROR, err
	}
	return IS_OK, nil
}

func CheckJoke(tableName string, id int) (int8, error) {
	// 判断是否有这个笑话存在
	if rss, err := DBSave.Querys("id", tableName, fmt.Sprint("id=", id)); err != nil {
		return IS_ERROR, err
	} else {
		if len(rss) > 0 {
			return IS_EXISTENCE, nil
		}
	}
	return IS_NULL, nil
}
