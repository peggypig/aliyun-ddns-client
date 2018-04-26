/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/26 13:44.
 */
package main

import (
	"os"
	"log"
	"aliyun-ddns-client/common"
	"aliyun-ddns-client/httpopt"
	"aliyun-ddns-client/util"
	"aliyun-ddns-client/model"
	"time"
	"strings"
)


var records []model.Record

var lastIp =""
var nowIp = ""

func main() {
	if len(os.Args)<=1 {
		common.PrintHelp()
		return
	}
	args := os.Args[1:]
	//获取当前出口IP
	nowIp = util.GetIp()
	log.Println("当前出口IP：",nowIp)
	//获取解析
	url := common.GetUrlAfterSignature(common.API_STRING,common.AccessKeySecret,common.GetDescribeDomainRecordsParams())
	records = httpopt.GetRecordDesc(url)
	if len(records)>0 {
		lastIp = records[0].Value
		log.Println("域名"+common.DomainName+"对应IP：",lastIp)
	}
	start := false
	for _ , arg := range args {
		if arg == "-ls" {
			//显示已经存在的解析
			common.PrintRecords(records)
		}
		if arg =="--help" || arg == "-help" || arg == "--h" || arg =="-h" {
			common.PrintHelp()
		}
		if arg == "-start" {
			start = true
		}
	}
	if start {
		log.Println("DDNS已经启动...")
		for {
			nowIp = util.GetIp()
			if nowIp != lastIp {
				log.Println("主机IP发生变化，修改域名解析记录")
				for _,record := range records {
					params := common.GetUpdateDomainRecordParams()
					params["RecordId"]= record.RecordId
					params["RR"]= record.RR
					params["Type"]= record.Type
					params["Value"]= nowIp
					url := common.GetUrlAfterSignature(common.API_STRING,common.AccessKeySecret,params)
					result := httpopt.UpdateRecord(url)
					log.Println("主机记录RR：",record.RR,"的修改结果：",result)
					if strings.Contains(result,"成功") {
						lastIp = nowIp
					}
				}

			}else {
				log.Println("主机出口IP没有发生变化，不需要修改解析记录")
			}
			//睡眠
			time.Sleep(time.Duration(int64(common.CycleTime))*time.Second)
		}
	}
}





