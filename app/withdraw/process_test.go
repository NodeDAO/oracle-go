package withdraw

import (
	"context"
	"github.com/NodeDAO/oracle-go/app"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProcessReport(t *testing.T) {
	ctx := context.Background()
	app.InitServer(ctx, "../../conf/config-dev.yaml")
	w := new(WithdrawHelper)
	err := w.ProcessReport(ctx)
	require.NoError(t, err)
}
