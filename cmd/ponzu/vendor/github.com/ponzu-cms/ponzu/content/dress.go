package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Dress struct {
	item.Item

	Name          string   `json:"name"`
	Price         int      `json:"price"`
	Stock         int      `json:"stock"`
	AverageRating float32  `json:"average_rating"`
	Photos        []string `json:"photos"`
	Description   string   `json:"description"`
	FitInfo       string   `json:"fit_info"`
	Detail        string   `json:"detail"`
	TryOnStatus   string   `json:"try_on_status"`
	Line          string   `json:"line"`
	Tags          []string `json:"tags"`
}

// MarshalEditor writes a buffer of html to edit a Dress within the CMS
// and implements editor.Editable
func (d *Dress) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(d,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Dress field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", d, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Input("Price", d, map[string]string{
				"label":       "Price",
				"type":        "text",
				"placeholder": "Enter the Price here",
			}),
		},
		editor.Field{
			View: editor.Input("Stock", d, map[string]string{
				"label":       "Stock",
				"type":        "text",
				"placeholder": "Enter the Stock here",
			}),
		},
		editor.Field{
			View: editor.Select("TryOnStatus", d, map[string]string{
				"label": "TryOnStatus",
			}, map[string]string{
				"true":  "True",
				"false": "False",
			}),
		},
		editor.Field{
			View: editor.Input("AverageRating", d, map[string]string{
				"label":       "AverageRating",
				"type":        "text",
				"placeholder": "Enter the AverageRating here",
			}),
		},
		editor.Field{
			View: editor.FileRepeater("Photos", d, map[string]string{
				"label":       "Photos",
				"placeholder": "Upload the Photos here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", d, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Detail", d, map[string]string{
				"label":       "Detail",
				"placeholder": "Enter the Detail here",
			}),
		},
		editor.Field{
			View: editor.Input("FitInfo", d, map[string]string{
				"label":       "FitInfo",
				"type":        "text",
				"placeholder": "Enter the FitInfo here",
			}),
		},
		editor.Field{
			View: reference.Select("Line", d, map[string]string{
				"label": "Line",
			},
				"Line",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.Tags("Tags", d, map[string]string{
				"label":       "Tags",
				"placeholder": "Add your tag",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Dress editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Dress"] = func() interface{} { return new(Dress) }
}

// String defines how a Dress is printed. Update it using more descriptive
// fields from the Dress struct type
func (d *Dress) String() string {
	return fmt.Sprintf("Dress: %s", d.Name)
}
