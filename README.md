# boilerplate

テンプレートレポジトリから、新規プロジェクトを作成するためのツールです。

## Installation

```
go install github.com/showcase-gig-platform/boilerplate@latest
```

## How to use

テンプレートレポジトリをローカルに準備し、書き出し先を指定して gen を開始します。

```
cd /YOUR_WORKING_DIR/
git clone https://github.com/showcase-gig-platform/scg-go-boilerplate
boilerplate gen --src ./scg-go-boilerplate --dst ./my-new-service
```

指定されたテンプレートレポジトリ内にある `boilerplate.yaml` ファイルを参照し、新プロジェクトで置換する文字列を調べます。

```yaml
project: BoilerPlate
targets:
  - name: repository name
    string: scg-go-boilerplate
  - name: service name
    string: boilerplate-service
ignore_prefixes:
  - ".git/"
  - ".idea/"
  - "boilerplate.yaml"
  - "boilerplate.md"
```

このような yaml がある場合、 gen 実行時に以下のように聞かれます。

```
New project name (example: BoilerPlate): MyNewService
New repository name (example: scg-go-boilerplate): mynew-server
New service name (example: boilerplate-service): mynew-service
```

project name は特殊なキーワードです。 `BoilerPlate` のように2単語以上をキャメルケースで指定します。

- BoilerPlate -> MyNewService
- boiler-plate -> my-new-service
- boiler_plate -> my_new_service

のように、 Snake, Kebab, Camel などのケースに対応してよしなに置換してくれます
(詳細な生成される変換ルールは `replace_rule.go` 内の `func NewReplaceRulesFromCamelCase()` を参照)

`targets` は単純な置換です。 project name より優先的に先に置換されます。
上記の例だと以下のように置換されます。

- scg-go-boilerplate -> mynew-server
- boilerplate-service -> mynew-service

置換対象は、ファイル内の文字列と、ファイル・ディレクトリ名の文字列です。

`ignore_prefixes` は置換対象から除外するパスの prefix を指定します。例えば `.git` などのコピーしたくないファイルを指定します。