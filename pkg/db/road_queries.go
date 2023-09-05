package db

const (
	addToKillList = `insert into manticore.road_card_data_kill(id)
		select max(passport_id) from hwm.hwt_road_data dat
		join cmn.cmn_doc doc on doc.end_date is null and dat.passport_id = doc.id
		where road_id = $1;`

	getNewID = `insert into cmn.cmn_doc (doctype_id, code, name, doc_date, beg_date, end_date)
  		values (4,'Карточка дороги',now(),now(),null,null)
  		returning id;`

	addToDelta = `insert into manticore.road_card_data_delta
			(id, road_id, value_of_the_road_gid, region_gid,
	  	full_name, road_number_full, egrad_number, is_checked,
	      road_category_gid, capacity, speed_limit, json_data)
		values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`
)
