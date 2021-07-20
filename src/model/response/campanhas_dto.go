package response

type CampanhasDTO struct {
	Data  []Campanha `json:"data"`
	Error Error      `json:"error"`
}

type Campanha struct {
	IdCampanha     int32           `json:"idCampanha"`
	Nome           string          `json:"nome"`
	DataInicio     CustomTime      `json:"dataInicio"`
	DataFim        CustomTime      `json:"dataFim"`
	IdTipoCampanha int32           `json:"idTipoCampanha"`
	TipoCampanha   string          `json:"tipoCampanha"`
	CnpjContrato   string          `json:"cnpjContrato"`
	Status         bool            `json:"status"`
	EntregaTipos   *[]EntregaTipos `json:"entregaTipos"`
}

type EntregaTipos struct {
	IdEntregaTipo int32  `json:"idEntregaTipo"`
	Nome          string `json:"nome"`
	Habilitado    bool   `json:"habilitado"`
}
