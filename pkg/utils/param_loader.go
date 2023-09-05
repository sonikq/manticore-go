package utils

import (
	"fmt"
)

func LoadParams(layout string, a ...interface{}) []byte {

	result := fmt.Sprintf(layout, a...)

	return []byte(result)
}

func LoadWFSParams(layout string, index string, limit int, maxMatches int, zoom int, scaleFactor float64, points []float64) []byte {
	/*var geom string

	if editMode == 0 {
		if zoom <= 7 {
			index = models.IdxWFS1000
			geom = "geom_1000"
		} else if zoom <= 9 {
			index = models.IdxWFS100
			geom = "geom_100"
		} else if zoom < 14 {
			index = models.IdxWFS10
			geom = "geom_10"
		} else {
			index = models.IdxWFSRaw
			geom = "geom"
		}
	} else {
		index = models.IdxWFSRaw
		geom = "geom"
	} */

	x1, y1, x2, y2 := points[0], points[1], points[2], points[3]

	if scaleFactor != 1 {
		calcCoordsBySF(&x1, &y1, &x2, &y2, scaleFactor)
	}

	result := fmt.Sprintf(
		layout, index, limit, maxMatches,
		zoom, x1, x2, y1, y2,
		zoom, x1, x2, y1, y2,
		zoom, x1, x2, y1, y2,
		zoom, x1, x2, y1, y2,
	)

	return []byte(result)
}

// calcCoordsBySF calculates coordinates depending on scale factor
func calcCoordsBySF(x1, y1, x2, y2 *float64, sf float64) {
	dx := (*x2 - *x1) * sf
	dy := (*y2 - *y1) * sf

	*x1 -= dx / 2
	*y1 -= dy / 2
	*x2 += dx / 2
	*y2 += dy / 2
}
