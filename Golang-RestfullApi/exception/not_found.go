package exception

type NotFoundLol struct {
	Gagal string
}

func NotFound(err string) NotFoundLol {
	return NotFoundLol{
		Gagal: err,
	}
}
