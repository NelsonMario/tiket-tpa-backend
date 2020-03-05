package query

import (
	"github.com/NelsonMario/graphql/query/type"
	"github.com/NelsonMario/graphql/resolvers"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object {

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{

			"admins": {
				Type:        graphql.NewList(types.GetAdminType()),
				Resolve:     resolvers.GetAllAdmin,
				Description: "Get All Admins",
			},
			"admin": {
				Type: graphql.NewList(types.GetAdminType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:     resolvers.GetAdmin,
				Description: "Get Single Users",
			},

			"sliders": {
				Type:        graphql.NewList(types.GetSliderType()),
				Resolve:     resolvers.GetAllSlider,
				Description: "Get All Slider",
			},

			"userById": {
				Type: graphql.NewList(types.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:     resolvers.GetUserById,
				Description: "Get Single Users by Id",
			},

			"userByEmailOrPhone": {
				Type: graphql.NewList(types.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetUserByEmailOrPhone,
				Description: "Get Single Users by Email or Phone",
			},

			"users": {
				Type:        graphql.NewList(types.GetUserType()),
				Resolve:     resolvers.GetAllUser,
				Description: "Get All Users",
			},

			"airlines": {
					Type: graphql.NewList(types.GetAirlineType()),
					Resolve: resolvers.GetAllAirline,
			},

			"airlineById": {
				Type: graphql.NewList(types.GetAirlineType()),
				Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.GetAirlineById,
				Description: "Get airline by id",
			},

			"airports": {
				Type:        graphql.NewList(types.GetAirportType()),
				Resolve:     resolvers.GetAllAirport,
				Description: "Get All Airport",
			},

			"distinctAirports": {
				Type:        graphql.NewList(types.GetAirportType()),
				Resolve:     resolvers.GetDistinctAirport,
				Description: "Get Distinct Airport",
			},

			"flights": {
				Type:	graphql.NewList(types.GetFlightType()),
				Resolve: resolvers.GetAllFlight,
			},

			"flightByFromAndTo": {
				Type:	graphql.NewList(types.GetFlightType()),
				Args:	graphql.FieldConfigArgument{
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"to": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.GetFlightByToAndFrom,
			},

			"flightBySchedule": {
				Type:	graphql.NewList(types.GetFlightType()),
				Args:	graphql.FieldConfigArgument{
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"to": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.GetFlightBySchedule,
			},

			"flightByOneSchedule": {
				Type:	graphql.NewList(types.GetFlightType()),
				Args:	graphql.FieldConfigArgument{
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"to": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"date": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.GetFlightByOneSchedule,
			},

			"flightRoutes": {
				Type:	graphql.NewList(types.GetFlightRoutesType()),
				Resolve: resolvers.GetAllFlightRoutes,
			},

			"flightFacilities": {
				Type:	graphql.NewList(types.GetFlightFacilityType()),
				Resolve: resolvers.GetAllFlightFacility,
			},

			"facilities": {
				Type:	graphql.NewList(types.GetFacilityType()),
				Resolve: resolvers.GetAllFacility,
			},

			"trains": {
				Type: graphql.NewList(types.GetTrainType()),
				Resolve: resolvers.GetAllTrain,
			},

			"stations": {
				Type: graphql.NewList(types.GetStationType()),
				Resolve: resolvers.GetAllStation,
			},

			"railroadRoutes": {
				Type: graphql.NewList(types.GetRailroadRoutesType()),
				Resolve: resolvers.GetAllRailroadRoutes,
			},

			"railroads": {
				Type:	graphql.NewList(types.GetRailroadType()),
				Resolve: resolvers.GetAllRailroad,
			},

			"railroadByFromAndTo": {
				Type:	graphql.NewList(types.GetRailroadType()),
				Args:	graphql.FieldConfigArgument{
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"to": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.GetRailroadByToAndFrom,
			},

			"railroadsByOneSchedule": {
				Type:	graphql.NewList(types.GetRailroadType()),
				Args:	graphql.FieldConfigArgument{
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"to": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"date": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.GetRailroadByOneSchedule,
			},

			"railroadsBySchedule": {
				Type:	graphql.NewList(types.GetRailroadType()),
				Args:	graphql.FieldConfigArgument{
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"to": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.GetRailroadBySchedule,
			},

			"distinctStation": {
				Type:        graphql.NewList(types.GetStationType()),
				Resolve:     resolvers.GetDistinctStation,
				Description: "Get Distinct Airport",
			},

			"rooms": {
				Type:        graphql.NewList(types.GetRoomType()),
				Resolve:     resolvers.GetAllRoom,
			},

			"hotels": {
				Type:        graphql.NewList(types.GetHotelType()),
				Resolve:     resolvers.GetAllHotel,
			},
			"hotel": {
				Type:        graphql.NewList(types.GetHotelType()),
				Args:	graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:     resolvers.GetHotelById,
			},
			"nearestHotels": {
				Type:        graphql.NewList(types.GetHotelType()),
				Args:	graphql.FieldConfigArgument{
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetNearestHotelByLocation,
			},

			"hotelByLocation": {
				Type:        graphql.NewList(types.GetHotelType()),
				Args:	graphql.FieldConfigArgument{
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetHotelByLocation,
			},

			"roomFacilities": {
				Type:        graphql.NewList(types.GetRoomFacilityType()),
				Resolve:     resolvers.GetAllRoomFacility,
			},

			"cars": {
				Type:        graphql.NewList(types.GetCarType()),
				Resolve:     resolvers.GetAllCar,
			},

			"vendors": {
				Type:        graphql.NewList(types.GetVendorType()),
				Resolve:     resolvers.GetAllVendor,
			},

			"carByLocation": {
				Type:        graphql.NewList(types.GetCarType()),
				Args:	graphql.FieldConfigArgument{
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetCarByLocation,
			},

			"location": {
				Type:        graphql.NewList(types.GetLocationType()),
				Resolve:     resolvers.GetAllLocation,
			},

			"blogs": {
				Type:        graphql.NewList(types.GetBlogType()),
				Resolve:     resolvers.GetAllBlog,
			},

			"blogById": {
				Type: graphql.NewList(types.GetBlogType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.GetSingleBlog,
				Description: "Get blog by id",
			},

			"events": {
				Type: graphql.NewList(types.GetEventType()),
				Resolve: resolvers.GetAllEvent,
			},

			"eventById": {
				Type: graphql.NewList(types.GetEventType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.GetEventById,
			},

			"eventsByCategory": {
				Type: graphql.NewList(types.GetEventType()),
				Args: graphql.FieldConfigArgument{
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.GetEventByCategory,
			},
			"nearestEvent": {
				Type: graphql.NewList(types.GetEventType()),
				Args: graphql.FieldConfigArgument{
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.GetNearestEventByLocation,
			},
			"eventByLocation": {
				Type: graphql.NewList(types.GetEventType()),
				Args: graphql.FieldConfigArgument{
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.GetEventByLocation,
			},

			"promos": {
				Type: graphql.NewList(types.GetPromoType()),
				Resolve: resolvers.GetAllPromo,
			},

			"promo": {
				Type: graphql.NewList(types.GetPromoType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.GetPromoById,
			},

			"banks": {
				Type: graphql.NewList(types.GetBankType()),
				Resolve: resolvers.GetAllBank,
			},
			"review": {
				Type: graphql.NewList(types.GetReviewType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.GetAllReview,
			},
			"tripReview": {
				Type: graphql.NewList(types.GetTripReviewType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.GetAllTripReview,
			},
			"notifyEmail":{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"content": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.NotifyEmail,
			},

			"orders": {
				Type: graphql.NewList(types.GetOrderType()),
				Resolve: resolvers.GetAllOrder,
			},
			"order": {
				Type: graphql.NewList(types.GetOrderType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.GetOrderById,
			},
		},
	})
}

