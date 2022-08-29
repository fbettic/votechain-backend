package sampledata

import (
	"github.com/fbettic/votechain-backend/pkg/dto"
)

var Options = map[string]*dto.Option{
	"1": {
		ID:          "1",
		Name:        "Elisa",
		Image:       "https://content.fakeface.rest/female_28_10b2cc0dce440f773d2ef9743179c110691a3974.jpg",
		Group:       "Republicanos",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus auctor in velit eget ultricies. Curabitur facilisis nibh eget tortor volutpat vestibulum. Ut pretium mattis sollicitudin. Sed sit amet erat ut odio dignissim aliquet. Suspendisse tempor magna et ante feugiat, eu cursus nunc ullamcorper. Duis sed ante a metus blandit varius eu et sapien. Praesent at mauris lacus.",
	},
	"2": {
		ID:          "2",
		Name:        "John",
		Image:       "https://content.fakeface.rest/female_67_473c955a1a59b825e4b97a6c4a306a13d39fe9ae.jpg",
		Group:       "Frente Progresista",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus auctor in velit eget ultricies. Curabitur facilisis nibh eget tortor volutpat vestibulum. Ut pretium mattis sollicitudin. Sed sit amet erat ut odio dignissim aliquet. Suspendisse tempor magna et ante feugiat, eu cursus nunc ullamcorper. Duis sed ante a metus blandit varius eu et sapien. Praesent at mauris lacus.",
	},
	"3": {
		ID:          "3",
		Name:        "Jane",
		Image:       "https://content.fakeface.rest/female_11_68beb3baf7f77b8dd2958a112adf6e0417cb5c72.jpg",
		Group:       "Libertad",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus auctor in velit eget ultricies. Curabitur facilisis nibh eget tortor volutpat vestibulum. Ut pretium mattis sollicitudin. Sed sit amet erat ut odio dignissim aliquet. Suspendisse tempor magna et ante feugiat, eu cursus nunc ullamcorper. Duis sed ante a metus blandit varius eu et sapien. Praesent at mauris lacus.",
	},
}
