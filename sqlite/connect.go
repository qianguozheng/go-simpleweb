package sqlite

import "database/sql"
import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"encoding/json"
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

func DoJob(mac string, t int, w http.ResponseWriter)  {
	switch t {
	case _Control:
		db, err := sql.Open("sqlite3", "./airdisk.db")
		checkErr(err)
		//stmt, err := db.Prepare("SELECT * FROM airdisk where mac=?")
		rows, err := db.Query(fmt.Sprintf("SELECT * FROM airdisk where mac=\"%s\"",mac))
		//checkErr(err)
		if err != nil{
			fmt.Println(err.Error())
			return
		}

		var Mac string
		var upgrade, control int
		if (rows.Next()) {
			err = rows.Scan(&Mac, &upgrade, &control)
		}
		checkErr(err)
		if control == 1 {
			stmt, err := db.Prepare("update airdisk set control where mac=?")
			checkErr(err)
			res, err := stmt.Exec(0, mac)
			checkErr(err)
			affect, err := res.RowsAffected()
			checkErr(err)
			//fmt.Println(affect)
			if affect != 1 {
				fmt.Println("Update failed", affect)
			}
		}
		db.Close()
		fmt.Println(mac, upgrade, control)

		if control == 1 {
			// Do repponse control message

			ctlJson := Control{Mac: mac, Switch:"on"}
			ctlSerilize, err := json.Marshal(ctlJson)
			if err != nil{
				fmt.Println(err.Error())
				return
			}
			w.Write(ctlSerilize)
		}
		break

	case _Upgrade:
		db, err := sql.Open("sqlite3", "./airdisk.db")
		checkErr(err)
		rows, err := db.Query(fmt.Sprintf("SELECT * FROM airdisk where mac=\"%s\"",mac))
		//rows, err := db.Query("SELECT * FROM airdisk where mac=\"hello\"")
		//rows, err := db.Query("SELECT * FROM airdisk")
		//checkErr(err)
		if err != nil{
			fmt.Println(err.Error())
			db.Close()
			return
		}
		var Mac string
		var upgrade, control int
		if (rows.Next()){
			err = rows.Scan(&Mac, &upgrade, &control)
		}
		checkErr(err)
		fmt.Println(mac, upgrade, control)

		if upgrade == 1 {
			//Do upgrade process, return data
			rows, err := db.Query(fmt.Sprintf("SELECT * FROM upgrade where mac=\"%s\"",mac))
			if err != nil{
				fmt.Println(err.Error())
				db.Close()
				return
			}

			var Mac, url, version, md5 string

			if (rows.Next()){
				err = rows.Scan(&url, &version, &md5, &Mac)
			}
			checkErr(err)

			stmt, err := db.Prepare("update airdisk set upgrade where mac=?")
			checkErr(err)
			res, err := stmt.Exec(0, mac)
			checkErr(err)
			affect, err := res.RowsAffected()
			checkErr(err)
			//fmt.Println(affect)
			if affect != 1 {
				fmt.Println("Update failed", affect)
			}

			db.Close()
			/***
				Mac: client judge no wrong return
				Url: firmware address
				Md5: firmware md5
				Ver: firmware version, compare with local version
			 */


			upgJson := Upgrade{Result: "OK",Mac: Mac,Url: url, Ver:version, Md5: md5}
			upgSerilize, err := json.Marshal(upgJson)
			if err != nil{
				fmt.Println(err.Error())
				return
			}
			w.Write(upgSerilize)
		} else {
			db.Close()
		}

		break
	}
}

func checkErr(err error)  {
	if nil != err{
		fmt.Println(err.Error())
		//panic(err)
	}
}