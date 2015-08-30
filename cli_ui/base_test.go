package cli_ui

import (
	"experiments/fivez/core"
	"io/ioutil"
	"os"
	"testing"
)

func TestDraw(t *testing.T) {
	game := core.NewGame()

	output := captureOutput(func() {
		draw(game)
	})

	shouldBeString := `/===========================\
|    0 |    0 |    0 |    0 |
|===========================|
|    0 |    0 |    2 |    0 |
|===========================|
|    0 |    0 |    0 |    0 |
|===========================|
|    0 |    0 |    0 |    0 |
|===========================|
| Next number: |      2     |
| Score: |             0    |
\===========================/`

	if output != shouldBeString {
		t.Errorf("Expected things to be drawn as %v, but got %v.", output, shouldBeString)
	}
}

func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}
