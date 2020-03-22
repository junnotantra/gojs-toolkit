package main

import (
	"encoding/json"

	"github.com/gopherjs/gopherjs/js"
	vue "github.com/oskca/gopherjs-vue"
)

type Model struct {
	*js.Object        // this is needed for bidirectional data bindings
	raw        string `js:"raw"`
	formatted  string `js:"formatted"`
	action     string `js:"action"`
	err        string `js:"err"`
}

func main() {
	m := &Model{
		Object: js.Global.Get("Object").New(),
	}

	m.formatted = ""
	m.raw = ""
	m.action = "0"
	m.err = ""

	// create the VueJS viewModel using a struct pointer
	vue.New("#app", m)
}

func (m *Model) Process() {
	switch m.action {
	case "1":
		m.uglify()
	default:
		m.prettify()
	}
}

// prettify will update formatted field on model with pretty-printed json from raw field
func (m *Model) prettify() {
	res, err := formatJson([]byte(m.raw), true, "\t")
	if err != nil {
		m.err = err.Error()
		return
	}

	m.formatted = string(res)
	m.err = ""
}

// uglify will update formatted field on model with pretty-printed json from raw field
func (m *Model) uglify() {
	res, err := formatJson([]byte(m.raw), false, "")
	if err != nil {
		m.err = err.Error()
		return
	}

	m.formatted = string(res)
	m.err = ""
}

func formatJson(raw []byte, wantPretty bool, indent string) ([]byte, error) {
	var (
		parsed map[string]interface{}
		err    error
	)

	err = json.Unmarshal(raw, &parsed)
	if err != nil {
		return nil, err
	}

	if wantPretty {
		if indent == "" {
			indent = "\t"
		}
		return json.MarshalIndent(parsed, "", indent)
	}
	return json.Marshal(parsed)
}
