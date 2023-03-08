package main

var students = []*Student{}

type Student struct{
	Id string 
	Name string
	Grade int
}


func SelectStudent(id string) *Student{
	for _, value := range students{
		if value.Id == id {
			return value
		}
	}
	return nil
}

func GetAllStudent() []*Student {
	return students
}

func init() {
	students = append(students, &Student{Id: "1", Name: "Match", Grade: 88})
	students = append(students, &Student{Id: "2", Name: "Rejel", Grade: 88})
	students = append(students, &Student{Id: "3", Name: "Angle", Grade: 88})
	students = append(students, &Student{Id: "4", Name: "Cristiano", Grade: 88})
	students = append(students, &Student{Id: "5", Name: "Lionel", Grade: 88})

}