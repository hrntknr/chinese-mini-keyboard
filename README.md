# chinese-mini-keyboard

## これは何？

怪しい中華製の左手デバイスの設定を行うためのツール。
linux/mac 上で動作します。

## 動作（すると思われる）デバイス

以下全て同一商品だと思われる。

- [SIKAI CASE 片手キーボード プログラマブルキーボード メカニカルキーボード USB 有線接続 ノブ 1 個付き 6 キー 左手ゲーミングキーボード ライブコンテンツ作成コントローラー カスタム可能なキー 片手 小型キーボード ショットカットキー プログラマー向き DIY キーボード Windows/VISTA/Linux/Mac に適用 オフィス/ゲーム/音楽/メディア/作業制御に最適 (6 キー, 黒)](https://www.amazon.co.jp/gp/product/B09Y5LCZRD)
- [マクロ-Bluetooth キーボード用のプログラム可能なキー 3/6,ホットスワップ,ミニメカニカル Bluetooth キーボード](https://ja.aliexpress.com/item/1005004748558819.html)
  - 作成にあたり動作検証したのはこちらで購入したもの。

![./image.jpg](./image.jpg)

## 利用方法

### ショートカットを設定する

#### 例

```
> sudo ./chinese-mini-keyboard key KEY1 ctrl+c
> sudo ./chinese-mini-keyboard key KEY2 ctrl+v
> sudo ./chinese-mini-keyboard key KEY3 ctrl+t a // ctrl+t → a (ctrl+t → ctrl+a とはならない)
> sudo ./chinese-mini-keyboard mediakey K1_LEFT voldown // つまみ左
> sudo ./chinese-mini-keyboard mediakey K1_CENTER mute // つまみ押
> sudo ./chinese-mini-keyboard mediakey K1_RIGHT voladd // つまみ右
> sudo ./chinese-mini-keyboard mousekey K1_LEFT wheeldown
> sudo ./chinese-mini-keyboard mousekey K1_RIGHT wheelup
```

#### key mode

キーボードのショートカットを設定できます。

```
> sudo ./chinese-mini-keyboard key [button] [keys ...]
```

> 備考：modifier を設定できるのは一番最初のキーのみです。

#### mediakey mode

メディアキーのショートカットを設定できます。(複数不可)

```
> sudo ./chinese-mini-keyboard mediakey [button] [mediakey]
```

#### mousekey mode

マウスのショートカットを設定できます。(複数不可)

```
> sudo ./chinese-mini-keyboard mediakey [button] [mousekey]
```

#### modifier 一覧

| key   | desc                                    |
| ----- | --------------------------------------- |
| ctrl  |                                         |
| shift |                                         |
| alt   |                                         |
| gui   | いわゆる windows キー（mac だと super） |

#### ボタン一覧

| key       | desc                 |
| --------- | -------------------- |
| KEY1      | キーボード１         |
| KEY2      | キーボード２         |
| KEY3      | キーボード３         |
| KEY4      | キーボード４         |
| KEY5      | キーボード５         |
| KEY6      | キーボード６         |
| KEY7      | キーボード７         |
| KEY8      | キーボード８         |
| KEY9      | キーボード９         |
| KEY10     | キーボード１０       |
| KEY11     | キーボード１１       |
| KEY12     | キーボード１２       |
| K1_Left   | スクロール１左回転   |
| K1_Center | スクロール１右回転   |
| K1_RIGHT  | スクロール１押し込み |
| K2_Left   | スクロール２左回転   |
| K2_Center | スクロール２右回転   |
| K2_RIGHT  | スクロール２押し込み |
| K3_Left   | スクロール３左回転   |
| K3_Center | スクロール３右回転   |
| K3_RIGHT  | スクロール３押し込み |

#### キー一覧

| key         | desc                            |
| ----------- | ------------------------------- |
| a ~ z       | キーボードの通常の a-z          |
| 0 ~ 9       | キーボードの通常の 0-9          |
| f1-f12      | キーボードの通常の function key |
| enter       | enter                           |
| esc         |                                 |
| backspace   |                                 |
| tab         |                                 |
| space       |                                 |
| -           |                                 |
| =           |                                 |
| \[          |                                 |
| \]          |                                 |
| \\          |                                 |
| #           |                                 |
| ;           |                                 |
| '           |                                 |
| `           |                                 |
| ,           |                                 |
| .           |                                 |
| /           |                                 |
| capslock    |                                 |
| printscreen |                                 |
| scrolllock  |                                 |
| pause       |                                 |
| insert      |                                 |
| home        |                                 |
| pageup      |                                 |
| delete      |                                 |
| end         |                                 |
| pagedown    |                                 |
| right       | 矢印キー                        |
| left        | 矢印キー                        |
| down        | 矢印キー                        |
| up          | 矢印キー                        |
| numlock     |                                 |
| kp-/        | テンキー                        |
| kp-\*       | テンキー                        |
| kp-+        | テンキー                        |
| kp-enter    | テンキー                        |
| kp-.        | テンキー                        |
| kp-1 ~ kp-9 | テンキーの 1-9                  |

#### メディアキー一覧

| key     | desc |
| ------- | ---- |
| play    |      |
| prev    |      |
| next    |      |
| mute    |      |
| volup   |      |
| voldown |      |

#### マウスキー一覧

| key         | desc |
| ----------- | ---- |
| leftclick   |      |
| rightclick  |      |
| centerclick |      |
| wheelup     |      |
| wheeldown   |      |

### LED のモードを変更する

> 備考: MODE3 以降、色指定に関しては中華ツールに UI 上に存在するものの、送っているパケットに差分は見られませんでした。(情報求む)

```
> sudo ./chinese-mini-keyboard led MODE0
> sudo ./chinese-mini-keyboard led MODE1
> sudo ./chinese-mini-keyboard led MODE2
```
