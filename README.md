# Student Service

## 概览

一个使用 Go 语言实现的简单学生管理后端服务。该项目采用了整洁架构（Clean Architecture）和命令查询职责分离（CQRS）模式。

## 技术栈

*   **语言**: Go
*   **Web 框架**: Gin
*   **数据库**: MySQL
*   **缓存**: Redis
*   **ORM**: GORM
*   **配置管理**: Viper

## 架构

*   **整洁架构 (Clean Architecture)**: 项目遵循整洁架构原则，将业务逻辑与外部依赖（如数据库、Web框架）分离。
    *   `internal/domain`: 核心领域模型和业务规则。
    *   `internal/app`: 应用层，包含命令（Commands）和查询（Queries）。
    *   `internal/adapters`: 适配器层，连接外部服务，如数据库仓库和 HTTP 服务器。
*   **CQRS**: 将系统的读（Query）和写（Command）操作分离，以优化性能和可伸缩性。

## API 端点

API 文档由 OpenAPI (`api/openapi.yaml`) 定义。

*   `POST /students`: 创建一个新学生
*   `GET /students`: 获取所有学生列表
*   `GET /students/{id}`: 根据 ID 获取单个学生信息
*   `PUT /students/{id}`: 根据 ID 更新学生信息
*   `DELETE /students/{id}`: 根据 ID 删除学生

## 快速开始

### 环境准备

*   Go (1.22+)
*   MySQL
*   Redis

### 安装与运行

1.  **克隆仓库**
    ```bash
    git clone https://github.com/LeonCheng0129/student_service.git
    cd student_service
    ```

2.  **配置**
    复制或重命名 `internal/common/configs/config.yaml.example` 为 `config.yaml`，并根据您的环境修改数据库和 Redis 的连接信息。

3.  **初始化数据库**
    连接到您的 MySQL 实例并执行 `scripts/init.sql` 脚本来创建所需的表。

4.  **安装依赖**
    ```bash
    go mod tidy
    ```

5.  **运行服务**
    ```bash
    go run cmd/main.go
    ```
    服务默认将在 `3456` 端口启动。