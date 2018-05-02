/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/26 13:58.
 */
package model

import "aliyun-ddns-client/util"

func GetCommonParam() []Param {
	var params []Param = make([]Param,7)
	params[0] = Param{ParamName:"Format",ParamValue:"JSON"}
	params[1] = Param{ParamName:"Version",ParamValue:"2015-01-09"}
	params[2] = Param{ParamName:"AccessKeyId"}
	params[3] = Param{ParamName:"SignatureMethod",ParamValue:"HMAC-SHA1"}
	params[4] = Param{ParamName:"Timestamp",ParamValue:util.GetTimeStr()}
	params[5] = Param{ParamName:"SignatureVersion",ParamValue:"1.0"}
	params[6] = Param{ParamName:"SignatureNonce",ParamValue:util.GetUUID()}
	return params
}




