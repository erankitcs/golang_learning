package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
	//fmt.Println(string(data))
	//return 1, nil
}

func Run(destination string) {
	log = stlog.New(fileLog(destination), "", stlog.LstdFlags)
}

func RegisterHandler() {
	http.HandleFunc("/log", func(rw http.ResponseWriter, r *http.Request) {
		msg, err := ioutil.ReadAll(r.Body)
		if err != nil && len(msg) == 0 {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		write(string(msg))
	})
}

func write(msg string) {
	log.Printf("%v\n", msg)
}
