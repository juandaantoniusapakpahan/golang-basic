package main


var students = []*Student{}

type Student struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Grade int `json:"grade"`
}

func GetStudents() []*Student{
	return students
}

func SelectStudent(id string) *Student{
	for _, value := range students {
		if value.Id == id{
			return value;
		}
	}
	return nil
}

func init() {
	students = append(students, &Student{Id: "1", Name: "Smith", Grade: 88})
	students = append(students, &Student{Id: "2", Name: "Jakop", Grade: 80})
	students = append(students, &Student{Id: "3", Name: "Salah", Grade: 90})
	students = append(students, &Student{Id: "4", Name: "Richard", Grade: 95})
}