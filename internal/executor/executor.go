package executor

import "ylambda/internal/model"

type Executor interface {
	Execute(function *model.Function) error
}
