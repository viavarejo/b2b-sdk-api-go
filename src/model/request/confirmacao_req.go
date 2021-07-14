// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    confirmacaoReq, err := UnmarshalConfirmacaoReq(bytes)
//    bytes, err = confirmacaoReq.Marshal()

package request

import "encoding/json"

func UnmarshalConfirmacaoReq(data []byte) (ConfirmacaoReq, error) {
	var r ConfirmacaoReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConfirmacaoReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConfirmacaoReq struct {
	IDCampanha         int64  `json:"idCampanha"`
	IDPedidoParceiro   int64  `json:"idPedidoParceiro"`
	Confirmado         bool   `json:"confirmado"`
	IDPedidoMktplc     string `json:"idPedidoMktplc"`
	Cancelado          bool   `json:"cancelado"`
	MotivoCancelamento string `json:"motivoCancelamento"`
	Parceiro           string `json:"parceiro"`
}
