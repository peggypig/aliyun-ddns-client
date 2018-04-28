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
)


var records []model.Record

var lastIp =""
var nowIp = ""
var dnsIp = ""

func main() {
	if len(os.Args)<=1 {
		common.PrintHelp()
		return
	}
	args := os.Args[1:]
	//获取当前出口IP
	nowIp = util.GetLocalIp()
	lastIp = nowIp
	log.Println("当前出口IP：",nowIp)
	//获取解析
	getDnsInfo()

	start := false
	for _ , arg := range args {
		if arg == "-ls" {
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
		startDDNS()
	}
}

func startDDNS()  {
	log.Println("DDNS已经启动...")
	for {
		getDnsInfo()
		nowIp = util.GetLocalIp()
		if nowIp != lastIp {
			log.Println("主机IP:",nowIp,"发生变化，上次IP：",lastIp,"，修改域名解析记录")
			updateDomainRecord()
			lastIp = nowIp
		}else if nowIp != dnsIp {
			log.Println("主机IP:",nowIp,"和解析记录IP:",dnsIp,"不一致，修改域名解析记录")
			updateDomainRecord()
		}else {
			log.Println("主机出口IP没有发生变化并且和DNS解析保持一致，不需要修改解析记录")
		}
		//睡眠
		time.Sleep(time.Duration(int64(common.CycleTime))*time.Second)
	}
}

func updateDomainRecord()  {
	for _,record := range records {
		params := common.GetUpdateDomainRecordParams()
		params["RecordId"]= record.RecordId
		params["RR"]= record.RR
		params["Type"]= record.Type
		params["Value"]= nowIp
		url := common.GetUrlAfterSignature(common.API_STRING,common.AccessKeySecret,params)
		result := httpopt.UpdateRecord(url)
		log.Println("主机记录RR：",record.RR,"的修改结果：",result)
	}
}




func getDnsInfo()  {
	url := common.GetUrlAfterSignature(common.API_STRING,common.AccessKeySecret,common.GetDescribeDomainRecordsParams())
	records = httpopt.GetRecordDesc(url)
	if len(records)>0 {
		dnsIp = records[0].Value
		log.Println("域名"+common.DomainName+"对应IP：",dnsIp)
	}else {
		log.Println("没有解析记录，DDNS不会生效，前往阿里云DNS添加相关解析")
	}
}




