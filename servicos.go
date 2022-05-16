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
	XMLName       xml.Name `xml:"urn:GetListaServicosXML"`
	EncodingStyle string   `xml:"soap:encodingStyle,attr"`
}

type GetListaServicosResponse struct {
	XMLName xml.Name `xml:"GetListaServicosXMLResponse"`
	Return  string   `xml:"return"`
}

type ServicoFortes struct {
	Codigo         string  `xml:"codigo"`
	Nome           string  `xml:"nome"`
	NomeApresentar string  `xml:"nome_apresentar"`
	Valor          string  `xml:"valor"`
	Iss            string  `xml:"iss"`
	Inss           string  `xml:"inss"`
	Irrf           string  `xml:"irrf"`
	PisCofinsCsll  string  `xml:"pis_cofins_csll"`
	AliqISS        string  `xml:"aliq_iss"`
	AliqINSS       float64 `xml:"aliq_inss"`
	AliqIRRF       float64 `xml:"aliq_irrf"`
	NomeMunicipio  float64 `xml:"nomemunicipio"`
	AliqPIS        float64 `xml:"aliq_pis"`
	AliqCOFINS     float64 `xml:"aliq_cofins"`
	AliqCSL        float64 `xml:"aliq_csl"`
}

func (i *iAG) GetListaServicos(request *GetListaServicosRequest, response *EnvelopeResponse) error {
	return i.client.Call("urn:AGIntf-IAG#GetListaServicosXML", request, response)
}
