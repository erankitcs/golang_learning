package employee

import "time"

var employees = []Employee{
	{
		ID:        962134,
		FirstName: "Jennifer",
		LastName:  "Wat",
		Position:  "CEO",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  30,
		TimeOff: []TimeOff{
			{
				Type:      TimeoffTypeHoliday,
				Amount:    8.,
				StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				Status:    TimeoffStatusTaken,
			}, {
				Type:      TimeoffTypePTO,
				Amount:    16.,
				StartDate: time.Date(2023, 8, 16, 0, 0, 0, 0, time.UTC),
				Status:    TimeoffStatusScheduled,
			}, {
				Type:      TimeoffTypePTO,
				Amount:    16.,
				StartDate: time.Date(2023, 12, 8, 0, 0, 0, 0, time.UTC),
				Status:    TimeoffStatusRequested,
			},
		},
	},
	{
		ID:        176158,
		FirstName: "Allison",
		LastName:  "Jane",
		Position:  "COO",
		StartDate: time.Now().Add(-4 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  20,
		TimeOff:   nil,
	},
	{
		ID:        160898,
		FirstName: "Aakar",
		LastName:  "Uppal",
		Position:  "CTO",
		StartDate: time.Now().Add(-6 * time.Hour * 24 * 365),
		TotalPTO:  20,
		TimeOff:   nil,
	},
	{
		ID:        297365,
		FirstName: "Jonathon",
		LastName:  "Anderson",
		Position:  "Worker Bee",
		StartDate: time.Now().Add(-12 * time.Hour * 24 * 365),
		TotalPTO:  30,
		TimeOff:   nil,
	},
}
