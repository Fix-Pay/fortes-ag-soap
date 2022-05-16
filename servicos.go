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
<AG><Servicos><Servico>
type ServicoFortes struct {
	AG struct {
		Servicos []struct {
			Servico struct {
				Codigo         string  `xml:"Codigo"`
				Nome           string  `xml:"Nome"`
				NomeApresentar string  `xml:"NomeApresentar"`
				Valor          string  `xml:"Valor"`
				Iss            string  `xml:"Iss"`
				Inss           string  `xml:"Inss"`
				Irrf           string  `xml:"Irrf"`
				PisCofinsCsll  string  `xml:"PisCofinsCsll"`
				AliqISS        string  `xml:"AliqISS"`
				AliqINSS       float64 `xml:"AliqINSS"`
				AliqIRRF       float64 `xml:"AliqIRRF"`
				NomeMunicipio  float64 `xml:"NomeMunicipio"`
				AliqPIS        float64 `xml:"AliqPIS"`
				AliqCOFINS     float64 `xml:"AliqCOFINS"`
				AliqCSL        float64 `xml:"AliqCSL"`
			} `xml:"Servico"`
		} `xml:"Servicos"`
	} `xml:"AG"`
}

func (i *iAG) GetListaServicos(request *GetListaServicosRequest, response *EnvelopeResponse) error {
	return i.client.Call("urn:AGIntf-IAG#GetListaServicosXML", request, response)
}
