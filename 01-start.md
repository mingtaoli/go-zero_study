# go 单体应用体验


## 建立框架

1. **建立文件夹 `01-start`**:
   - 打开命令行工具。
   - 输入命令 `mkdir 01-start` 来创建一个名为 `01-start` 的文件夹。
   - 使用命令 `cd 01-start` 进入这个新建的文件夹。

2. **使用 `goctl` 创建 API 项目**:
   - 在 `01-start` 文件夹内，运行命令 `goctl api new firstapi`。
   - 这将创建一个名为 `firstapi` 的新目录，并在其中生成一些初始文件，包括 API 定义文件和 Go 源代码。

3. **进入项目文件夹**:
   - 使用命令 `cd firstapi` 进入新创建的 `firstapi` 项目文件夹。

4. **使用 `go mod tidy`**:
   - 在 `firstapi` 文件夹内，运行 `go mod tidy` 命令。这会下载并安装所有必要的依赖。

5. **运行 API 服务**:
   - 运行 `go run .` 或 `go run firstapi.go`（取决于生成的文件名）来启动 API 服务。

6. **测试 API**:
   - 打开另一个命令行窗口或使用 Postman 这样的工具来测试您的 API。
   - 发送请求到 API 的端口，通常是 `http://localhost:8888/from/me`（除非您更改了默认端口）。


## 修改路由和逻辑