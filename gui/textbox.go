package gui

import (
	"fmt"
)

// {{.Textbox id value}} adds a textbox to the document.
// value is the initial text in the box.
func (t *Templ) TextBox(id string, value string) string {
	return t.textbox("text", id, nil, value, "size=12")
}

// {{.Numbox id value}} adds a textbox for numbers to the document.
// value is the initial text in the box.
func (t *Templ) NumBox(id string, value float64) string {
	return t.textbox("text", id, &floatData{value}, value, "size=12")
}

// {{.IntBox id value}} adds a textbox for integer numbers to the document.
// value is the initial text in the box.
func (t *Templ) IntBox(id string, value int) string {
	return t.textbox("number", id, &intData{value}, value, "min=0 max=99999") // hack to get width right
}

// general textbox with data model (nil = default interfaceData)
func (t *Templ) textbox(typ string, id string, dm data, value interface{}, extraAttr string) string {
	d := (*Doc)(t)
	el := d.addElem(id)
	if dm != nil {
		el.data = dm
	}
	el.update = func(id string) jsCall {
		return jsCall{F: "setTextbox", Args: []interface{}{id, el.value()}}
	}
	return fmt.Sprintf(`<input type=%v class=TextBox id=%v value="%v" onchange="notifytextbox('%v')" onfocus="notifyfocus('%v')" onblur="notifyblur('%v')" %v />`, typ, id, value, id, id, id, extraAttr)
}
