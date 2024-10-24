package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"golang-project/project/Weather-API/be/activity"

	"github.com/go-redis/redis/v8"
	"github.com/rs/cors"
)

// Tạo một context toàn cục để làm việc với Redis
var ctx = context.Background()

// Hàm kết nối đến Redis
func connectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Địa chỉ Redis
		Password: "",               // Mật khẩu Redis, nếu có
		DB:       0,                // Sử dụng DB 0
	})
	return rdb
}

// Hàm xử lý logic lấy dữ liệu thời tiết từ cache
func getWeatherWithCache(w http.ResponseWriter, r *http.Request) {
	rdb := connectRedis()

	// Lấy các tham số từ query
	local := r.URL.Query().Get("local")
	startDate := r.URL.Query().Get("startdate")
	endDate := r.URL.Query().Get("enddate")

	// Tạo cache key bao gồm các tham số
	cacheKey := "weather:" + local + ":" + startDate + ":" + endDate

	// Kiểm tra xem dữ liệu thời tiết có trong Redis cache không
	cachedWeather, err := rdb.Get(ctx, cacheKey).Result()
	if errors.Is(err, redis.Nil) {
		// Cache không có dữ liệu, lấy dữ liệu từ API
		weather := activity.GetDataWeather(local, startDate, endDate)

		// Convert dữ liệu weather thành JSON string để lưu trong Redis
		weatherJson, _ := json.Marshal(weather)

		// Lưu dữ liệu vào Redis với TTL (Time To Live) là 60 giây
		rdb.Set(ctx, cacheKey, weatherJson, 60*time.Second)

		// Trả về phản hồi cho client
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weather)
	} else if err != nil {
		// Xử lý lỗi khi kết nối với Redis
		log.Println("Redis error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		// Cache có dữ liệu, trả về dữ liệu từ Redis
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedWeather)) // Gửi dữ liệu cache về client
	}
}

func main() {
	// Tạo HTTP server và đăng ký endpoint "/api/weather" sử dụng hàm getWeatherWithCache
	mux := http.NewServeMux()
	mux.HandleFunc("/api/weather", getWeatherWithCache)

	// Cấu hình CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Bọc HTTP handler với CORS handler
	handler := c.Handler(mux)

	// Server chạy tại cổng 8080
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
