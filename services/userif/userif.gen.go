// generated from spec version: 1.0
package userif

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
	XMLName                          xml.Name `xml:"GetInfoResponse"`
	NewUpgradeAvailable              bool     `xml:"NewUpgradeAvailable"`
	NewPasswordRequired              bool     `xml:"NewPasswordRequired"`
	NewPasswordUserSelectable        bool     `xml:"NewPasswordUserSelectable"`
	NewWarrantyDate                  string   `xml:"NewWarrantyDate"`
	NewX_AVM_DE_Version              string   `xml:"NewX_AVM-DE_Version"`
	NewX_AVM_DE_DownloadURL          string   `xml:"NewX_AVM-DE_DownloadURL"`
	NewX_AVM_DE_InfoURL              string   `xml:"NewX_AVM-DE_InfoURL"`
	NewX_AVM_DE_UpdateState          string   `xml:"NewX_AVM-DE_UpdateState"`
	NewX_AVM_DE_BuildType            string   `xml:"NewX_AVM-DE_BuildType"`
	NewX_AVM_DE_SetupAssistantStatus bool     `xml:"NewX_AVM-DE_SetupAssistantStatus"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetInfoRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
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

type X_AVM_DE_CheckUpdateRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_CheckUpdateRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_CheckUpdateResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_CheckUpdateResponse"`
}

func (client *ServiceClient) X_AVM_DE_CheckUpdate() error {
	in := &X_AVM_DE_CheckUpdateRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_CheckUpdateRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_CheckUpdateRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_CheckUpdateResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_CheckUpdateResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_CheckUpdateResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_CheckUpdate", soapRequest, soapResponse)
}

type X_AVM_DE_DoUpdateRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_DoUpdateRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_DoUpdateResponse struct {
	XMLName                 xml.Name `xml:"X_AVM-DE_DoUpdateResponse"`
	NewUpgradeAvailable     bool     `xml:"NewUpgradeAvailable"`
	NewX_AVM_DE_UpdateState string   `xml:"NewX_AVM-DE_UpdateState"`
}

func (client *ServiceClient) X_AVM_DE_DoUpdate(out *X_AVM_DE_DoUpdateResponse) error {
	in := &X_AVM_DE_DoUpdateRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_DoUpdateRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_DoUpdateRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_DoUpdateResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_DoUpdateResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DoUpdate", soapRequest, soapResponse)
}

type X_AVM_DE_DoPrepareCGIRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_DoPrepareCGIRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_DoPrepareCGIResponse struct {
	XMLName               xml.Name `xml:"X_AVM-DE_DoPrepareCGIResponse"`
	NewX_AVM_DE_CGI       string   `xml:"NewX_AVM-DE_CGI"`
	NewX_AVM_DE_SessionID string   `xml:"NewX_AVM-DE_SessionID"`
}

func (client *ServiceClient) X_AVM_DE_DoPrepareCGI(out *X_AVM_DE_DoPrepareCGIResponse) error {
	in := &X_AVM_DE_DoPrepareCGIRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_DoPrepareCGIRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_DoPrepareCGIRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_DoPrepareCGIResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_DoPrepareCGIResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DoPrepareCGI", soapRequest, soapResponse)
}

type X_AVM_DE_DoManualUpdateRequest struct {
	XMLName                    xml.Name `xml:"u:X_AVM-DE_DoManualUpdateRequest"`
	XMLNameSpace               string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_AllowDowngrade bool     `xml:"NewX_AVM-DE_AllowDowngrade"`
	NewX_AVM_DE_DownloadURL    string   `xml:"NewX_AVM-DE_DownloadURL"`
}

type X_AVM_DE_DoManualUpdateResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_DoManualUpdateResponse"`
}

func (client *ServiceClient) X_AVM_DE_DoManualUpdate(in *X_AVM_DE_DoManualUpdateRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_DoManualUpdateRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_DoManualUpdateRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_DoManualUpdateResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_DoManualUpdateResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_DoManualUpdateResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_DoManualUpdate", soapRequest, soapResponse)
}

type X_AVM_DE_GetInternationalConfigRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetInternationalConfigRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetInternationalConfigResponse struct {
	XMLName                  xml.Name `xml:"X_AVM-DE_GetInternationalConfigResponse"`
	NewX_AVM_DE_Language     string   `xml:"NewX_AVM-DE_Language"`
	NewX_AVM_DE_Country      string   `xml:"NewX_AVM-DE_Country"`
	NewX_AVM_DE_Annex        string   `xml:"NewX_AVM-DE_Annex"`
	NewX_AVM_DE_LanguageList string   `xml:"NewX_AVM-DE_LanguageList"`
	NewX_AVM_DE_CountryList  string   `xml:"NewX_AVM-DE_CountryList"`
	NewX_AVM_DE_AnnexList    string   `xml:"NewX_AVM-DE_AnnexList"`
}

func (client *ServiceClient) X_AVM_DE_GetInternationalConfig(out *X_AVM_DE_GetInternationalConfigResponse) error {
	in := &X_AVM_DE_GetInternationalConfigRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetInternationalConfigRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetInternationalConfigRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetInternationalConfigResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetInternationalConfigResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetInternationalConfig", soapRequest, soapResponse)
}

type X_AVM_DE_SetInternationalConfigRequest struct {
	XMLName              xml.Name `xml:"u:X_AVM-DE_SetInternationalConfigRequest"`
	XMLNameSpace         string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_Language string   `xml:"NewX_AVM-DE_Language"`
	NewX_AVM_DE_Country  string   `xml:"NewX_AVM-DE_Country"`
	NewX_AVM_DE_Annex    string   `xml:"NewX_AVM-DE_Annex"`
}

type X_AVM_DE_SetInternationalConfigResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetInternationalConfigResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetInternationalConfig(in *X_AVM_DE_SetInternationalConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetInternationalConfigRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetInternationalConfigRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetInternationalConfigResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetInternationalConfigResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetInternationalConfigResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetInternationalConfig", soapRequest, soapResponse)
}

type X_AVM_DE_GetInfoRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetInfoResponse struct {
	XMLName                      xml.Name `xml:"X_AVM-DE_GetInfoResponse"`
	NewX_AVM_DE_AutoUpdateMode   string   `xml:"NewX_AVM-DE_AutoUpdateMode"`
	NewX_AVM_DE_UpdateTime       string   `xml:"NewX_AVM-DE_UpdateTime"`
	NewX_AVM_DE_LastFwVersion    string   `xml:"NewX_AVM-DE_LastFwVersion"`
	NewX_AVM_DE_LastInfoUrl      string   `xml:"NewX_AVM-DE_LastInfoUrl"`
	NewX_AVM_DE_CurrentFwVersion string   `xml:"NewX_AVM-DE_CurrentFwVersion"`
	NewX_AVM_DE_UpdateSuccessful string   `xml:"NewX_AVM-DE_UpdateSuccessful"`
}

func (client *ServiceClient) X_AVM_DE_GetInfo(out *X_AVM_DE_GetInfoResponse) error {
	in := &X_AVM_DE_GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetInfoRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetInfoRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetInfoResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetInfoResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetInfo", soapRequest, soapResponse)
}

type X_AVM_DE_SetConfigRequest struct {
	XMLName                    xml.Name `xml:"u:X_AVM-DE_SetConfigRequest"`
	XMLNameSpace               string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_AutoUpdateMode string   `xml:"NewX_AVM-DE_AutoUpdateMode"`
}

type X_AVM_DE_SetConfigResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetConfigResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetConfig(in *X_AVM_DE_SetConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetConfigRequest]{
		XMLNameSpace:     "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetConfigRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetConfigResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetConfigResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetConfigResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetConfig", soapRequest, soapResponse)
}
