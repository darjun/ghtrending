## ghtrending

API to fetch github trending


## usage

get package:

```cmd
$ go get -u github.com/darjun/ghtrending
```

fetch trending repositories：

```golang
import "github.com/darjun/ghtrending"

t := ghtrending.New()
repos, err := t.FetchRepos()
```

fetch trending developers:

```golang
t := ghtrending.New()
developers, err := t.FetchDevelopers()
```

use options：

```golang
// Go weekly trending
t := ghtrending.New(ghtrending.WithWeekly(), ghtrending.WithLanguage("Go"))

// C++ monthly trending
t := ghtrending.New(ghtrending.WithMonthly(), ghtrending.WithLanguage("C++"))

// 中文
t := ghtrending.New(ghtrending.WithSpokenLanguageCode("cn"))
```

you don't need to new object:

```golang
repos := ghtrending.TrendingRepositories()
developers := ghtrending.TrendingDevelopers()
```