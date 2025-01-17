// generated from spec version: 1.0
package x_storage

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetInfoResponse struct {
	XMLName          xml.Name `xml:"GetInfoResponse"`
	NewFTPEnable     bool     `xml:"NewFTPEnable"`
	NewFTPStatus     string   `xml:"NewFTPStatus"`
	NewSMBEnable     bool     `xml:"NewSMBEnable"`
	NewFTPWANEnable  bool     `xml:"NewFTPWANEnable"`
	NewFTPWANSSLOnly bool     `xml:"NewFTPWANSSLOnly"`
	NewFTPWANPort    uint16   `xml:"NewFTPWANPort"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type RequestFTPServerWANRequest struct {
	XMLName      xml.Name `xml:"u:RequestFTPServerWANRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type RequestFTPServerWANResponse struct {
	XMLName           xml.Name `xml:"RequestFTPServerWANResponse"`
	NewFTPWANPort     uint16   `xml:"NewFTPWANPort"`
	NewFTPWANLifetime uint32   `xml:"NewFTPWANLifetime"`
}

func (client *ServiceClient) RequestFTPServerWAN(out *RequestFTPServerWANResponse) error {
	in := &RequestFTPServerWANRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "RequestFTPServerWAN", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetFTPServerRequest struct {
	XMLName      xml.Name `xml:"u:SetFTPServerRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewFTPEnable bool     `xml:"NewFTPEnable"`
}

type SetFTPServerResponse struct {
	XMLName xml.Name `xml:"SetFTPServerResponse"`
}

func (client *ServiceClient) SetFTPServer(in *SetFTPServerRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetFTPServerResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetFTPServer", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetFTPServerWANRequest struct {
	XMLName          xml.Name `xml:"u:SetFTPServerWANRequest"`
	XMLNameSpace     string   `xml:"xmlns:u,attr"`
	NewFTPWANEnable  bool     `xml:"NewFTPWANEnable"`
	NewFTPWANSSLOnly bool     `xml:"NewFTPWANSSLOnly"`
}

type SetFTPServerWANResponse struct {
	XMLName xml.Name `xml:"SetFTPServerWANResponse"`
}

func (client *ServiceClient) SetFTPServerWAN(in *SetFTPServerWANRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetFTPServerWANResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetFTPServerWAN", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetSMBServerRequest struct {
	XMLName      xml.Name `xml:"u:SetSMBServerRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewSMBEnable bool     `xml:"NewSMBEnable"`
}

type SetSMBServerResponse struct {
	XMLName xml.Name `xml:"SetSMBServerResponse"`
}

func (client *ServiceClient) SetSMBServer(in *SetSMBServerRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetSMBServerResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetSMBServer", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetUserInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetUserInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetUserInfoResponse struct {
	XMLName                           xml.Name `xml:"GetUserInfoResponse"`
	NewEnable                         bool     `xml:"NewEnable"`
	NewUsername                       string   `xml:"NewUsername"`
	NewX_AVM_DE_NetworkAccessReadOnly bool     `xml:"NewX_AVM-DE_NetworkAccessReadOnly"`
}

func (client *ServiceClient) GetUserInfo(out *GetUserInfoResponse) error {
	in := &GetUserInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetUserInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetUserConfigRequest struct {
	XMLName                           xml.Name `xml:"u:SetUserConfigRequest"`
	XMLNameSpace                      string   `xml:"xmlns:u,attr"`
	NewEnable                         bool     `xml:"NewEnable"`
	NewPassword                       string   `xml:"NewPassword"`
	NewX_AVM_DE_NetworkAccessReadOnly bool     `xml:"NewX_AVM-DE_NetworkAccessReadOnly"`
}

type SetUserConfigResponse struct {
	XMLName xml.Name `xml:"SetUserConfigResponse"`
}

func (client *ServiceClient) SetUserConfig(in *SetUserConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetUserConfigResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetUserConfig", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
