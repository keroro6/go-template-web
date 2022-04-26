package middleware

import (
	"bytes"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net"
	"net/http"
)

var LocalIp string

type addr struct {
	RemoteIP string
	LocalIP  string
}

type bodyLogWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	//w.WriteString()
	//return w.Write(b)
	return w.ResponseWriter.Write(b)
}

//func (w bodyLogWriter) WriteString(s string) (int, error) {
//	w.body.WriteString(s)
//	return w.WriteString(s)
//}
func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Errorf("Oops, InterfaceAddrs err:\n%v", err)
	}
	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				LocalIp = ipNet.IP.String()
				return
			}
		}
	}
}
func getClientIp() {
	//ip := exnet.ClientPublicIP(r)
	//if ip == "" {
	//	ip = exnet.ClientIP(r)
	//}
}

func TraceLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//start := time.Now()
		//path := r.URL.Path
		reqBuf, _ := ioutil.ReadAll(r.Body)
		reqBuf1 := ioutil.NopCloser(bytes.NewBuffer(reqBuf))
		reqBuf2 := ioutil.NopCloser(bytes.NewBuffer(reqBuf))
		r.Body = reqBuf2
		reqBody, _ := ioutil.ReadAll(reqBuf1)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: w}
		w = blw
		next(w, r)
		//end := time.Now()
		//latency := end.Sub(start)

		//m := map[string]interface{}{
		//	"elapse":   latency.Nanoseconds() / int64(time.Millisecond),
		//	"path":     path,
		//	//"body":     string(reqBody),
		//	"response": blw.body.String(),
		//}
		//bd, _ := json.Marshal(m)
		logx.Infof("request:%s response:%s", string(reqBody), string(blw.body.String()))
	}
}
