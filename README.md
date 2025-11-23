# Go Mono - Golang 多模块项目示例（真正的动态库对比）

这是一个使用 **Go Workspace** 管理的多模块单体仓库（Monorepo）项目，展示了：
- ✅ 静态链接 vs **真正的动态链接**（.dylib）
- ✅ Go Workspace 技术的实际应用  
- ✅ C-shared buildmode 生成动态库
- ✅ CGO 调用动态库的完整实现

## 项目结构

```
go-mono/
├── go.work                # Go Workspace 配置文件
├── packages/              # 库模块目录
│   ├── common/           # 通用库（日志、工具函数）
│   │   ├── logger.go     # Go 包实现
│   │   ├── utils.go
│   │   └── cshared/      # C 导出封装（用于动态库）
│   ├── mathlib/          # 数学计算库
│   │   ├── calculator.go # Go 包实现
│   │   └── cshared/      # C 导出封装（用于动态库）
│   └── stringlib/        # 字符串处理库
│       ├── processor.go  # Go 包实现
│       └── cshared/      # C 导出封装（用于动态库）
├── apps/                 # 应用模块目录
│   ├── static-app/       # 静态链接版本（直接依赖 Go 包）
│   └── dynamic-app/      # 动态链接版本（通过 CGO 调用 .dylib）
├── scripts/              # 构建脚本
│   └── build.sh          # 自动化构建脚本（生成真正的动态库）
├── md/                   # 项目文档
│   ├── 项目结构说明.md   # 详细的项目结构和功能说明
│   └── 构建测试总结.md   # 真正的动态库对比测试结果
└── build/                # 编译输出目录（自动生成）
    ├── static/           # 静态链接版本
    │   └── static-app
    └── dynamic/          # 动态链接版本
        ├── dynamic-app
        └── lib/          # 动态库文件
            ├── libcommon.dylib
            ├── libmathlib.dylib
            └── libstringlib.dylib
```

## 快速开始

### 1. 编译项目

```bash
./scripts/build.sh
```

编译脚本会：
- 清理之前的构建产物
- 编译标准版本应用（包含调试信息）
- 编译精简版本应用（去除调试信息）
- 生成运行脚本
- 显示文件大小对比

### 2. 运行应用

**运行标准编译版本:**
```bash
./build/static/run.sh
```

**运行精简编译版本:**
```bash
./build/dynamic/run.sh
```

两个版本的功能完全相同，输出结果一致。

## 编译结果对比

### 静态链接版本
| 组件 | 大小 | 说明 |
|------|------|------|
| static-app | 2.3 MB | 所有代码编译进可执行文件 |
| **总计** | **2.3 MB** | 单一文件，独立运行 |

### 动态链接版本
| 组件 | 大小 | 说明 |
|------|------|------|
| dynamic-app | 2.4 MB | 通过 CGO 调用动态库 |
| libcommon.dylib | 1.7 MB | 通用库动态库 |
| libmathlib.dylib | 1.6 MB | 数学库动态库 |
| libstringlib.dylib | 1.7 MB | 字符串库动态库 |
| **总计** | **7.7 MB** | 需要所有文件才能运行 |

### 关键发现

⚠️ **动态链接反而更大**: 在这个例子中，动态链接总体积是静态链接的 **3.15倍** (215% 更大)

**原因:**
- 每个 .dylib 包含独立的 Go Runtime
- CGO 桥接代码开销
- 符号表导出要求
- 动态库格式元数据

## 功能特性

### 库模块

- **common**: 提供日志记录和基础工具函数
- **mathlib**: 提供数学计算功能（加法、乘法、阶乘、求最大值）
- **stringlib**: 提供字符串处理功能（反转、连接、大写转换、单词计数）

### 应用功能

两个应用执行相同的操作：
- 数学运算：加法、乘法、阶乘、求最大值
- 字符串处理：反转、连接、大写转换、单词统计
- 带时间戳的日志输出
- 结果摘要展示

## 技术文档

详细的技术文档位于 `md/` 目录：

- **[项目结构说明.md](md/项目结构说明.md)**: 详细介绍各个模块的功能和依赖关系
- **[构建测试总结.md](md/构建测试总结.md)**: 编译测试结果、性能对比和技术分析

## 主要亮点

1. **Go Workspace**: 使用 go.work 统一管理多个模块
2. **真正的动态库**: 生成实际的 .dylib 文件，可被其他程序调用
3. **CGO 实践**: 完整的 Go 到 C 互操作实现
4. **对比测试**: 静态 vs 动态链接的真实数据对比
5. **自动化构建**: 一键完成所有动态库和应用的编译
6. **详细文档**: 深入的技术分析和最佳实践建议

## 技术栈

- **Go 1.21+**: 支持 Workspace 特性
- **CGO**: C 和 Go 的互操作
- **buildmode=c-shared**: 生成 C 共享库
- **macOS**: 生成 .dylib 动态库（Linux 下为 .so）

## 学习要点

通过这个项目，你可以学习到：

### Go Workspace
- 如何创建和使用 go.work 文件
- 多模块项目的统一管理
- 模块间的依赖关系处理

### 动态库技术
- 使用 `-buildmode=c-shared` 生成动态库
- 为 Go 代码创建 C 导出接口
- CGO 的指针传递限制和解决方案
- 动态库的加载和路径管理

### 对比分析
- 静态链接 vs 动态链接的实际差异
- Go Runtime 在动态库中的行为
- 何时选择静态链接，何时考虑动态链接
- 文件大小和部署复杂度的权衡

### 工程实践
- 自动化构建脚本编写
- 跨平台动态库处理（.so vs .dylib）
- 模块化设计和代码组织
- 技术文档的编写

## 验证动态链接

查看应用依赖的动态库：

**macOS:**
```bash
otool -L build/dynamic/dynamic-app
```

**Linux:**
```bash
ldd build/dynamic/dynamic-app
```

**输出示例（macOS）:**
```
build/dynamic/dynamic-app:
    libcommon.dylib (compatibility version 0.0.0)
    libmathlib.dylib (compatibility version 0.0.0)
    libstringlib.dylib (compatibility version 0.0.0)
    /System/Library/Frameworks/CoreFoundation.framework/...
    /usr/lib/libSystem.B.dylib
```

✅ 确认应用依赖我们生成的三个 .dylib 文件！

## 环境要求

- Go 1.21 或更高版本（支持 Workspace）
- macOS / Linux（动态库生成）
- CGO 支持（通常默认启用）
- Bash Shell（用于运行构建脚本）

**注意:** Windows 用户需要修改脚本以生成 .dll 文件

## 许可证

查看 [LICENSE](LICENSE) 文件了解详情。

---

**项目创建日期**: 2025年11月23日  
**作者**: AoEiuV020
测试golang多模块项目，
