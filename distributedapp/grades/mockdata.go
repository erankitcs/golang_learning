package grades

func init() {
	students = []Student{
		Student{
			ID:        1,
			FirstName: "Averill",
			LastName:  "Simen",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 94,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		Student{
			ID:        2,
			FirstName: "Marge",
			LastName:  "Garrard",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 100,
				},
				Grade{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 100,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		Student{
			ID:        3,
			FirstName: "Sydnie",
			LastName:  "Barber",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 77,
				},
				Grade{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 0,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 65,
				},
			},
		},
		Student{
			ID:        4,
			FirstName: "Louie",
			LastName:  "Easton",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 88,
				},
				Grade{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 93,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 84,
				},
			},
		},
		Student{
			ID:        5,
			FirstName: "Kylee",
			LastName:  "Attwood",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 95,
				},
				Grade{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 100,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 97,
				},
			},
		},
	}
}
