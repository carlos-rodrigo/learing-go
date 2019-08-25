package integers

import "testing"

func TestAdder(t *testing.T) {

	t.Run("sum 2 plus 2 and expect 4 as result", func(t *testing.T) {
		sum := Add(2, 2)
		expect := 4

		if sum != expect {
			t.Errorf("expected %d but got %d", expect, sum)
		}
	})

	t.Run("sum 2 plus 3 and expect 5 as result", func(t *testing.T) {
		sum := Add(2, 3)
		expect := 5

		if sum != expect {
			t.Errorf("expected %d but got %d", expect, sum)
		}
	})
}
