package dbops

import (
	"Streaming_websites/api/defs"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?,?,?)")
	if err != nil {
		return err
	}

	fmt.Printf("The session insert SQL hat prepared")
	_, err = stmtIns.Exec(sid, ttlStr, uname)
	if err != nil {
		return err
	}

	fmt.Print("session insert success!")

	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id =?")
	if err != nil {
		return nil, err
	}

	var ttl string
	var uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSession() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var (
			id     string
			ttlstr string
			uname  string
		)

		if err = rows.Scan(&id, &ttlstr, &uname); err != nil {
			log.Printf("retrieve sessions error: %s ", err)
			break
		}

		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			ss := &defs.SimpleSession{Username: uname, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id: %s, ttl: %d ", id, ss.TTL)
		}
	}

	defer stmtOut.Close()
	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _, err = stmtOut.Query(sid); err != nil {
		return err
	}

	defer stmtOut.Close()
	return nil
}
