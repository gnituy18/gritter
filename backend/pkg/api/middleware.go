package api

import (
	"net/http"

	"github.com/fasthttp/session/v2"
	"github.com/fasthttp/session/v2/providers/memory"
	routing "github.com/qiangxue/fasthttp-routing"
	"go.uber.org/zap"

	"gritter/pkg/context"
	"gritter/pkg/log"
)

var (
	sess *session.Session
)

func init() {
	cfg := session.NewDefaultConfig()
	sess = session.New(cfg)

	provider, err := memory.New(memory.Config{})
	if err != nil {
		log.Global().Fatal(err)
	}

	if err := sess.SetProvider(provider); err != nil {
		log.Global().Fatal(err)
	}
}

func injectContext(rctx *routing.Context) error {
	rctx.Set("ctx", context.Background())
	rctx.Next()
	return nil
}

func logRequest(rctx *routing.Context) error {
	ctx := rctx.Get("ctx").(context.Context)
	ctx.With(zap.String("req", rctx.String())).Info("req")
	rctx.Next()
	return nil
}

func mustAuthUser(rctx *routing.Context) error {
	ctx := rctx.Get("ctx").(context.Context)

	store, err := sess.Get(rctx.RequestCtx)
	if err != nil {
		ctx.Error("session.Session.Get failed in api.GetAuthUser")
		return err
	}
	defer saveStore(rctx, store)

	val := store.Get("userID")
	userID, ok := val.(string)
	if !ok || userID == "" {
		JSON(rctx, http.StatusUnauthorized, nil)
		rctx.Abort()
		return nil
	}

	rctx.Set("userID", userID)

	rctx.Next()

	return nil
}

func saveStore(rctx *routing.Context, store *session.Store) {
	if err := sess.Save(rctx.RequestCtx, store); err != nil {
		log.Global().Error(err)
	}
}
