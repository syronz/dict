# dict
[![Go Report Card](https://goreportcard.com/badge/github.com/syronz/dict)](https://goreportcard.com/report/github.com/syronz/dict)

Dict is a Go package for translating terms

## Installation
1. By running below go command you can install it, if you use 'go mod' you can skip this step

```sh
go get github.com/syronz/dict
```

2. Import it in your code:

```go
import "github.com/syronz/dict"
```

## Terms
Dict support JSON and TOML, we recommend toml because it prevent duplication.
In version 1.1 this package just support English, Kurdish and Arabic. The style of terms.toml should be like below

terms.toml
```toml
["route not found"]
en = 'route not found'
ku = 'route booni nie'
ar = 'الطريق غير موجود'

[unauthorized]
en = 'unauthorized'
ku = 'daxl naboo'
ar = 'غير مصرح'

["%v shouldn't be more than %v"]
en = "%v shouldn't be more than %v"
ku = "%v nabe la %v ziatr be"
ar = "يجب ألا يكون٪ v أكثر من٪ v"
```

For loading terms just one time run the below method

```go
dict.Init("./terms.toml", true)
```

## Translation

### Simple term
For tranlate simple sentence/word without params it be like below

```go
translated := dict.T("route not found", dict.Ku)
fmt.Println("route not found: ", translated)
```

Output:
```sh
route not found:  route booni nie
```

### Terms with param
For translate sentences with params, dynamic terms. you can pass arbitrary params as arguments
In case you want translate params should cast them as `dict.R`

```go
phrase := "%v shouldn't be more than %v"
translated = dict.T(phrase, dict.Ku, dict.R("age"), 18)
fmt.Println(phrase+": ", translated)
```

Output:
```sh
%v shouldn't be more than %v:  taman nabe la 18 ziatr be
```

### Not exist terms
If the word not exist inside the TOML or JSON file, around the output '!!!' would be added. For example

```go
translated = dict.T("term not exist", dict.Ku)
fmt.Println("term not exist: ", translated)
```

Output:
```sh
term not exist:  !!! term not exist !!!
```

For prevent adding '!!!' around not existed terms we can use SafeTranslate, in this case empty and false will be returned

```go
var ok bool
translated, ok = dict.SafeTranslate("term not exist", dict.Ku)
fmt.Println("term not exist: ", translated, ok)
```

Output: 
```sh
term not exist:    false
```

