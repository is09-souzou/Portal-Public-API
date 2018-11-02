package handler

import (
	"fmt"
	"strconv"

	"github.com/is09-souzou/Portal-Public-Api/src/model"
)

// ListWork list work struct
type ListWork struct {
	Limit             *int             `json:"limit"`
	ExclusiveStartKey *string          `json:"exclusiveStartKey"`
	Option            *WorkQueryOption `json:"option"`
}

// WorkQueryOption work query option struct
type WorkQueryOption struct {
	Tags   *[]string `json:"tags"`
	Word   *string   `json:"word"`
	UserID *string   `json:"userId"`
}

// ListWorkHandle List Work Handle
func ListWorkHandle(arg ListWork) (WorkConnection, error) {

	svc, err := model.GetSVC()

	if err != nil {
		return WorkConnection{}, err
	}

	limit := int64(10)
	if arg.Limit != nil {
		limit = int64(*arg.Limit)
	}

	var workList model.ScanWorkListResult
	if arg.Option != nil && arg.Option.Tags != nil {
		workList, err = model.ScanWorkListByTags(svc, limit, arg.ExclusiveStartKey, *arg.Option.Tags)
	} else {
		workList, err = model.ScanWorkList(svc, limit, arg.ExclusiveStartKey)
	}

	if err != nil {
		fmt.Println("Got error calling ListWorkHandle:")
		fmt.Println(err.Error())
		return WorkConnection{}, err
	}

	items := []Work{}

	for _, i := range workList.Items {
		item := Work{}

		item.ID = i.ID
		item.UserID = i.UserID
		item.Title = i.Title
		item.Tags = i.Tags
		item.ImageURL = i.ImageURL
		item.Description = i.Description
		createdAt, _ := strconv.Atoi(i.CreatedAt)
		item.CreatedAt = createdAt

		items = append(items, item)
	}

	return WorkConnection{items, workList.ExclusiveStartKey}, nil
}
