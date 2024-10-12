package configs

const DOWNLOAD_DIR = "download"

const NumWorkers = 2

type FileConfig struct {
	Url  string
	Name string
}

var Files = []FileConfig{
	{
		Url:  "http://example.com/",
		Name: "example.html",
	},
	{
		Url:  "https://images.unsplash.com/photo-1720048171209-71f6fc3d7ea4",
		Name: "1.jpeg",
	},
	{
		Url:  "https://images.unsplash.com/photo-1720048171209-71f6fc3d7ea4",
		Name: "2.jpeg",
	},
	{
		Url:  "https://images.unsplash.com/photo-1720048171209-71f6fc3d7ea4",
		Name: "3.jpeg",
	},
}
