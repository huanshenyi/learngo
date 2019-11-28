go pathの書き直し

```text
export GOPATH=path
mkdir path/src
```
goland使用する場合

```
Preferences /Go /GOPATH / ProjectGOPATH
で、現在設定してるGOPATHを追加
```

localでgomodule閉じる

```text
export GO111MODULE=off
```

## GOVENDOR (今はもう使わない)
同じパッケージの違うバージョンの使用
プロジェクト内に `vendor`フォルダを新規して
管理用ツール

```
glide, dep,go dep...など
```

# go mod