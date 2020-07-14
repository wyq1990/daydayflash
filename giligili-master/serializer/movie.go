package serializer

import "giligili/model"

// Movie 视频序列化器
type Movie struct {
	ID           uint   `json:"id"`
	CreatedAt    int64  `json:"created_at"`
	DoubanId     int    `json:"doubanId"`
	MovieName    string `json:"movieName"`
	MovieContent string `json:"movieContent"`
	MovieImg     string `json:"movieImg"`
}

// BuildMovie 序列化电影
func BuildMovie(item model.Movie) Movie {
	return Movie{
		ID:           item.ID,
		DoubanId:     item.DoubanId,
		MovieName:    item.MovieName,
		MovieImg:     item.MovieImg,
		MovieContent: item.MovieContent,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

// BuildMovies 序列化电影列表
func BuildMovies(items []model.Movie) (movies []Movie) {
	for _, item := range items {
		movie := BuildMovie(item)
		movies = append(movies, movie)
	}
	return movies
}
