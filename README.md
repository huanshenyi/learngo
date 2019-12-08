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

# go cover-test
使い方
```text
go tool cover
```

```text
go test -coverprofile=c.out
```

htmlでカバレージ率を見る
```text
go tool cover -html=c.out
```

# ベンチマークテスト(性能テスト)

```text
func BenchmarkSubstr(b *testing.B)  {
	s := "pwwkew"
	ans := 8
	
	for i := 0;i<b.N;i++{
		actual := LengthOfNonRepeatingSubStr(s)
		if actual != ans{
			b.Errorf("got %d for input %s;" + "expected %d",actual,s,ans)
		}	
	}
}
```
コマンドラインで実行
```text
go test -bench .
```
```text
b.ResetTimer()
```

二進数のテスト結果ファイルを生成する

```text
go test -bench . -cpuprofile cpu.out
```
結果ファイルを見る
```text
go tool pprof cpu.out
```
webで性能グラフを見る 
ライブラリダウンロード  http://www.graphviz.org/  binファイルをpathに追加
```text
web
```

## httpサーバーテスト関連

```text
\learngo\errhanding\filelistingserver\errwrapper_test.go
```