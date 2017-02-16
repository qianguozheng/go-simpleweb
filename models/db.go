package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

const (
	_Upgrade = iota
	_Control
)

type Upgrade struct {
	Result string `json:result`
	Mac string `json:mac`
	Url string `json:url`
	Md5 string `json:md5`
	Ver string `json:version`
}

type Control struct {
	Mac string `json:mac`
	Switch string `json:switch`
}

type Error struct {
	msg string `json:msg`
}


var db *sql.DB
var m *sync.Mutex
func InitDB(path string) {
	var err error
	db, err = sql.Open("sqlite3", path)
	//defer db.Close()

	if err != nil{
		//return nil, errors.New("Open database failed")
		fmt.Println("Open database failed", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	fmt.Println("Database open success")

	m = new(sync.Mutex)

}

func disableUpgrade(mac string){
	m.Lock()
	_, err :=db.Exec("update airdisk set upgrade=0 where mac=$1", mac)
	m.Unlock()
	if err != nil{
		log.Println(err.Error())
	}
}
func disableControl(mac string)  {
	m.Lock()
	_, err := db.Exec("update airdisk set control=0 where mac=$1", mac)
	m.Unlock()
	if err != nil{
		log.Println(err.Error())
	}
}
func DoJob(mac string, t int) (interface{}, error) {
	//db, err := sql.Open("sqlite3", "./airdisk.db")
	//defer db.Close()
	//
	//if err != nil{
	//	return nil, errors.New("Open database failed")
	//}
	switch t {
	case _Control:
		rows, err := db.Query(fmt.Sprintf("SELECT * FROM airdisk where mac=\"%s\"",mac))

		if err != nil{
			log.Println(err.Error())
			return nil, err
		}

		var Mac string
		var upgrade, control int
		if (rows.Next()) {
			err = rows.Scan(&Mac, &upgrade, &control)
		}
		rows.Close()
		if err != nil{
			log.Println(err.Error())
			return nil, err
		}

		log.Println(mac, upgrade, control)

		if control == 1 {
			// Do repponse control message
			defer disableControl(mac)
			ctlJson := Control{Mac: mac, Switch:"on"}
			return ctlJson, nil
		}
		break

	case _Upgrade:
		rows, err := db.Query(fmt.Sprintf("SELECT * FROM airdisk where mac=\"%s\"",mac))

		if err != nil{
			log.Println(err.Error())
			return nil, err
		}

		var Mac string
		var upgrade, control int
		if (rows.Next()){
			err = rows.Scan(&Mac, &upgrade, &control)
		}
		rows.Close()
		if err != nil{
			log.Println(err.Error())
			return nil, err
		}
		log.Println(mac, upgrade, control)

		if upgrade == 1 {
			//Do upgrade process, return data
			defer disableUpgrade(mac)
			rows, err := db.Query(fmt.Sprintf("SELECT * FROM upgrade where mac=\"%s\"",mac))
			if err != nil{
				log.Println(err.Error())
				return nil, err
			}
			defer rows.Close()
			var Mac, url, version, md5 string

			if (rows.Next()){
				err = rows.Scan(&url, &version, &md5, &Mac)
			}
			if err != nil{
				log.Println(err.Error())
				return nil, err
			}
			/***
				Mac: client judge no wrong return
				Url: firmware address
				Md5: firmware md5
				Ver: firmware version, compare with local version
			 */

			upgJson := Upgrade{Result: "OK",Mac: Mac,Url: url, Ver:version, Md5: md5}

			return upgJson, nil
		} else {
			return nil, nil
		}
		break
	}
	return nil, nil
}
