package session

import (
	"Streaming_websites/api/dbops"
	"Streaming_websites/api/defs"
	"Streaming_websites/api/utils"
	"log"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

func GenerateNewSessionId(uname string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000 // Severside session valid time : 30 min
	log.Printf("The session's ttl is %d\n", ttl)
	ss := &defs.SimpleSession{Username: uname, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, uname)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, exist := sessionMap.Load(sid)
	if exist {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			return "", true
		} else {
			return ss.(*defs.SimpleSession).Username, false
		}
	}
	return "", true
}

func DeleteExpiredSession(sid string) {
	if _, exist := sessionMap.Load(sid); exist {
		if _, expired := IsSessionExpired(sid); expired {
			sessionMap.Delete(sid)
			dbops.DeleteSession(sid)
		}
	}

}
