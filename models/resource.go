package models

type Resource struct {
	ID                              string            `json:"_id"`
	ParentID                        string            `json:"parentId"`
	Modified                        int64             `json:"modified"`
	Created                         int64             `json:"created"`
	Name                            string            `json:"name"`
	Description                     string            `json:"description,omitempty"`
	Certificates                    []interface{}     `json:"certificates,omitempty"`
	Type                            string            `json:"_type"`
	Data                            DataUrl           `json:"data,omitempty"`
	Color                           interface{}       `json:"color,omitempty"`
	Cookies                         []interface{}     `json:"cookies,omitempty"`
	Environment                     struct{}          `json:"environment,omitempty"`
	URL                             string            `json:"url,omitempty"`
	Method                          string            `json:"method,omitempty"`
	Body                            EntityBody        `json:"body,omitempty"`
	Parameters                      []QueryParameters `json:"parameters,omitempty"`
	Headers                         []EntityHeader    `json:"headers,omitempty"`
	Authentication                  struct{}          `json:"authentication,omitempty"`
	SettingStoreCookies             bool              `json:"settingStoreCookies,omitempty"`
	SettingSendCookies              bool              `json:"settingSendCookies,omitempty"`
	SettingDisableRenderRequestBody bool              `json:"settingDisableRenderRequestBody,omitempty"`
	SettingEncodeURL                bool              `json:"settingEncodeUrl,omitempty"`
	InsomniaParams                  []string          `json:"insomnia_params,omitempty"`
}

type EntityBody struct {
	MimeType string        `json:"mimeType"`
	Params   []EntityParam `json:"params"`
}

type EntityParam struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	ID       string `json:"id"`
	Disabled bool   `json:"disabled"`
}
type QueryParameters struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type EntityHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Also check if any header is valid

func (resource Resource) HasParamsInHeader() bool {

}

// Define method to check if the header values are allowed to be part of OAS
// According to OpenAPI 3.0 "Accept", "Content-Type" and "Authorization" are not allowed

func (entity EntityHeader) IsAllowed() bool {
	restrictedFields := []string{"Accept", "Content-Type", "Authorization"}
	_, found := Find(restrictedFields, entity.Name)
	if found {
		return false
	} else {
		return true
	}
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

type DataUrl struct {
	BaseURL string `json:"base_url"`
	RefURL  string `json:"ref_url"`
}
