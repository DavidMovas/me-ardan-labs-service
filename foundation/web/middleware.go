package web

type MidFunc func(handler HandlerFunc) HandlerFunc

func wrapMiddleware(mw []MidFunc, handler HandlerFunc) HandlerFunc {
	for i := len(mw) - 1; i >= 0; i-- {
		mwFunc := mw[i]
		if mwFunc != nil {
			handler = mwFunc(handler)
		}
	}

	return handler
}
