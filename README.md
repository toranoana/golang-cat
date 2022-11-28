
# golang-cat

GO言語で作成したハイライトできるcatコマンド

導入: `go install github.com/yuki-kano-lab/golang-cat` またはreleaseからバイナリの取得

実行方法: `golang-cat [options...] target_file_path`

デプロイ:
```
gobump (major|minor|patch) -w -v -r .
git tag v$(make app-version)
git push origin v$(make app-version)
```
