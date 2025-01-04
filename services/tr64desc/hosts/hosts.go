// generated from spec version: 1.0
package hosts

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetHostNumberOfEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetHostNumberOfEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetHostNumberOfEntriesResponse struct {
	XMLName                xml.Name `xml:"GetHostNumberOfEntriesResponse"`
	NewHostNumberOfEntries uint16   `xml:"NewHostNumberOfEntries"`
}

func (client *ServiceClient) GetHostNumberOfEntries(out *GetHostNumberOfEntriesResponse) error {
	in := &GetHostNumberOfEntriesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[GetHostNumberOfEntriesRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetHostNumberOfEntriesRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetHostNumberOfEntriesResponse]{
		Body: &tr064.SOAPResponseBody[GetHostNumberOfEntriesResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetHostNumberOfEntries", soapRequest, soapResponse)
}

type GetSpecificHostEntryRequest struct {
	XMLName       xml.Name `xml:"u:GetSpecificHostEntryRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMACAddress string   `xml:"NewMACAddress"`
}

type GetSpecificHostEntryResponse struct {
	XMLName               xml.Name `xml:"GetSpecificHostEntryResponse"`
	NewIPAddress          string   `xml:"NewIPAddress"`
	NewAddressSource      string   `xml:"NewAddressSource"`
	NewLeaseTimeRemaining int32    `xml:"NewLeaseTimeRemaining"`
	NewInterfaceType      string   `xml:"NewInterfaceType"`
	NewActive             bool     `xml:"NewActive"`
	NewHostName           string   `xml:"NewHostName"`
}

func (client *ServiceClient) GetSpecificHostEntry(in *GetSpecificHostEntryRequest, out *GetSpecificHostEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetSpecificHostEntryRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetSpecificHostEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetSpecificHostEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetSpecificHostEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetSpecificHostEntry", soapRequest, soapResponse)
}

type GetGenericHostEntryRequest struct {
	XMLName      xml.Name `xml:"u:GetGenericHostEntryRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetGenericHostEntryResponse struct {
	XMLName               xml.Name `xml:"GetGenericHostEntryResponse"`
	NewIPAddress          string   `xml:"NewIPAddress"`
	NewAddressSource      string   `xml:"NewAddressSource"`
	NewLeaseTimeRemaining int32    `xml:"NewLeaseTimeRemaining"`
	NewMACAddress         string   `xml:"NewMACAddress"`
	NewInterfaceType      string   `xml:"NewInterfaceType"`
	NewActive             bool     `xml:"NewActive"`
	NewHostName           string   `xml:"NewHostName"`
}

func (client *ServiceClient) GetGenericHostEntry(in *GetGenericHostEntryRequest, out *GetGenericHostEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[GetGenericHostEntryRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[GetGenericHostEntryRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[GetGenericHostEntryResponse]{
		Body: &tr064.SOAPResponseBody[GetGenericHostEntryResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "GetGenericHostEntry", soapRequest, soapResponse)
}

type X_AVM_DE_GetInfoRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetInfoResponse struct {
	XMLName                          xml.Name `xml:"X_AVM-DE_GetInfoResponse"`
	NewX_AVM_DE_FriendlynameMinChars uint16   `xml:"NewX_AVM-DE_FriendlynameMinChars"`
	NewX_AVM_DE_FriendlynameMaxChars uint16   `xml:"NewX_AVM-DE_FriendlynameMaxChars"`
	NewX_AVM_DE_HostnameMinChars     uint16   `xml:"NewX_AVM-DE_HostnameMinChars"`
	NewX_AVM_DE_HostnameMaxChars     uint16   `xml:"NewX_AVM-DE_HostnameMaxChars"`
	NewX_AVM_DE_HostnameAllowedChars string   `xml:"NewX_AVM-DE_HostnameAllowedChars"`
}

func (client *ServiceClient) X_AVM_DE_GetInfo(out *X_AVM_DE_GetInfoResponse) error {
	in := &X_AVM_DE_GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetInfoRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
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

type X_AVM_DE_GetChangeCounterRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetChangeCounterRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetChangeCounterResponse struct {
	XMLName                   xml.Name `xml:"X_AVM-DE_GetChangeCounterResponse"`
	NewX_AVM_DE_ChangeCounter uint32   `xml:"NewX_AVM-DE_ChangeCounter"`
}

func (client *ServiceClient) X_AVM_DE_GetChangeCounter(out *X_AVM_DE_GetChangeCounterResponse) error {
	in := &X_AVM_DE_GetChangeCounterRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetChangeCounterRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetChangeCounterRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetChangeCounterResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetChangeCounterResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetChangeCounter", soapRequest, soapResponse)
}

type X_AVM_DE_SetHostNameByMACAddressRequest struct {
	XMLName       xml.Name `xml:"u:X_AVM-DE_SetHostNameByMACAddressRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMACAddress string   `xml:"NewMACAddress"`
	NewHostName   string   `xml:"NewHostName"`
}

type X_AVM_DE_SetHostNameByMACAddressResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetHostNameByMACAddressResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetHostNameByMACAddress(in *X_AVM_DE_SetHostNameByMACAddressRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetHostNameByMACAddressRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetHostNameByMACAddressRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetHostNameByMACAddressResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetHostNameByMACAddressResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetHostNameByMACAddressResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetHostNameByMACAddress", soapRequest, soapResponse)
}

type X_AVM_DE_GetAutoWakeOnLANByMACAddressRequest struct {
	XMLName       xml.Name `xml:"u:X_AVM-DE_GetAutoWakeOnLANByMACAddressRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMACAddress string   `xml:"NewMACAddress"`
}

type X_AVM_DE_GetAutoWakeOnLANByMACAddressResponse struct {
	XMLName           xml.Name `xml:"X_AVM-DE_GetAutoWakeOnLANByMACAddressResponse"`
	NewAutoWOLEnabled bool     `xml:"NewAutoWOLEnabled"`
}

func (client *ServiceClient) X_AVM_DE_GetAutoWakeOnLANByMACAddress(in *X_AVM_DE_GetAutoWakeOnLANByMACAddressRequest, out *X_AVM_DE_GetAutoWakeOnLANByMACAddressResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetAutoWakeOnLANByMACAddressRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetAutoWakeOnLANByMACAddressRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetAutoWakeOnLANByMACAddressResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetAutoWakeOnLANByMACAddressResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetAutoWakeOnLANByMACAddress", soapRequest, soapResponse)
}

type X_AVM_DE_SetAutoWakeOnLANByMACAddressRequest struct {
	XMLName           xml.Name `xml:"u:X_AVM-DE_SetAutoWakeOnLANByMACAddressRequest"`
	XMLNameSpace      string   `xml:"xmlns:u,attr"`
	NewMACAddress     string   `xml:"NewMACAddress"`
	NewAutoWOLEnabled bool     `xml:"NewAutoWOLEnabled"`
}

type X_AVM_DE_SetAutoWakeOnLANByMACAddressResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetAutoWakeOnLANByMACAddressResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetAutoWakeOnLANByMACAddress(in *X_AVM_DE_SetAutoWakeOnLANByMACAddressRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetAutoWakeOnLANByMACAddressRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetAutoWakeOnLANByMACAddressRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetAutoWakeOnLANByMACAddressResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetAutoWakeOnLANByMACAddressResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetAutoWakeOnLANByMACAddressResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetAutoWakeOnLANByMACAddress", soapRequest, soapResponse)
}

type X_AVM_DE_WakeOnLANByMACAddressRequest struct {
	XMLName       xml.Name `xml:"u:X_AVM-DE_WakeOnLANByMACAddressRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMACAddress string   `xml:"NewMACAddress"`
}

type X_AVM_DE_WakeOnLANByMACAddressResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_WakeOnLANByMACAddressResponse"`
}

func (client *ServiceClient) X_AVM_DE_WakeOnLANByMACAddress(in *X_AVM_DE_WakeOnLANByMACAddressRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_WakeOnLANByMACAddressRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_WakeOnLANByMACAddressRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_WakeOnLANByMACAddressResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_WakeOnLANByMACAddressResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_WakeOnLANByMACAddressResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_WakeOnLANByMACAddress", soapRequest, soapResponse)
}

type X_AVM_DE_GetSpecificHostEntryByIPRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetSpecificHostEntryByIPRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIPAddress string   `xml:"NewIPAddress"`
}

type X_AVM_DE_GetSpecificHostEntryByIPResponse struct {
	XMLName                             xml.Name `xml:"X_AVM-DE_GetSpecificHostEntryByIPResponse"`
	NewMACAddress                       string   `xml:"NewMACAddress"`
	NewActive                           bool     `xml:"NewActive"`
	NewHostName                         string   `xml:"NewHostName"`
	NewInterfaceType                    string   `xml:"NewInterfaceType"`
	NewX_AVM_DE_Port                    uint32   `xml:"NewX_AVM-DE_Port"`
	NewX_AVM_DE_Speed                   uint32   `xml:"NewX_AVM-DE_Speed"`
	NewX_AVM_DE_UpdateAvailable         bool     `xml:"NewX_AVM-DE_UpdateAvailable"`
	NewX_AVM_DE_UpdateSuccessful        string   `xml:"NewX_AVM-DE_UpdateSuccessful"`
	NewX_AVM_DE_InfoURL                 string   `xml:"NewX_AVM-DE_InfoURL"`
	NewX_AVM_DE_MACAddressList          string   `xml:"NewX_AVM-DE_MACAddressList"`
	NewX_AVM_DE_Model                   string   `xml:"NewX_AVM-DE_Model"`
	NewX_AVM_DE_URL                     string   `xml:"NewX_AVM-DE_URL"`
	NewX_AVM_DE_Guest                   bool     `xml:"NewX_AVM-DE_Guest"`
	NewX_AVM_DE_RequestClient           bool     `xml:"NewX_AVM-DE_RequestClient"`
	NewX_AVM_DE_VPN                     bool     `xml:"NewX_AVM-DE_VPN"`
	NewX_AVM_DE_WANAccess               string   `xml:"NewX_AVM-DE_WANAccess"`
	NewX_AVM_DE_Disallow                bool     `xml:"NewX_AVM-DE_Disallow"`
	NewX_AVM_DE_IsMeshable              bool     `xml:"NewX_AVM-DE_IsMeshable"`
	NewX_AVM_DE_Priority                bool     `xml:"NewX_AVM-DE_Priority"`
	NewX_AVM_DE_FriendlyName            string   `xml:"NewX_AVM-DE_FriendlyName"`
	NewX_AVM_DE_FriendlyNameIsWriteable bool     `xml:"NewX_AVM-DE_FriendlyNameIsWriteable"`
}

func (client *ServiceClient) X_AVM_DE_GetSpecificHostEntryByIP(in *X_AVM_DE_GetSpecificHostEntryByIPRequest, out *X_AVM_DE_GetSpecificHostEntryByIPResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetSpecificHostEntryByIPRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetSpecificHostEntryByIPRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetSpecificHostEntryByIPResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetSpecificHostEntryByIPResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetSpecificHostEntryByIP", soapRequest, soapResponse)
}

type X_AVM_DE_HostsCheckUpdateRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_HostsCheckUpdateRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_HostsCheckUpdateResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_HostsCheckUpdateResponse"`
}

func (client *ServiceClient) X_AVM_DE_HostsCheckUpdate() error {
	in := &X_AVM_DE_HostsCheckUpdateRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_HostsCheckUpdateRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_HostsCheckUpdateRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_HostsCheckUpdateResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_HostsCheckUpdateResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_HostsCheckUpdateResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_HostsCheckUpdate", soapRequest, soapResponse)
}

type X_AVM_DE_HostDoUpdateRequest struct {
	XMLName       xml.Name `xml:"u:X_AVM-DE_HostDoUpdateRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMACAddress string   `xml:"NewMACAddress"`
}

type X_AVM_DE_HostDoUpdateResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_HostDoUpdateResponse"`
}

func (client *ServiceClient) X_AVM_DE_HostDoUpdate(in *X_AVM_DE_HostDoUpdateRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_HostDoUpdateRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_HostDoUpdateRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_HostDoUpdateResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_HostDoUpdateResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_HostDoUpdateResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_HostDoUpdate", soapRequest, soapResponse)
}

type X_AVM_DE_SetPrioritizationByIPRequest struct {
	XMLName              xml.Name `xml:"u:X_AVM-DE_SetPrioritizationByIPRequest"`
	XMLNameSpace         string   `xml:"xmlns:u,attr"`
	NewIPAddress         string   `xml:"NewIPAddress"`
	NewX_AVM_DE_Priority bool     `xml:"NewX_AVM-DE_Priority"`
}

type X_AVM_DE_SetPrioritizationByIPResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetPrioritizationByIPResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetPrioritizationByIP(in *X_AVM_DE_SetPrioritizationByIPRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetPrioritizationByIPRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetPrioritizationByIPRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetPrioritizationByIPResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetPrioritizationByIPResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetPrioritizationByIPResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetPrioritizationByIP", soapRequest, soapResponse)
}

type X_AVM_DE_GetHostListPathRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetHostListPathRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetHostListPathResponse struct {
	XMLName                  xml.Name `xml:"X_AVM-DE_GetHostListPathResponse"`
	NewX_AVM_DE_HostListPath string   `xml:"NewX_AVM-DE_HostListPath"`
}

func (client *ServiceClient) X_AVM_DE_GetHostListPath(out *X_AVM_DE_GetHostListPathResponse) error {
	in := &X_AVM_DE_GetHostListPathRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetHostListPathRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetHostListPathRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetHostListPathResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetHostListPathResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetHostListPath", soapRequest, soapResponse)
}

type X_AVM_DE_GetMeshListPathRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetMeshListPathRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetMeshListPathResponse struct {
	XMLName                  xml.Name `xml:"X_AVM-DE_GetMeshListPathResponse"`
	NewX_AVM_DE_MeshListPath string   `xml:"NewX_AVM-DE_MeshListPath"`
}

func (client *ServiceClient) X_AVM_DE_GetMeshListPath(out *X_AVM_DE_GetMeshListPathResponse) error {
	in := &X_AVM_DE_GetMeshListPathRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetMeshListPathRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetMeshListPathRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetMeshListPathResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetMeshListPathResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetMeshListPath", soapRequest, soapResponse)
}

type X_AVM_DE_GetFriendlyNameRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetFriendlyNameRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetFriendlyNameResponse struct {
	XMLName                  xml.Name `xml:"X_AVM-DE_GetFriendlyNameResponse"`
	NewX_AVM_DE_FriendlyName string   `xml:"NewX_AVM-DE_FriendlyName"`
}

func (client *ServiceClient) X_AVM_DE_GetFriendlyName(out *X_AVM_DE_GetFriendlyNameResponse) error {
	in := &X_AVM_DE_GetFriendlyNameRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_GetFriendlyNameRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_GetFriendlyNameRequest]{
			In: in,
		},
	}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_GetFriendlyNameResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_GetFriendlyNameResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetFriendlyName", soapRequest, soapResponse)
}

type X_AVM_DE_SetFriendlyNameRequest struct {
	XMLName                  xml.Name `xml:"u:X_AVM-DE_SetFriendlyNameRequest"`
	XMLNameSpace             string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_FriendlyName string   `xml:"NewX_AVM-DE_FriendlyName"`
}

type X_AVM_DE_SetFriendlyNameResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetFriendlyNameResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetFriendlyName(in *X_AVM_DE_SetFriendlyNameRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetFriendlyNameRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetFriendlyNameRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetFriendlyNameResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetFriendlyNameResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetFriendlyNameResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetFriendlyName", soapRequest, soapResponse)
}

type X_AVM_DE_SetFriendlyNameByIPRequest struct {
	XMLName                  xml.Name `xml:"u:X_AVM-DE_SetFriendlyNameByIPRequest"`
	XMLNameSpace             string   `xml:"xmlns:u,attr"`
	NewIPAddress             string   `xml:"NewIPAddress"`
	NewX_AVM_DE_FriendlyName string   `xml:"NewX_AVM-DE_FriendlyName"`
}

type X_AVM_DE_SetFriendlyNameByIPResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetFriendlyNameByIPResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetFriendlyNameByIP(in *X_AVM_DE_SetFriendlyNameByIPRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetFriendlyNameByIPRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetFriendlyNameByIPRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetFriendlyNameByIPResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetFriendlyNameByIPResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetFriendlyNameByIPResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetFriendlyNameByIP", soapRequest, soapResponse)
}

type X_AVM_DE_SetFriendlyNameByMACRequest struct {
	XMLName                  xml.Name `xml:"u:X_AVM-DE_SetFriendlyNameByMACRequest"`
	XMLNameSpace             string   `xml:"xmlns:u,attr"`
	NewMACAddress            string   `xml:"NewMACAddress"`
	NewX_AVM_DE_FriendlyName string   `xml:"NewX_AVM-DE_FriendlyName"`
}

type X_AVM_DE_SetFriendlyNameByMACResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetFriendlyNameByMACResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetFriendlyNameByMAC(in *X_AVM_DE_SetFriendlyNameByMACRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := &tr064.SOAPRequest[X_AVM_DE_SetFriendlyNameByMACRequest]{
		XMLNameSpace:     tr064.XMLNameSpace,
		XMLEncodingStyle: tr064.XMLEncodingStyle,
		Body: &tr064.SOAPRequestBody[X_AVM_DE_SetFriendlyNameByMACRequest]{
			In: in,
		},
	}
	out := &X_AVM_DE_SetFriendlyNameByMACResponse{}
	soapResponse := &tr064.SOAPResponse[X_AVM_DE_SetFriendlyNameByMACResponse]{
		Body: &tr064.SOAPResponseBody[X_AVM_DE_SetFriendlyNameByMACResponse]{
			Out: out,
		},
	}
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetFriendlyNameByMAC", soapRequest, soapResponse)
}