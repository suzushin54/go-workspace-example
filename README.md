# go-workspace-example

複数のサービスと連携するプロジェクトの検証用サンプル。

特定の SaaS, そしてそのサービスごとに異なるコンテナを用意して利用する想定。

その際の共通パッケージの管理方法を検証した。


## go.work example
```go
go 1.22

use (
	./cmd/saas-a/feature-a
	./cmd/saas-a/feature-b
	.
)
```