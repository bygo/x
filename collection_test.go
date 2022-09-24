package x

import (
	"reflect"
	"testing"
)

func TestCollection_Collect(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		want  *Collection[V]
	}

	for _, tt := range []Args[int]{
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
		t.Run(tt.name, func(t *testing.T) {
			if got := Collect(tt.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Append(t *testing.T) {
	type Args[V comparable] struct {
		name  string
		items []V
		input []V
		want  *Collection[V]
	}

	for _, tt := range []Args[int]{
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
		t.Run(tt.name, func(t *testing.T) {
			if got := Collect(tt.items).Append(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect.Append() = %v, want %v", got, tt.want)
			}
		})
	}

	type User struct {
		ID int
	}
	for _, tt := range []Args[User]{
		{
			name:  "Struct_1",
			items: []User{{1}},
			input: []User{{2}, {3}, {4}},
			want:  Collect([]User{{1}, {2}, {3}, {4}}),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := Collect(tt.items).Append(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect.Append() = %v, want %v", got, tt.want)
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

	for _, tt := range []Args[int]{
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
		t.Run(tt.name, func(t *testing.T) {
			if got := Collect(tt.items).Replace(tt.old, tt.new, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect().Replace() = %v, want %v", got, tt.want)
			}
		})
	}

	type User struct {
		ID int
	}
	for _, tt := range []Args[User]{
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
		t.Run(tt.name, func(t *testing.T) {
			if got := Collect(tt.items).Replace(tt.old, tt.new, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect().Replace() = %v, want %v", got, tt.want)
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

	for _, tt := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1, 2, 2, 3},
			input: []int{1, 2},
			want:  Collect([]int{3}),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := Collect(tt.items).Diff(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect().Replace() = %v, want %v", got, tt.want)
			}
		})
	}

	type User struct {
		ID int
	}
	for _, tt := range []Args[User]{
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
		t.Run(tt.name, func(t *testing.T) {
			if got := Collect(tt.items).Diff(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect().Diff() = %v, want %v", got, tt.want)
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

	for _, tt := range []Args[int]{
		{
			name:    "Int_1",
			items:   []int{1, 2, 3, 4},
			want:    []int{1, 2, 3, 4},
			wantKey: []int{0, 1, 2, 3},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			gotKey := []int{}
			Collect(tt.items).ForEach(func(v int, k int) {
				got = append(got, v)
				gotKey = append(gotKey, k)
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect().ForEach().values = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(gotKey, tt.wantKey) {
				t.Errorf("Collect().ForEach().keys = %v, want %v", gotKey, tt.wantKey)
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

	for _, tt := range []Args[int]{
		{
			name:    "Int_1",
			items:   []int{1, 2, 3, 4},
			want:    []int{4, 3, 2, 1},
			wantKey: []int{3, 2, 1, 0},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			gotKey := []int{}
			Collect(tt.items).ForEachRight(func(v int, k int) {
				got = append(got, v)
				gotKey = append(gotKey, k)
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect().ForEachRight().values = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(gotKey, tt.wantKey) {
				t.Errorf("Collect().ForEachRight().keys = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}

func TestCollection_Map(t *testing.T) {
	type Args[V comparable] struct {
		name    string
		items   []V
		want    []V
		wantKey []V
	}

	for _, tt := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1, 2, 3, 4},
			want:  []int{2, 3, 4, 8},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := Collect(tt.items).Map(func(v int, k int) int {
				if k == 3 {
					return v * 2
				}
				return v + 1
			}).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect().Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Filter(t *testing.T) {
	type Args[V comparable] struct {
		name    string
		items   []V
		want    []V
		wantKey []V
	}

	for _, tt := range []Args[int]{
		{
			name:  "Int_1",
			items: []int{1, 2, 3, 4},
			want:  []int{2},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := Collect(tt.items).Filter(func(v int, k int) bool {
				return v == 2
			}).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect().Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
