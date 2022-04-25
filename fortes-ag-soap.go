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

type UrnGetListaClientes struct {
	XMLName       xml.Name `xml:"urn:getListaClientes"`
	EncodingStyle string   `xml:"soap:encodingStyle,attr"`

	ConsiderarCancelados bool `xml:"ConsiderarCancelados,omitempty"`
}

type UrnGetListaClientesResponse struct {
	XMLName  xml.Name `xml:"ns1:getListaClientesResponse"`
	XmlnsNs1 string   `xml:"xmlns:ns1,attr"`

	Return string `xml:"return"`
}

type IAGExample interface {
	GetListaClientesExample(request *UrnGetListaClientes) (*string, error)
}

type iAGExample struct {
	client *soap.Client
}

func (i *iAGExample) GetListaClientesExample(request *UrnGetListaClientes) (*string, error) {
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
