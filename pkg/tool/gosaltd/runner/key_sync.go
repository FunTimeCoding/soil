package runner

func (r *Runner) keySync() error {
	return r.attempt(
		func() error {
			keys, e := r.salt.Keys()

			if e != nil {
				return e
			}

			accepted := make(map[string]bool, len(keys.Minions))

			for _, m := range keys.Minions {
				accepted[m] = true
			}

			for _, m := range keys.MinionsDenied {
				if accepted[m] {
					r.logger.Structured("key_delete", "minion", m)

					if _, f := r.salt.DeleteKey(m); f != nil {
						return f
					}
				}
			}

			for _, m := range keys.MinionsPre {
				r.logger.Structured("key_accept", "minion", m)

				if _, f := r.salt.AcceptKey(m); f != nil {
					return f
				}
			}

			return nil
		},
	)
}
