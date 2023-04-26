# copilot-handson

GIthubCopilotワークショップ

プレゼンテーション 30m
QA 30m
hands-on 60m

### プレゼンテーション
提案の採用率 35%
copilotによる新しいコードの割合 46%

- コピペではなく合成コード
- 公開コードと一致する150文字以上の提案をブロック
  - `Suggestions matching public code`のこと
  - OSSサジェスト　厳しいほうのポリシー適用される
    - エディタ単位でのコンテキスト取得なので、将来的にもrepo単位などの対応はしないだろう
  - 150文字というのはデフォルト値
- for Businessはそもそもまるっと学習に使わない

- copilot X
  - Chat
  - PullReq
    - 要約とtest不足もみてkるえる
  - Docs
    - 学習方法、コンテキストなど諸々調整中
  - CLI

**エディタ単位でのコンテキスト**

### QA

- ドメイントロールみたいな生成してrepoだけ抱えていくようなの
  - 生成コード量が増えていくと偶然中身が一致することあるよね？
    - copilotは確率論で生成しているので、input同じでも重複することはほぼ〜〜ない。

- ビジネスの利用料
  - 従量課金

- 重複した場合
  - 個人利用　解約される
  - for Businessが有効化される

- for Businessの利点
  - `Allow GitHub to use my code snippets for product improvements`がdefaultで無効
  - 裁判などあったらgithubがサポートはいる

- chat(X)とcopilotのクオリティの差
  - modelは確かに違う。 chatのほうがbaseはgpt4なので拡がりはあるが、copilotは速さ重視。

- 障害？あるの？
  - いまのところ障害無し（すごい）
  - 状態はgithubのstatusで見れる。保証などいまのとこない。

- 将来的にチューニングして提案を寄せるなどのエンプラ対応も検討はしている。


### QK

### hands-on
じゃんけんゲームつくろうぜ

codespaces初体験。立ち上げ結構かかるね。

codetour便利

同じinputしたけど、たしかに別のコードが生成された。

golangにしつつ仕様からcopilotさんにミッチリ書いてもらったが、main関数に入れ込む要件が重すぎて生成動かない
→適当に機能分割できそうな関数単位にするコメントを書く。あとはもうほぼフルオートでコード生成できた。
みたいなモノを発表してTシャツげとぉー

