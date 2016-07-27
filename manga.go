package manga

type Backend interface {
	Name() string
	Search(string) ([]SearchResult, error)
	Manga(string) (MangaResult, error)
	Chapter(string) ([]string, error)
}

type SearchResult struct {
	ID              string
	Title           string
	Image           string
	Status          string
	Genres          []string
	LastChapterDate string
	Views           int64
}

type ChapterInfo struct {
	Number string
	Date   string
	Title  string
	ID     string
}

type MangaResult struct {
	Title           string
	Image           string
	Status          string
	Genres          []string
	LastChapterDate string
	Views           int64
	Description     string
	NumChapters     int64
	Chapters        []ChapterInfo
}
