package ui

import (
	"github.com/daylioti/docker-commander/config"
	"github.com/daylioti/docker-commander/docker"
	"github.com/gizak/termui"
)

type UI struct {
	Cmd *Commands
}

func (ui *UI) Init(cnf *config.Config, dockerClient *docker.Docker) {

	termui.Body.AddRows(
		termui.NewRow(),
		termui.NewRow(),
		termui.NewRow(),
	)

	ui.Cmd = &Commands{}
	ui.Cmd.Init(cnf, dockerClient)
}

func StringColor(text string, color string) string {
	return "[" + text + "](" + color + ")"
}
