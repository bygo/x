package x

import (
	"math"
	"reflect"
	"testing"
)

func TestCollection_Collect(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		want  *Collection[V]
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1, 2, 3},
			want:  &Collection[int]{items: []int{1, 2, 3}},
		},
		{
			name:  "Int_2",
			items: []int{},
			want:  &Collection[int]{items: []int{}},
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			if got := Collect(args.items); !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_Merge(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		input []V
		want  *Collection[V]
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1, 2},
			input: []int{3},
			want:  Collect([]int{1, 2, 3}),
		},
		{
			name:  "Int_2",
			items: []int{},
			input: nil,
			want:  Collect([]int{}),
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			if got := Collect(args.items).Merge(args.input); !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect.Merge() = %v, want %v", got, args.want)
			}
		})
	}

	type User struct {
		ID int
	}
	for _, v := range []Args[User]{
		{
			name:  "Struct_1",
			items: []User{{1}},
			input: []User{{2}, {3}, {4}},
			want:  Collect([]User{{1}, {2}, {3}, {4}}),
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			if got := Collect(args.items).Merge(args.input); !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect.Merge() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_Replace(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		old   V
		new   V
		n     int
		want  *Collection[V]
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1, 2, 3},
			old:   1,
			new:   2,
			n:     1,
			want:  Collect([]int{2, 2, 3}),
		},
		{
			name:  "Int_2",
			items: []int{1, 1, 3},
			old:   1,
			new:   2,
			n:     2,
			want:  Collect([]int{2, 2, 3}),
		},
		{
			name:  "Int_3",
			items: []int{1, 1, 1},
			old:   1,
			new:   2,
			n:     -1,
			want:  Collect([]int{2, 2, 2}),
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			if got := Collect(args.items).Replace(args.old, args.new, args.n); !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Replace() = %v, want %v", got, args.want)
			}
		})
	}

	type User struct {
		ID int
	}
	for _, v := range []Args[User]{
		{
			name:  "Struct_1",
			items: []User{{1}},
			old:   User{1},
			new:   User{2},
			n:     0,
			want:  Collect([]User{{1}}),
		},
		{
			name:  "Struct_2",
			items: []User{{1}, {1}, {3}, {1}},
			old:   User{1},
			new:   User{2},
			n:     -1,
			want:  Collect([]User{{2}, {2}, {3}, {2}}),
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			if got := Collect(args.items).Replace(args.old, args.new, args.n); !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Replace() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_Diff(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		input []V
		want  *Collection[V]
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1, 2, 2, 3},
			input: []int{1, 2},
			want:  Collect([]int{3}),
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			if got := Collect(args.items).Diff(args.input); !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Replace() = %v, want %v", got, args.want)
			}
		})
	}

	type User struct {
		ID int
	}
	for _, v := range []Args[User]{
		{
			name:  "Struct_1",
			items: []User{{1}, {2}},
			input: []User{{1}},
			want:  Collect([]User{{2}}),
		},
		{
			name:  "Struct_2",
			items: []User{{1}, {1}, {3}, {1}},
			input: []User{{3}},
			want:  Collect([]User{{1}, {1}, {1}}),
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			if got := Collect(args.items).Diff(args.input); !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Diff() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_DiffBy(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		input []V
		want  *Collection[V]
	}

	for _, v := range []Args[float64]{
		{
			name:  "Int_1",
			items: []float64{1.1, 2.9, 3.1, 5.55},
			input: []float64{5.1, 2.7},
			want:  Collect([]float64{1.1, 3.1}),
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			if got := Collect(args.items).DiffBy(args.input, math.Floor); !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Replace() = %v, want %v", got, args.want)
			}
		})
	}

	type User struct {
		ID int
	}
	for _, v := range []Args[User]{
		{
			name:  "Struct_1",
			items: []User{{1}, {2}},
			input: []User{{1}},
			want:  Collect([]User{{2}}),
		},
		{
			name:  "Struct_2",
			items: []User{{1}, {1}, {3}, {1}},
			input: []User{{3}},
			want:  Collect([]User{{1}, {1}, {1}}),
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			if got := Collect(args.items).Diff(args.input); !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Diff() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_ForEach(t *testing.T) {
	type Args[V comparable] struct {
		name    string
		items   []V
		want    []V
		wantKey []V
	}

	for _, v := range []Args[int]{
		{
			name:    "Int_1",
			items:   []int{1, 2, 3, 4},
			want:    []int{1, 2, 3, 4},
			wantKey: []int{0, 1, 2, 3},
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			got := []int{}
			gotKey := []int{}
			Collect(args.items).ForEach(func(v int, k int) {
				got = append(got, v)
				gotKey = append(gotKey, k)
			})
			if !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().ForEach().values = %v, want %v", got, args.want)
			}

			if !reflect.DeepEqual(gotKey, args.wantKey) {
				t.Errorf("Collect().ForEach().keys = %v, want %v", gotKey, args.wantKey)
			}
		})
	}
}

func TestCollection_ForEachRight(t *testing.T) {
	type Args[V comparable] struct {
		name    string
		items   []V
		want    []V
		wantKey []V
	}

	for _, v := range []Args[int]{
		{
			name:    "Int_1",
			items:   []int{1, 2, 3, 4},
			want:    []int{4, 3, 2, 1},
			wantKey: []int{3, 2, 1, 0},
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			got := []int{}
			gotKey := []int{}
			Collect(args.items).ForEachRight(func(v int, k int) {
				got = append(got, v)
				gotKey = append(gotKey, k)
			})
			if !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().ForEachRight().values = %v, want %v", got, args.want)
			}

			if !reflect.DeepEqual(gotKey, args.wantKey) {
				t.Errorf("Collect().ForEachRight().keys = %v, want %v", gotKey, args.wantKey)
			}
		})
	}
}

func TestCollection_Map(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		want  []V
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1, 2, 3, 4},
			want:  []int{2, 3, 4, 8},
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			got := Collect(args.items).Map(func(v int, k int) int {
				if k == 3 {
					return v * 2
				}
				return v + 1
			}).ToSlice()
			if !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Map() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_Filter(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		want  []V
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1, 2, 3, 4},
			want:  []int{2},
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			got := Collect(args.items).Filter(func(v int, k int) bool {
				return v == 2
			}).ToSlice()
			if !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Filter() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_ToSlice(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		want  []V
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			got := Collect(args.items).ToSlice()
			if !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().ToSlice() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_Chunk(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		size  int
		items []V
		want  [][]V
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			size:  1,
			items: []int{1, 2, 3, 4, 5},
			want:  [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			name:  "Int_2",
			size:  2,
			items: []int{1, 2, 3, 4, 5},
			want:  [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:  "Int_3",
			size:  3,
			items: []int{1, 2, 3, 4, 5},
			want:  [][]int{{1, 2, 3}, {4, 5}},
		},
		{
			name:  "Int_4",
			size:  4,
			items: []int{1, 2, 3, 4, 5},
			want:  [][]int{{1, 2, 3, 4}, {5}},
		},
		{
			name:  "Int_5",
			size:  5,
			items: []int{1, 2, 3, 4, 5},
			want:  [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name:  "Int_6",
			size:  6,
			items: []int{1, 2, 3, 4, 5},
			want:  [][]int{{1, 2, 3, 4, 5}},
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			got := Collect(args.items).Chunk(args.size)
			if !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Chunk() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_Sum(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		want  int
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1},
			want:  1,
		},
		{
			name:  "Int_2",
			items: []int{1, 2},
			want:  3,
		},
		{
			name:  "Int_3",
			items: []int{1, 2, 3},
			want:  6,
		},
		{
			name:  "Int_4",
			items: []int{1, 2, 3, 4, 5},
			want:  15,
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			got := Collect(args.items).Sum(func(v int) int {
				return v
			})
			if !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Sum() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_Avg(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		want  int
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1},
			want:  1,
		},
		{
			name:  "Int_2",
			items: []int{1, 2},
			want:  1,
		},
		{
			name:  "Int_3",
			items: []int{1, 2, 3},
			want:  2,
		},
		{
			name:  "Int_4",
			items: []int{1, 2, 3, 4, 5},
			want:  3,
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			got := Collect(args.items).Avg(func(v int) int {
				return v
			})
			if !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Avg() = %v, want %v", got, args.want)
			}
		})
	}
}

func TestCollection_Reverse(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		want  []V
	}

	for _, v := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1},
			want:  []int{1},
		},
		{
			name:  "Int_2",
			items: []int{1, 2},
			want:  []int{2, 1},
		},
		{
			name:  "Int_3",
			items: []int{1, 2, 3},
			want:  []int{3, 2, 1},
		},
	} {
		var args = v
		t.Run(args.name, func(t *testing.T) {
			got := Collect(args.items).Reverse().ToSlice()
			if !reflect.DeepEqual(got, args.want) {
				t.Errorf("Collect().Reverse() = %v, want %v", got, args.want)
			}
		})
	}
}
