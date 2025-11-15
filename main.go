package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 统计数据结构
type Stats struct {
	TotalRequests int64            `json:"total_requests"`
	ImageCounts   map[string]int64 `json:"image_counts"`
	LastUpdated   time.Time        `json:"last_updated"`
}

// 全局变量
var (
	stats     Stats
	statsMutex sync.RWMutex
	dataFile  string
	imageDir  string
)

func init() {
	// 初始化路径 - 使用相对路径或从环境变量获取
	basePath := getBasePath()
	dataFile = filepath.Join(basePath, "data.json")
	imageDir = filepath.Join(basePath, "resources")
	
	// 初始化统计
	stats = Stats{
		ImageCounts: make(map[string]int64),
		LastUpdated: time.Now(),
	}
	
	// 确保目录存在
	if err := ensureDirectories(); err != nil {
		log.Fatalf("创建目录失败: %v", err)
	}
	
	// 加载已有的统计数据
	loadStats()
	
	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())
}

// 获取基础路径
func getBasePath() string {
	// 尝试从环境变量获取
	if envPath := os.Getenv("RANDOM_IMAGE_API_PATH"); envPath != "" {
		return envPath
	}
	
	// 尝试获取当前工作目录
	if wd, err := os.Getwd(); err == nil {
		// 检查当前目录是否有 resources 文件夹
		if _, err := os.Stat(filepath.Join(wd, "resources")); err == nil {
			return wd
		}
		// 检查上一级目录
		parent := filepath.Dir(wd)
		if _, err := os.Stat(filepath.Join(parent, "resources")); err == nil {
			return parent
		}
	}
	
	// 默认路径
	return "/random-image-api"
}

// 确保必要的目录存在
func ensureDirectories() error {
	// 创建资源目录（如果不存在）
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		return fmt.Errorf("创建资源目录失败: %v", err)
	}
	
	// 创建数据文件目录（如果不存在）
	dataDir := filepath.Dir(dataFile)
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("创建数据目录失败: %v", err)
	}
	
	log.Printf("✅ 目录检查完成")
	log.Printf("📸 图片目录: %s", imageDir)
	log.Printf("📊 数据文件: %s", dataFile)
	
	return nil
}

func loadStats() {
	statsMutex.Lock()
	defer statsMutex.Unlock()
	
	// 检查文件是否存在
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		log.Printf("统计文件不存在，将创建新文件: %s", dataFile)
		// 创建空的统计文件
		saveStats()
		return
	}
	
	// 读取文件
	data, err := os.ReadFile(dataFile)
	if err != nil {
		log.Printf("读取统计文件失败: %v", err)
		return
	}
	
	// 解析JSON
	if err := json.Unmarshal(data, &stats); err != nil {
		log.Printf("解析统计文件失败: %v", err)
		return
	}
	
	log.Printf("已加载统计: 总请求数=%d, 图片数量=%d", stats.TotalRequests, len(stats.ImageCounts))
}

func saveStats() {
	statsMutex.Lock()
	defer statsMutex.Unlock()
	
	stats.LastUpdated = time.Now()
	
	// 转换为JSON
	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		log.Printf("序列化统计数据失败: %v", err)
		return
	}
	
	// 写入文件
	if err := os.WriteFile(dataFile, data, 0644); err != nil {
		log.Printf("写入统计文件失败: %v", err)
		return
	}
}

func getRandomImage() (string, string, error) {
	// 支持的图片格式
	allowedExt := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, 
		".gif": true, ".bmp": true, ".webp": true,
	}
	
	// 读取目录
	entries, err := os.ReadDir(imageDir)
	if err != nil {
		return "", "", fmt.Errorf("读取图片目录失败: %v", err)
	}
	
	// 筛选图片文件
	var images []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := filepath.Ext(entry.Name())
		if allowedExt[ext] {
			images = append(images, entry.Name())
		}
	}
	
	if len(images) == 0 {
		// 创建示例图片或返回错误
		return "", "", fmt.Errorf("图片目录为空，请在 %s 目录中添加图片文件", imageDir)
	}
	
	// 随机选择
	randomImage := images[rand.Intn(len(images))]
	return randomImage, filepath.Join(imageDir, randomImage), nil
}

func updateStats(imageName string) {
	statsMutex.Lock()
	defer statsMutex.Unlock()
	
	// 更新统计
	stats.TotalRequests++
	stats.ImageCounts[imageName]++
	
	log.Printf("统计更新: 图片=%s, 总请求=%d, 该图片请求=%d", 
		imageName, stats.TotalRequests, stats.ImageCounts[imageName])
	
	// 异步保存到文件（避免阻塞请求处理）
	go saveStats()
}

func randomImageHandler(w http.ResponseWriter, r *http.Request) {
	// 获取随机图片
	imageName, imagePath, err := getRandomImage()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
			"image_dir": imageDir,
		})
		return
	}
	
	// 更新统计
	updateStats(imageName)
	
	// 设置正确的Content-Type
	ext := filepath.Ext(imageName)
	switch ext {
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	case ".bmp":
		w.Header().Set("Content-Type", "image/bmp")
	case ".webp":
		w.Header().Set("Content-Type", "image/webp")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	
	// 返回图片
	http.ServeFile(w, r, imagePath)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	statsMutex.RLock()
	defer statsMutex.RUnlock()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	// 获取当前图片列表
	images, _, _ := getRandomImage()
	imageCount := 0
	if images != "" {
		entries, _ := os.ReadDir(imageDir)
		for _, entry := range entries {
			if !entry.IsDir() {
				imageCount++
			}
		}
	}
	
	response := map[string]interface{}{
		"message": "随机图片API",
		"endpoints": map[string]string{
			"random_image": "/random-image",
			"stats":        "/stats",
		},
		"image_dir":   imageDir,
		"data_file":   dataFile,
		"image_count": imageCount,
		"status":      "ready",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// 设置路由
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/random-image", randomImageHandler)
	http.HandleFunc("/stats", statsHandler)
	
	// 启动服务器
	port := "15555"
	addr := ":" + port
	
	log.Printf("🚀 服务器启动中...")
	log.Printf("📡 监听端口: %s (TCP/UDP)", port)
	log.Printf("📸 图片目录: %s", imageDir)
	log.Printf("📊 数据文件: %s", dataFile)
	log.Printf("🌐 访问地址: http://localhost:%s/random-image", port)
	
	// 检查图片目录内容
	entries, err := os.ReadDir(imageDir)
	if err != nil {
		log.Printf("⚠️ 警告: 无法读取图片目录: %v", err)
	} else {
		imageCount := 0
		for _, entry := range entries {
			if !entry.IsDir() {
				imageCount++
			}
		}
		log.Printf("📁 图片目录包含 %d 个文件", imageCount)
	}
	
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}