package xo

func PopulationsByOffsetLimit(db XODB, offset, limit int) ([]*Population, error) {
	// sql query
	const sqlstr = `SELECT ` +
		`id, pref_code, pref_name, era_name, era_year, year, population, male_population, female_population ` +
		`FROM populationdb.population ` +
		`LIMIT ?, ?`

	// run query
	XOLog(sqlstr, offset, offset, limit)

	q, err := db.Query(sqlstr, offset, limit)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	var res []*Population

	for q.Next() {
		p := Population{}
		err = q.Scan(&p.ID, &p.PrefCode, &p.PrefName, &p.EraName, &p.EraYear, &p.Year, &p.Population, &p.MalePopulation, &p.FemalePopulation)
		if err != nil {
			return nil, err
		}
		res = append(res, &p)
	}
	return res, nil
}
