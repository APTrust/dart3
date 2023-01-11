package common

type Choice struct {
	Name     string
	Value    string
	Selected bool
}

type Field struct {
	//ID             string
	Name           string
	Label          string
	Value          string
	Required       bool
	Help           string
	Hide           bool
	Error          string
	Choices        []Choice
	CssClasses     []string
	Attrs          map[string]string
	Changed        bool
	Cast           string // Type to cast to, if field is not string. Options: number, bool
	FormGroupClass string // set to "form-group-hidden" if SystemMustSet or tag has default value - https://github.com/APTrust/dart/blob/213e0a9acde2c73fcc8430ca436263b806228f8f/ui/forms/job_tags_form.js#L80
}

func NewField(name, label, value string, required bool) *Field {
	return &Field{
		Name:       name,
		Label:      label,
		Value:      value,
		Required:   required,
		Choices:    make([]Choice, 0),
		CssClasses: make([]string, 0),
		Attrs:      make(map[string]string),
	}
}

type Form struct {
	ID             string
	ObjectID       string
	Fields         map[string]*Field
	SubmitFunction string
	CancelFunction string
	DeleteFunction string
	UserCanDelete  bool
}

func NewForm(formId, objectId string) *Form {
	return &Form{
		ID:       formId,
		ObjectID: objectId,
		Fields:   make(map[string]*Field),
	}
}

func (f *Form) AddField(name, label, value string, required bool) *Field {
	f.Fields[name] = NewField(name, label, value, required)
	return f.Fields[name]
}
