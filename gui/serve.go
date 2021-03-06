package gui

import (
	"encoding/json"
	"net/http"
	"time"
)

// ServeHTTP implements http.Handler.
func (d *Doc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.keepAlive = time.Now()
	switch r.Method {
	default:
		http.Error(w, "not allowed: "+r.Method+" "+r.URL.Path, http.StatusForbidden)
	case "GET":
		d.serveContent(w, r)
	case "POST":
		d.serveRefresh(w, r)
	case "PUT":
		d.serveEvent(w, r)
	}
}

func (d *Doc) KeepAlive() time.Time {
	return d.keepAlive
}

// serves the html content.
func (d *Doc) serveContent(w http.ResponseWriter, r *http.Request) {
	w.Write(d.htmlCache)
}

// HTTP handler for event notifications by button clicks etc
func (d *Doc) serveEvent(w http.ResponseWriter, r *http.Request) {
	var ev event
	check(json.NewDecoder(r.Body).Decode(&ev))
	el := d.elem(ev.ID)
	el.setValue(ev.Arg)
	if el.onevent != nil {
		el.onevent()
	}
}

type event struct {
	ID  string
	Arg interface{}
}

// HTTP handler for refreshing the dynamic elements
func (d *Doc) serveRefresh(w http.ResponseWriter, r *http.Request) {
	d.onRefresh()
	calls := make([]jsCall, 0, len(d.elems))
	for id, el := range d.elems {
		calls = append(calls, el.update(id))
	}
	check(json.NewEncoder(w).Encode(calls))
}

// javascript call
type jsCall struct {
	F    string        // function to call
	Args []interface{} // function arguments
}
