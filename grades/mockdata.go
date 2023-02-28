package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Week1 Homework",
					Type:  GradeHomework,
					Score: 90,
				},
				{
					Title: "Quiz2",
					Type:  GradeQuiz,
					Score: 94,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 90,
				},
				{
					Title: "Week1 Homework",
					Type:  GradeHomework,
					Score: 86,
				},
				{
					Title: "Quiz2",
					Type:  GradeQuiz,
					Score: 99,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Jenny",
			LastName:  "Davids",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 77,
				},
				{
					Title: "Week1 Homework",
					Type:  GradeHomework,
					Score: 75,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 80,
				},
			},
		},
	}
}
