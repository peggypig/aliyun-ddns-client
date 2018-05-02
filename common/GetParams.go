/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/26 14:24.
 */
package common

import (
	"aliyun-ddns-client/model"
)

func GetDescribeDomainRecordsParams() map[string]string {
	var params map[string]string = make(map[string]string)
	commonParam := model.GetCommonParam()
	descParam := model.GetDescribeDomainRecordsParam()
	for _, value :=range commonParam {
		if value.ParamName == "AccessKeyId" {
			params[value.ParamName] = AccessKeyId
		}else {
			params[value.ParamName] = value.ParamValue
		}
	}
	for _,value := range descParam {
		if value.ParamName == "DomainName" {
			params[value.ParamName] = DomainName
		}else {
			params[value.ParamName] = value.ParamValue
		}
	}
	return params
}


func GetUpdateDomainRecordParams() map[string]string {
	var params map[string]string = make(map[string]string)
	commonParam := model.GetCommonParam()
	updateParam := model.GetUpdateDomainRecordParams()
	for _, value :=range commonParam {
		if value.ParamName == "AccessKeyId" {
			params[value.ParamName] = AccessKeyId
		}else {
			params[value.ParamName] = value.ParamValue
		}
	}
	for _,value := range updateParam {
		params[value.ParamName] = value.ParamValue
	}
	if len(RecordConfig.RecordId) > 0 {
		params["RecordId"] = RecordConfig.RecordId
	}
	if len(RecordConfig.RR) > 0 {
		params["RR"] = RecordConfig.RR
	}
	if len(RecordConfig.Type) > 0 {
		params["Type"] = RecordConfig.Type
	}
	if len(RecordConfig.Value) > 0 {
		params["Value"] = RecordConfig.Value
	}
	if len(RecordConfig.TTL) > 0 {
		params["TTL"] = RecordConfig.TTL
	}
	if len(RecordConfig.Priority) > 0 {
		params["Priority"] = RecordConfig.Priority
	}
	if len(RecordConfig.Line) > 0 {
		params["Line"] = RecordConfig.Line
	}
	return params
}
