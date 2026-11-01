# StockAI

StockAI 是一个股票分析与智能决策辅助产品。当前处于 V0.1 工程初始化阶段，仅包含后端、前端、基础依赖和健康检查联通。

## 技术栈

- 后端：Go、Gin、GORM
- 前端：React、TypeScript、Vite
- 数据库：MySQL 8.0
- 缓存：Redis
- 依赖编排：Docker Compose

## 目录结构

```text
StockAI/
├── backend/                 # Go 后端
│   ├── cmd/server/          # 服务入口
│   ├── internal/            # 内部实现
│   └── migrations/          # 数据库迁移目录
├── frontend/                # React 前端
├── docs/                    # 设计与进度文档
└── docker-compose.yml       # MySQL 与 Redis 编排
```

## 本地启动

要求本机已安装 Go 1.26+、Node.js 20+、npm 和 Docker。

### 1. 启动 MySQL 与 Redis

```bash
docker compose up -d
```

等待 `stockai-mysql` 与 `stockai-redis` 进入 healthy 状态：

```bash
docker compose ps
```

Compose 默认把 MySQL 映射到宿主机 `3307`、Redis 映射到 `6380`，避免与常见本地实例冲突。如需使用其他端口，可设置 `MYSQL_PORT` 与 `REDIS_PORT` 后启动，并同步配置后端环境变量。

### 2. 启动后端

```bash
cd backend
go mod download
go run ./cmd/server
```

后端默认监听 `http://localhost:8080`。健康检查地址：

```bash
curl http://localhost:8080/api/v1/health
```

### 3. 启动前端

```bash
cd frontend
npm install
npm run dev
```

浏览器打开 `http://localhost:5173`。Vite 会把 `/api` 请求代理到后端。

## 配置

后端通过环境变量覆盖默认配置，本地默认值与 Compose 一致。

| 环境变量 | 默认值 | 说明 |
| --- | --- | --- |
| `SERVER_ADDR` | `:8080` | 后端监听地址 |
| `GIN_MODE` | `release` | Gin 运行模式 |
| `MYSQL_HOST` | `127.0.0.1` | MySQL 地址 |
| `MYSQL_PORT` | `3307` | MySQL 端口 |
| `MYSQL_USER` | `stockai` | MySQL 用户 |
| `MYSQL_PASSWORD` | `stockai_pass` | MySQL 密码 |
| `MYSQL_DATABASE` | `stockai` | MySQL 数据库 |
| `MYSQL_MIGRATIONS_DIR` | `./migrations` | 迁移目录 |
| `REDIS_ADDR` | `127.0.0.1:6380` | Redis 地址 |
| `REDIS_PASSWORD` | `stockai_redis` | Redis 密码 |
| `REDIS_DB` | `0` | Redis 数据库编号 |

## 构建与测试

```bash
cd backend
go build ./...
go test ./...
```

```bash
cd frontend
npm run typecheck
npm run build
```

## 阶段说明

当前仅实现 V0.1 工程初始化，不包含 V0.2 及之后的业务功能。

