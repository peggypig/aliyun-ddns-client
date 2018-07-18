/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/25 17:29.
 */
package common

import (
	"github.com/Unknwon/com"
	"aliyun-ddns-client/model"
	"aliyun-ddns-client/util"
)

func GetUrlAfterSignature(api , privateKey string ,params map[string]string) (url string) {
	//存储参数的键
	var  keys []string  = make([]string,len(params))

	flag :=0
	for key , _ := range params {
		keys[flag] = key
		flag ++
	}
	paramsNum := len(keys)
	//根据参数名称升序
	for i := 0;i<paramsNum-1 ;i++  {
		for j := i+1; j<paramsNum; j++ {
			if keys[i] >keys[j] {
				temp := keys[i]
				keys[i] = keys[j]
				keys[j] = temp
			}

		}
	}

	//构建排序好的参数列表
	var newParams []model.Param = make([]model.Param , len(keys))
	for index , value := range keys{
		param  := model.Param{}
		param.ParamName = com.UrlEncode(value)
		param.ParamValue = com.UrlEncode(params[value])
		newParams[index] = param
	}

	canonicalizedQueryString := ""
	for index,value := range newParams {
		if index != len(newParams)-1 {
			canonicalizedQueryString += value.ParamName+EQUAL_STRING+value.ParamValue+HTTP_AND_STRING
		}else {
			canonicalizedQueryString += value.ParamName+EQUAL_STRING+value.ParamValue
		}
	}


	stringToSign := HTTP_METHOD_GET_STRING + HTTP_AND_STRING +com.UrlEncode(HTTP_PATH_STRING) +HTTP_AND_STRING + com.UrlEncode(canonicalizedQueryString)


	httpStr := api + HTTP_PARAM_STRING

	for _,value := range newParams {
		httpStr += value.ParamName+EQUAL_STRING+value.ParamValue+HTTP_AND_STRING
	}
	httpStr +=SIGNATURE_STRING+EQUAL_STRING+com.UrlEncode(util.CalcHMACSHA(privateKey+HTTP_AND_STRING,stringToSign))

	return  httpStr
}
