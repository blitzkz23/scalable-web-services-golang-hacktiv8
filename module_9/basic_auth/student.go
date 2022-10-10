package main

var students = []*Student{}

type Student struct {
	ID    string
	Name  string
	Grade int32
}

func init() {
	students = append(students, &Student{ID: "1", Name: "John Wick", Grade: 2})
	students = append(students, &Student{ID: "2", Name: "John Doe", Grade: 3})
	students = append(students, &Student{ID: "3", Name: "Jane Doe", Grade: 4})
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, student := range students {
		if student.ID == id {
			return student
		}
	}
	return nil
}
