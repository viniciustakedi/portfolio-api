package flashcards

import (
	"fmt"
	"net/http"
	"strconv"
	response "portfolio/api/utils"

	"github.com/gin-gonic/gin"
)

type FlashcardsController struct {
	svc *FlashcardsService
}

func NewFlashcardsController(svc *FlashcardsService) *FlashcardsController {
	return &FlashcardsController{svc: svc}
}

func validLanguage(s string) bool {
	return s == "en" || s == "es"
}

func validPath(s string) bool {
	return s == "beginner" || s == "intermediate" || s == "advanced"
}

// List GET /flashcards?language=&path=&limit=&skip=
func (c *FlashcardsController) List(ctx *gin.Context) {
	language := ctx.Query("language")
	path := ctx.Query("path")
	if !validLanguage(language) || !validPath(path) {
		response.Error(ctx, "Query parameters language (en|es) and path (beginner|intermediate|advanced) are required.", http.StatusBadRequest)
		return
	}

	limit := int64(20)
	if v := ctx.Query("limit"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil && n > 0 && n <= 100 {
			limit = n
		}
	}
	skip := int64(0)
	if v := ctx.Query("skip"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil && n >= 0 {
			skip = n
		}
	}

	data, err := c.svc.List(ctx.Request.Context(), language, path, limit, skip)
	if err != nil {
		fmt.Println(err.Error())
		response.Error(ctx, "Could not load flashcards.")
		return
	}

	total, err := c.svc.Count(ctx.Request.Context(), language, path)
	if err != nil {
		fmt.Println(err.Error())
		response.Error(ctx, "Could not load flashcards.")
		return
	}

	response.Data(ctx, gin.H{
		"cards": data,
		"total": total,
		"limit": limit,
		"skip":  skip,
	}, "Flashcards loaded.", http.StatusOK)
}

// ListPaths GET /flashcards/paths?language=
func (c *FlashcardsController) ListPaths(ctx *gin.Context) {
	language := ctx.Query("language")
	if !validLanguage(language) {
		response.Error(ctx, "Query parameter language (en|es) is required.", http.StatusBadRequest)
		return
	}

	paths, counts, err := c.svc.ListPaths(ctx.Request.Context(), language)
	if err != nil {
		fmt.Println(err.Error())
		response.Error(ctx, "Could not load paths.")
		return
	}

	type pathOut struct {
		PathDocument
		TotalCards int64 `json:"totalCards"`
	}
	out := make([]pathOut, 0, len(paths))
	for _, p := range paths {
		n := counts[p.Level]
		if n == 0 {
			n = int64(p.TotalCards)
		}
		out = append(out, pathOut{PathDocument: p, TotalCards: n})
	}

	response.Data(ctx, out, "Paths loaded.", http.StatusOK)
}

// GetByID GET /flashcards/:id
func (c *FlashcardsController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	doc, err := c.svc.GetByID(ctx.Request.Context(), id)
	if err != nil {
		response.Error(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	if doc == nil {
		response.Error(ctx, "Flashcard not found.", http.StatusNotFound)
		return
	}
	response.Data(ctx, doc, "Flashcard loaded.", http.StatusOK)
}

// Create POST /flashcards (admin)
func (c *FlashcardsController) Create(ctx *gin.Context) {
	raw, ok := ctx.MustGet("payload").(*CreateFlashcardPayload)
	if !ok {
		response.Error(ctx, "Invalid payload.")
		return
	}

	doc, err := c.svc.Create(ctx.Request.Context(), raw)
	if err != nil {
		fmt.Println(err.Error())
		response.Error(ctx, "Could not create flashcard.")
		return
	}
	response.Data(ctx, doc, "Flashcard created.", http.StatusCreated)
}

// Delete DELETE /flashcards/:id (admin)
func (c *FlashcardsController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	ok, err := c.svc.Delete(ctx.Request.Context(), id)
	if err != nil {
		response.Error(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		response.Error(ctx, "Flashcard not found.", http.StatusNotFound)
		return
	}
	response.Message(ctx, "Flashcard deleted.", http.StatusOK)
}
