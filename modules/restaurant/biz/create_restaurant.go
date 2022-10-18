package restaurantbiz

import (
	"context"
	restaurantmodel "simple-rest-api/modules/restaurant/model"
)

type CreateRestaurantStore interface {
	InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := biz.store.InsertRestaurant(ctx, data); err != nil {
		return err
	}
	return nil
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}
