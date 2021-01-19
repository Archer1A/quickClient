package field

type HttpRequestField struct {
	Method string
	Path string
    Response interface{}
	ContentType string
    RequestField *[]FieldDef
}


