package common

import "net/http"

type FilterHandle func(rw http.ResponseWriter, req *http.Request) error

// 拦截器
type Filter struct {
	// 用来存储需要拦截的URI
	filterMap map[string]FilterHandle
}

func NewFilter() *Filter {
	return &Filter{filterMap: make(map[string]FilterHandle)}
}

func (f *Filter) RegisterFilterUri(uri string, handler FilterHandle) {
	f.filterMap[uri] = handler
}

// 根据uri获取handle
func (f *Filter) GetFilterHandler(uri string) FilterHandle {
	return f.filterMap[uri]
}

type WebHandler func(rw http.ResponseWriter, req *http.Request)

// 执行拦截器
func (f *Filter) Handle(handler WebHandler) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		for path, handle := range f.filterMap {
			if path == r.RequestURI {
				// 执行拦截业务逻辑
				err := handle(rw, r)
				if err != nil {
					rw.Write([]byte(err.Error()))
					return
				}
				break
			}
		}

		handler(rw, r)
	}
}
