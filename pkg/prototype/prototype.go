package prototype

type Cinema struct {
	Title          string
	DownloadedFile string
}

type Prototype struct {
	Cinema Cinema
}

func (r Prototype) NewCinemaCopy() Cinema {
	newCopy := Cinema{}
	newCopy.Title = r.Cinema.Title
	newCopy.DownloadedFile = r.Cinema.DownloadedFile

	return newCopy
}
