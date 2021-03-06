package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathURL struct {
	Path string
	URL  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		pathFromRequest := r.URL.Path

		if p := pathsToUrls[pathFromRequest]; p != "" {
			http.Redirect(w, r, p, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	pathMap, err := YmlToMap(yml)

	if err != nil {
		return nil, err
	}

	return MapHandler(pathMap, fallback), nil

}

// YmlToMap will take yml bytestring and return a map out of it.
func YmlToMap(yamlBytes []byte) (map[string]string, error) {
	ymlMap := make(map[string]string)

	var pathURLSlice []pathURL
	err := yaml.Unmarshal(yamlBytes, &pathURLSlice)

	if err != nil {
		return nil, err
	}

	for _, val := range pathURLSlice {
		ymlMap[val.Path] = val.URL
	}

	return ymlMap, nil
}
