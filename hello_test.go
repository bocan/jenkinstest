package main
import "testing"
import "os"


func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}
