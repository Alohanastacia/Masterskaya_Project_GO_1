package test

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/processors/mocks"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func TestCreateComplaintshandlers(t *testing.T) {
	//не представляю как написать с моками
	cases := []entity.CreateComplaint{
		{
			Priority:    "Срочно",
			Description: "bla bla bla",
			Category:    "Ka",
			Created_at:  time.Now(),
		},
		{
			Priority:    "",
			Description: "",
			Category:    "",
			Created_at:  time.Now(),
		},
	}
	for _, tc := range cases {
		ctrl := gomock.NewController(t)
		m := mocks.NewMockComplaintsRepository(ctrl)
		m.EXPECT().CreateComplaints(tc).Return()

	}

}
