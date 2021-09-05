package api

import (
	"io"
	"mail2rss/internal/config"
	"mail2rss/internal/feed"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	config.InitConfig()
	tag := strings.TrimPrefix(r.URL.Path, "/")
	resp, err := getMail(tag)
	if err != nil {
		writeError(w, err)
		return
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		writeError(w, err)
		return
	}
	defer resp.Body.Close()
	rss, err := feed.MakeRSS(string(data), tag)
	if err != nil {
		writeError(w, err)
		return
	}
	_, _ = w.Write([]byte(rss))
}

func getMail(tag string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, config.EndPoint, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("apikey", config.APIKEY)
	q.Add("namespace", config.NAMESPACE)
	q.Add("pretty", "true")
	if tag != "" {
		q.Add("tag", tag)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}
