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
	XMLName  xml.Name `xml:"AG"`
	Servicos Servicos `xml:"Servicos"`
}

type Servicos struct {
	XMLName xml.Name  `xml:"Servicos"`
	Servico []Servico `xml:"Servico"`
}

type Servico struct {
	XMLName        xml.Name `xml:"Servico"`
	Codigo         string   `xml:"Codigo"`
	Nome           string   `xml:"Nome"`
	NomeApresentar string   `xml:"NomeApresentar"`
	Valor          string   `xml:"Valor"`
	ISS            string   `xml:"ISS"`
	INSS           string   `xml:"INSS"`
	IRRF           string   `xml:"IRRF"`
	PISCOFINSCSLL  string   `xml:"PISCOFINSCSLL"`
	AliqISS        string   `xml:"AliqISS"`
	AliqINSS       string   `xml:"AliqINSS"`
	AliqIRRF       string   `xml:"AliqIRRF"`
	AliqPIS        string   `xml:"AliqPIS"`
	AliqCOFINS     string   `xml:"AliqCOFINS"`
	AliqCSL        string   `xml:"AliqCSL"`
}

func (i *iAG) GetListaServicos(request *GetListaServicosRequest, response *EnvelopeResponse) error {
	return i.client.Call("urn:AGIntf-IAG#GetListaServicosXML", request, response)
}
