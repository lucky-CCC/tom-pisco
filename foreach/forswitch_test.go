package foreach

import (
	"fmt"
	"testing"
)

func TestForSwitch(t *testing.T) {
	t.Log(foreach())
}

func foreach() error {
	for {
		switch 1 {
		case 1:
			return fmt.Errorf("hello")
		}
		return fmt.Errorf("no")
	}
}
