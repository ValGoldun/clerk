package problem

import (
	"encoding/json"
	"errors"
	"github.com/ValGoldun/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"io"
	"net/http"
)

type Clerk struct {
	logger logger.Logger
}

func New(logger logger.Logger) Clerk {
	return Clerk{
		logger: logger,
	}
}

func (c Clerk) Problem(ctx *gin.Context, err error, metadata Metadata) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case *json.UnmarshalTypeError:
		c.clientProblem(ctx, errors.New("invalid json type"), metadata)
		return
	case *json.SyntaxError:
		c.clientProblem(ctx, errors.New("invalid json"), metadata)
		return
	case validator.ValidationErrors:
		var fields []Field
		for _, field := range e {
			fields = append(fields, Field{Key: field.Field(), Error: field.Tag()})
		}
		c.clientProblemWithFields(ctx, errors.New("validation error"), fields, metadata)
		return
	default:
		if errors.Is(err, io.EOF) {
			c.clientProblem(ctx, errors.New("empty body"), metadata)
			return
		}
		c.serverProblem(ctx, err, metadata)
		return
	}
}
func (c Clerk) serverProblem(ctx *gin.Context, err error, metadata Metadata) {
	c.logger.Error(err.Error(), metadata.LoggerFields()...)

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, Problem{Error: "server problem", Metadata: metadata})
}

func (c Clerk) clientProblem(ctx *gin.Context, err error, metadata Metadata) {
	c.logger.Warn(err.Error(), metadata.LoggerFields()...)

	ctx.AbortWithStatusJSON(http.StatusBadRequest, Problem{Error: err.Error(), Metadata: metadata})
}

func (c Clerk) clientProblemWithFields(ctx *gin.Context, err error, fields []Field, metadata Metadata) {
	c.logger.Warn(err.Error(), metadata.LoggerFields()...)

	ctx.AbortWithStatusJSON(http.StatusBadRequest, Problem{Error: err.Error(), Fields: fields, Metadata: metadata})
}
