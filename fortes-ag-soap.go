package fortes_ag_soap

import (
	"encoding/xml"
	"github.com/Fix-Pay/gowsdl/soap"
	"time"
)

var _ time.Time
var _ xml.Name

type RequestGetListaClientes struct {
	XMLName              xml.Name `xml:"http://schemas.xmlsoap.org/soap/encoding/ soap:encodingStyle"`
	EncodingStyle        string   `xml:"soap:encodingStyle,http://schemas.xmlsoap.org/soap/encoding/"`
	ConsiderarCancelados bool     `xml:"xsi:type ConsiderarCancelados"`
}

type GetListaClientesRequest struct {
	XMLName       xml.Name `xml:"urn:getListaClientes"`
	EncodingStyle string   `xml:"soap:encodingStyle,attr"`

	ConsiderarCancelados bool `xml:"ConsiderarCancelados,omitempty"`
}

type EnvelopeResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body
}

type Body struct {
	XMLName                  xml.Name `xml:"Body"`
	GetListaClientesResponse GetListaClientesResponse
}

type GetListaClientesResponse struct {
	XMLName xml.Name `xml:"getListaClientesResponse"`
	Return  string   `xml:"return"`
}

type Client struct {
	Codigo         string `json:"codigo"`
	Cnpj           string `json:"cnpj"`
	RazaoSocial    string `json:"razao social"`
	IM             string `json:"im"`
	EndLogradouro  string `json:"endlogradouro"`
	EndComplemento string `json:"endcomplemento"`
	EndNumero      string `json:"endnumero"`
	EndBairro      string `json:"endbairro"`
	EndCep         string `json:"endcep"`
	EndUF          string `json:"enduf"`
	EndMunicipio   string `json:"endmunicipio"`
	NomeMunicipio  string `json:"nomemunicipio"`
	Telefone       string `json:"telefone"`
	Email          string `json:"email"`
}

type IAGExample interface {
	GetListaClientesExample(request *GetListaClientesRequest, response *EnvelopeResponse) error
	Call(soapAction string, request *interface{}, response *interface{}) error
}

type iAGExample struct {
	client *soap.Client
}

func (i *iAGExample) GetListaClientesExample(request *GetListaClientesRequest, response *EnvelopeResponse) error {
	return i.client.Call("urn:AGIntf-IAG#getListaClientes", request, response)
}

func (i *iAGExample) Call(soapAction string, request *interface{}, response *interface{}) error {
	return i.client.Call(soapAction, request, response)
}

func NewIAGExample(client *soap.Client) IAGExample {
	return &iAGExample{
		client: client,
	}
}
