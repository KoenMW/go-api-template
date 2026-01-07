package migration

import "context"

type Migration interface {
	Up(context.Context) error
}
