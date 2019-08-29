package iterator

import "testing"

func TestIteration(t *testing.T) {

	t.Run("recieve a letter and repeat it five times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %s but got %s", expected, repeated)
		}
	})

	t.Run("recieve a character and the number of times that the character should be repeated", func(t *testing.T) {
		repeated := RepeatTimes("a", 7)
		expected := "aaaaaaa"

		if repeated != expected {
			t.Errorf("expected %s but got %s", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
