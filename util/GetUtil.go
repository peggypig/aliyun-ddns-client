/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/26 13:49.
 */
package util

import (
	"log"
	"fmt"
	"time"
	"github.com/satori/go.uuid"
	"crypto/hmac"
	"crypto/sha1"
	"github.com/Unknwon/com"
	"os/exec"
	"strings"
	"regexp"
)





func GetTimeStr() string {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	timeStr := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ",
		year, mon, day, hour, min, sec)
	return timeStr
}

func GetUUID() string {
	uu ,_ := uuid.NewV4()
	return  uu.String()
}



func CalcHMACSHA(key string,content string) string {
	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(content))
	return  com.Base64Encode(string(mac.Sum(nil)))           //hex.EncodeToString(mac.Sum(nil))
}

func GetLocalIp()  string{
	ip := ""
	var cmd *exec.Cmd
	cmd = exec.Command("curl", "-s" ,"ip.cn")
	var strBody  = ""
	if curlBody , err := cmd.Output();err != nil {
		log.Println(err)
	}else {
		strBody = string(curlBody)
		ipReg := regexp.MustCompile("(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)")
		ip = ipReg.FindStringSubmatch(strBody)[0]
	}
	return strings.Trim(ip," ")
}



