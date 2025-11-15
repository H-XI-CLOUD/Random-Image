# 随机图像API

轻量级Go服务，从目录返回随机图像

仅支持Linux

![默认端口: 15555](https://img.shields.io/badge/Port-15555-blue)
![Go语言](https://img.shields.io/badge/Language-Go-green)

## 项目简介

随机图像API是一个使用Go语言编写的轻量级Web服务，它从指定目录中随机选择图像并通过HTTP API返回。

### 主要功能

- 从目录随机返回图像
- 支持JPG、PNG、GIF、BMP、WebP格式
- 自动统计请求次数
- 轻量级设计，资源占用低

## 系统要求

> **注意：** 此程序仅支持Linux操作系统。

- **操作系统:** Linux
- **Go版本:** 1.16或更高
- **默认端口:** 15555 (TCP/UDP)
- **内存:** 至少64MB RAM

## 安装与运行

### 1. 环境准备

```bash
# 安装Go (Ubuntu/Debian)
sudo apt update
sudo apt install golang-go
```

### 2. 创建目录结构

```bash
mkdir -p /random-image-api/resources
cd /random-image-api
```

### 3. 添加图像文件

```bash
# 将图像文件放入resources目录
cp /path/to/images/* /random-image-api/resources/
```

### 4. 运行服务

```bash
# 使用默认端口
go run main.go

# 使用自定义端口
PORT=8080 go run main.go
```

## API使用

### 端点

| 端点 | 方法 | 描述 |
|------|------|------|
| `/` | GET | 返回API基本信息 |
| `/random-image` | GET | 返回随机图像 |
| `/stats` | GET | 返回请求统计数据 |

### 使用示例

```bash
# 获取随机图像
curl http://localhost:15555/random-image --output image.jpg

# 查看统计数据
curl http://localhost:15555/stats

# 查看API信息
curl http://localhost:15555/
```

## 故障排除

### 正常启动信息

```
🚀 服务器启动中...
📡 监听端口: 15555 (TCP/UDP)
📸 图片目录: /random-image-api/resources
📊 数据文件: /random-image-api/data.json
🌐 访问地址: http://localhost:15555/random-image
```

### 常见问题

#### 图片目录不存在

**错误信息:**
```
读取图片目录失败: open /random-image-api/resources: no such file or directory
```

**解决方案:**
```bash
mkdir -p /random-image-api/resources
```

#### 端口被占用

**错误信息:**
```
启动服务器失败: listen tcp :15555: bind: address already in use
```

**解决方案:**
```bash
# 使用其他端口
PORT=15556 go run main.go

# 或查找占用进程
sudo lsof -i :15555
```

#### 权限不足

**错误信息:**
```
读取图片目录失败: permission denied
```

**解决方案:**
```bash
chmod 755 /random-image-api/resources
```

## 项目结构

```
/random-image-api/
├── main.go          # 主程序
├── go.mod          # Go模块文件
├── data.json       # 统计数据(自动生成)
└── resources/      # 图像目录
    ├── image1.jpg
    ├── image2.png
    └── ...
```

## 许可证

本项目使用MIT开源许可证。

---
千年云/盒情盒里的蛋饺/BILIEDP1145 &copy; 2025
