// generated from spec version: 1.0
package lanconfigsecurity

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
	XMLName                             xml.Name `xml:"GetInfoResponse"`
	NewMaxCharsPassword                 uint16   `xml:"NewMaxCharsPassword"`
	NewMinCharsPassword                 uint16   `xml:"NewMinCharsPassword"`
	NewAllowedCharsPassword             string   `xml:"NewAllowedCharsPassword"`
	NewAllowedCharsUsername             string   `xml:"NewAllowedCharsUsername"`
	NewX_AVM_DE_IsDefaultPasswordActive bool     `xml:"NewX_AVM-DE_IsDefaultPasswordActive"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetInfoRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetInfoResponse]{
		Body: &tr064.SOAPResponseBody[GetInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", soapRequest, soapResponse)
}

type X_AVM_DE_GetCurrentUserRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetCurrentUserRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetCurrentUserResponse struct {
	XMLName                       xml.Name `xml:"X_AVM-DE_GetCurrentUserResponse"`
	NewX_AVM_DE_CurrentUsername   string   `xml:"NewX_AVM-DE_CurrentUsername"`
	NewX_AVM_DE_CurrentUserRights string   `xml:"NewX_AVM-DE_CurrentUserRights"`
}

func (client *ServiceClient) X_AVM_DE_GetCurrentUser(out *X_AVM_DE_GetCurrentUserResponse) error {
	in := &X_AVM_DE_GetCurrentUserRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetCurrentUserRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetCurrentUserRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetCurrentUserResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetCurrentUserResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetCurrentUser", soapRequest, soapResponse)
}

type X_AVM_DE_GetAnonymousLoginRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetAnonymousLoginRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetAnonymousLoginResponse struct {
	XMLName                           xml.Name `xml:"X_AVM-DE_GetAnonymousLoginResponse"`
	NewX_AVM_DE_AnonymousLoginEnabled bool     `xml:"NewX_AVM-DE_AnonymousLoginEnabled"`
	NewX_AVM_DE_ButtonLoginEnabled    bool     `xml:"NewX_AVM-DE_ButtonLoginEnabled"`
}

func (client *ServiceClient) X_AVM_DE_GetAnonymousLogin(out *X_AVM_DE_GetAnonymousLoginResponse) error {
	in := &X_AVM_DE_GetAnonymousLoginRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetAnonymousLoginRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetAnonymousLoginRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetAnonymousLoginResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetAnonymousLoginResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetAnonymousLogin", soapRequest, soapResponse)
}

type SetConfigPasswordRequest struct {
	XMLName      xml.Name `xml:"u:SetConfigPasswordRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewPassword  string   `xml:"NewPassword"`
}

type SetConfigPasswordResponse struct {
	XMLName xml.Name `xml:"SetConfigPasswordResponse"`
}

func (client *ServiceClient) SetConfigPassword(in *SetConfigPasswordRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[SetConfigPasswordRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[SetConfigPasswordRequest]{
			In: in,
		},
	}
	out := &SetConfigPasswordResponse{}
	soapResponse := &tr064.SOAPResponse[SetConfigPasswordResponse]{
		Body: &tr064.SOAPResponseBody[SetConfigPasswordResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "SetConfigPassword", soapRequest, soapResponse)
}

type X_AVM_DE_GetUserListRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetUserListRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetUserListResponse struct {
	XMLName              xml.Name `xml:"X_AVM-DE_GetUserListResponse"`
	NewX_AVM_DE_UserList string   `xml:"NewX_AVM-DE_UserList"`
}

func (client *ServiceClient) X_AVM_DE_GetUserList(out *X_AVM_DE_GetUserListResponse) error {
	in := &X_AVM_DE_GetUserListRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetUserListRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetUserListRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetUserListResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetUserListResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetUserList", soapRequest, soapResponse)
}