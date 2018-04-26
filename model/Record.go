/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/26 14:08.
 */
package model

type Record struct {
	RecordId string `json:"RecordId"`
	RR       string `json:"RR"`
	Type     string `json:"Type"`
	Value    string `json:"Value"`
	TTL      string `json:"-"`
	Priority string `json:"-"`
	Line     string `json:"Line"`
}