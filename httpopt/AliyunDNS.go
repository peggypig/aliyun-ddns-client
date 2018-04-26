/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/26 10:49.
 */
package httpopt

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"aliyun-ddns-client/model"
)

func UpdateRecord(urlAfterSignature string)string  {
	result := ""
	resp , err :=http.Get(urlAfterSignature)
	defer  resp.Body.Close()
	if err != nil {
		log.Println(err)
	}else {
		body , err :=ioutil.ReadAll(resp.Body)
		if err != nil{
			log.Println(err)
		}else {
			updateResp := model.UpdateResp{}
			err :=json.Unmarshal(body,&updateResp)
			if err == nil {
				log.Println(updateResp)
				result += "修改解析记录成功"
			}else {
				result="修改解析记录失败，请求结果如下（请到参照阿里云“修改解析记录错误码”进行修改）："+string(body)
			}
		}
	}
	return  result
}


func GetRecordDesc(urlAfterSignature string)[]model.Record {
	resp , err :=http.Get(urlAfterSignature)
	records := []model.Record{}
	defer  resp.Body.Close()
	if err != nil {
		log.Println(err)
	}else {
		body , err :=ioutil.ReadAll(resp.Body)
		if err != nil{
			log.Println(err)
		}else {
			descResp := model.DescResp{}
			json.Unmarshal(body,&descResp)
			for _,value :=range descResp.DomainRecords.Records {
				records = append(records, value)
			}
		}
	}
	return records
}
