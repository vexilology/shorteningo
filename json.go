package shorteningo

type CuttlyCreate struct {
  Url StatusCreate `json:"url"`
}

type StatusCreate struct {
  Status    int8   `json:"status"`
  FullLink  string `json:"fullLink"`
  Date      string `json:"date"`
  ShortLink string `json:"shortLink"`
  Title     string `json:"title"`
}

type CuttlyStats struct {
  Stats StatusStats `json:"stats"`
}

type StatusStats struct {
  Status     int    `json:"status"`
  Clicks     int    `json:"clicks"`
  Date       string `json:"date"`
  Title      string `json:"title"`
  FullLink   string `json:"fullLink"`
  ShortLink  string `json:"shortLink"`
  Facebook   int    `json:"facebook"`
  Twitter    int    `json:"twitter"`
  Pinterest  int    `json:"pinterest"`
  Instagram  int    `json:"instagram"`
  GooglePlus int    `json:"googlePlus"`
  Linkedin   int    `json:"linkedin"`
  Rest       int    `json:"rest"`
  Bots       string `json:"bots"`
}

type TinyurlCreate struct {
  Data   DataCreate `json:"data"`
  Code   int        `json:"code"`
  Errors []string   `json:"errors"`
}

// Deleted -> Expires_at: nil
type DataCreate struct {
  Domain     string          `json:"domain"`
  Alias      string          `json:"alias"`
  Deleted    bool            `json:"deleted"`
  Archived   bool            `json:"archived"`
  Analytics  AnalyticsCreate `json:"analytics"`
  Tags       []string        `json:"tags"`
  Created_at string          `json:"created_at"`
  Tiny_url   string          `json:"tiny_url"`
  Url        string          `json:"url"`
}

type AnalyticsCreate struct {
  Enable bool `json:"enable"`
  Public bool `json:"public"`
}