package viewmodels

import ()

type standLocator struct {
	Title  string
	Active string
}

func GetStandLocator() standLocator {
	result := standLocator{
		Title:  "Lemonade Stand Supply - Stand Locator",
		Active: "stand_locator",
	}

	return result
}

type standLocation struct {
	Lat   float32
	Lng   float32
	Title string
}

//52.586244, lng: -1.982797
func GetStandLocations() []standLocation {
	result := []standLocation{
		standLocation{
			Lat:   52.58217,
			Lng:   -1.98275,
			Title: "Matthew's stand",
		},
		standLocation{
			Lat:   52.58206,
			Lng:   -1.988,
			Title: "Alice's stand",
		},
		standLocation{
			Lat:   52.59205,
			Lng:   -1.9883,
			Title: "Kara's stand",
		},
		standLocation{
			Lat:   52.5814,
			Lng:   -1.979,
			Title: "Fred's stand",
		},
		standLocation{
			Lat:   52.5812,
			Lng:   -1.989,
			Title: "Jake's stand",
		},
		standLocation{
			Lat:   52.5741,
			Lng:   -1.9893,
			Title: "Wallace's stand",
		},
		standLocation{
			Lat:   52.5816,
			Lng:   -1.9895,
			Title: "Gromit's stand",
		},
		standLocation{
			Lat:   52.581,
			Lng:   -1.97,
			Title: "Kirk's stand",
		},
		standLocation{
			Lat:   52.581,
			Lng:   -1.9802,
			Title: "Lorelei's stand",
		},
		standLocation{
			Lat:   52.5912,
			Lng:   -1.9799,
			Title: "Rebecca's stand",
		},
		standLocation{
			Lat:   52.588,
			Lng:   -1.99025,
			Title: "Chris's stand",
		},
		standLocation{
			Lat:   52.5823,
			Lng:   -1.98025,
			Title: "Carson's stand",
		},
	}

	return result
}
