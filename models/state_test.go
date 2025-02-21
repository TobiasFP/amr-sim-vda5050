package models

import "testing"

func TestFlyTowardsNode(t *testing.T) {
	state := GetDefaultState("1234")
	// We start at 150, 150.
	// Lets move towrads 140,160
	// For now, we move 1 unit towards each [x,y]
	state.FlyTowardsNode(140, 160, 1)
	if state.AgvPosition.X != 149 {
		t.Errorf("got %.2f, wanted %.2f", state.AgvPosition.X, 149.0)
	}
	if state.AgvPosition.Y != 151 {
		t.Errorf("got %.2f, wanted %.2f", state.AgvPosition.X, 148.0)
	}
}
