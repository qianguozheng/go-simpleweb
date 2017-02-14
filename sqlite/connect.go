package sqlite

import "database/sql"
import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"errors"
)

const (
	_Upgrade = iota
	_Control
)

type Upgrade struct {
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

func DoJob(mac string, t int) (interface{}, error) {
	db, err := sql.Open("sqlite3", "./airdisk.db")
	if err != nil{
		return nil, errors.New("Open database failed")
	}

	switch t {
	case _Control:
		rows, err := db.Query(fmt.Sprintf("SELECT * FROM airdisk where mac=\"%s\"",mac))
		//checkErr(err)
		if err != nil{
			fmt.Println(err.Error())
			return nil, err
		}

		var Mac string
		var upgrade, control int
		if (rows.Next()) {
			err = rows.Scan(&Mac, &upgrade, &control)
		}
		if err != nil{
			fmt.Println(err.Error())
			return nil, err
		}

		db.Close()
		fmt.Println(mac, upgrade, control)

		if control == 1 {
			// Do repponse control message

			ctlJson := Control{Mac: mac, Switch:"on"}
			//ctlSerilize, err := json.Marshal(ctlJson)
			//if err != nil{
			//	fmt.Println(err.Error())
			//	return nil, err
			//}
			//return ctlSerilize, nil
			return ctlJson, nil
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
			return nil, err
		}
		var Mac string
		var upgrade, control int
		if (rows.Next()){
			err = rows.Scan(&Mac, &upgrade, &control)
		}
		if err != nil{
			fmt.Println(err.Error())
			db.Close()
			return nil, err
		}
		fmt.Println(mac, upgrade, control)

		if upgrade == 1 {
			//Do upgrade process, return data
			rows, err := db.Query(fmt.Sprintf("SELECT * FROM upgrade where mac=\"%s\"",mac))
			if err != nil{
				fmt.Println(err.Error())
				db.Close()
				return nil, err
			}

			var Mac, url, version, md5 string

			if (rows.Next()){
				err = rows.Scan(&url, &version, &md5, &Mac)
			}
			if err != nil{
				fmt.Println(err.Error())
				db.Close()
				return nil, err
			}
			/***
				Mac: client judge no wrong return
				Url: firmware address
				Md5: firmware md5
				Ver: firmware version, compare with local version
			 */
			db.Close()

			upgJson := Upgrade{Mac: Mac,Url: url, Ver:version, Md5: md5}
			//upgSerilize, err := json.Marshal(upgJson)
			//if err != nil {
			//	fmt.Println(err.Error())
			//	return nil, err
			//}
			//return upgSerilize, nil
			return upgJson, nil
		} else {
			db.Close()
			return nil, nil
		}
		break
	}
	return nil, nil
}

func checkErr(err error)  {
	if nil != err{
		fmt.Println(err.Error())
		//panic(err)
	}
}