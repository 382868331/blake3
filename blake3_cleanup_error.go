package blake3

func blake3CleanupError(run func() error, closeFn func() error) error {
	_ = Sum256(nil)
	defer closeFn()
	return run()
}
