# twitter-bulk-deleter

特定の期間内のツイートを削除するコマンドラインアプリケーションです。

Goで実装されています。バイナリが必要なら適当にビルドしてください。

## usage

### twitter developer account

Twitterの開発者用アカウントおよびアプリ登録が必要なので、各自申請してください。

アプリ権限は「Read and write」が必要です。

API keyとAPI secret keyを控え、Access tokenおよびAccess token secretを発行してください（自分用のtokenはダッシュボードから発行可能）。

### get tweet archive

Twitterから取得できるツイートアーカイブが必要です。

公式にログインし、「設定とプライバシー」→「Twitterデータ」→「アーカイブをダウンロード（要パスワード）」から、ツイートアーカイブをダウンロードしてください。

アーカイブは30日に1回リクエストできます。ダウンロードできるようになるまで数日程度かかります。

アーカイブをダウンロード後、zipファイルを解凍し、dataフォルダ内の `tweet.js` を控えておき、このプロジェクトのルートディレクトリ（このREADMEと同じ階層）に配置します。

ファイルの冒頭の `window.YTD.tweet.part0 = ` という部分を消しておき、ファイル全体がjson（配列）形式になるようにしてください。

### setting .env file

.env.sampleをもとに、必要事項を追記してください。

ファイルは実行ディレクトリに設置します。

```
$ cp .env.sample .env
```

* ACCESS_TOKEN
  * アプリのAPI key
* ACCESS_TOKEN_SECRET
  * アプリのAPI secret key
* CONSUMER_KEY
  * ユーザごとのAccess token
* CONSUMER_SECRET
  * ユーザごとのAccess token secret

### execute

ドライラン

```
$ go run cmd/deleter/main.go
or
$ go run cmd/deleter/main.go 'yyyy-mm-dd hh:mm:ss' 'yyyy-mm-dd hh:mm:ss'
```

実行

```
$ go run cmd/deleter/main.go execute
or
$ go run cmd/deleter/main.go 'yyyy-mm-dd hh:mm:ss' 'yyyy-mm-dd hh:mm:ss' execute
```

実行すると開始時刻、終了時刻と確認の入力が促されるので従ってください（コマンド実行時に日時を入力する場合は確認のみ）。

終了時刻は開始時刻より後にする必要があります。`2006-01-02 15:04:05` のような形式で入力してください。

コマンド実行時に日時を入力する場合、日時をクォーテーションで囲ってください。

## note

大体1秒で5ツイートくらい削除できるようです。気長にお待ち下さい。

ツイートデータを丸々読んで処理するので、あまりに投稿データが多いとメモリ不足等で落ちるかもしれません。

アーカイブデータ取得後のツイートを削除することはできません。

同一の期間を指定して複数回実行しても削除に失敗するだけで害はありませんが、削除リクエスト自体はほぼ同等の時間がかかります。

アプリ化すると色々大変そうなのであまりやる気がありません。

これをベースに改造したりWebサービスとして公開するなどお好きにどうぞ。