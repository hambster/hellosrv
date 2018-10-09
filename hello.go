package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

const (
	// LogDir ...
	LogDir = "/hello_log"
	// DefaultAddr ...
	DefaultAddr = "127.0.0.1"
	// LogDelay ...
	LogDelaySec = 10
)

func logMsg(msg ...interface{}) {
	fmt.Printf("%s\n", msg...)
}

func getLogPath() string {
	ret := "127.0.0.1"
	if ifaces, err := net.Interfaces(); nil == err {
	outer:
		for _, iface := range ifaces {
			if addrs, err := iface.Addrs(); nil == err {
				for _, addr := range addrs {
					tmp := addr.String()
					fmt.Printf("[?] tmp - %s\n", tmp)
					if -1 != strings.Index(tmp, ":") {
						continue
					}

					slashIdx := strings.Index(tmp, "/")
					if -1 == slashIdx {
						continue
					}
					tmp = string(tmp[0:slashIdx])
					fmt.Printf("[-] tmp - %s\n", tmp)
					if tmp != DefaultAddr {
						fmt.Printf("[v] tmp - %s\n", tmp)
						ret = tmp
						break outer
					}
				}
			}
		}
	}

	ret = ret + ".log"
	return ret
}

// LogWriter ...
func LogWriter() {
	logPath := path.Join(LogDir, getLogPath())
	fileHandle, err := os.Create(logPath)
	if nil != err {
		logMsg("failed to open file - %s", err.Error())
		return
	}

	cnt := 0
	for {
		cnt++
		select {
		case <-time.After(time.Second * LogDelaySec):
			fileHandle.WriteString(fmt.Sprintf("dummy log %d\n", cnt))
		}
	}
}

func main() {
	go LogWriter()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello~\n"))
	})

	http.ListenAndServe(":8080", nil)
}
