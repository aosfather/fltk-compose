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

type _Model_Row []string
type TableModel struct {
	Cells    [][]*Cell
	colCount int
	cols     []string
	rows     []_Model_Row
}

func (m *TableModel) Count() int {
	return len(m.rows)
}

func (m *TableModel) AddRow(row ...string) {
	r := make(_Model_Row, m.colCount)
	copy(r, row)
	m.rows = append(m.rows, r)
}

func (m *TableModel) Cell(r, c int) string {
	if c >= m.colCount || r < 0 || r >= m.Count() {
		return ""
	}

	return m.rows[r][c]
}

type TableView interface {
	AddRow(row ...string)
}

type _Table struct {
	Rect
	_Component
	wrap  *fltk.TableRow
	model TableModel
}

func (t *_Table) SetOptions(cols []string) {
	t.model.cols = cols
	t.model.colCount = len(t.model.cols)
}
func (t *_Table) _Render() {
	t.applyModifier()
	if t.wrap == nil {
		t.wrap = fltk.NewTableRow(t.x, t.y, t.width, t.height)
		t.wrap.SetRowCount(0)
		t.wrap.SetColumnCount(t.model.colCount)
		t.wrap.EnableColumnHeaders()
		// t.wrap.EnableRowHeaders()
		t.wrap.AllowColumnResizing()
		// t.wrap.SetColumnWidth(0, 20)
		// t.wrap.SetColumnWidth(1, 20)
		// t.wrap.SetColumnWidth(2, 180)
		t.wrap.AllowRowResizing()
		t.wrap.SetDrawCellCallback(t.drawCell)
	}
}
func (t *_Table) AddRow(row ...string) {
	t.model.AddRow(row...)
	t.wrap.SetRowCount(t.model.Count())
	// t.wrap.Redraw()
}

func (t *_Table) Bind(o *BindObj[TableView]) *_Table {
	if o != nil {
		o.getter = func() TableView { return t }
		o.setter = func(TableView) {}
	}
	return t
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
		fltk.Draw(t.model.cols[col], x, y, w, h, fltk.ALIGN_CENTER)
	case fltk.ContextCell:
		// loc := CellLoc{Row: row, Col: col}
		fltk.SetDrawFont(fltk.HELVETICA, 14)
		fltk.DrawBox(fltk.FLAT_BOX, x, y, w, h, fltk.LIGHT1)
		fltk.DrawBox(fltk.FLAT_BOX, x+1, y+1, w-2, h-2, fltk.WHITE)
		fltk.SetDrawColor(fltk.BLACK)
		if t.model.Count() == 0 {
			return
		}
		fltk.Draw(t.model.rows[row][col], x, y, w, h, fltk.ALIGN_CENTER)
	}
}

func Table(m ...Modifier) *_Table {
	t := &_Table{}
	t.modifiers = m
	t.self = t
	return t
}
