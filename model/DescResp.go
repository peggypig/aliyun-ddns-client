/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/26 14:59.
 */
package model

type DescResp struct {
	PageNumber    string `json:"PageNumber,string"`
	TotalCount    string `json:"TotalCount,string"`
	PageSize      string `json:"PageSize,string"`
	RequestId     string `json:"RequestId"`
	DomainRecords Domain `json:"DomainRecords"`
}
