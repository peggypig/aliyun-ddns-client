/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/26 14:25.
 */
package model

func GetUpdateDomainRecordParams() []Param  {
	var params []Param = make([]Param,8)
	params[0] =Param{ParamName:"Action",ParamValue:"UpdateDomainRecord"}
	params[1] =Param{ParamName:"RecordId"}
	params[2] =Param{ParamName:"RR",ParamValue:"@"}
	params[3] =Param{ParamName:"Type"}
	params[4] =Param{ParamName:"Value"}
	params[5] =Param{ParamName:"TTL",ParamValue:"600"}
	params[6] =Param{ParamName:"Priority",ParamValue:"10"}
	params[7] =Param{ParamName:"Line",ParamValue:"default"}
	return params
}
