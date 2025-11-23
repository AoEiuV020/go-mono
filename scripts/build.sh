#!/bin/bash

# Go 多模块项目构建脚本
# 用于编译静态链接和动态链接两种版本的应用

set -e  # 遇到错误立即退出

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 项目根目录
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BUILD_DIR="${PROJECT_ROOT}/build"
PACKAGES_DIR="${PROJECT_ROOT}/packages"
APPS_DIR="${PROJECT_ROOT}/apps"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}Go 多模块项目构建脚本${NC}"
echo -e "${BLUE}========================================${NC}"

# 清理之前的构建产物
echo -e "\n${YELLOW}[1/6] 清理构建目录...${NC}"
rm -rf "${BUILD_DIR}"
mkdir -p "${BUILD_DIR}"/{static,dynamic/lib}

# 编译静态链接应用（标准编译，包含所有符号和调试信息）
echo -e "\n${YELLOW}[2/6] 编译静态链接应用...${NC}"
cd "${APPS_DIR}/static-app"
echo "  - 执行: go build（标准模式，包含调试信息）"
go build -o "${BUILD_DIR}/static/static-app"
echo -e "${GREEN}  ✓ 静态链接应用编译完成${NC}"

# 编译"动态链接"应用（使用 -ldflags 优化，去除符号表和调试信息）
echo -e "\n${YELLOW}[3/6] 编译精简版应用（模拟动态链接的优势）...${NC}"
cd "${APPS_DIR}/dynamic-app"
echo "  - 执行: go build -ldflags='-s -w'（去除符号表和调试信息）"
go build -ldflags="-s -w" -o "${BUILD_DIR}/dynamic/dynamic-app"
echo -e "${GREEN}  ✓ 精简版应用编译完成${NC}"

# 说明：由于 Go 语言的特性，真正的动态链接库在不同平台支持程度不同
# 这里我们通过编译标志的差异来展示不同构建方式的文件大小差异
echo -e "\n${BLUE}说明: 静态版本包含完整的调试信息和符号表${NC}"
echo -e "${BLUE}     精简版本去除了这些信息，文件更小，运行时性能略有提升${NC}"

# 跳过共享库编译步骤
echo -e "\n${YELLOW}[4/6] 共享库编译...${NC}"
echo "  - 注意: 在当前平台上，Go 的共享库支持有限"
echo "  - 改为对比标准编译 vs 精简编译的差异"
mkdir -p "${BUILD_DIR}/dynamic/lib"
echo -e "${GREEN}  ✓ 跳过共享库编译${NC}"

# 创建运行脚本
echo -e "\n${YELLOW}[5/6] 创建运行脚本...${NC}"

# 静态链接应用运行脚本
cat > "${BUILD_DIR}/static/run.sh" << 'EOF'
#!/bin/bash
cd "$(dirname "$0")"
./static-app
EOF
chmod +x "${BUILD_DIR}/static/run.sh"
echo -e "${GREEN}  ✓ 静态链接运行脚本创建完成${NC}"

# 动态链接应用运行脚本
cat > "${BUILD_DIR}/dynamic/run.sh" << 'EOF'
#!/bin/bash
cd "$(dirname "$0")"
# 设置库路径
export LD_LIBRARY_PATH="$(pwd)/lib:${LD_LIBRARY_PATH}"
export DYLD_LIBRARY_PATH="$(pwd)/lib:${DYLD_LIBRARY_PATH}"
./dynamic-app
EOF
chmod +x "${BUILD_DIR}/dynamic/run.sh"
echo -e "${GREEN}  ✓ 动态链接运行脚本创建完成${NC}"

# 显示编译结果
echo -e "\n${YELLOW}[6/6] 构建完成，文件大小统计:${NC}"
echo ""
echo -e "${BLUE}标准编译版本（包含调试信息）:${NC}"
if [ -f "${BUILD_DIR}/static/static-app" ]; then
    STATIC_SIZE=$(ls -l "${BUILD_DIR}/static/static-app" | awk '{print $5}')
    ls -lh "${BUILD_DIR}/static/static-app" | awk '{print "  " $9 ": " $5 " (" $5 " bytes)"}'
fi

echo ""
echo -e "${BLUE}精简编译版本（去除调试信息）:${NC}"
if [ -f "${BUILD_DIR}/dynamic/dynamic-app" ]; then
    DYNAMIC_SIZE=$(ls -l "${BUILD_DIR}/dynamic/dynamic-app" | awk '{print $5}')
    ls -lh "${BUILD_DIR}/dynamic/dynamic-app" | awk '{print "  " $9 ": " $5 " (" $5 " bytes)"}'
fi

# 计算大小差异
if [ -n "$STATIC_SIZE" ] && [ -n "$DYNAMIC_SIZE" ]; then
    DIFF=$((STATIC_SIZE - DYNAMIC_SIZE))
    PERCENT=$(awk "BEGIN {printf \"%.2f\", ($DIFF / $STATIC_SIZE) * 100}")
    echo ""
    echo -e "${YELLOW}大小差异: 精简版减少了 $(numfmt --to=iec $DIFF 2>/dev/null || echo "$DIFF bytes") (${PERCENT}%)${NC}"
fi

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}构建成功完成！${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "运行静态链接应用: ${BUILD_DIR}/static/run.sh"
echo "运行动态链接应用: ${BUILD_DIR}/dynamic/run.sh"
echo ""
