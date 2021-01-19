package field

type LocationType int32

const (
	Path LocationType = iota
	Query
)

type FieldDef struct {
	LocationType LocationType
	JsonTag      string
	Name         string
	KindName     string
}

func NewFieldDef() *FieldDef {
	return &FieldDef{}
}

func (f *FieldDef)WithLocationType(location LocationType)*FieldDef  {
	f.LocationType = location
	return f
}

func (f *FieldDef)WithJsonTag(tag string)*FieldDef  {
	f.JsonTag = tag
	return f
}

func (f *FieldDef)WithName(name string)*FieldDef  {
	f.Name = name
	return f
}
func (f *FieldDef)WithKindName(kind string)*FieldDef  {
	f.KindName = kind
	return f
}