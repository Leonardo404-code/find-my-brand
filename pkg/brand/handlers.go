package brand

import "net/http"

type Handlers interface {
	Search(w http.ResponseWriter, r *http.Request)
}
