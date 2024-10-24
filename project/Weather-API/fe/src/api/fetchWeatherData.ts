async function fetchWeatherData(local: string, startdate: string, enddate: string) {
    try {
        const response = await fetch(`http://localhost:8080/api/weather?local=${local}&startdate=${startdate}&enddate=${enddate}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching weather data:', error);
        throw error;
    }
}

export default fetchWeatherData;