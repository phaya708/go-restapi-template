# Go REST API Template
## 起動方法
```
docker compose build
```

```
docker compose up
```


## その他コマンド一覧
※`server`ディレクトリ上で実行可能
### DB関連
#### seed
```
make seed
```
#### reset
```
make reset
```

### テスト関連
#### 特定のテストを実行
```
make test path=[テスト対象のパス]
```
#### すべてのテストを実行
```
make test-all
```

## 参考サイト
- https://www.youtube.com/watch?v=BvzjpAe3d4g
- https://zenn.dev/a_ichi1/articles/4b113d4c46857a