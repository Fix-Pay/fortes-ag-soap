package fortes_ag_soap

import (
	"encoding/xml"
	"time"
)

var _ time.Time
var _ xml.Name

type RequestGetListaServicos struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/encoding/ soap:encodingStyle"`
	EncodingStyle string   `xml:"soap:encodingStyle,http://schemas.xmlsoap.org/soap/encoding/"`
}

type GetListaServicosRequest struct {
	XMLName       xml.Name `xml:"urn:getListaServicos"`
	EncodingStyle string   `xml:"soap:encodingStyle,attr"`
}

type GetListaServicosResponse struct {
	XMLName xml.Name `xml:"getListaServicosResponse"`
	Return  string   `xml:"return"`
}

type ServicoFortes struct {
	Codigo         string  `json:"codigo"`
	Nome           string  `json:"nome"`
	NomeApresentar string  `json:"nome_apresentar"`
	Valor          float64 `json:"valor"`
	Iss            int     `json:"iss"`
	Inss           int     `json:"inss"`
	Irrf           int     `json:"irrf"`
	PisCofinsCsll  int     `json:"pis_cofins_csll"`
	AliqISS        float64 `json:"aliq_iss"`
	AliqINSS       float64 `json:"aliq_inss"`
	AliqIRRF       float64 `json:"aliq_irrf"`
	NomeMunicipio  float64 `json:"nomemunicipio"`
	AliqPIS        float64 `json:"aliq_pis"`
	AliqCOFINS     float64 `json:"aliq_cofins"`
	AliqCSL        float64 `json:"aliq_csl"`
}

func (i *iAG) GetListaServicos(request *GetListaServicosRequest, response *EnvelopeResponse) error {
	return i.client.Call("urn:AGIntf-IAG#getListaServicosFull", request, response)
}
