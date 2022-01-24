package main

import (
	"testing"
)

//третье задание
func BenchmarkFirstSetAdd(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(10)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1.0)
			}
		})
	})
}

func BenchmarkFirstSetHas(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(90)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1.0)
			}
		})
	})
}

func BenchmarkSecondSetAdd(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(50)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1.0)
			}
		})
	})
}

func BenchmarkSecondSetHas(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(50)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1.0)
			}
		})
	})
}

func BenchmarkThirdSetAdd(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(90)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1.0)
			}
		})
	})
}

func BenchmarkThirdSetHas(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(10)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1.0)
			}
		})
	})
}

//RWMutex

func BenchmarkRwFirstSetAdd(b *testing.B) {
	var set = NewRwSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(10)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddRw(1.0)
			}
		})
	})
}

func BenchmarkRwFirstSetHas(b *testing.B) {
	var set = NewRwSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(90)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasRw(1.0)
			}
		})
	})
}

func BenchmarkRwSecondSetAdd(b *testing.B) {
	var set = NewRwSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(50)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddRw(1.0)
			}
		})
	})
}

func BenchmarkRwSecondSetHas(b *testing.B) {
	var set = NewRwSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(50)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasRw(1.0)
			}
		})
	})
}

func BenchmarkRwThirdSetAdd(b *testing.B) {
	var set = NewRwSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(90)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddRw(1.0)
			}
		})
	})
}

func BenchmarkRwThirdSetHas(b *testing.B) {
	var set = NewRwSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(10)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasRw(1.0)
			}
		})
	})
}
