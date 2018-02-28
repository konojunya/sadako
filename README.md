# sadako

git pullをフックして貞子の音楽を流す。

![](https://raw.githubusercontent.com/konojunya/sadako-git-pull/master/media/sadako.png)

## Usage

first:

```
$ go get github.com/konojunya/sadako
```

貞子をgitリポジトリに住ませる

```
$ sadako set
```

貞子をgitリポジトリから抹消する

```
$ sadako rm
```

## Development

```
$ dep ensure
```

sadakoの音を流すスクリプトをbuild

```
$ make build
```

sadakoコマンドの実行ファイルをbuildする

```
$ make install
```

## Auther

[@konojunya](https://twitter.com/konojunya)

## Collaborator

- [@kinokoruumu](https://github.com/kinokoruumu)
- [@mattn](https://github.com/mattn)
