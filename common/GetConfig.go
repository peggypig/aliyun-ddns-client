/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/26 14:13.
 */
package common

import (
	"github.com/Unknwon/goconfig"
	"log"
	"strconv"
	"aliyun-ddns-client/model"
)

var CycleTime int

var AccessKeyId string
var AccessKeySecret string
var DomainName string

var RecordConfig model.Record


func init()  {
	cf , err := goconfig.LoadConfigFile("./ddns.conf")
	if err != nil {
		panic("没有找到./ddns.conf的配置文件")
	}else {
		accessKeyMap,err  :=cf.GetSection("AccessKey")
		if err != nil {
			panic("没有AccessKey相关配置")
		}else {
			if value , ok :=  accessKeyMap["AccessKeyId"];ok {
				AccessKeyId=value
			}else {
				panic("没有AccessKeyId相关配置")
			}
			if value , ok := accessKeyMap["AccessKeySecret"];ok {
				AccessKeySecret = value
			}else {
				panic("没有AccessKeySecret相关配置")
			}
		}

		paramMap , err := cf.GetSection("Param")
		if err != nil {
			log.Println("没有Param相关配置,将采用原始配置信息")
		}else {
			if value , ok := paramMap["RecordId"];ok {
				RecordConfig.RecordId = value;
			}
			if value , ok := paramMap["RR"];ok {
				RecordConfig.RR = value;
			}
			if value , ok := paramMap["Type"];ok {
				RecordConfig.Type = value;
			}
			if value , ok := paramMap["Value"];ok {
				RecordConfig.Value = value;
			}

			if value , ok := paramMap["TTL"];ok {
				RecordConfig.TTL= value;
			}
			if value , ok := paramMap["Priority"];ok {
				RecordConfig.Priority = value;
			}
			if value , ok := paramMap["Line"];ok {
				RecordConfig.Line = value;
			}
		}
		localConfigMap,err := cf.GetSection("Local")
		if err != nil {
			log.Println(err)
		}else {

			if value ,ok := localConfigMap["CycleTime"];ok {
				CycleTime, err  = strconv.Atoi(value)
				if err != nil {
					log.Println(err)
					CycleTime = 300
				}
			}
			if value ,ok := localConfigMap["DomainName"];ok {
				DomainName = value
			}else {
				panic("Local下面缺少DomainName的配置")
			}
		}
	}
}
