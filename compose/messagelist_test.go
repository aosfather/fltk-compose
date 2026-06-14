package compose_test

import (
	"fmt"
	"testing"

	"github.com/aosfather/fltk-compose/compose"
)

func TestMessage(t *testing.T) {

	s := compose.NewStyle()
	fmt.Println(s.Bold().Large().ToFormat())
	s = compose.NewStyle()
	fmt.Println(s.Bold().Italic().Center().ToFormat())
	s = compose.NewStyle()
	fmt.Println(s.Bold().Italic().Underline().Center().ToFormat())
	s = compose.NewStyle()
	fmt.Println(s.Bold().Italic().Underline().Center().Color(12).ToFormat())
	fmt.Println(s.Bold().Italic().Underline().Center().Color(12).Font(15).Size(20).ToFormat())
	fmt.Println(compose.NewStyle().Color(10).Medium().ToFormat())
}
