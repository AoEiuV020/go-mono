# Go Mono - Go 多模块项目（静态链接 vs 动态库对比）

使用 **Go Workspace** 管理的多模块项目，对比静态链接与动态链接的实际差异。

## 项目结构

```
go-mono/
├── go.work                # Go Workspace 配置
├── packages/              # 库模块
│   ├── common/           # 通用库
│   │   ├── logger.go     # 业务逻辑
│   │   ├── utils.go
│   │   └── cshared/      # C导出层（动态库用）
│   ├── mathlib/          # 数学计算库
│   │   ├── calculator.go
│   │   └── cshared/
│   └── stringlib/        # 字符串处理库
│       ├── processor.go
│       └── cshared/
├── apps/
│   ├── static-app/       # 静态链接版本（直接import Go包）
│   └── dynamic-app/      # 动态链接版本（CGO调用.dylib）
├── scripts/
│   └── build.sh          # 构建脚本
└── build/                # 构建输出
    ├── static/
    │   └── static-app
    └── dynamic/
        ├── dynamic-app
        └── lib/
            ├── libcommon.dylib
            ├── libmathlib.dylib
            └── libstringlib.dylib
```

## 快速开始

### 构建

```bash
./scripts/build.sh
```

### 运行

```bash
# 静态链接版本
./build/static/run.sh

# 动态链接版本
./build/dynamic/run.sh
```

两个版本功能完全相同，输出结果一致。

## 构建产物对比

详见 [构建测试总结.md](md/构建测试总结.md)

### 静态链接
- **static-app**: 2.3M
- **总计**: 2.3M

### 动态链接
- **dynamic-app**: 2.4M
- **libcommon.dylib**: 1.7M
- **libmathlib.dylib**: 1.7M
- **libstringlib.dylib**: 1.7M
- **总计**: 7.4M

⚠️ 动态链接版本比静态链接大 **215%**

## 技术要点

### Go Workspace
- 使用 `go.work` 统一管理 5 个模块
- 无需在 `go.mod` 中使用 `replace` 指令
- 模块间依赖自动解析

### 动态库生成
- 使用 `-buildmode=c-shared` 生成 `.dylib`（macOS）或 `.so`（Linux）
- 每个库的 `cshared/` 目录包含 C 导出层
- C 导出层只做类型转换，调用原始 Go 代码（确保静态和动态版本逻辑一致）

### CGO 集成
- `dynamic-app` 通过 CGO 调用动态库的 C 接口
- 使用 ID 映射管理 Go 对象（避免 CGO 指针限制）
- `rpath` 设置为 `@executable_path/lib`（相对路径）

## 验证动态链接

```bash
otool -L build/dynamic/dynamic-app
```

输出应显示对三个 `.dylib` 的依赖。

## 架构设计

### 代码复用
- **业务逻辑**: 在 `packages/*/` 的 Go 包中实现
- **静态链接**: 直接 import Go 包
- **动态链接**: cshared 层包装 Go 包，通过 CGO 调用

关键：两种方式执行**完全相同**的代码。

### 模块依赖
```
common (基础库)
  ↓
mathlib, stringlib (依赖 common)
  ↓
static-app (import Go包)
dynamic-app (CGO调用.dylib)
```

## 环境要求

- Go 1.21+（Workspace 支持）
- CGO 支持
- macOS / Linux

---

**创建日期**: 2025年11月23日
