# StockAI 启动与停止命令

## 访问地址

- 前端：<http://127.0.0.1:5173>
- 后端健康检查：<http://127.0.0.1:8080/api/v1/health>
- MySQL 宿主机端口：`3307`
- Redis 宿主机端口：`6380`

## 启动

### 1. 启动 MySQL 与 Redis

```bash
cd /Users/hangj/Projects/StockAI
docker compose up -d mysql redis
```

确认两个容器都已进入 `healthy` 状态：

```bash
docker compose ps
```

### 2. 启动后端

新开一个终端：

```bash
cd /Users/hangj/Projects/StockAI/backend
go run ./cmd/server
```

看到以下日志表示后端已启动：

```text
msg="server listening" addr=:8080
```

### 3. 启动前端

再开一个终端：

```bash
cd /Users/hangj/Projects/StockAI/frontend
npm run dev
```

看到 Vite 输出后，打开：

```text
http://127.0.0.1:5173
```

## 停止

### 停止后端或前端

在对应的后端或前端终端中按：

```text
Ctrl+C
```

### 停止 MySQL 与 Redis

停止并删除容器：

```bash
cd /Users/hangj/Projects/StockAI
docker compose down
```

只停止容器但保留容器：

```bash
docker compose stop
```

## 清理当前已运行的进程

如果端口被之前启动的进程占用，可以执行：

```bash
lsof -ti :5173 | xargs kill
lsof -ti :8080 | xargs kill
cd /Users/hangj/Projects/StockAI
docker compose down
```

## 说明

- `docker compose down` 默认保留 MySQL 和 Redis 的数据卷。
- 不要随意增加 `-v` 参数，否则会删除数据库和 Redis 数据。
- 如果本机其他服务占用了 `3307` 或 `6380`，可以用 `MYSQL_PORT`、`REDIS_PORT` 环境变量覆盖宿主机端口。

