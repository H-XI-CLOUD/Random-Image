<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>随机图像API - 简洁文档</title>
    <style>
        :root {
            --light: #C3CDD6;
            --dark: #2B4C6F;
            --white: #FFFFFF;
        }
        
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, sans-serif;
            line-height: 1.6;
            color: #333;
            background-color: var(--white);
            padding: 0;
        }
        
        .container {
            max-width: 1000px;
            margin: 0 auto;
            padding: 20px;
        }
        
        header {
            background: var(--dark);
            color: var(--white);
            padding: 40px 20px;
            text-align: center;
            margin-bottom: 30px;
        }
        
        header h1 {
            font-size: 2.2rem;
            margin-bottom: 10px;
            font-weight: 600;
        }
        
        header p {
            font-size: 1.1rem;
            opacity: 0.9;
            max-width: 600px;
            margin: 0 auto;
        }
        
        .badge {
            display: inline-block;
            background: var(--light);
            color: var(--dark);
            padding: 4px 12px;
            border-radius: 4px;
            font-size: 0.8rem;
            margin: 5px;
            font-weight: 500;
        }
        
        section {
            margin-bottom: 40px;
            padding: 20px;
            background: var(--white);
            border: 1px solid var(--light);
            border-radius: 8px;
        }
        
        h2 {
            color: var(--dark);
            margin-bottom: 15px;
            padding-bottom: 10px;
            border-bottom: 1px solid var(--light);
            font-weight: 600;
        }
        
        h3 {
            color: var(--dark);
            margin: 20px 0 10px;
            font-weight: 600;
        }
        
        .card {
            background: var(--light);
            border-radius: 6px;
            padding: 15px;
            margin: 15px 0;
        }
        
        code {
            background: var(--dark);
            color: var(--white);
            padding: 2px 6px;
            border-radius: 3px;
            font-family: 'Monaco', 'Menlo', monospace;
            font-size: 0.9rem;
        }
        
        pre {
            background: var(--dark);
            color: var(--white);
            padding: 15px;
            border-radius: 6px;
            overflow-x: auto;
            margin: 15px 0;
            line-height: 1.5;
            font-size: 0.9rem;
        }
        
        table {
            width: 100%;
            border-collapse: collapse;
            margin: 20px 0;
            font-size: 0.9rem;
        }
        
        th, td {
            padding: 10px 12px;
            text-align: left;
            border-bottom: 1px solid var(--light);
        }
        
        th {
            background: var(--light);
            color: var(--dark);
            font-weight: 600;
        }
        
        .api-endpoint {
            background: var(--light);
            padding: 10px 15px;
            border-radius: 6px;
            margin: 10px 0;
            font-family: monospace;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        
        .method {
            background: var(--dark);
            color: var(--white);
            padding: 4px 8px;
            border-radius: 3px;
            font-weight: 500;
            font-size: 0.8rem;
        }
        
        .steps {
            margin: 20px 0;
        }
        
        .step {
            display: flex;
            margin-bottom: 15px;
            align-items: flex-start;
        }
        
        .step:before {
            content: counter(step);
            counter-increment: step;
            background: var(--dark);
            color: var(--white);
            width: 24px;
            height: 24px;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            margin-right: 12px;
            flex-shrink: 0;
            font-size: 0.8rem;
            font-weight: 500;
        }
        
        .status {
            display: inline-block;
            padding: 4px 8px;
            border-radius: 3px;
            font-weight: 500;
            margin: 5px;
            font-size: 0.8rem;
        }
        
        .status.info {
            background: var(--light);
            color: var(--dark);
        }
        
        .status.error {
            background: #FFE6E6;
            color: #B00020;
        }
        
        .status.warning {
            background: #FFF3E0;
            color: #E65100;
        }
        
        footer {
            background: var(--dark);
            color: var(--white);
            text-align: center;
            padding: 20px;
            margin-top: 40px;
            font-size: 0.9rem;
        }
        
        ul, ol {
            padding-left: 20px;
            margin: 10px 0;
        }
        
        li {
            margin-bottom: 8px;
        }
        
        @media (max-width: 768px) {
            .container {
                padding: 15px;
            }
            
            header {
                padding: 30px 15px;
            }
            
            header h1 {
                font-size: 1.8rem;
            }
            
            section {
                padding: 15px;
            }
            
            .api-endpoint {
                flex-direction: column;
                align-items: flex-start;
            }
            
            .method {
                margin-top: 8px;
            }
        }
    </style>
</head>
<body>
    <header>
        <div class="container">
            <h1>随机图像API</h1>
            <p>轻量级Go服务，从目录随机返回图像</p>
            <div style="margin-top: 15px;">
                <span class="badge">仅支持Linux</span>
                <span class="badge">默认端口: 15555</span>
                <span class="badge">Go语言</span>
            </div>
        </div>
    </header>
    
    <div class="container">
        <section id="overview">
            <h2>项目简介</h2>
            <p>一个简单的Go语言API服务，从指定目录随机返回图像文件，并记录请求统计。</p>
            
            <div class="card">
                <h3>主要功能</h3>
                <ul>
                    <li>从目录随机返回图像 (JPG, PNG, GIF, BMP, WebP)</li>
                    <li>自动记录请求统计到JSON文件</li>
                    <li>轻量级设计，资源占用低</li>
                    <li>简单的RESTful API</li>
                </ul>
            </div>
        </section>
        
        <section id="requirements">
            <h2>系统要求</h2>
            
            <div class="card">
                <h3>基本要求</h3>
                <ul>
                    <li><strong>操作系统:</strong> Linux (Ubuntu, CentOS, Debian等)</li>
                    <li><strong>Go版本:</strong> 1.16+</li>
                    <li><strong>端口:</strong> 15555 (默认) 或其他可用端口</li>
                </ul>
            </div>
        </section>
        
        <section id="installation">
            <h2>安装与启动</h2>
            
            <h3>环境准备</h3>
            <div class="steps">
                <div class="step">
                    <div>
                        <p><strong>安装Go环境</strong></p>
                        <pre># Ubuntu/Debian
sudo apt install golang-go</pre>
                    </div>
                </div>
                
                <div class="step">
                    <div>
                        <p><strong>创建项目目录</strong></p>
                        <pre>mkdir -p /random-image-api/resources
cd /random-image-api</pre>
                    </div>
                </div>
                
                <div class="step">
                    <div>
                        <p><strong>添加图像文件</strong></p>
                        <pre># 复制图像到resources目录
cp /path/to/images/* /random-image-api/resources/</pre>
                    </div>
                </div>
            </div>
            
            <h3>启动服务</h3>
            <div class="steps">
                <div class="step">
                    <div>
                        <p><strong>保存代码为main.go</strong></p>
                        <pre># 将Go代码保存为main.go文件</pre>
                    </div>
                </div>
                
                <div class="step">
                    <div>
                        <p><strong>初始化模块</strong></p>
                        <pre>go mod init random-image-api</pre>
                    </div>
                </div>
                
                <div class="step">
                    <div>
                        <p><strong>运行服务</strong></p>
                        <pre>go run main.go</pre>
                        <p>或编译后运行:</p>
                        <pre>go build -o random-image-api
./random-image-api</pre>
                    </div>
                </div>
            </div>
        </section>
        
        <section id="usage">
            <h2>API使用</h2>
            
            <h3>API端点</h3>
            
            <div class="api-endpoint">
                <span>http://localhost:15555/</span>
                <span class="method">GET</span>
            </div>
            <p>返回API基本信息。</p>
            
            <div class="api-endpoint">
                <span>http://localhost:15555/random-image</span>
                <span class="method">GET</span>
            </div>
            <p>返回随机图像。</p>
            
            <div class="api-endpoint">
                <span>http://localhost:15555/stats</span>
                <span class="method">GET</span>
            </div>
            <p>返回请求统计数据。</p>
            
            <h3>使用示例</h3>
            
            <h4>获取随机图像</h4>
            <pre>curl http://localhost:15555/random-image --output random.jpg</pre>
            
            <h4>查看统计数据</h4>
            <pre>curl http://localhost:15555/stats</pre>
            
            <h4>查看API信息</h4>
            <pre>curl http://localhost:15555/</pre>
        </section>
        
        <section id="troubleshooting">
            <h2>故障排除</h2>
            
            <h3>正常启动信息</h3>
            <div class="card">
                <pre>🚀 服务器启动中...
📡 监听端口: 15555 (TCP/UDP)
📸 图片目录: /random-image-api/resources
📊 数据文件: /random-image-api/data.json
✅ 目录检查完成</pre>
                <p>这些信息表示服务启动正常。</p>
            </div>
            
            <h3>常见问题</h3>
            
            <div class="card">
                <h4><span class="status warning">警告</span> 图片目录为空</h4>
                <pre>⚠️ 警告: 无法读取图片目录
📁 图片目录包含 0 个文件</pre>
                <p><strong>解决方案:</strong></p>
                <ul>
                    <li>创建目录: <code>mkdir -p /random-image-api/resources</code></li>
                    <li>添加图像文件到目录</li>
                </ul>
            </div>
            
            <div class="card">
                <h4><span class="status error">错误</span> 端口被占用</h4>
                <pre>启动服务器失败: listen tcp :15555: bind: address already in use</pre>
                <p><strong>解决方案:</strong></p>
                <ul>
                    <li>使用其他端口: <code>PORT=15556 go run main.go</code></li>
                    <li>查找占用进程: <code>sudo lsof -i :15555</code></li>
                </ul>
            </div>
            
            <div class="card">
                <h4><span class="status warning">警告</span> 权限不足</h4>
                <pre>读取图片目录失败: permission denied</pre>
                <p><strong>解决方案:</strong></p>
                <ul>
                    <li>更改权限: <code>chmod 755 /random-image-api/resources</code></li>
                </ul>
            </div>
            
            <h3>服务检查</h3>
            <table>
                <tr>
                    <th>命令</th>
                    <th>用途</th>
                </tr>
                <tr>
                    <td><code>ps aux | grep random-image-api</code></td>
                    <td>检查服务是否运行</td>
                </tr>
                <tr>
                    <td><code>netstat -tulpn | grep 15555</code></td>
                    <td>检查端口监听</td>
                </tr>
                <tr>
                    <td><code>curl http://localhost:15555/</code></td>
                    <td>测试API响应</td>
                </tr>
            </table>
        </section>
        
        <section id="structure">
            <h2>项目结构</h2>
            
            <pre>
/random-image-api/
├── main.go
├── go.mod
├── data.json
└── resources/
    ├── image1.jpg
    ├── image2.png
    └── ...
            </pre>
        </section>
        
        <section id="license">
            <h2>许可证</h2>
            
            <div class="card">
                <h3>MIT License</h3>
                <p>允许自由使用、修改和分发。</p>
            </div>
        </section>
    </div>
    
    <footer>
        <div class="container">
            <p>随机图像API &copy; 2025 - 基于Go语言开发的轻量级图像服务</p>
        </div>
    </footer>

    <script>
        // 平滑滚动
        document.querySelectorAll('a[href^="#"]').forEach(anchor => {
            anchor.addEventListener('click', function(e) {
                e.preventDefault();
                const targetId = this.getAttribute('href');
                if (targetId === '#') return;
                
                const targetElement = document.querySelector(targetId);
                if (targetElement) {
                    window.scrollTo({
                        top: targetElement.offsetTop - 20,
                        behavior: 'smooth'
                    });
                }
            });
        });
    </script>
</body>
</html>
