package compose

import (
	"fltk"
	"fmt"
)

type CellRow int

func (r CellRow) String() string {
	return fmt.Sprintf("%d", r)
}

type CellCol int

func (r CellCol) String() string {
	return fmt.Sprintf("%d", r)
}

type CellLoc struct {
	Row CellRow
	Col CellCol
}

func (loc CellLoc) String() string {
	return fmt.Sprintf("[%d,%d]", loc.Row, loc.Col)
}

type CellData interface {
	Display() string
	Raw() any
}
type Cell struct {
	Loc      CellLoc
	Data     CellData
	RawValue string
}
type TableModel struct {
	Cells [][]*Cell
}

type _Table struct {
	Rect
	_Component
	wrap     *fltk.TableRow
	rowCount int
	colCount int
	model    TableModel
}

func (t *_Table) _Render() {
	t.applyModifier()
	t.rowCount = 10
	t.colCount = 3
	if t.wrap == nil {
		t.wrap = fltk.NewTableRow(t.x, t.y, t.width, t.height)
		t.wrap.SetRowCount(t.rowCount)
		t.wrap.SetColumnCount(t.colCount)
		t.wrap.EnableColumnHeaders()
		t.wrap.EnableRowHeaders()
		t.wrap.AllowColumnResizing()
		t.wrap.AllowRowResizing()
		t.wrap.SetDrawCellCallback(t.drawCell)
	}
}

func (t *_Table) drawCell(tc fltk.TableContext, i, j, x, y, w, h int) {
	row := CellRow(i)
	col := CellCol(j)

	switch tc {
	case fltk.ContextRowHeader:
		fltk.SetDrawFont(fltk.HELVETICA_BOLD, 14)
		fltk.DrawBox(fltk.UP_BOX, x, y, w, h, fltk.BACKGROUND_COLOR)
		fltk.SetDrawColor(fltk.BLACK)
		fltk.Draw(row.String(), x, y, w, h, fltk.ALIGN_CENTER)
	case fltk.ContextColHeader:
		fltk.SetDrawFont(fltk.HELVETICA_BOLD, 14)
		fltk.DrawBox(fltk.UP_BOX, x, y, w, h, fltk.BACKGROUND_COLOR)
		fltk.SetDrawColor(fltk.BLACK)
		fltk.Draw(col.String(), x, y, w, h, fltk.ALIGN_CENTER)
	case fltk.ContextCell:
		// loc := CellLoc{Row: row, Col: col}
		fltk.SetDrawFont(fltk.HELVETICA, 14)
		fltk.DrawBox(fltk.FLAT_BOX, x, y, w, h, fltk.BLACK)
		fltk.DrawBox(fltk.FLAT_BOX, x+1, y+1, w-2, h-2, fltk.WHITE)
		fltk.SetDrawColor(fltk.BLACK)
		if len(t.model.Cells) == 0 {
			return
		}
		fltk.Draw(t.model.Cells[row][col].RawValue, x, y, w, h, fltk.ALIGN_CENTER)
	}
}

func Table(m ...Modifier) Component {
	t := &_Table{}
	t.modifiers = m
	t.self = t
	return t
}
