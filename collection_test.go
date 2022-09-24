package x

import (
	"reflect"
	"testing"
)

func TestCollection_Collect(t *testing.T) {
	type CollectArgs[V comparable] struct {
		name  string
		items []V
		want  *Collection[V]
	}

	for _, tt := range []CollectArgs[int]{
		{
			name:  "Int_1",
			items: []int{1, 2, 3},
			want:  &Collection[int]{items: []int{1, 2, 3}},
		},
		{
			name:  "Int_1",
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
	type AppendArgs[V comparable] struct {
		name  string
		items []V
		input []V
		want  *Collection[V]
	}

	for _, tt := range []AppendArgs[int]{
		{
			name:  "Int_1",
			items: []int{1, 2},
			input: []int{3},
			want:  Collect([]int{1, 2, 3}),
		},
		{
			name:  "Int_1",
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
	for _, tt := range []AppendArgs[User]{
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
	type ReplaceArgs[V comparable] struct {
		name  string
		items []V
		old   V
		new   V
		n     int
		want  *Collection[V]
	}

	for _, tt := range []ReplaceArgs[int]{
		{
			name:  "Int1",
			items: []int{1, 2, 3},
			old:   1,
			new:   2,
			n:     1,
			want:  Collect([]int{2, 2, 3}),
		},
		{
			name:  "Int2",
			items: []int{1, 1, 3},
			old:   1,
			new:   2,
			n:     2,
			want:  Collect([]int{2, 2, 3}),
		},
		{
			name:  "Int3",
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
	for _, tt := range []ReplaceArgs[User]{
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
	type DiffArgs[V comparable] struct {
		name  string
		items []V
		input []V
		want  *Collection[V]
	}

	for _, tt := range []DiffArgs[int]{
		{
			name:  "Int1",
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
	for _, tt := range []DiffArgs[User]{
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
