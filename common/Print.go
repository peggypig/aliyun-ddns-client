/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/26 15:54.
 */
package common

import (
	"aliyun-ddns-client/model"
	"log"
)

func PrintRecords(records []model.Record)  {
	for index,record :=range records {
		log.Println("==========记录",index+1,"==========")
		log.Println("记录编号：\t",record.RecordId)
		log.Println("记录值：\t",record.Value)
		log.Println("记录类型：\t",record.Type)
		log.Println("记录线路：\t",record.Line)
		log.Println("记录主机：\t",record.RR)
	}
}
func PrintHelp()  {
	log.Println("----------------------aliyun-ddns-client 帮助指南----------------------")
	log.Println("参数\t\t\t\t\t\t\t\t\t说明")
	log.Println("-h 或 -help 或 --h 或 --help\t\t\t显示帮助")
	log.Println("-ls\t\t\t\t\t\t\t\t\t显示已经存在的解析")
	log.Println("-start\t\t\t\t\t\t\t\t开启DDNS")
	log.Println("----------------------aliyun-ddns-client 帮助指南----------------------")
}

