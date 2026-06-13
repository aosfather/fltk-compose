package attr

import "github.com/aosfather/fltk-compose/compose"

func Point(x, y int) compose.Modifier {
	return compose.Point(x, y)
}

func Size(w, h int) compose.Modifier {
	return compose.Size(w, h)
}

func Title(t string) compose.Modifier {
	return compose.Title(t)
}

func Options(op []string) compose.Modifier {
	return compose.Options(op)
}
