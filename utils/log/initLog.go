package log

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log = logrus.New()

type bodyLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func SetupLogger() {
	Log.Println("Setup Logger start")
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash("./log/log"),
		MaxSize:    5,
		MaxBackups: 10,
		MaxAge:     20,
		Compress:   true,
	}

	multiwriter := io.MultiWriter(lumberjackLogger, os.Stderr)
	formatter := new(logrus.JSONFormatter)

	formatter.TimestampFormat = "02-01-2006 15:05:05"
	Log.SetFormatter(formatter)
	Log.SetOutput(multiwriter)
	Log.Println("Setup Log Finish")
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	return buf.String()

}

func (wr bodyLogWrite) Write(b []byte) (int, error) {
	wr.body.Write(b)
	return wr.ResponseWriter.Write(b)
}

func WriteLogReq(c *gin.Context) {
	if c.Request.Method == "POST" || c.Request.Method == "GET" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
		buf, err := io.ReadAll(c.Request.Body)
		if err != nil {
			Log.Println("Error read all", err)
		}

		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))

		re := regexp.MustCompile(`\r?\n`)
		var request = re.ReplaceAllString(readBody(rdr1), "")
		if strings.Contains(c.FullPath(), "login") {
			jsonMap := make(map[string]interface{})
			json.Unmarshal([]byte(request), &jsonMap)
			delete(jsonMap, "password")
			req, _ := json.Marshal(jsonMap)
			request = string(req)
		}
		Log.WithFields(logrus.Fields{
			"logType":     "Request",
			"url":         c.Request.URL.Path,
			"method":      c.Request.Method,
			"requestId":   requestid.Get(c),
			"userAgent":   c.Request.UserAgent(),
			"requestBody": request,
		}).Info()
		c.Request.Body = rdr2
	} else {
		if !(c.FullPath() == "/") {
			Log.WithFields(logrus.Fields{
				"logType":   "Request",
				"url":       c.Request.URL.Path,
				"method":    c.Request.Method,
				"requestId": requestid.Get(c),
				"userAgent": c.Request.UserAgent(),
			}).Info()
		}
	}
}

func WriteLogResp(c *gin.Context, resp string) {
	latency, _ := c.Get("Latency")
	if !(c.FullPath() == "/") {
		RespString := ""
		if len(resp) <= 300 {
			RespString = resp
		} else {
			RespString = strings.TrimSpace(resp[:300])
		}
		Log.WithFields(logrus.Fields{
			"logType":      "Response",
			"url":          c.Request.URL.Path,
			"method":       c.Request.Method,
			"requestId":    requestid.Get(c),
			"userAgent":    c.Request.UserAgent(),
			"latency":      latency.(string),
			"responseBody": RespString,
		}).Info()
	}
}

func RequestLoggerActivity() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		WriteLogReq(c)
		blw := &bodyLogWrite{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		dur := time.Since(t)
		c.Set("Latency", dur.String())
		WriteLogResp(c, blw.body.String())

	}
}
