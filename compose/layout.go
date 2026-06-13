package compose

type Layout interface {
	Component
	SetParentRect(*Rect)
}
type _ParentRect struct {
	parentRect *Rect
}

func (pr *_ParentRect) SetParentRect(r *Rect) {
	pr.parentRect = r
}

// 转换为绝对坐标
func (pr *_ParentRect) absoluteRect(r *Rect) Rect {
	t := Rect{}
	t.x = pr.parentRect.x + r.x
	t.y = pr.parentRect.y + r.y
	t.width = r.width
	t.height = r.height
	if t.x+r.width > pr.parentRect.x+pr.parentRect.width {
		t.width = pr.parentRect.x + pr.parentRect.width - t.x
	}

	if t.y+r.height > pr.parentRect.y+pr.parentRect.height {
		t.height = pr.parentRect.y + pr.parentRect.height - t.y
	}
	return t
}

// 轴上组件排列方式
type AxisAlignment byte

const (
	AA_Start        AxisAlignment = 0 //从上开始排列
	AA_Center       AxisAlignment = 1 //从下开始排列
	AA_End          AxisAlignment = 2 //从中间开始排列
	AA_SpaceBetween AxisAlignment = 3 //首尾子组件贴边，中间均匀分布
	AA_SpaceAround  AxisAlignment = 4 //子组件周围均匀分布间距。
	AA_SpaceEvenly  AxisAlignment = 5 //子组件和间距完全均匀分布
	AA_Stretch      AxisAlignment = 6 //拉伸子组件至填满水平空间
)

// 垂直布局，x相等，y均匀间隔，width小于等于，
type _Column struct {
	_ParentRect
	Rect
	_Component
	mainAxisAlignment  AxisAlignment
	crossAxisAlignment AxisAlignment
}

func (col *_Column) _Render() {
	col.applyModifier()
	x := col.x
	y := col.y
	if col.parentRect != nil {
		x += col.parentRect.x
		y += col.parentRect.y
	}
	for index, child := range col.children {
		child.AddModifier(Point(x, y+index*25))
		child.AddModifier(Height(25))
		child._Render()
	}

}

func Column(m ...Modifier) Component {
	col := &_Column{}
	col.modifiers = m
	col.self = col
	return col
}

type _Row struct {
	_ParentRect
	Rect
	_Component
	mainAxisAlignment  AxisAlignment
	crossAxisAlignment AxisAlignment
}

func (row *_Row) _Render() {
	row.applyModifier()
	x := row.x
	y := row.y
	if row.parentRect != nil {
		x += row.parentRect.x
		y += row.parentRect.y
	}
	for index, child := range row.children {
		child.AddModifier(Point(x+index*100, y))
		child.AddModifier(Width(98))
		child._Render()
	}

}

func Row(m ...Modifier) Component {
	row := &_Row{}
	row.modifiers = m
	row.self = row
	return row
}
