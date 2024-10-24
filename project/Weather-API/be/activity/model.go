package activity

type Hour struct {
	Datetime      string   `json:"datetime"`
	DatetimeEpoch int64    `json:"datetimeEpoch"`
	Temp          float64  `json:"temp"`
	Feelslike     float64  `json:"feelslike"`
	Humidity      float64  `json:"humidity"`
	Dew           float64  `json:"dew"`
	Precip        *float64 `json:"precip"`
	//Precipprob     int      `json:"precipprob"`
	//Snow           int      `json:"snow"`
	//Snowdepth      int      `json:"snowdepth"`
	//Preciptype     *string  `json:"preciptype"`
	//Windgust       *float64 `json:"windgust"`
	//Windspeed      float64  `json:"windspeed"`
	//Winddir        float64  `json:"winddir"`
	//Pressure       float64  `json:"pressure"`
	//Visibility     float64  `json:"visibility"`
	//Cloudcover     float64  `json:"cloudcover"`
	//Solarradiation float64  `json:"solarradiation"`
	//Solarenergy    float64  `json:"solarenergy"`
	//Uvindex        int      `json:"uvindex"`
	//Severerisk     int      `json:"severerisk"`
	Conditions string   `json:"conditions"`
	Icon       string   `json:"icon"`
	Stations   []string `json:"stations"`
	Source     string   `json:"source"`
}

type Day struct {
	Datetime      string  `json:"datetime"`
	DatetimeEpoch int64   `json:"datetimeEpoch"`
	Tempmax       float64 `json:"tempmax"`
	Tempmin       float64 `json:"tempmin"`
	Temp          float64 `json:"temp"`
	Feelslikemax  float64 `json:"feelslikemax"`
	Feelslikemin  float64 `json:"feelslikemin"`
	Feelslike     float64 `json:"feelslike"`
	Dew           float64 `json:"dew"`
	Humidity      float64 `json:"humidity"`
	Precip        float64 `json:"precip"`
	//Precipprob     int      `json:"precipprob"`
	//Precipcover    int      `json:"precipcover"`
	//Preciptype     *string  `json:"preciptype"`
	//Snow           int      `json:"snow"`
	//Snowdepth      int      `json:"snowdepth"`
	//Windgust       float64  `json:"windgust"`
	//Windspeed      float64  `json:"windspeed"`
	//Winddir        float64  `json:"winddir"`
	//Pressure       float64  `json:"pressure"`
	//Cloudcover     float64  `json:"cloudcover"`
	//Visibility     float64  `json:"visibility"`
	//Solarradiation float64  `json:"solarradiation"`
	//Solarenergy    float64  `json:"solarenergy"`
	//Uvindex        int      `json:"uvindex"`
	//Severerisk     int      `json:"severerisk"`
	Sunrise      string   `json:"sunrise"`
	SunriseEpoch int64    `json:"sunriseEpoch"`
	Sunset       string   `json:"sunset"`
	SunsetEpoch  int64    `json:"sunsetEpoch"`
	Moonphase    float64  `json:"moonphase"`
	Conditions   string   `json:"conditions"`
	Description  string   `json:"description"`
	Icon         string   `json:"icon"`
	Stations     []string `json:"stations"`
	Source       string   `json:"source"`
	Hours        []Hour   `json:"hours"`
}

type Station struct {
	Distance  float64 `json:"distance"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UseCount  int     `json:"useCount"`
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Quality   int     `json:"quality"`
	//Contribution int     `json:"contribution"`
}

type CurrentConditions struct {
	Datetime      string   `json:"datetime"`
	DatetimeEpoch int64    `json:"datetimeEpoch"`
	Temp          float64  `json:"temp"`
	Feelslike     float64  `json:"feelslike"`
	Humidity      float64  `json:"humidity"`
	Dew           float64  `json:"dew"`
	Precip        *float64 `json:"precip"`
	//Precipprob     int      `json:"precipprob"`
	//Snow           int      `json:"snow"`
	//Snowdepth      int      `json:"snowdepth"`
	//Preciptype     *string  `json:"preciptype"`
	//Windgust       *float64 `json:"windgust"`
	//Windspeed      float64  `json:"windspeed"`
	//Winddir        float64  `json:"winddir"`
	//Pressure       float64  `json:"pressure"`
	//Visibility     float64  `json:"visibility"`
	//Cloudcover     float64  `json:"cloudcover"`
	//Solarradiation float64  `json:"solarradiation"`
	//Solarenergy    float64  `json:"solarenergy"`
	//Uvindex        int      `json:"uvindex"`
	//Severerisk     int      `json:"severerisk"`
	Conditions   string   `json:"conditions"`
	Icon         string   `json:"icon"`
	Stations     []string `json:"stations"`
	Source       string   `json:"source"`
	Sunrise      string   `json:"sunrise"`
	SunriseEpoch int64    `json:"sunriseEpoch"`
	Sunset       string   `json:"sunset"`
	SunsetEpoch  int64    `json:"sunsetEpoch"`
	Moonphase    float64  `json:"moonphase"`
}

type WeatherData struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Timezone        string  `json:"timezone"`
	//Tzoffset          string             `json:"tzoffset"`
	Description       string             `json:"description"`
	Days              []Day              `json:"days"`
	Alerts            []interface{}      `json:"alerts"`
	Stations          map[string]Station `json:"stations"`
	CurrentConditions CurrentConditions  `json:"currentConditions"`
}
