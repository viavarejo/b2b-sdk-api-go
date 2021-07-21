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
	IDCampanha         int64  `json:"idCampanha,omitempty"`
	IDPedidoParceiro   int64  `json:"idPedidoParceiro,omitempty"`
	Confirmado         bool   `json:"confirmado,omitempty"`
	IDPedidoMktplc     string `json:"idPedidoMktplc,omitempty"`
	Cancelado          bool   `json:"cancelado,omitempty"`
	MotivoCancelamento string `json:"motivoCancelamento,omitempty"`
	Parceiro           string `json:"parceiro,omitempty"`
}
