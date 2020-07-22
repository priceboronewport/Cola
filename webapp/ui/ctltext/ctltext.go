package ctltext

import (
	".."
	"../../../element"
)

type CtlText struct {
	Label *element.Element
	Input *element.Element
}

func New(pg *ui.Page, label string, id string) *CtlText {
	ctl := CtlText{Label: element.New("label"), Input: element.New("input")}
	ctl.Label.Attributes["class"] = "ctl"
	ctl.Label.InnerHTML = label
	ctl.Input.Attributes["class"] = "ctl"
	ctl.Input.Attributes["id"] = id
	ctl.Input.Attributes["name"] = id
	ctl.Input.Attributes["type"] = "text"
	pg.AddStylesheet("/res/css/ctl.css")
	return &ctl
}

func (ctl *CtlText) OuterHTML() string {
	if ctl.Label.InnerHTML == "" {
		return ctl.Input.OuterHTML()
	}
	return ctl.Label.OuterHTML() + "<br/>" + ctl.Input.OuterHTML()
}
