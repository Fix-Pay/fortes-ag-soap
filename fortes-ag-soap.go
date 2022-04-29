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

type TCliente struct {
	Codigo string `xml:"codigo,omitempty" json:"codigo,omitempty"`
	Nome string `xml:"nome,omitempty" json:"nome,omitempty"`
	NomeFantasia string `xml:"nomeFantasia,omitempty" json:"nomeFantasia,omitempty"`
	Cnpjcpf string `xml:"cnpjcpf,omitempty" json:"cnpjcpf,omitempty"`
	Cep string `xml:"cep,omitempty" json:"cep,omitempty"`
	Logradouro string `xml:"logradouro,omitempty" json:"logradouro,omitempty"`
	Bairro string `xml:"bairro,omitempty" json:"bairro,omitempty"`
	Numero string `xml:"numero,omitempty" json:"numero,omitempty"`
	Complemento string `xml:"complemento,omitempty" json:"complemento,omitempty"`
	Uf string `xml:"uf,omitempty" json:"uf,omitempty"`
	Cidade string `xml:"cidade,omitempty" json:"cidade,omitempty"`
	Ddd string `xml:"ddd,omitempty" json:"ddd,omitempty"`
	Fone string `xml:"fone,omitempty" json:"fone,omitempty"`
	DddFax string `xml:"dddFax,omitempty" json:"dddFax,omitempty"`
	Fax string `xml:"fax,omitempty" json:"fax,omitempty"`
	Identificador string `xml:"identificador,omitempty" json:"identificador,omitempty"`
	DddCelular string `xml:"dddCelular,omitempty" json:"dddCelular,omitempty"`
	Celular string `xml:"celular,omitempty" json:"celular,omitempty"`
	Email string `xml:"email,omitempty" json:"email,omitempty"`
	Contato string `xml:"contato,omitempty" json:"contato,omitempty"`
	Site string `xml:"site,omitempty" json:"site,omitempty"`
	Rg string `xml:"rg,omitempty" json:"rg,omitempty"`
	Ie string `xml:"ie,omitempty" json:"ie,omitempty"`
	Grupo string `xml:"grupo,omitempty" json:"grupo,omitempty"`
	AgenteCobrador string `xml:"agenteCobrador,omitempty" json:"agenteCobrador,omitempty"`
	Receita string `xml:"receita,omitempty" json:"receita,omitempty"`
	ContaContabil string `xml:"contaContabil,omitempty" json:"contaContabil,omitempty"`
	Representante string `xml:"Representante,omitempty" json:"Representante,omitempty"`
	Vencimento int32 `xml:"Vencimento,omitempty" json:"Vencimento,omitempty"`
	Etiqueta int32 `xml:"Etiqueta,omitempty" json:"Etiqueta,omitempty"`
	GeraNFAuto int32 `xml:"GeraNFAuto,omitempty" json:"GeraNFAuto,omitempty"`
	RetemISS int32 `xml:"RetemISS,omitempty" json:"RetemISS,omitempty"`
	RetemINSS int32 `xml:"RetemINSS,omitempty" json:"RetemINSS,omitempty"`
	AliquotaISS float64 `xml:"AliquotaISS,omitempty" json:"AliquotaISS,omitempty"`
	CentroResultado string `xml:"CentroResultado,omitempty" json:"CentroResultado,omitempty"`
	IM string `xml:"IM,omitempty" json:"IM,omitempty"`
	CampoExtra1 string `xml:"CampoExtra1,omitempty" json:"CampoExtra1,omitempty"`
	CampoExtra2 string `xml:"CampoExtra2,omitempty" json:"CampoExtra2,omitempty"`
	CampoExtra3 string `xml:"CampoExtra3,omitempty" json:"CampoExtra3,omitempty"`
	CampoExtra4 string `xml:"CampoExtra4,omitempty" json:"CampoExtra4,omitempty"`
	CampoExtra5 string `xml:"CampoExtra5,omitempty" json:"CampoExtra5,omitempty"`
	CampoExtra6 string `xml:"CampoExtra6,omitempty" json:"CampoExtra6,omitempty"`
	CampoExtra7 string `xml:"CampoExtra7,omitempty" json:"CampoExtra7,omitempty"`
	CampoExtra8 string `xml:"CampoExtra8,omitempty" json:"CampoExtra8,omitempty"`
	CampoExtra9 string `xml:"CampoExtra9,omitempty" json:"CampoExtra9,omitempty"`
	CampoExtra10 string `xml:"CampoExtra10,omitempty" json:"CampoExtra10,omitempty"`
	//ServicosCliente *TServicosCliente `xml:"ServicosCliente,omitempty" json:"ServicosCliente,omitempty"`
	IDWS string `xml:"IDWS,omitempty" json:"IDWS,omitempty"`
	EhCLIPagador int32 `xml:"EhCLIPagador,omitempty" json:"EhCLIPagador,omitempty"`
}

type IAGExample interface {
	GetListaClientesExample(request *GetListaClientesRequest, response *EnvelopeResponse) error
}

type iAGExample struct {
	client *soap.Client
}

func (i *iAGExample) GetListaClientesExample(request *GetListaClientesRequest, response *EnvelopeResponse) error {
	return i.client.Call("urn:AGIntf-IAG#getListaClientes", request, response)
}

func NewIAGExample(client *soap.Client) IAGExample {
	return &iAGExample{
		client: client,
	}
}
