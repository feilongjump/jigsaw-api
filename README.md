# Jigsaw API

基于 Gin 框架的轻量 API 服务，为 [Jigsaw](https://github.com/feilongjump/jigsaw) 项目提供配套接口支持。

### 项目目录

jigsaw-api/
├── api/                # 接口层（处理 HTTP 请求，参数绑定与基础验证）
│   ├── handler/        # 请求处理器（接收请求并调用 Service 层）
│   └── router/         # 路由定义（注册 HTTP 路由）
├── application/        # 应用层（业务逻辑编排与数据传输对象 DTO）
│   └── note/           # Note 模块应用逻辑
│       ├── dto/        # 数据传输对象（Request/Response 结构体）
│       └── service.go  # 业务服务逻辑实现
├── domain/             # 领域层（核心业务模型与接口定义）
│   ├── entity/         # 业务实体（与数据库表映射的结构体）
│   └── repo/           # 存储库接口（定义数据持久化契约）
├── infrastructure/     # 基础设施层（外部依赖的具体实现）
│   ├── config/         # 配置文件初始化（Viper 封装）
│   ├── db/             # 数据库连接与迁移（GORM 封装）
│   └── repo_impl/      # 存储库接口实现（具体的 SQL 操作）
├── pkg/                # 公共工具层（跨模块的共享包）
│   ├── carbon/         # 时间处理库扩展
│   ├── err_code/       # 错误码定义
│   ├── gin_util/       # Gin 框架工具（参数绑定辅助、解析器）
│   ├── logger/         # 日志工具（Zap 封装）
│   ├── response/       # 统一响应结构处理
│   └── validator/      # 自定义参数验证器（支持 I18n 翻译）
├── dockerfiles/        # Docker 容器定义
│   └── mysql/          # MySQL 容器配置
├── .air.toml           # Air 热重载配置
├── .env.example        # Docker 环境变量模板
├── .gitignore          # Git 忽略规则文件
├── .jigsaw.toml        # 项目环境配置文件
├── docker-compose.yml  # Docker 容器编排
├── go.mod              # Go 依赖管理
├── go.sum              # 依赖锁文件
├── main.go             # 项目入口文件
└── README.md           # 项目说明文档

### 错误码规范

| 错误码 | 类型 | 说明 | 示例 |
| --- | --- | --- | --- |
| 0 | 成功 | 所有接口成功统一返回 0 | 0 |
| 1000 - 1999 | 通用错误 | 服务器、数据库、框架错误、参数校验错误 | 1001（系统异常）、1002（数据库连接失败） |
| 2000 - 9999 | 业务级错误 | 按业务模块划分 | 每个模块从0-499，例如用户模块从2000-2499 |

> 预留 xx0 作为错误兜底码，例如 1000、2000、2500 等，具体根据模块划分。
