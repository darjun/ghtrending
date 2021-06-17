package ghtrending

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Repository represent a repository in the trending list.
type Repository struct {
	Author  string
	Name    string
	Link    string
	Desc    string
	Lang    string
	Stars   int
	Forks   int
	Add     int
	BuiltBy []string
}

// Developer represent a developer in the developer trending list.
type Developer struct {
	Name        string
	Username    string
	PopularRepo string
	Desc        string
}

// Fetcher defines methods to fetch trending repos and developers
type Fetcher interface {
	FetchRepos() ([]*Repository, error)
	FetchDevelopers() ([]*Developer, error)
}

type trending struct {
	opts options
}

func loadOptions(opts ...option) options {
	o := options{
		GitHubURL: "http://github.com",
	}
	for _, option := range opts {
		option(&o)
	}

	return o
}

// New returns a Fetcher
func New(opts ...option) Fetcher {
	return &trending{
		opts: loadOptions(opts...),
	}
}

// FetchRepos fetch all repositories from  GitHub trending.
func (t trending) FetchRepos() ([]*Repository, error) {
	resp, err := http.Get(fmt.Sprintf("%s/trending/%s?spoken_language_code=%s&since=%s", t.opts.GitHubURL, t.opts.Language, t.opts.SpokenLang, t.opts.DateRange))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	repos := make([]*Repository, 0, 10)
	doc.Find(".Box .Box-row").Each(func(i int, s *goquery.Selection) {
		repo := &Repository{}

		// author name link
		titleSel := s.Find("h1 a")
		repo.Author = strings.Trim(titleSel.Find("span").Text(), "/\n ")
		repo.Name = strings.TrimSpace(titleSel.Contents().Last().Text())
		relativeLink, _ := titleSel.Attr("href")
		if len(relativeLink) > 0 {
			repo.Link = t.opts.GitHubURL + relativeLink
		}

		// desc
		repo.Desc = strings.TrimSpace(s.Find("p").Text())

		var langIdx, addIdx, builtByIdx int
		spanSel := s.Find("div>span")
		if spanSel.Size() == 2 {
			// language not exist
			langIdx = -1
			addIdx = 1
		} else {
			builtByIdx = 1
			addIdx = 2
		}

		// language
		if langIdx >= 0 {
			repo.Lang = strings.TrimSpace(spanSel.Eq(langIdx).Text())
		} else {
			repo.Lang = "unknown"
		}

		// add
		addParts := strings.SplitN(strings.TrimSpace(spanSel.Eq(addIdx).Text()), " ", 2)
		repo.Add, _ = strconv.Atoi(addParts[0])

		// builtby
		spanSel.Eq(builtByIdx).Find("a>img").Each(func(i int, img *goquery.Selection) {
			src, _ := img.Attr("src")
			repo.BuiltBy = append(repo.BuiltBy, src)
		})

		// stars forks
		aSel := s.Find("div>a")
		starStr := strings.TrimSpace(aSel.Eq(-2).Text())
		star, _ := strconv.Atoi(strings.Replace(starStr, ",", "", -1))
		repo.Stars = star
		forkStr := strings.TrimSpace(aSel.Eq(-1).Text())
		fork, _ := strconv.Atoi(strings.Replace(forkStr, ",", "", -1))
		repo.Forks = fork

		repos = append(repos, repo)
	})

	return repos, nil
}

// FetchRepos fetch all developers from  GitHub trending.
func (t trending) FetchDevelopers() ([]*Developer, error) {
	resp, err := http.Get(fmt.Sprintf("%s/trending/developers?lanugage=%s&since=%s", t.opts.GitHubURL, t.opts.Language, t.opts.DateRange))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	developers := make([]*Developer, 0, 10)
	doc.Find(".Box .Box-row").Each(func(i int, s *goquery.Selection) {
		developer := &Developer{}

		// name username
		developer.Name = strings.TrimSpace(s.Last().Find("div>div>h1>a").Text())
		developer.Username = strings.TrimSpace(s.Last().Find("div>div>p>a").Text())

		// popular repo
		developer.PopularRepo = strings.TrimSpace(s.Last().Find("div>div>article>h1>a").Text())
		developer.Desc = strings.TrimSpace(s.Last().Find("div>div>article").Children().Last().Text())

		developers = append(developers, developer)
	})

	return developers, nil
}

// TrendingRepositories fetch all repositories from  GitHub trending.
func TrendingRepositories(opts ...option) ([]*Repository, error) {
	return New(opts...).FetchRepos()
}

// TrendingRepositories fetch all developers from  GitHub trending.
func TrendingDevelopers(opts ...option) ([]*Developer, error) {
	return New(opts...).FetchDevelopers()
}
