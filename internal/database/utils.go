package database

func SetDefaultPageLimit(page *int64, limit *int64) {
	if *page < 0 {
		*page = 0
	}

	if *limit < 0 || *limit > 50 {
		*limit = 20
	}
}
