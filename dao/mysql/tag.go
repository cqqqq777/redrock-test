package mysql

import (
	g "redrock-test/global"
	"redrock-test/model"
)

func TagList() (data []*model.Tag, err error) {
	data = make([]*model.Tag, 0)
	rows, err := g.Mdb.Query("select tid,tag from tags")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			tid int64
			tag string
		)
		err = rows.Scan(&tid, &tag)
		if err != nil {
			return nil, err
		}
		data = append(data, &model.Tag{Tid: tid, Tag: tag})
	}
	return
}
