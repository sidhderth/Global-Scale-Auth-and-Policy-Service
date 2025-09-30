package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/open-policy-agent/opa/rego"
)

type OPAMiddleware struct {
	Query rego.PreparedEvalQuery
}

func NewOPAMiddleware(policyPath string) (*OPAMiddleware, error) {
	ctx := context.Background()

	policy, err := os.ReadFile(policyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read policy file: %w", err)
	}

	r := rego.New(
		rego.Query("data.authz.allow"),
		rego.Module(policyPath, string(policy)),
	)

	pq, err := r.PrepareForEval(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	return &OPAMiddleware{Query: pq}, nil
}

func (m *OPAMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		claims, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "claims not found"})
			return
		}

		input := map[string]interface{}{
			"claims": claims,
			"method": c.Request.Method,
			"path":   strings.Split(strings.Trim(c.Request.URL.Path, "/"), "/"),
		}

		results, err := m.Query.Eval(ctx, rego.EvalInput(input))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "policy evaluation error"})
			return
		}

		if len(results) == 0 || !results.Allowed() {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied"})
			return
		}

		c.Next()
	}
}