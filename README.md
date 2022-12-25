### GOPATHを確認し、正しいディレクトリに当プロジェクトを配置する
まずは以下のようにGOPATHを調べる
```
miyoshishinya@miyoshiyanoAir2 chat % go env GOPATH
/Users/miyoshishinya/go
```
→`/Users/miyoshishinya/go`だとわかる
このディレクトリ配下にsrcフォルダを作成する

```
cd /Users/miyoshishinya/go
mkdir src
```

そのsrc配下に当プロジェクトをcloneしてくる

もしGOPATHが設定されていない場合は`~/.zshrc`にGOPATHの設定を書き込む
ちなみに僕のGOPATHの設定は以下のようになっている
```
miyoshishinya@miyoshiyanoAir2 chat % cat ~/.zshrc
GOPATH=$HOME/go
PATH=$PATH:$HOME/go/bin:$GOPATH/bin:
PATH=$PATH:/Users/miyoshishinya/Desktop/OwaraiGo/src:
PATH=$PATH://Users/miyoshishinya/go:

miyoshishinya@miyoshiyanoAir2 chat % echo $HOME
/Users/miyoshishinya
```

からの`source ~/.zshrc`を忘れず行う


ディレクトリに以下のファイルとフォルダがあるのを確認
```
go_docker
├── docker-compose.yml
├── app
│   ├── Dockerfile
│   └── src
│       ├── article
│       │   └── article.go
│       ├── go.mod
│       ├── //go.sum
│       └── main.go
└── mysql
    ├── .env
    ├── Dockerfile
    └── init
        └── create_table.sh
```

`docker network create golang_chat_app_network`を行ってネットワークを手動で作成する
create_table.shファイルのパーミッションを645にする（末尾を5にすることでこのテーブルを作成してくれるシェルをgoが実行できるようにする）

参考：https://zenn.dev/ajapa/articles/443c396a2c5dd1