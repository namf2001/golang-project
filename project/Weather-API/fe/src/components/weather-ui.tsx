import { useEffect, useState } from 'react'
import { Sun, Moon, Cloud, CloudRain, Wind, Droplets, Search, Sunrise, Sunset, Thermometer, Calendar } from 'lucide-react'
import { Calendar as CalendarIcon } from "lucide-react"
import { DateRange } from "react-day-picker"
import { format } from "date-fns"
import { Switch } from "./ui/switch"
import { Popover, PopoverContent, PopoverTrigger } from "./ui/popover"
import { Button } from "./ui/button"
import { cn } from "../lib/utils"
import { Input } from "./ui/input"
import { Calendar as CalendarComponent } from "./ui/calendar"
import fetchWeatherData from '../api/fetchWeatherData'

export default function WeatherUI() {
  const [isCelsius, setIsCelsius] = useState(true)
  const [apiResponse, setApiResponse] = useState<any>(null)
  const [searchQuery, setSearchQuery] = useState('london')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [date, setDate] = useState<DateRange | undefined>({
    from: new Date(),
    to: new Date(new Date().setDate(new Date().getDate() + 6))
  })

  const fetchData = async () => {
    setLoading(true);
    setError(null);

    try {
      const startdate = date?.from ? format(date.from, 'yyyy-MM-dd') : '';
      const enddate = date?.to ? format(date.to, 'yyyy-MM-dd') : '';
      const data = await fetchWeatherData(searchQuery, startdate, enddate);
      setApiResponse(data);
    } catch (error) {
      if (error instanceof Error) {
        setError(error.message);
      } else {
        setError(String(error));
      }
      console.error('Error fetching data:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  const handleSearch = () => {
    fetchData();
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (!apiResponse) {
    return null;
  }

  const weatherData = {
    location: apiResponse.resolvedAddress,
    currentTemp: apiResponse.currentConditions.temp,
    description: apiResponse.currentConditions.conditions,
    humidity: apiResponse.currentConditions.humidity,
    windSpeed: apiResponse.currentConditions.windspeed,
    windDirection: apiResponse.currentConditions.winddir,
    uvIndex: apiResponse.currentConditions.uvindex,
    pressure: apiResponse.currentConditions.pressure,
    rainChance: apiResponse.days[0].precipprob,
    hourlyForecast: apiResponse.days[0].hours.slice(0, 5).map((hour: { datetime: string; temp: number; icon: string }) => ({
      time: hour.datetime.slice(0, 5),
      temp: hour.temp,
      icon: getWeatherIcon(hour.icon),
    })),
    weeklyForecast: apiResponse.days.map((day: { datetime: string; tempmax: number; tempmin: number; icon: string }) => ({
      day: new Date(day.datetime).toLocaleDateString('en-US', { weekday: 'long' }),
      high: day.tempmax,
      low: day.tempmin,
      icon: getWeatherIcon(day.icon),
    })),
    sunrise: apiResponse.currentConditions.sunrise,
    sunset: apiResponse.currentConditions.sunset,
    lastUpdated: new Date(apiResponse.currentConditions.datetimeEpoch * 1000).toLocaleTimeString(),
  }

  function getWeatherIcon(iconName: string) {
    switch (iconName) {
      case 'clear-day':
      case 'clear-night':
        return <Sun className="w-6 h-6" />
      case 'partly-cloudy-day':
      case 'partly-cloudy-night':
      case 'cloudy':
        return <Cloud className="w-6 h-6" />
      case 'rain':
        return <CloudRain className="w-6 h-6" />
      default:
        return <Sun className="w-6 h-6" />
    }
  }

  const convertTemp = (temp: number) => {
    if (isCelsius) return temp
    return Math.round((temp * 9) / 5 + 32)
  }

  const getBackgroundClass = () => {
    switch (weatherData.description.toLowerCase()) {
      case 'clear':
        return 'bg-gradient-to-br from-yellow-300 to-orange-400'
      case 'partly cloudy':
      case 'overcast':
        return 'bg-gradient-to-br from-blue-200 to-gray-300'
      case 'rain':
        return 'bg-gradient-to-br from-gray-400 to-blue-600'
      default:
        return 'bg-gradient-to-br from-blue-400 to-purple-500'
    }
  }

  return (
    <div className={`min-h-screen ${getBackgroundClass()} text-white p-8`}>
      <div className="max-w-4xl mx-auto">
        <header className="flex flex-col md:flex-row justify-between items-center mb-8 gap-4">
          <h1 className="text-4xl font-bold">Weather Forecast</h1>
          <div className="flex items-center space-x-4">
            <div className="flex items-center space-x-2">
              <span>°C</span>
              <Switch checked={!isCelsius} onCheckedChange={() => setIsCelsius(!isCelsius)} />
              <span>°F</span>
            </div>
            <Popover>
              <PopoverTrigger asChild>
                <Button
                  variant={"outline"}
                  className={cn(
                    "w-[280px] justify-start text-left font-normal text-black",
                    !date && "text-muted-foreground"
                  )}
                  onCanPlayCapture={handleSearch}
                >
                  <CalendarIcon className="mr-2 h-4 w-4" />
                  {date?.from ? (
                    date.to ? (
                      <>
                        {format(date.from, "LLL dd, y")} -{" "}
                        {format(date.to, "LLL dd, y")}
                      </>
                    ) : (
                      format(date.from, "LLL dd, y")
                    )
                  ) : (
                    <span>Pick a date range</span>
                  )}
                </Button>
              </PopoverTrigger>
              <PopoverContent className="w-auto p-0" align="start">
                <CalendarComponent
                  initialFocus
                  mode="range"
                  defaultMonth={date?.from}
                  selected={date}
                  onSelect={setDate}
                  numberOfMonths={2}
                />
              </PopoverContent>
            </Popover>
          </div>
        </header>

        <div className="flex items-center space-x-4 mb-8">
          <Input
            type="text"
            placeholder="Search location..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="bg-white/20 border-none text-white placeholder-white/70"
          />
          <Button variant="outline" size="icon" onClick={handleSearch} className='text-black'>
            <Search className="h-4 w-4" />
          </Button>
        </div>

        <main>
          <div className="bg-white/10 rounded-lg p-6 mb-8">
            <h2 className="text-2xl font-semibold mb-2">Weather in {weatherData.location}</h2>
            <div className="flex items-center justify-between">
              <div>
                <p className="text-6xl font-bold">{convertTemp(weatherData.currentTemp)}°{isCelsius ? 'C' : 'F'}</p>
                <p className="text-xl">{weatherData.description}</p>
              </div>
              <div className="text-8xl">
                {getWeatherIcon(apiResponse.currentConditions.icon)}
              </div>
            </div>
          </div>

          <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8">
            <div className="bg-white/10 rounded-lg p-4 flex items-center">
              <Droplets className="w-6 h-6 mr-2" />
              <div>
                <p className="text-sm">Humidity</p>
                <p className="font-semibold">{weatherData.humidity}%</p>
              </div>
            </div>
            <div className="bg-white/10 rounded-lg p-4 flex items-center">
              <Wind className="w-6 h-6 mr-2" />
              <div>
                <p className="text-sm">Wind</p>
                <p className="font-semibold">{weatherData.windSpeed} km/h {weatherData.windDirection}°</p>
              </div>
            </div>
            <div className="bg-white/10 rounded-lg p-4 flex items-center">
              <Sun className="w-6 h-6 mr-2" />
              <div>
                <p className="text-sm">UV Index</p>
                <p className="font-semibold">{weatherData.uvIndex}</p>
              </div>
            </div>
            <div className="bg-white/10 rounded-lg p-4 flex items-center">
              <CloudRain className="w-6 h-6 mr-2" />
              <div>
                <p className="text-sm">Chance of Rain</p>
                <p className="font-semibold">{weatherData.rainChance}%</p>
              </div>
            </div>
          </div>

          <div className="bg-white/10 rounded-lg p-6 mb-8">
            <h3 className="text-xl font-semibold mb-4">Hourly Forecast</h3>
            <div className="flex justify-between">
              {weatherData.hourlyForecast.map((hour, index) => (
                <div key={index} className="text-center">
                  <p className="text-sm">{hour.time}</p>
                  {hour.icon}
                  <p className="font-semibold">{convertTemp(hour.temp)}°</p>
                </div>
              ))}
            </div>
          </div>

          <div className="bg-white/10 rounded-lg p-6 mb-8">
            <h3 className="text-xl font-semibold mb-4">7-Day Forecast</h3>
            {weatherData.weeklyForecast.map((day, index) => (
              <div key={index} className="flex justify-between items-center mb-2">
                <p className="w-24">{day.day}</p>
                {day.icon}
                <p className="w-24 text-right">
                  {convertTemp(day.high)}° / {convertTemp(day.low)}°
                </p>
              </div>
            ))}
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-8">
            <div className="bg-white/10 rounded-lg p-4 flex items-center">
              <Sunrise className="w-6 h-6 mr-2" />
              <div>
                <p className="text-sm">Sunrise</p>
                <p className="font-semibold">{weatherData.sunrise}</p>
              </div>
            </div>
            <div className="bg-white/10 rounded-lg p-4 flex items-center">
              <Sunset className="w-6 h-6 mr-2" />
              <div>
                <p className="text-sm">Sunset</p>
                <p className="font-semibold">{weatherData.sunset}</p>
              </div>
            </div>
          </div>

          <p className="text-sm text-center">Last updated: {weatherData.lastUpdated}</p>
        </main>
      </div>
    </div>
  )
}