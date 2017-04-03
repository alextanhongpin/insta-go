package photosvc

import (
	"errors"
)

type Service struct{}

func (s Service) All(request interface{}) ([]Photo, error) {
	var iter *mgo.Iter

	req := request.(allRequest)
	res := []Photo{}

	ds := common.NewDataStore()
	defer ds.Close()

	c := ds.C("photos")
	if req.Query == "" {
		iter = c.Find(nil).Iter()
	} else {
		iter = c.Find(req.Query).Iter()
	}

	var result Photo
	for iter.Next(&result) {
		res = append(res, result)
	}
	return res, nil
}
