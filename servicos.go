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
	XMLName       xml.Name `xml:"urn:GetListaServicosFull"`
	EncodingStyle string   `xml:"soap:encodingStyle,attr"`
}

type GetListaServicosResponse struct {
	XMLName xml.Name `xml:"GetListaServicosFullResponse"`
	Return  string   `xml:"return"`
}

type ServicoFortes struct {
	Item []string `xml:"item"`
}

func (i *iAG) GetListaServicos(request *GetListaServicosRequest, response *EnvelopeResponse) error {
	return i.client.Call("urn:AGIntf-IAG#GetListaServicosFull", request, response)
}
