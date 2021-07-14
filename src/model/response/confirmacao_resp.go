// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    confirmacaoResp, err := UnmarshalConfirmacaoResp(bytes)
//    bytes, err = confirmacaoResp.Marshal()

package response

import "encoding/json"

func UnmarshalConfirmacaoResp(data []byte) (ConfirmacaoResp, error) {
	var r ConfirmacaoResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConfirmacaoResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConfirmacaoResp struct {
	Data  Data  `json:"data"`
	Error Error `json:"error"`
}

type Data struct {
	PedidoConfirmado bool `json:"pedidoConfirmado"`
	PedidoCancelado  bool `json:"pedidoCancelado"`
}

type Error struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Fields  []Field `json:"fields"`
}

type Field struct {
	Field   string `json:"field"`
	Value   string `json:"value"`
	Message string `json:"message"`
}
