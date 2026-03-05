# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 常用命令

### 后端 (Go)
- **运行**: `go run server/cmd/api/main.go`
- **构建**: `go build -o server/bin/api server/cmd/api/main.go`
- **测试**: `go test ./server/...`
- **整理依赖**: `go mod tidy` (在 server 目录下运行)

### 前端 (Vue 3)
- **开发**: `npm run dev` (在 client 目录下运行)
- **构建**: `npm run build` (在 client 目录下运行)
- **Lint**: `npm run lint` (如果配置了)

### 部署 (Docker)
- **一键启动**: `docker-compose up -d`
- **重新构建**: `docker-compose up --build -d`
- **停止**: `docker-compose down`
- **查看日志**: `docker-compose logs -f`

## 项目架构

### 高层设计
本项目是一个支持多种棋牌游戏的 Web 平台，采用解耦的游戏引擎架构。
- **通信层**: 使用 WebSocket (Gorilla) 实现服务端与客户端的实时双向通信。
- **引擎层 (`server/internal/engine`)**: 管理房间、桌子和玩家状态。它是通用的，不包含具体游戏规则。
- **游戏规则层 (`server/internal/games`)**: 包含不同游戏（如长沙麻将）的具体逻辑，通过接口或组合接入引擎。
- **前端层 (`client`)**: 使用 Vue 3 构建的单页应用，针对移动端（微信浏览器）进行了 UI 优化。

### 核心结构
- `server/cmd/api`: 程序入口，负责路由和 WebSocket 升级。
- `server/internal/engine/table.go`: 定义 `Table` (房间) 和 `Player` 模型，处理广播逻辑。
- `server/internal/games/mahjong`: 长沙麻将的具体实现，包括洗牌、判胡算法。
- `server/internal/socket/hub.go`: WebSocket 连接池管理，处理消息的分发与心跳。
- `client/src/games`: 存放不同游戏的 Vue 视图组件。
- `client/src/socket`: 处理前端 WebSocket 的重连、心跳和消息分发。

## 开发规范
- **错误处理**: 后端应返回清晰的错误日志；前端应有断线重连提示。
- **状态管理**: 游戏局势状态应以服务器为准，客户端通过 WebSocket 增量更新。
- **移动端适配**: 所有 UI 必须优先考虑手机竖屏和微信浏览器兼容性。
- **命名规范**: 代码注释及变量命名使用英文或清晰的拼音，用户提示及 CLAUDE.md 交流使用简体中文。
