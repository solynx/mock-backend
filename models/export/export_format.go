package export_format

type Url struct {
	Raw  string   `json:"raw"`
	Host []string `json:"host"`
	Port string   `json:"port"`
	Path []string `json:"path"`
}
type OriginalRequest struct {
	Method string `json:"method"`
	Header string `json:"header"`
	Url    Url    `json:"url"`
}
type Response struct {
	Name            string          `json:"name"`
	PreviewLanguage string          `json:"_postman_previewlanguage"`
	Header          []string        `json:"header"`
	Cookie          []string        `json:"cookie"`
	Body            string          `json:"body"`
	OriginalRequest OriginalRequest `json:"originalRequest"`
}

// req
type Request struct {
	Name          string        `json:"name"`
	Method        string        `json:"method"`
	Header        interface{}   `json:"header"`
	Response      []Response    `json:"response"`
	BodyOfRequest BodyOfRequest `json:"body"`
}
type RequestOfCollection struct {
	Name                    string             `json:"name"`
	ProtocolProfileBehavior disableBodyPruning `json:"protocolProfileBehavior"`
	Request                 SingleRequest      `json:"request"`
	Response                []Response         `json:"response"`
}
type disableBodyPruning struct {
	BodyPruning bool `json:"disableBodyPruning"`
}
type SingleRequest struct {
	Method        string        `json:"method"`
	Header        interface{}   `json:"header"`
	BodyOfRequest BodyOfRequest `json:"body"`
	Url           Url_Req_Col   `json:"url"`
}
type Url_Req_Col struct {
	Raw      string      `json:"raw"`
	Protocol string      `json:"protocol"`
	Host     []string    `json:"host"`
	Port     string      `json:"port"`
	Path     []string    `json:"path"`
	Query    interface{} `json:"query"`
}
type BodyOfRequest struct {
	Mode    string           `json:"mode"`
	Raw     string           `json:"raw"`
	Options OptionsOfRequest `json:"options"`
}
type OptionsOfRequest struct {
	Raw RawLang `json:"raw"`
}
type RawLang struct {
	Language string `json:"language"`
}

// folder
type Folder struct {
	Name string        `json:"name"`
	Item []interface{} `json:"item"`
}

// collection
type CollectionJson struct {
	Collection Collection1
}
type Collection1 struct {
	Info     Info          `json:"info"`
	Item     []interface{} `json:"item"`
	Variable interface{}   `json:"variable,"`
}
type Info struct {
	Id       string `json:"_postman_id"`
	Name     string `json:"name"`
	Schema   string `json:"schema"`
	ExportId int32  `json:"_exporter_id"`
}
type Item struct {
	Request Request
	Folder  Folder
}
