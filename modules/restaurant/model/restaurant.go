package restaurantmodel

import "time"

type Restaurant struct {
	Id        int    `json:"id" gorm:"primaryKe y"`
	OwnerId   int    `json:"owner_id" gorm:"owner_id"`
	Name      string `json:"name" gorm:"name"`
	Addr      string `json:"addr" gorm:"addr"`
	CityId    int    `json:"city_id" gorm:"city_id"`
	Status    int    `json:"status" gorm:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RestaurantCreate struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Addr string `json:"addr"`
}

type RestaurantUpdate struct {
	Id   int     `json:"id"`
	Name *string `json:"name"`
	Addr *string `json:"addr"`
}
