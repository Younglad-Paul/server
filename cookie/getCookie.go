
package cookie

import (
    "context"
    "net/http"
)

type contextKey string

const ResponseWriterKey = contextKey("responseWriter")

func GetResponseWriter(ctx context.Context) http.ResponseWriter {
    return ctx.Value(ResponseWriterKey).(http.ResponseWriter)
}