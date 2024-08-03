package matrice

import (
	"testing"
)

func TestAddRow(t *testing.T) {
	m, err := NewMatrice([]Row{{1, 2, 3}, {4, 5, 6}}, 2, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = m.AddRow(Row{7, 8, 9})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if m.R != 3 {
		t.Fatalf("expected R to be 3, got %d", m.R)
	}

	if len(m.Rows) != 3 {
		t.Fatalf("expected 3 rows, got %d", len(m.Rows))
	}

	err = m.AddRow(Row{1, 2})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestAdd(t *testing.T) {
	m1, err := NewMatrice([]Row{{2, 5, 6}}, 1, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	m2, err := NewMatrice([]Row{{3, 5, 7}}, 1, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	result, err := m1.Add(&m2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedRows := []Row{{5, 10, 13}}
	for i := range result.Rows {
		for j := range result.Rows[i] {
			if result.Rows[i][j] != expectedRows[i][j] {
				t.Fatalf("expected %d, got %d", expectedRows[i][j], result.Rows[i][j])
			}
		}
	}

	m3, err := NewMatrice([]Row{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 3, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = m1.Add(&m3)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSubtract(t *testing.T) {
	m1, err := NewMatrice([]Row{{7, 8, 9}, {10, 11, 12}}, 2, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	m2, err := NewMatrice([]Row{{1, 2, 3}, {4, 5, 6}}, 2, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	result, err := m1.Subtract(&m2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedRows := []Row{{6, 6, 6}, {6, 6, 6}}
	for i := range result.Rows {
		for j := range result.Rows[i] {
			if result.Rows[i][j] != expectedRows[i][j] {
				t.Fatalf("expected %d, got %d", expectedRows[i][j], result.Rows[i][j])
			}
		}
	}

	m3, err := NewMatrice([]Row{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 3, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = m1.Subtract(&m3)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestMultiply(t *testing.T) {

	testCases := []struct {
		mS1 struct {
			R    int
			C    int
			Rows []Row
		}
		mS2 struct {
			R    int
			C    int
			Rows []Row
		}
		Name string

		expected                    []Row
		expectedMultiplicationError bool
	}{
		{

			Name: "Test for succes",

			mS1: struct {
				R    int
				C    int
				Rows []Row
			}{
				R: 2,
				C: 3,
				Rows: []Row{
					{1, 4, -2},
					{3, 5, -6},
				},
			},
			mS2: struct {
				R    int
				C    int
				Rows []Row
			}{
				R: 3,
				C: 4,
				Rows: []Row{
					{5, 2, 8, -1},
					{3, 6, 4, 5},
					{-2, 9, 7, -3},
				},
			},
			expected: []Row{
				{21, 8, 10, 25},
				{42, -18, 2, 40},
			},
			expectedMultiplicationError: false,
		},
		{
			Name: "Test for Failure, when we try multiply incompatible matrices",

			mS1: struct {
				R    int
				C    int
				Rows []Row
			}{
				R: 3,
				C: 4,
				Rows: []Row{
					{5, 2, 8, -1},
					{3, 6, 4, 5},
					{-2, 9, 7, -3},
				},
			},
			mS2: struct {
				R    int
				C    int
				Rows []Row
			}{
				R: 2,
				C: 3,
				Rows: []Row{
					{1, 4, -2},
					{3, 5, -6},
				},
			},

			expected:                    []Row{},
			expectedMultiplicationError: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			m1, err := NewMatrice(tc.mS1.Rows, tc.mS1.R, tc.mS1.C)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			m2, err := NewMatrice(tc.mS2.Rows, tc.mS2.R, tc.mS2.C)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			result, err := m1.Multiply(&m2)
			if tc.expectedMultiplicationError {
				if err == nil {
					t.Fatal("expected error, got nil")

				}
				return
			}
			if err != nil {

				t.Fatalf("expected no error, got %v", err)
			}
			for i := range result.Rows {
				for j := range result.Rows[i] {
					if result.Rows[i][j] != tc.expected[i][j] {
						t.Fatalf("expected %d, got %d", tc.expected[i][j], result.Rows[i][j])
					}
				}
			}
		})
	}

}

func TestNewMatrice(t *testing.T) {
	rows := []Row{{1, 2, 3}, {4, 5, 6}}
	m, err := NewMatrice(rows, 2, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if m.R != 2 || m.C != 3 {
		t.Fatalf("expected R=2 and C=3, got R=%d and C=%d", m.R, m.C)
	}

	rows = []Row{{1, 2, 3}, {4, 5}}
	_, err = NewMatrice(rows, 2, 3)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
