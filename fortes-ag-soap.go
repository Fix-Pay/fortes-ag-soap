package fortes_ag_soap

import (
	"encoding/xml"
	"github.com/Fix-Pay/gowsdl/soap"
	"time"
)

var _ time.Time
var _ xml.Name

type RequestClientExample struct {
	XMLName                  xml.Name `xml:"http://example.com/stockquote.xsd RequestClientExample"`
	ConsiderarCancelados     bool     `xml:"ConsiderarCancelados,omitempty" json:"ConsiderarCancelados,omitempty"`
	CPFCNPJ                  string   `xml:"CPFCNPJ,omitempty" json:"CPFCNPJ,omitempty"`
	DetalhaContatosEServicos bool     `xml:"DetalhaContatosEServicos,omitempty" json:"DetalhaContatosEServicos,omitempty"`
}

type IAGExample interface {
	GetListaClientesExample(request *RequestClientExample) (*string, error)
}

type iAGExample struct {
	client *soap.Client
}

func (i *iAGExample) GetListaClientesExample(request *RequestClientExample) (*string, error) {
	response := new(string)
	err := i.client.Call("urn:AGIntf-IAG#getListaClientes", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewIAGExample(client *soap.Client) IAGExample {
	return &iAGExample{
		client: client,
	}
}
