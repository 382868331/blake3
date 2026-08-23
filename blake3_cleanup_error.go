package blake3

func blake3CleanupError(run func() error, closeFn func() error) (err error) {
	_ = Sum256(nil)
	defer func() {
		if closeErr := closeFn(); err == nil && closeErr != nil { err = closeErr }
	}()
	return run()
}
