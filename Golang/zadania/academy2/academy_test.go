package academy

import (
	//"github.com/grupawp/akademia-programowania/Golang/zadania/academy2/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGradeStudent(t *testing.T) {

	t.Run("Student_Mock not found", func(t *testing.T) {
		repository := NewRepository(t)
		repository.On("Get", "John").Return(nil, ErrStudentNotFound)
		err := GradeStudent(repository, "John")
		assert.Equal(t, nil, err)
	})

	t.Run("Stay at the same year", func(t *testing.T) {
		repository := NewRepository(t)

		student := Sophomore{
			name:       "Bartosz",
			grades:     []int{2, 2, 3, 2, 2},
			project:    2,
			attendance: []bool{true, false, true, true, false},
		}

		repository.On("Get", "Bartosz").Return(student, nil)
		repository.On("Save", "Bartosz", student.Year()).Return(nil)

		err := GradeStudent(repository, "Bartosz")

		assert.Equal(t, nil, err)
		repository.AssertNotCalled(t, "Save", student.Year()+1)
	})

	t.Run("Go to the next year", func(t *testing.T) {
		repository := NewRepository(t)

		student := Sophomore{
			name:       "Bartosz",
			grades:     []int{5, 5, 5, 5, 4},
			project:    4,
			attendance: []bool{true, true, true, true, false},
		}

		repository.On("Get", "Bartosz").Return(student, nil)
		repository.On("Save", "Bartosz", student.Year()+1).Return(nil)

		err := GradeStudent(repository, "Bartosz")

		assert.Equal(t, nil, err)
		repository.AssertCalled(t, "Save", student.Name(), student.Year()+1)
	})

	t.Run("Graduate", func(t *testing.T) {
		repository := NewRepository(t)
		student := NewStudent(t)

		repository.On("Get", "Bartosz").Return(student, nil)
		student.On("FinalGrade").Return(5)
		student.On("Name").Return("Bartosz")
		student.On("Year").Return(uint8(3))
		repository.On("Graduate", "Bartosz").Return(nil)

		err := GradeStudent(repository, "Bartosz")

		assert.Equal(t, nil, err)
		repository.AssertCalled(t, "Graduate", student.Name())
	})

	t.Run("Bad attendance", func(t *testing.T) {
		repository := NewRepository(t)

		student := Sophomore{
			name:       "Wojtek",
			grades:     []int{5, 5, 5, 4, 4},
			project:    4,
			attendance: []bool{false, false, false, false, false},
		}

		repository.On("Get", "Wojtek").Return(student, nil)
		repository.On("Save", "Wojtek", student.Year()).Return(nil)

		err := GradeStudent(repository, "Wojtek")

		assert.Equal(t, nil, err)
		repository.AssertNotCalled(t, "Save", student.Year()+1)
	})

	t.Run("Invalid Grade", func(t *testing.T) {
		repository := NewRepository(t)
		student := NewStudent(t)

		repository.On("Get", "Bartosz").Return(student, nil)
		student.On("FinalGrade").Return(6)

		err := GradeStudent(repository, "Bartosz")

		assert.Equal(t, ErrInvalidGrade, err)
	})
}
