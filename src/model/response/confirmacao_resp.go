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
	Data  Confirmacao `json:"data"`
	Error Error       `json:"error"`
}

type Confirmacao struct {
	PedidoConfirmado bool `json:"pedidoConfirmado"`
	PedidoCancelado  bool `json:"pedidoCancelado"`
}
