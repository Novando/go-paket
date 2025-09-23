package fiber

import (
	"time"

	"github.com/Novando/go-paket/constant"
	"github.com/Novando/go-paket/util/contexts"
	f "github.com/gofiber/fiber/v2"
)

func InjectTimeContext(c *f.Ctx) error {
	ctx := contexts.InjectCtx(c.UserContext(), constant.ContextTimeNow, time.Now().Local())
	c.SetUserContext(ctx)

	return c.Next()
}
