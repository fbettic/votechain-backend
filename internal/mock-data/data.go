package sampledata

import (
	"github.com/fbettic/votechain-backend/pkg/dto"
)

var Options = map[string]*dto.Option{
	"1": {
		ID:           "1",
		Name:         "Elisa Williams",
		Image:        "https://content.fakeface.rest/female_28_10b2cc0dce440f773d2ef9743179c110691a3974.jpg",
		Group:        "Republicanos",
		Presentation: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus auctor in velit eget ultricies.",
		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus auctor in velit eget ultricies. Curabitur facilisis nibh eget tortor volutpat vestibulum. Ut pretium mattis sollicitudin. Sed sit amet erat ut odio dignissim aliquet. Suspendisse tempor magna et ante feugiat, eu cursus nunc ullamcorper. Duis sed ante a metus blandit varius eu et sapien. Praesent at mauris lacus.",
		Proposals: []dto.Proposal{
			{
				Number: 1,
				Text:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			},
			{
				Number: 2,
				Text:   "Vivamus auctor in velit eget ultricies.",
			},
			{
				Number: 3,
				Text:   "Consectetur adipiscing elit. Vivamus auctor in velit eget ultricies.",
			},
			{
				Number: 4,
				Text:   "Curabitur facilisis nibh eget tortor volutpat vestibulum.",
			},
			{
				Number: 5,
				Text:   "Facilisis nibh eget tortor volutpat vestibulum. ",
			},
			
		},
	},
	"2": {
		ID:           "2",
		Name:         "John Jonson",
		Image:        "https://content.fakeface.rest/female_67_473c955a1a59b825e4b97a6c4a306a13d39fe9ae.jpg",
		Group:        "Frente Progresista",
		Presentation: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus auctor in velit eget ultricies.",
		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus auctor in velit eget ultricies. Curabitur facilisis nibh eget tortor volutpat vestibulum. Ut pretium mattis sollicitudin. Sed sit amet erat ut odio dignissim aliquet. Suspendisse tempor magna et ante feugiat, eu cursus nunc ullamcorper. Duis sed ante a metus blandit varius eu et sapien. Praesent at mauris lacus.",
		Proposals: []dto.Proposal{
			{
				Number: 1,
				Text:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			},
			{
				Number: 2,
				Text:   "Vivamus auctor in velit eget ultricies.",
			},
			{
				Number: 3,
				Text:   "Consectetur adipiscing elit. Vivamus auctor in velit eget ultricies.",
			},
			{
				Number: 4,
				Text:   "Curabitur facilisis nibh eget tortor volutpat vestibulum.",
			},
			{
				Number: 5,
				Text:   "Facilisis nibh eget tortor volutpat vestibulum. ",
			},
		},
	},
	"3": {
		ID:           "3",
		Name:         "Jane Jefferson",
		Image:        "https://content.fakeface.rest/female_11_68beb3baf7f77b8dd2958a112adf6e0417cb5c72.jpg",
		Group:        "Libertad",
		Presentation: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus auctor in velit eget ultricies.",
		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus auctor in velit eget ultricies. Curabitur facilisis nibh eget tortor volutpat vestibulum. Ut pretium mattis sollicitudin. Sed sit amet erat ut odio dignissim aliquet. Suspendisse tempor magna et ante feugiat, eu cursus nunc ullamcorper. Duis sed ante a metus blandit varius eu et sapien. Praesent at mauris lacus.",
		Proposals: []dto.Proposal{
			{
				Number: 1,
				Text:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			},
			{
				Number: 2,
				Text:   "Vivamus auctor in velit eget ultricies.",
			},
			{
				Number: 3,
				Text:   "Consectetur adipiscing elit. Vivamus auctor in velit eget ultricies.",
			},
			{
				Number: 4,
				Text:   "Curabitur facilisis nibh eget tortor volutpat vestibulum.",
			},
			{
				Number: 5,
				Text:   "Facilisis nibh eget tortor volutpat vestibulum. ",
			},
		},
	},
}

var Users = map[int]*dto.User{
	20586568548: {
		Name:        "Martin G.",
		Cuit:        20586568548,
		Password:    "123456",
		AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiTWF0aWFzIEd1dGllcnJleiIsImN1aXQiOjIwNTg2NTY4NTQ4fQ.Au1ind4DRn2izrgHQSNQKmEI3GCMTb7GwBezsckv3nA",
		HasVoted: false,
	},
	18668524894: {
		Name:        "Jorge G.",
		Cuit:        18668524894,
		Password:    "123456",
		AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9yZ2UgR29uemFsZXoiLCJjdWl0IjoxODY2ODUyNDg5NH0.k4ZSsaJ7WeKZwNyTte0ukGGs59ALURVqGVPQmeKBbys",
		HasVoted: false,
	},
	20405049768: {
		Name:        "Jose A.",
		Cuit:        20405049768,
		Password:    "123456",
		AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9zZSBBLiIsImN1aXQiOjIwNDA1MDQ5NzY4fQ.d7Bn5_DpbLxN2a7uVhujHQXmyYqs_qaFMMZTxxlUg2A",
		HasVoted: false,
	},
}
