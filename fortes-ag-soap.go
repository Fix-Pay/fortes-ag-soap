package fortes_ag_soap

import (
	"encoding/xml"
	"github.com/Fix-Pay/gowsdl/soap"
	"time"
)

var _ time.Time
var _ xml.Name

type RequestGetListaClientes struct {
	XMLName              xml.Name `xml:"urn:getListaClientes"`
	EncodingStyle        string   `xml:"http://schemas.xmlsoap.org/soap/encoding/ soap:encodingStyle"`
	ConsiderarCancelados bool     `xml:"xsi:type ConsiderarCancelados"`
}

// `xml:"xsi:type,xsd:boolean ConsiderarCancelados,omitempty"`

type IAGExample interface {
	GetListaClientesExample(request *RequestGetListaClientes) (*string, error)
}

type iAGExample struct {
	client *soap.Client
}

func (i *iAGExample) GetListaClientesExample(request *RequestGetListaClientes) (*string, error) {
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
