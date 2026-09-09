package transport

// func JSONResponse(msg string, err error) errorResponse {
// 	return errorResponse{
// 		Message: msg,
// 		Details: err.Error(),
// 	}
// }

type service interface {
	HandleSaveURL(url string, alias string) (string, string, string, error)
	HandlePatchURL(newUrl string, alias string) (string, string, error)
	HandleDeleteURL(alias string) (error)
	HandleRedirectURL(alias string) (string, error)
}

type HTTPHandlers struct {
	service service
}

func InitHTTPHandlers(service service) *HTTPHandlers {
	return &HTTPHandlers{service: service}
}


// func New(urlRepo UrlRepository) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		logger := r.Context().Value(middleware.LoggerContextKey).(*slog.Logger)
// 		logger.Debug("Вызов /url New обработчика")
// 		var req SaveURLRequest
// 		if err := render.DecodeJSON(r.Body, &req); err != nil {
// 			render.JSON(w, r, ErrorResponse("failed to decode request body", err))
// 			return 
// 		}

// 		if err := validator.New().Struct(req); err != nil {
// 			render.JSON(w, r, ErrorResponse("failed to validate request body", err))
// 			return 
// 		}


// 		if req.Alias == "" {
// 			req.Alias = randomizeAlias(6)
// 		}

// 		surl, err := urlRepo.SaveURL(req.Url, req.Alias)
// 		if err != nil {
// 			render.JSON(w, r, ErrorResponse("failed to short url", err))
// 			return 
// 		}

// 		render.JSON(w, r, SaveURLResponse{
// 			Url: surl,
// 		})

// 	}
// }