package people

import (
	"context"
	"log/slog"
	"time"
)

func RunWorker(ctx context.Context, ch <-chan *People, repo PeopleRepository, batch int) {
	slog.Debug("Starting worker...")
	defer slog.Debug("Finishing worker...")
	i := 0
	peoples := make([]*People, 0, batch)
	tick := time.NewTicker(1 * time.Second)

	for {
		select {
		case p, ok := <-ch:
			if p.Id != "" {
				peoples = append(peoples, p)
			}
			if i == batch || !ok {
				if err := repo.CreateMany(ctx, peoples); err != nil {
					slog.Error(err.Error())
				}
				i = 0
				peoples = make([]*People, 0, batch)
			}
			i++

			if !ok {
				return
			}

		case <-tick.C:
			if len(peoples) > 0 {
				if err := repo.CreateMany(ctx, peoples); err != nil {
					slog.Error(err.Error())
				}
			}
			i = 0
			peoples = make([]*People, 0, batch)
		}
	}
}
