package tui

import (
	"fmt"

	"github.com/RickardAhlstedt/GitDash/repo"
	"github.com/RickardAhlstedt/GitDash/style"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Render(statuses []*repo.RepoStatus, sortBy string) {
	app := tview.NewApplication()

	table := tview.NewTable().
		SetBorders(false).
		SetFixed(1, 0)

	headers := []string{"Name", "Branch", "↑", "↓", "Dirty", "Path"}
	for i, h := range headers {
		table.SetCell(0, i, tview.NewTableCell(h).
			SetTextColor(style.Color("header")).
			SetAlign(tview.AlignCenter).
			SetSelectable(false).
			SetAttributes(tcell.AttrBold))
	}

	for row, s := range statuses {
		dirtyIcon := ""

		if s.Dirty {
			dirtyIcon = "✴"
		}

		table.SetCell(row+1, 0, tview.NewTableCell(s.Name))
		table.SetCell(row+1, 1, tview.NewTableCell(s.Branch).
			SetTextColor(style.Color("branch")))
		table.SetCell(row+1, 2, tview.NewTableCell(fmt.Sprintf("%d", s.Ahead)).
			SetTextColor(ifColor(s.Ahead > 0, style.Color("ahead"), style.Color("path"))))
		table.SetCell(row+1, 3, tview.NewTableCell(fmt.Sprintf("%d", s.Behind)).
			SetTextColor(ifColor(s.Behind > 0, style.Color("behind"), style.Color("path"))))
		table.SetCell(row+1, 4, tview.NewTableCell(dirtyIcon).
			SetTextColor(ifColor(s.Dirty, style.Color("dirty"), style.Color("clean"))))
		table.SetCell(row+1, 5, tview.NewTableCell(s.Path).
			SetTextColor(style.Color("path")))
	}

	table.SetTitle(fmt.Sprintf(" gitdash - sorted by: %s ", sortBy)).
		SetBorder(true)

	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func ifColor(cond bool, yes, no tcell.Color) tcell.Color {
	if cond {
		return yes
	}
	return no
}
