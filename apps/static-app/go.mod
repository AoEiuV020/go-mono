module github.com/AoEiuV020/go-mono/apps/static-app

go 1.21

require (
	github.com/AoEiuV020/go-mono/packages/common v0.0.0
	github.com/AoEiuV020/go-mono/packages/mathlib v0.0.0
	github.com/AoEiuV020/go-mono/packages/stringlib v0.0.0
)

replace (
	github.com/AoEiuV020/go-mono/packages/common => ../../packages/common
	github.com/AoEiuV020/go-mono/packages/mathlib => ../../packages/mathlib
	github.com/AoEiuV020/go-mono/packages/stringlib => ../../packages/stringlib
)
