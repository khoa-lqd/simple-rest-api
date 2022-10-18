package restaurantstorage

import (
	"context"
	restaurantmodel "simple-rest-api/modules/restaurant/model"
)

// InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error -- from biz

func (store *sqlStore) InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {

	if err := store.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
