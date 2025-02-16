package altalabs

type Devices []Device

type Device struct {
	Type          any    `json:"type,omitempty"`
	Vlan          int    `json:"vlan,omitempty"`
	IP            any    `json:"ip,omitempty"`
	UlRate        int    `json:"ulRate,omitempty"`
	DlRate        any    `json:"dlRate,omitempty"`
	IgnoreHotspot any    `json:"ignoreHotspot,omitempty"`
	IgnoreSched   any    `json:"ignoreSched,omitempty"`
	IgnoreFilter  any    `json:"ignoreFilter,omitempty"`
	Icon          string `json:"icon,omitempty"`
	Upnp          any    `json:"upnp,omitempty"`
	UpnpPorts     any    `json:"upnpPorts,omitempty"`
	Siteid        string `json:"siteid,omitempty"`
	ID            string `json:"id"` // Required
	Wired         bool   `json:"wired,omitempty"`
}

type ListDeviceRequest struct {
	SiteName string `json:"siteName"`
}

func (a *AltaClient) ListDevices(siteName string) (Devices, error) {
	siteURL := "device/list"

	req := ListDeviceRequest{
		SiteName: siteName,
	}

	var devices = make(Devices, 0)

	if err := a.getRequest(siteURL, req, &devices); err != nil {
		return nil, err
	}

	return devices, nil
}

// TODO: Update devices
