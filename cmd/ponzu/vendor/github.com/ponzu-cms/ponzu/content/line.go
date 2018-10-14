package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Line struct {
	item.Item

	Name string `json:"name"`
}

// MarshalEditor writes a buffer of html to edit a Line within the CMS
// and implements editor.Editable
func (l *Line) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(l,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Line field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", l, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Line editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Line"] = func() interface{} { return new(Line) }
}

// String defines how a Line is printed. Update it using more descriptive
// fields from the Line struct type
func (l *Line) String() string {
	return fmt.Sprintf("Line: %s", l.Name)
}
