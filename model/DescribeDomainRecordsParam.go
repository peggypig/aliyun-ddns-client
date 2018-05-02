/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/26 14:04.
 */
package model

func GetDescribeDomainRecordsParam() []Param {
	var params []Param = make([]Param,7)
	params[0] = Param{ParamName:"Action",ParamValue:"DescribeDomainRecords"}
	params[1] = Param{ParamName:"DomainName"}
	return params
}



