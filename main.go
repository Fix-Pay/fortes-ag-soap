package fortes_ag_soap

import (
	"encoding/xml"
	"github.com/Fix-Pay/gowsdl/soap"
	"time"
)

var _ time.Time
var _ xml.Name

type EnvelopeResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body
}

type Body struct {
	XMLName                  xml.Name `xml:"Body"`
	GetListaClientesResponse GetListaClientesResponse
}

type IAG interface {
	GetListaClientes(request *GetListaClientesRequest, response *EnvelopeResponse) error
	GetListaServicos(request *GetListaServicosRequest, response *EnvelopeResponse) error
	Call(soapAction string, request, response interface{}) error
}

type iAG struct {
	client *soap.Client
}

func (i *iAG) Call(soapAction string, request, response interface{}) error {
	return i.client.Call(soapAction, request, response)
}

func NewIAG(client *soap.Client) IAG {
	return &iAG{
		client: client,
	}
}
